# Resumo Técnico: Zip Revival by Jair Lima

**Objetivo:** Este projeto é um clone dos utilitários clássicos de terminal `zip` e `unzip`, projetado para alta performance e compatibilidade de comandos (flags), ideal para uso no Windows e outros sistemas sem a necessidade de wrappers ou dependência de ferramentas com interfaces distintas (como 7-Zip).

**Stack Tecnológico:**
- Linguagem: Go (Golang)
- Arquitetura: Binários separados (`zip.exe`, `unzip.exe`, `unrar.exe`, `pkzip.exe`, `pkunzip.exe`, `untar.exe`, `unlha.exe`, `uniso.exe`, `uncab.exe`, `un7z.exe`) para fácil substituição via `PATH`.
- Bibliotecas: `github.com/yeka/zip` para suporte a senhas/criptografia; `github.com/mholt/archiver/v3` e `nwaples/rardecode` para RAR/TAR; `github.com/koron-go/lha` para LHA; `github.com/kdomanski/iso9660` para ISO; `github.com/bogdannbv/go-cabfile` para CAB e `github.com/bodgit/sevenzip` para 7Z. O suporte ao ARJ não foi implementado devido à falta de bibliotecas modernas.

**Estrutura de Diretórios:**
- `/cmd/zip`: Ponto de entrada para o binário do `zip`. Responsável pelo parsing da linha de comando.
- `/cmd/unzip`: Ponto de entrada para o binário do `unzip`. Responsável pelo parsing da linha de comando.
- `/internal/archiver`: Lógica "core" isolada para compactação (`zip.go`) e extração (`unzip.go`).

**Instruções de Manutenção e Continuidade:**
- **Compatibilidade:** Qualquer nova funcionalidade deve observar a sintaxe do utilitário `zip` e `unzip` original dos ambientes Unix/DOS para manter a "memória muscular" do usuário no terminal.
- **Permissões de GitHub:** Se versionado remotamente, os repositórios criados a partir deste código fonte devem sempre ser configurados como **Privados** (regra mandatória do usuário).
- **Testes:** Novas flags ou comandos devem possuir cobertura de testes (`go test ./...`) garantindo paridade com as ferramentas clássicas.