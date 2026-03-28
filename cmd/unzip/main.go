package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jairlima/zip-revival/internal/archiver"
)

func main() {
	var (
		destDir  string
		list     bool
		over     bool
		never    bool
		password string
	)

	flag.StringVar(&destDir, "d", "", "extract files into exdir")
	flag.BoolVar(&list, "l", false, "list files (short format)")
	flag.BoolVar(&over, "o", false, "overwrite existing files without prompting")
	flag.BoolVar(&never, "n", false, "never overwrite existing files")
	flag.StringVar(&password, "P", "", "password to decrypt")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "UnZip Revival by Jair Lima\n")
		fmt.Fprintf(os.Stderr, "Usage: unzip [-Z] [-opts[modifiers]] file[.zip] [list] [-x xlist] [-d exdir]\n")
		fmt.Fprintf(os.Stderr, "  Default action is to extract files in list, except those in xlist, to exdir;\n")
		fmt.Fprintf(os.Stderr, "  file[.zip] may be a wildcard.  -Z => ZipInfo mode (\"unzip -Z\" for usage).\n")
		fmt.Fprintf(os.Stderr, "  -l  list files (short format)        -d  extract files into exdir\n")
		fmt.Fprintf(os.Stderr, "  -o  overwrite files WITHOUT prompting  -n  never overwrite files\n")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	src := args[0]
	// Se não tiver extensão, tentar com .zip
	if _, err := os.Stat(src); os.IsNotExist(err) {
		if len(src) < 4 || src[len(src)-4:] != ".zip" {
			src += ".zip"
		}
	}

	opts := archiver.UnzipOptions{
		DestDir:   destDir,
		ListOnly:  list,
		Overwrite: 0,
	}

	if over {
		opts.Overwrite = 1
	} else if never {
		opts.Overwrite = 2
	}

	err := archiver.UnzipFile(src, &opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unzip error: %v\n", err)
		os.Exit(1)
	}
}