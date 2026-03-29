package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/bodgit/sevenzip"
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

	r, err := sevenzip.OpenReader(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "un7z error: %v\n", err)
		os.Exit(1)
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}
		os.MkdirAll(filepath.Dir(path), os.ModePerm)

		rc, err := f.Open()
		if err != nil {
			fmt.Fprintf(os.Stderr, "un7z open error: %v\n", err)
			continue
		}

		out, err := os.Create(path)
		if err != nil {
			rc.Close()
			continue
		}

		_, err = io.Copy(out, rc)
		out.Close()
		rc.Close()
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "un7z copy error: %v\n", err)
		} else {
			fmt.Printf("  Extracting: %s\n", f.Name)
		}
	}
	fmt.Println("Extracted", src, "successfully.")
}