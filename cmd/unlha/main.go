package main

import (
	"flag"
	"fmt"
	zipsafe "github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
	"io"
	"os"
	"path/filepath"

	"github.com/koron-go/lha"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNLHA Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: unlha archive.lzh [dest_dir]\n")
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
		fmt.Fprintf(os.Stderr, "unlha error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	r := lha.NewReader(f)
	ignoradas := 0
	for {
		hdr, err := r.NextHeader()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "unlha read error: %v\n", err)
			os.Exit(1)
		}

		path, err := zipsafe.SafeJoin(destDir, hdr.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unlha: entrada ignorada (%v)\n", err)
			ignoradas++
			continue
		}

		if hdr.Method == "-lhd-" || len(hdr.Name) > 0 && hdr.Name[len(hdr.Name)-1] == '/' {
			os.MkdirAll(path, 0755)
			continue
		}

		os.MkdirAll(filepath.Dir(path), 0755)

		out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unlha extract error: %v\n", err)
			continue
		}

		_, err = r.Decode(out)
		out.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "unlha copy error: %v\n", err)
		} else {
			fmt.Printf("  inflating: %s\n", hdr.Name)
		}
	}
	if ignoradas > 0 {
		fmt.Fprintf(os.Stderr, "unlha: %d entrada(s) com caminho inseguro foram ignoradas.\n", ignoradas)
		os.Exit(2)
	}
}
