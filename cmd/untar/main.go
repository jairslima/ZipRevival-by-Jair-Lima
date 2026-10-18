package main

import (
	"flag"
	"fmt"
	"os"

	zipsafe "github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNTAR Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: untar [options] archive.tar [destdir]\n")
		fmt.Fprintf(os.Stderr, "Extracts TAR, TAR.GZ, TAR.BZ2 and TAR.XZ. Symbolic and hard links are never created.\n")
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

	_, ignoradas, err := zipsafe.ExtractTar(src, dest, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "untar error: %v\n", err)
		os.Exit(1)
	}
	if ignoradas > 0 {
		fmt.Fprintf(os.Stderr, "untar: %d entrada(s) foram ignoradas por segurança.\n", ignoradas)
		os.Exit(2)
	}
	fmt.Println("Extracted", src, "successfully.")
}
