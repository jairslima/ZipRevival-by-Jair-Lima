package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kdomanski/iso9660/util"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UNISO Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: uniso archive.iso [dest_dir]\n")
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
		fmt.Fprintf(os.Stderr, "uniso error: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	err = util.ExtractImageToDirectory(f, destDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uniso extract error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Extracted", src, "successfully.")
}
