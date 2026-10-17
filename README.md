# Zip Revival by Jair Lima

Uma suíte completa de utilitários de linha de comando desenvolvida em Go, focada na clonagem e aprimoramento de ferramentas clássicas de arquivamento (como zip, unzip, ar, 	ar, lha, etc.), garantindo extrema velocidade, compatibilidade de sintaxe para preservar a memória muscular, e independência (binários nativos autônomos sem necessidade de instaladores complexos).

## Utilitários Incluídos

*   **zip.exe e unzip.exe**: Operações completas de ZIP. Suporta compressão máxima (-9), recursividade (-r), movimentação (-m) e **criptografia por senha (-P)**.
*   **pkzip.exe e pkunzip.exe**: Wrappers nostálgicos com a sintaxe idêntica à do DOS.
*   **unrar.exe**: Extração de arquivos proprietários .rar.
*   **untar.exe**: Extração de arquivos de fita Unix (.tar, .tar.gz, .tar.bz2, .xz).
*   **un7z.exe**: Extração autônoma do moderno e eficiente formato .7z.
*   **unlha.exe**: Extração de arquivos clássicos japoneses .lzh e .lha (método LH5).
*   **uniso.exe**: Extração direta do conteúdo de imagens de CD/DVD .iso (ISO 9660).
*   **uncab.exe**: Extração de contêineres de drivers e instaladores Windows .cab (Cabinet).

## Compilar

Requer [Go](https://go.dev/dl/) instalado.

```bash
go build -o zip.exe ./cmd/zip
go build -o unzip.exe ./cmd/unzip
# o mesmo para: pkzip, pkunzip, unrar, untar, un7z, unlha, uniso, uncab
```

Para compilar tudo de uma vez no PowerShell:

```powershell
Get-ChildItem cmd -Directory | ForEach-Object { go build -o "$($_.Name).exe" "./cmd/$($_.Name)" }
```

Copie os `.exe` gerados para uma pasta do `PATH`. Os binários não ficam no repositório.

## Bibliotecas

- ZIP com senha: `github.com/yeka/zip`
- TAR, RAR e 7z: `github.com/mholt/archiver/v3` e `github.com/bodgit/sevenzip`
- LHA/LZH: `github.com/koron-go/lha`
- ISO 9660: `github.com/kdomanski/iso9660`
- CAB: `github.com/bogdannbv/go-cabfile`

O formato ARJ ficou de fora por falta de bibliotecas estáveis em Go.

## Licença

MIT, veja `LICENSE`. Cada biblioteca acima mantém a sua própria licença.
