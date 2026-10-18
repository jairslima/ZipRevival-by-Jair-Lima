package main

import (
	"flag"
	"fmt"
	zipsafe "github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
	"io"
	"os"
	"path/filepath"

	"github.com/bogdannbv/go-cabfile/cabfile"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNCAB Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: uncab archive.cab [dest_dir]\n")
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

	f, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uncab error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	cab, err := cabfile.New(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uncab parse error: %v\n", err)
		os.Exit(1)
	}

	files := cab.FileList()
	ignoradas := 0
	for _, fileName := range files {
		reader, err := cab.Content(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uncab content error %s: %v\n", fileName, err)
			continue
		}

		path, err := zipsafe.SafeJoin(destDir, fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uncab: entrada ignorada (%v)\n", err)
			ignoradas++
			continue
		}
		os.MkdirAll(filepath.Dir(path), 0755)

		out, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uncab create error %s: %v\n", fileName, err)
			continue
		}

		_, err = io.Copy(out, reader)
		out.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "uncab copy error %s: %v\n", fileName, err)
		} else {
			fmt.Printf("  inflating: %s\n", fileName)
		}
	}
	if ignoradas > 0 {
		fmt.Fprintf(os.Stderr, "uncab: %d entrada(s) com caminho inseguro foram ignoradas.\n", ignoradas)
		os.Exit(2)
	}
}
