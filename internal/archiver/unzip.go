package archiver

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/yeka/zip"
)

// UnzipOptions define as opções de extração baseadas nas flags clássicas.
type UnzipOptions struct {
	DestDir   string
	ListOnly  bool
	Overwrite int // 0 = prompt, 1 = always (-o), 2 = never (-n)
	Password  string
}

// UnzipFile extrai ou lista o conteúdo de um arquivo ZIP.
func UnzipFile(src string, opts *UnzipOptions) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo zip %s: %w", src, err)
	}
	defer r.Close()

	if opts.ListOnly {
		fmt.Printf("Archive:  %s\n", src)
		fmt.Println("  Length      Date    Time    Name")
		fmt.Println("---------  ---------- -----   ----")
		var totalSize uint64
		for _, f := range r.File {
			fmt.Printf("%9d  %s   %s\n", f.UncompressedSize64, f.ModTime().Format("2006-02-01 15:04"), f.Name)
			totalSize += f.UncompressedSize64
		}
		fmt.Println("---------                     -------")
		fmt.Printf("%9d                     %d files\n", totalSize, len(r.File))
		return nil
	}

	dest := opts.DestDir
	if dest == "" {
		dest = "."
	}

	// Criar o diretório de destino, se não existir
	if dest != "." {
		os.MkdirAll(dest, 0755)
	}

	for _, f := range r.File {
		err := extractFile(f, dest, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

func extractFile(f *zip.File, destDir string, opts *UnzipOptions) error {
	if f.IsEncrypted() && opts.Password != "" {
		f.SetPassword(opts.Password)
	}
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// Segurança contra ZipSlip
	path := filepath.Join(destDir, f.Name)
	cleanDest := filepath.Clean(destDir)
	cleanPath := filepath.Clean(path)
	
	if destDir != "." {
		if !strings.HasPrefix(cleanPath, cleanDest+string(os.PathSeparator)) {
			return fmt.Errorf("caminho de arquivo ilegal: %s", path)
		}
	} else {
		if strings.HasPrefix(cleanPath, "..") {
			return fmt.Errorf("caminho de arquivo ilegal: %s", path)
		}
	}

	if f.FileInfo().IsDir() {
		os.MkdirAll(path, f.Mode())
		return nil
	}

	os.MkdirAll(filepath.Dir(path), 0755)

	if _, err := os.Stat(path); err == nil {
		if opts.Overwrite == 2 {
			return nil // Never overwrite (-n)
		} else if opts.Overwrite == 0 {
			// Prompt básico
			fmt.Printf("replace %s? [y]es, [n]o, [A]ll, [N]one: ", path)
			var resp string
			fmt.Scanln(&resp)
			resp = strings.ToLower(strings.TrimSpace(resp))
			if resp == "a" || resp == "all" {
				opts.Overwrite = 1 // Transforma em Always (-o) para as próximas passadas
			} else if resp == "n" || resp == "none" || resp == "no" {
				return nil
			} else if resp != "y" && resp != "yes" {
				return nil
			}
		}
	}

	fmt.Printf("  inflating: %s\n", f.Name)
	fDest, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer fDest.Close()

	_, err = io.Copy(fDest, rc)
	return err
}
