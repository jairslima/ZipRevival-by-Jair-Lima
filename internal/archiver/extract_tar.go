// ZipRevival by Jair Lima: extração segura de TAR, TAR.GZ, TAR.BZ2 e TAR.XZ.
package archiver

import (
	"archive/tar"
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
)

// ExtractTar extrai um arquivo TAR (puro, gzip, bzip2 ou xz, detectado pelos primeiros bytes) para
// destDir. Todo caminho passa por SafeJoin. Links simbólicos e físicos NÃO são criados (são o vetor
// clássico de fuga da pasta de destino); cada um é ignorado e contado em ignoradas.
func ExtractTar(src, destDir string, log io.Writer) (extraidos, ignoradas int, err error) {
	f, err := os.Open(src)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	cabecalho, _ := br.Peek(6)
	var r io.Reader = br
	switch {
	case len(cabecalho) >= 2 && cabecalho[0] == 0x1f && cabecalho[1] == 0x8b:
		gz, e := gzip.NewReader(br)
		if e != nil {
			return 0, 0, e
		}
		defer gz.Close()
		r = gz
	case len(cabecalho) >= 3 && string(cabecalho[:3]) == "BZh":
		r = bzip2.NewReader(br)
	case len(cabecalho) >= 6 && string(cabecalho) == "\xfd7zXZ\x00":
		xr, e := xz.NewReader(br)
		if e != nil {
			return 0, 0, e
		}
		r = xr
	}

	tr := tar.NewReader(r)
	for {
		h, e := tr.Next()
		if e == io.EOF {
			return extraidos, ignoradas, nil
		}
		if e != nil {
			return extraidos, ignoradas, e
		}
		alvo, e := SafeJoin(destDir, h.Name)
		if e != nil {
			fmt.Fprintf(log, "entrada ignorada (%v)\n", e)
			ignoradas++
			continue
		}
		switch h.Typeflag {
		case tar.TypeDir:
			if e := os.MkdirAll(alvo, 0o755); e != nil {
				return extraidos, ignoradas, e
			}
		case tar.TypeReg:
			if e := os.MkdirAll(filepath.Dir(alvo), 0o755); e != nil {
				return extraidos, ignoradas, e
			}
			out, e := os.OpenFile(alvo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(h.Mode)&0o755|0o600)
			if e != nil {
				return extraidos, ignoradas, e
			}
			_, e = io.Copy(out, tr)
			out.Close()
			if e != nil {
				return extraidos, ignoradas, e
			}
			extraidos++
			fmt.Fprintf(log, "  extracting: %s\n", h.Name)
		case tar.TypeSymlink, tar.TypeLink:
			fmt.Fprintf(log, "link ignorado (não criado por segurança): %s -> %s\n", h.Name, h.Linkname)
			ignoradas++
		default:
			ignoradas++
		}
	}
}
