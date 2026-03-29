package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mholt/archiver/v3"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UN7Z Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: un7z archive.7z [dest_dir]\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	src := args[0]
	destDir := "."
	if len(args) > 1 {
		destDir = args[1]
	}

	z := archiver.NewSevenZip()
	err := z.Unarchive(src, destDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "un7z error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Extracted", src, "successfully.")
}