package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mholt/archiver/v3"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNRAR Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: unrar <command> [-<switch 1> -<switch N>] archive [files...]\n")
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

	rar := archiver.NewRar()

	switch cmd {
	case "x", "e":
		// Para simplificar, o 'x' e 'e' inicialmente extrairão para a pasta atual
		err := rar.Unarchive(src, ".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "unrar error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Extracting from", src, "OK")
	case "l":
		err := rar.Walk(src, func(f archiver.File) error {
			fmt.Printf("%10d %s\n", f.Size(), f.Name())
			return nil
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unrar error: %v\n", err)
			os.Exit(1)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}
}