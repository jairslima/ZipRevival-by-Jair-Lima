package archiver

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/yeka/zip"
)

// ZipOptions define as opções de compactação similares às flags originais.
type ZipOptions struct {
	Recursive        bool
	Move             bool
	CompressionLevel int // 0 para Store, >0 para Deflate
	Password         string
}

// ZipFiles compacta uma lista de arquivos ou diretórios em um arquivo de destino.
func ZipFiles(dest string, files []string, opts ZipOptions) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo zip %s: %w", dest, err)
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	var filesToArchive []string
	var dirsToArchive []string

	// Levantar todos os arquivos baseando-se na lista de entrada e na flag Recursive (-r)
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Printf("zip warning: name not matched: %s\n", file)
			continue
		}

		if info.IsDir() {
			if opts.Recursive {
				err = filepath.Walk(file, func(path string, walkInfo os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if walkInfo.IsDir() {
						dirsToArchive = append(dirsToArchive, path)
					} else {
						filesToArchive = append(filesToArchive, path)
					}
					return nil
				})
				if err != nil {
					return err
				}
			} else {
				// Comportamento do zip clássico: adiciona o diretório vazio, mas não o seu conteúdo
				dirsToArchive = append(dirsToArchive, file)
			}
		} else {
			filesToArchive = append(filesToArchive, file)
		}
	}

	// Adicionar diretórios (apenas entradas vazias para manter a estrutura)
	for _, path := range dirsToArchive {
		err := addDirToZip(archive, path)
		if err != nil {
			return err
		}
	}

	// Adicionar arquivos e comprimir
	for _, path := range filesToArchive {
		err := addFileToZip(archive, path, opts)
		if err != nil {
			return err
		}
		fmt.Printf("  adding: %s\n", filepath.ToSlash(path))
	}

	// Operação -m (Move): apaga os originais se não houve erros
	if opts.Move {
		for _, path := range filesToArchive {
			os.Remove(path)
		}
		// Apagar diretórios em ordem reversa para garantir que estão vazios
		for i := len(dirsToArchive) - 1; i >= 0; i-- {
			os.Remove(dirsToArchive[i])
		}
	}

	return nil
}

func addDirToZip(zipWriter *zip.Writer, dirPath string) error {
	info, err := os.Stat(dirPath)
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Para indicar um diretório no ZIP, o nome deve terminar com barra
	header.Name = filepath.ToSlash(dirPath) + "/"
	header.Method = zip.Store

	_, err = zipWriter.CreateHeader(header)
	return err
}

func addFileToZip(zipWriter *zip.Writer, filePath string, opts ZipOptions) error {
	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filepath.ToSlash(filePath)
	
	if opts.CompressionLevel == 0 {
		header.Method = zip.Store
	} else {
		header.Method = zip.Deflate
	}

	if opts.Password != "" {
		header.SetPassword(opts.Password)
		header.SetEncryptionMethod(zip.StandardEncryption)
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)
	return err
}
