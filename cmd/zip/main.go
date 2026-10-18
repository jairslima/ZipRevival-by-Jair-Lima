package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
)

func main() {
	var (
		recursive bool
		move      bool
		comp0     bool
		comp1     bool
		comp9     bool
		password  string
	)

	flag.BoolVar(&recursive, "r", false, "Recurse into directories")
	flag.BoolVar(&move, "m", false, "Move files into zip (delete OS files)")
	flag.BoolVar(&comp0, "0", false, "Store only")
	flag.BoolVar(&comp1, "1", false, "Compress faster")
	flag.BoolVar(&comp9, "9", false, "Compress better")
	flag.StringVar(&password, "P", "", "Password to encrypt")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Copyright (c) 2026 Zip Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: zip [-options] [-b path] [-t mmddyyyy] [-n suffixes] [zipfile list] [-xi list]\n")
		fmt.Fprintf(os.Stderr, "  The default action is to add or replace zipfile entries from list, which\n")
		fmt.Fprintf(os.Stderr, "  can include the special name - to compress standard input.\n")
		fmt.Fprintf(os.Stderr, "  -r  recurse into directories\n")
		fmt.Fprintf(os.Stderr, "  -m  move into zipfile (delete OS files)\n")
		fmt.Fprintf(os.Stderr, "  -0  store only\n")
		fmt.Fprintf(os.Stderr, "  -1  compress faster\n")
		fmt.Fprintf(os.Stderr, "  -9  compress better\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	dest := args[0]
	// Adiciona a extensão .zip se não tiver
	if len(dest) < 4 || dest[len(dest)-4:] != ".zip" {
		dest += ".zip"
	}

	files := args[1:]

	opts := archiver.ZipOptions{
		Recursive:        recursive,
		Move:             move,
		CompressionLevel: 8, // Default
		Password:         password,
	}

	if comp0 {
		opts.CompressionLevel = 0
	} else if comp1 {
		opts.CompressionLevel = 1
	} else if comp9 {
		opts.CompressionLevel = 9
	}

	err := archiver.ZipFiles(dest, files, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "zip error: %v\n", err)
		os.Exit(1)
	}
}
