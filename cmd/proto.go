package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var packageRe, packageRefRe, importRe *regexp.Regexp

var REPO string = "gopkg.in/freddierice/go-hproto.v1"

type Proto struct {
	Name    string
	OldPath string
	NewPath string
	Package *ProtoLine
	Imports []*ProtoLine
}

type ProtoLine struct {
	Num      int
	Contents string
}

func init() {
	packageRe = regexp.MustCompile("^package (.*);")
	importRe = regexp.MustCompile("^import \"(.*)\";")
	packageRefRe = regexp.MustCompile("hadoop\\.(common|yarn|hdfs|mapreduce)")
}

func GetProtos(hpath string) ([]*Proto, error) {

	// get the file locations
	protoFiles, err := GetProtoFiles(hpath)
	if err != nil {
		return nil, err
	}

	// collect the metadata
	protos := make([]*Proto, len(protoFiles))
	for i, pf := range protoFiles {
		protos[i] = &Proto{
			Name:    filepath.Base(pf),
			OldPath: pf,
		}
	}

	// pick information from the file contents
	for _, p := range protos {
		buf, err := ioutil.ReadFile(p.OldPath)
		if err != nil {
			return nil, err
		}
		r := bytes.NewBuffer(buf)

		lineNum := 0
		for {
			line, err := r.ReadString('\n')
			if err != nil && err != io.EOF {
				return nil, err
			}

			matches := packageRe.FindStringSubmatch(line)
			if len(matches) != 0 {
				p.Package = &ProtoLine{
					Num:      lineNum,
					Contents: matches[1],
				}
			}

			matches = importRe.FindStringSubmatch(line)
			if len(matches) != 0 {
				p.Imports = append(p.Imports, &ProtoLine{
					Num:      lineNum,
					Contents: matches[1],
				})
			}

			if err == io.EOF {
				break
			}
			lineNum++
		}
	}

	return protos, nil
}

// Rename each component of the protos
func RenameProtos(protos []*Proto) {

	importMap := map[string]string{}

	// rename the packages and set new paths
	for _, p := range protos {
		stem := strings.Replace(p.Package.Contents, "hadoop.", "", 1)
		p.Package.Contents = "hproto." + stem

		extra := "/"
		stemIter := p.OldPath
		for {
			stemIter = filepath.Dir(stemIter)
			parentDir := filepath.Base(stemIter)
			if parentDir != "proto" {
				extra = "/" + parentDir + extra
			} else {
				break
			}
		}
		p.NewPath = strings.Replace(stem, ".", "/", -1) + extra + p.Name
		importMap[p.Name] = REPO + "/" + p.NewPath
	}

	for _, p := range protos {
		for _, imp := range p.Imports {
			if newImp, ok := importMap[imp.Contents]; ok {
				imp.Contents = newImp
			}
		}
	}
}

func ConvertLines(r *bytes.Buffer, w *bytes.Buffer, n int) {
	for ; n != 0; n-- {
		line, err := r.ReadString('\n')

		matches := packageRefRe.FindStringSubmatch(line)
		if len(matches) != 0 {
			w.WriteString(strings.Replace(line, "hadoop."+matches[1],
				"hproto."+matches[1], -1))
		} else {
			w.WriteString(line)
		}

		if err != nil {
			break
		}
	}
}

func TrashLines(r *bytes.Buffer, n int) {
	for ; n != 0; n-- {
		if _, err := r.ReadBytes('\n'); err != nil {
			break
		}
	}
}

func WriteProtos(protos []*Proto) error {
	for _, p := range protos {
		buf, err := ioutil.ReadFile(p.OldPath)
		if err != nil {
			return err
		}
		r := bytes.NewBuffer(buf)
		w := &bytes.Buffer{}

		// write out syntax so the compiler does not complain
		w.WriteString("syntax = \"proto2\";\n")

		ConvertLines(r, w, p.Package.Num)
		TrashLines(r, 1)
		w.WriteString("package " + p.Package.Contents + ";\n")
		currLine := p.Package.Num + 1

		for _, imp := range p.Imports {
			ConvertLines(r, w, imp.Num-currLine)
			TrashLines(r, 1)
			w.WriteString("import \"" + imp.Contents + "\";\n")
			currLine = imp.Num + 1
		}

		ConvertLines(r, w, -1)

		if err := os.MkdirAll(filepath.Dir(p.NewPath), 0755); err != nil {
			return err
		}
		if err := ioutil.WriteFile(p.NewPath, w.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

// GetProtoFiles finds all the non-test proto files in a directory.
func GetProtoFiles(hpath string) ([]string, error) {
	paths := []string{}
	err := filepath.Walk(hpath,
		func(p string, info os.FileInfo, err error) error {
			if strings.HasSuffix(p, ".proto") &&
				!strings.Contains(p, "/test/") {
				paths = append(paths, p)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

// CompileProtos compiles the gathered protos and deletes the proto files. This
// function assumes that we are in the root of the project and in the basename
// directory of the REPO.
func CompileProtos(protos []*Proto) {
	rlen := len(strings.Split(REPO, "/"))
	newdir := strings.Repeat("../", rlen)
	os.Chdir(newdir)

	// collect the protobuf files that have the same package
	protoDirs := map[string][]string{}
	for _, p := range protos {
		protoDir := filepath.Dir(p.NewPath)
		protoDirs[protoDir] = append(protoDirs[protoDir], REPO+"/"+p.NewPath)
	}

	// build compilation commands
	cmds := []*exec.Cmd{}
	for _, paths := range protoDirs {
		args := make([]string, len(paths)+1)
		args[0] = "--go_out=plugins=grpc:."
		copy(args[1:], paths)
		cmds = append(cmds, exec.Command("protoc", args...))
	}

	// compile the files
	for _, c := range cmds {
		c.Start()
	}
	for _, c := range cmds {
		c.Wait()
	}

	// remove the proto files
	for _, p := range protos {
		os.Remove(REPO + "/" + p.NewPath)
	}
}
