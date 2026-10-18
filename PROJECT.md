# ZipRevival by Jair Lima

## Descrição
Suíte de utilitários de linha de comando em Go que reproduzem programas clássicos de compactação e extração: `zip`, `unzip`, `pkzip`, `pkunzip`, `unrar`, `untar`, `un7z`, `unlha`, `uniso` e `uncab`.

## Stack
Go 1.26+. Dependências: `yeka/zip` (ZIP com senha), `bodgit/sevenzip`, `koron-go/lha`, `kdomanski/iso9660`, `bogdannbv/go-cabfile`, `nwaples/rardecode/v2`, `ulikunitz/xz`.

## Estrutura
- `cmd/<comando>/main.go`: um executável por comando.
- `internal/archiver/`: `pathsafe.go` (SafeJoin, contra Zip Slip), `extract_tar.go`, `extract_rar.go`, `zip.go`, `unzip.go` e os testes.
- `.oss-scanner/`: `Dockerfile` e `threat_model.md` para o OSS Scanner da Anthropic.

## Comandos
```
go build -o bin/ ./cmd/...
go test ./...
go vet ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

## Decisões
- Todo extrator passa o nome da entrada por `archiver.SafeJoin`; entradas inseguras são ignoradas e o comando sai com código 2.
- `untar` e `unrar` nunca criam links. O RAR tem limite de dicionário (512 MiB) e de tamanho por arquivo (8 GiB).
- `mholt/archiver/v3` foi removido (path traversal sem correção, GO-2024-2698 e GO-2025-3605).
- Os binários não ficam no repositório; compile com o comando acima.

## Estado atual
Correções de segurança aplicadas em 2026-10-18 (ver README, seção Segurança). `govulncheck` sem vulnerabilidades de módulos. Inscrição no OSS Scanner preparada, ainda não enviada.

## Próximos passos
Abrir o PR de inscrição em `anthropics/oss-scanner` (precisa assinar o CLA); limitar o tamanho total extraído de ZIP, 7z, LHA, CAB e ISO; cobrir `un7z`, `unlha`, `uncab` e `uniso` com testes de arquivos maliciosos.

## Problemas conhecidos
`uniso` usa `util.ExtractImageToDirectory` da biblioteca `kdomanski/iso9660` sem checagem própria de caminho. `unzip` e `zip` não limitam o tamanho total extraído.
