// ZipRevival by Jair Lima: junção segura de caminhos para extração de arquivos compactados.
package archiver

import (
	"fmt"
	"path/filepath"
	"strings"
)

// SafeJoin junta destDir com o nome de uma entrada de arquivo compactado e garante que o resultado
// fica DENTRO de destDir. Rejeita nomes absolutos, com letra de unidade, com "..", com bytes nulos ou
// que escapem da pasta de destino (ataque "Zip Slip"). Barras invertidas são tratadas como separador
// também fora do Windows, porque arquivos criados no Windows as usam.
func SafeJoin(destDir, name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("nome de entrada vazio")
	}
	if strings.ContainsRune(name, 0) {
		return "", fmt.Errorf("nome de entrada com byte nulo")
	}
	n := strings.ReplaceAll(name, "\\", "/")
	if strings.HasPrefix(n, "/") || filepath.IsAbs(name) || filepath.VolumeName(name) != "" || (len(n) >= 2 && n[1] == ':') {
		return "", fmt.Errorf("caminho absoluto não permitido: %q", name)
	}
	for _, parte := range strings.Split(n, "/") {
		if parte == ".." {
			return "", fmt.Errorf("caminho com \"..\" não permitido: %q", name)
		}
	}
	base, err := filepath.Abs(destDir)
	if err != nil {
		return "", err
	}
	alvo := filepath.Join(base, filepath.FromSlash(n))
	rel, err := filepath.Rel(base, alvo)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("caminho fora da pasta de destino: %q", name)
	}
	return alvo, nil
}
