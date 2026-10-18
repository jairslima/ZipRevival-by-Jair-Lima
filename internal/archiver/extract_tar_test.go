package archiver

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func montarTar(t *testing.T, entradas []tar.Header, conteudos map[string]string, gz bool) string {
	t.Helper()
	var buf bytes.Buffer
	var w io.Writer = &buf
	var gw *gzip.Writer
	if gz {
		gw = gzip.NewWriter(&buf)
		w = gw
	}
	tw := tar.NewWriter(w)
	for _, h := range entradas {
		h := h
		if h.Typeflag == tar.TypeReg {
			h.Size = int64(len(conteudos[h.Name]))
		}
		if err := tw.WriteHeader(&h); err != nil {
			t.Fatal(err)
		}
		if h.Typeflag == tar.TypeReg {
			tw.Write([]byte(conteudos[h.Name]))
		}
	}
	tw.Close()
	if gw != nil {
		gw.Close()
	}
	arq := filepath.Join(t.TempDir(), "teste.tar")
	if err := os.WriteFile(arq, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	return arq
}

func TestExtractTarBloqueiaFugaELinks(t *testing.T) {
	pai := t.TempDir()
	dest := filepath.Join(pai, "saida")
	os.MkdirAll(dest, 0o755)
	arq := montarTar(t, []tar.Header{
		{Name: "ok/a.txt", Typeflag: tar.TypeReg, Mode: 0o644},
		{Name: "../fugiu.txt", Typeflag: tar.TypeReg, Mode: 0o644},
		{Name: "/abs.txt", Typeflag: tar.TypeReg, Mode: 0o644},
		{Name: "elo", Typeflag: tar.TypeSymlink, Linkname: "../../etc"},
		{Name: "duro", Typeflag: tar.TypeLink, Linkname: "ok/a.txt"},
	}, map[string]string{"ok/a.txt": "ola", "../fugiu.txt": "mal", "/abs.txt": "mal"}, true)

	var log bytes.Buffer
	extraidos, ignoradas, err := ExtractTar(arq, dest, &log)
	if err != nil {
		t.Fatalf("erro inesperado: %v", err)
	}
	if extraidos != 1 || ignoradas != 4 {
		t.Errorf("extraidos=%d ignoradas=%d, esperado 1 e 4. Log: %s", extraidos, ignoradas, log.String())
	}
	if b, err := os.ReadFile(filepath.Join(dest, "ok", "a.txt")); err != nil || string(b) != "ola" {
		t.Errorf("arquivo legítimo não foi extraído: %v %q", err, b)
	}
	if _, err := os.Stat(filepath.Join(pai, "fugiu.txt")); err == nil {
		t.Error("FALHA DE SEGURANÇA: ../fugiu.txt foi criado fora da pasta de destino")
	}
	if _, err := os.Lstat(filepath.Join(dest, "elo")); err == nil {
		t.Error("link simbólico foi criado")
	}
}

func TestExtractTarSemCompressao(t *testing.T) {
	dest := t.TempDir()
	arq := montarTar(t, []tar.Header{{Name: "pasta/", Typeflag: tar.TypeDir, Mode: 0o755}, {Name: "pasta/x.txt", Typeflag: tar.TypeReg, Mode: 0o644}},
		map[string]string{"pasta/x.txt": "conteudo"}, false)
	n, ign, err := ExtractTar(arq, dest, io.Discard)
	if err != nil || n != 1 || ign != 0 {
		t.Fatalf("n=%d ign=%d err=%v", n, ign, err)
	}
}
