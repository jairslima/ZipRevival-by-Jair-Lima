package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
)

func main() {
	var (
		extract bool
		destDir string
	)

	flag.BoolVar(&extract, "extract", false, "Extract files")
	flag.StringVar(&destDir, "d", "", "Destination directory")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "PKUNZIP Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: pkunzip [options] zipfile\n")
		fmt.Fprintf(os.Stderr, "  -extract  Extract files (default)\n")
		fmt.Fprintf(os.Stderr, "  -d        Destination directory\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	src := args[0]
	if _, err := os.Stat(src); os.IsNotExist(err) {
		if len(src) < 4 || src[len(src)-4:] != ".zip" {
			src += ".zip"
		}
	}

	opts := archiver.UnzipOptions{
		DestDir:   destDir,
		ListOnly:  false,
		Overwrite: 0, // Prompt default
	}

	err := archiver.UnzipFile(src, &opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pkunzip error: %v\n", err)
		os.Exit(1)
	}
}
