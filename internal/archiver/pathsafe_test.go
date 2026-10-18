package archiver

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestSafeJoinAceitaNomesNormais(t *testing.T) {
	dest := t.TempDir()
	for _, nome := range []string{"a.txt", "pasta/b.txt", "pasta\\c.txt", "./d.txt", "a/./b", "pasta/sub/", "arquivo..txt", "..sombra"} {
		got, err := SafeJoin(dest, nome)
		if err != nil {
			t.Errorf("SafeJoin(%q) rejeitou nome legítimo: %v", nome, err)
			continue
		}
		if !strings.HasPrefix(got, dest) {
			t.Errorf("SafeJoin(%q) = %q, fora de %q", nome, got, dest)
		}
	}
}

func TestSafeJoinRejeitaFugas(t *testing.T) {
	dest := t.TempDir()
	for _, nome := range []string{
		"", "../fora.txt", "a/../../fora.txt", "..\\fora.txt", "a\\..\\..\\fora.txt", "/etc/passwd", "\\windows\\x",
		"C:\\Windows\\x", "C:/x", "c:x", "a/..", "..", "ok\x00/../x",
	} {
		if got, err := SafeJoin(dest, nome); err == nil {
			t.Errorf("SafeJoin(%q) deveria falhar, mas devolveu %q", nome, got)
		}
	}
}

func TestSafeJoinResultadoFicaDentroDoDestino(t *testing.T) {
	dest := t.TempDir()
	got, err := SafeJoin(dest, "x/y/../z.txt")
	if err == nil {
		t.Fatalf("esperava erro por conter .., recebeu %q", got)
	}
	got, err = SafeJoin(dest, "x/y/z.txt")
	if err != nil || filepath.Dir(filepath.Dir(got)) != filepath.Join(dest, "x") && !strings.HasPrefix(got, dest) {
		t.Fatalf("resultado inesperado: %q %v", got, err)
	}
}
