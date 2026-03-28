package main
import (
	"fmt"
	"os"
	"github.com/yeka/zip"
)
func main() {
	f, _ := os.Create("test.zip")
	w := zip.NewWriter(f)
	info, _ := f.Stat()
	header, _ := zip.FileInfoHeader(info)
	header.SetPassword("pass") // Does SetPassword exist?
	w.CreateHeader(header)
	w.Close()
	fmt.Println("OK")
}
