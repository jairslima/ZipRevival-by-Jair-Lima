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

## Instalação

Copie os executáveis (.exe) gerados para qualquer diretório listado na variável de ambiente PATH do seu sistema operacional. No desenvolvimento original, eles foram alocados em $env:USERPROFILE\ZipRevivalBin.

## Notas de Desenvolvimento (Para IAs Futuras)
Este repositório foi construído seguindo diretrizes estritas do arquivo GEMINI.md.

*   **Bibliotecas Recomendadas**:
    *   ZIP (c/ Senhas): github.com/yeka/zip
    *   TAR/RAR/7Z (Multi-formato): github.com/mholt/archiver/v3
    *   LHA/LZH: github.com/koron-go/lha
    *   ISO: github.com/kdomanski/iso9660
    *   CAB: github.com/bogdannbv/go-cabfile
*   **Limitações**: O formato ARJ foi intencionalmente excluído do projeto devido à ausência de bibliotecas estáveis e ativas em Go moderno, garantindo a robustez do restante da suíte.
*   **Privacidade**: O versionamento foi feito num repositório estritamente Privado, conforme preferência global do usuário.

Desenvolvido para **Jair Lima**.
