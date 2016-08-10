package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: \n")
	fmt.Fprintf(os.Stderr, "  %v <path to hadoop source>\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	hpath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatalf("could not find absolute file path for argument")
	}

	olddir, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get the current working directory")
	}

	if err := os.Chdir(filepath.Dir(os.Args[0]) + "/../"); err != nil {
		log.Fatalf("could not change the working directory to the root path " +
			"of this project")
	}

	protos, err := GetProtos(hpath)
	if err != nil {
		os.Chdir(olddir)
		log.Fatalf("could not collect the proto files: %v", err)
	}

	RenameProtos(protos)

	if err := WriteProtos(protos); err != nil {
		os.Chdir(olddir)
		log.Fatalf("could not write the protos out: %v", err)
	}

	CompileProtos(protos)

	os.Chdir(olddir)
}
