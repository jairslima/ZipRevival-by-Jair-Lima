package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mholt/archiver/v3"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNTAR Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: untar [options] archive.tar\n")
		fmt.Fprintf(os.Stderr, "Extracts TAR, TAR.GZ, TAR.BZ2, TAR.XZ, etc.\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	src := args[0]
	dest := "."
	if len(args) > 1 {
		dest = args[1]
	}

	err := archiver.Unarchive(src, dest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "untar error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Extracted", src, "successfully.")
}