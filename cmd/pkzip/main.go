package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
)

func main() {
	var (
		add       bool
		recursive bool
		move      bool
	)

	flag.BoolVar(&add, "add", false, "Add files")
	flag.BoolVar(&recursive, "r", false, "Recurse directories")
	flag.BoolVar(&move, "m", false, "Move files")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "PKZIP Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: pkzip [options] zipfile [files...]\n")
		fmt.Fprintf(os.Stderr, "  -add  Add files (default)\n")
		fmt.Fprintf(os.Stderr, "  -r    Recurse directories\n")
		fmt.Fprintf(os.Stderr, "  -m    Move files (delete originals)\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	dest := args[0]
	if !strings.HasSuffix(strings.ToLower(dest), ".zip") {
		dest += ".zip"
	}

	var files []string
	if len(args) > 1 {
		files = args[1:]
	} else {
		// No files specified, behave like a helper
		fmt.Println("pkzip: no files specified")
		os.Exit(1)
	}

	opts := archiver.ZipOptions{
		Recursive:        recursive,
		Move:             move,
		CompressionLevel: 8,
	}

	err := archiver.ZipFiles(dest, files, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pkzip error: %v\n", err)
		os.Exit(1)
	}
}
