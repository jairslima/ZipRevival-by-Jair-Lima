// ZipRevival by Jair Lima: extração e listagem seguras de RAR.
package archiver

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/nwaples/rardecode/v2"
)

// Limites contra "bombas" de descompressão (um RAR pode pedir um dicionário de vários GB).
const (
	MaxDicionarioRar     int64 = 512 << 20 // 512 MiB
	MaxTamanhoArquivoRar int64 = 8 << 30   // 8 GiB por arquivo extraído
)

func opcoesRar() []rardecode.Option {
	return []rardecode.Option{rardecode.MaxDictionarySize(MaxDicionarioRar), rardecode.MaxReadFileSize(MaxTamanhoArquivoRar)}
}

// ExtractRar extrai um RAR em destDir. Se semCaminhos for true ("unrar e"), usa só o nome base de
// cada arquivo. Links são ignorados e contados em ignoradas.
func ExtractRar(src, destDir string, semCaminhos bool, log io.Writer) (extraidos, ignoradas int, err error) {
	rc, err := rardecode.OpenReader(src, opcoesRar()...)
	if err != nil {
		return 0, 0, err
	}
	defer rc.Close()
	for {
		h, e := rc.Next()
		if e == io.EOF {
			return extraidos, ignoradas, nil
		}
		if e != nil {
			return extraidos, ignoradas, e
		}
		nome := h.Name
		if semCaminhos {
			nome = path.Base(filepath.ToSlash(nome))
		}
		alvo, e := SafeJoin(destDir, nome)
		if e != nil {
			fmt.Fprintf(log, "entrada ignorada (%v)\n", e)
			ignoradas++
			continue
		}
		if h.LinkType != 0 {
			fmt.Fprintf(log, "link ignorado (não criado por segurança): %s\n", h.Name)
			ignoradas++
			continue
		}
		if h.IsDir {
			if !semCaminhos {
				if e := os.MkdirAll(alvo, 0o755); e != nil {
					return extraidos, ignoradas, e
				}
			}
			continue
		}
		if e := os.MkdirAll(filepath.Dir(alvo), 0o755); e != nil {
			return extraidos, ignoradas, e
		}
		out, e := os.OpenFile(alvo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
		if e != nil {
			return extraidos, ignoradas, e
		}
		_, e = io.Copy(out, rc)
		out.Close()
		if e != nil {
			return extraidos, ignoradas, e
		}
		extraidos++
		fmt.Fprintf(log, "  extracting: %s\n", h.Name)
	}
}

// ListRar escreve "tamanho nome" de cada entrada.
func ListRar(src string, w io.Writer) error {
	rc, err := rardecode.OpenReader(src, opcoesRar()...)
	if err != nil {
		return err
	}
	defer rc.Close()
	for {
		h, e := rc.Next()
		if e == io.EOF {
			return nil
		}
		if e != nil {
			return e
		}
		fmt.Fprintf(w, "%10d %s\n", h.UnPackedSize, h.Name)
	}
}
