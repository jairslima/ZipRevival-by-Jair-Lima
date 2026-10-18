# Threat model: ZipRevival

## O que é e por onde entra dado não confiável
Conjunto de utilitários de linha de comando em Go que reproduzem programas clássicos de compactação
(`zip`, `unzip`, `pkzip`, `pkunzip`, `unrar`, `untar`, `un7z`, `unlha`, `uniso`, `uncab`).
**Todo arquivo compactado passado aos extratores deve ser tratado como hostil.** Os pontos de entrada
são o argumento `archive` de cada comando em `cmd/*/main.go` e o pacote `internal/archiver`.

## O que mais importa
1. **Extração fora da pasta de destino** (path traversal / "Zip Slip"): nomes com `..`, caminhos
   absolutos, letra de unidade (`C:`), barras invertidas e bytes nulos. Todos os extratores devem
   passar os nomes por `archiver.SafeJoin` (`internal/archiver/pathsafe.go`). `unzip` tem uma
   verificação própria em `internal/archiver/unzip.go`.
2. **Links**: `untar` e `unrar` nunca criam link simbólico ou físico. `unzip`, `un7z`, `unlha` e
   `uncab` gravam sempre arquivos comuns. Qualquer caminho em que um link seja criado e depois
   atravessado é uma falha.
3. **Esgotamento de recursos** (bombas de descompressão, dicionários RAR enormes, tamanhos
   declarados falsos): o RAR tem limites (`MaxDicionarioRar`, `MaxTamanhoArquivoRar`). Os demais
   formatos ainda não têm limite de tamanho total extraído.
4. **Corrupção de memória nas bibliotecas de terceiros** (`bodgit/sevenzip`, `koron-go/lha`,
   `bogdannbv/go-cabfile`, `kdomanski/iso9660`, `yeka/zip`, `nwaples/rardecode/v2`) quando
   chamadas por este código.

## Como exercitar
- Build em `/src/bin/<comando>`; exemplo: `/src/bin/unzip arquivo.zip destino`, `/src/bin/untar a.tar.gz destino`,
  `/src/bin/un7z a.7z destino`, `/src/bin/unlha a.lzh destino`, `/src/bin/uncab a.cab destino`,
  `/src/bin/uniso a.iso destino`, `/src/bin/unrar x a.rar destino`.
- `go test ./internal/archiver/` cobre `SafeJoin`, ZIP e TAR com entradas maliciosas.

## Severidade
- Escrita fora da pasta de destino, ou criação de link que leve a isso: **alta** (crítica se permitir
  sobrescrever arquivo executável ou de configuração do usuário).
- Travamento, estouro de memória ou consumo de disco ilimitado com um arquivo pequeno: **média**.
- Erros de lógica sem efeito de segurança (mensagens, formatação, opções ignoradas): fora de escopo.

## Deixar de lado
- Vulnerabilidades da biblioteca padrão do Go e da toolchain (por exemplo `net/http`): este código não
  usa rede.
- A criptografia de `zip -P` (ZipCrypto/AES do `yeka/zip`) é fraca por definição do formato.
