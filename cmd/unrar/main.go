package main

import (
	"flag"
	"fmt"
	"os"

	zipsafe "github.com/jairslima/ZipRevival-by-Jair-Lima/internal/archiver"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNRAR Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: unrar <command> archive [destdir]\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  e  Extract files without archived paths\n")
		fmt.Fprintf(os.Stderr, "  x  Extract files with full path\n")
		fmt.Fprintf(os.Stderr, "  l  List archive contents\n")
	}

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	src := os.Args[2]
	dest := "."
	if len(os.Args) > 3 {
		dest = os.Args[3]
	}

	switch cmd {
	case "x", "e":
		_, ignoradas, err := zipsafe.ExtractRar(src, dest, cmd == "e", os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unrar error: %v\n", err)
			os.Exit(1)
		}
		if ignoradas > 0 {
			fmt.Fprintf(os.Stderr, "unrar: %d entrada(s) foram ignoradas por segurança.\n", ignoradas)
			os.Exit(2)
		}
		fmt.Println("Extracting from", src, "OK")
	case "l":
		if err := zipsafe.ListRar(src, os.Stdout); err != nil {
			fmt.Fprintf(os.Stderr, "unrar error: %v\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}
