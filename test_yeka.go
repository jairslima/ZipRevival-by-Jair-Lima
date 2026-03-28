package main
import (
	"fmt"
	"os"
	"github.com/yeka/zip"
)
func main() {
	f, _ := os.Create("test.zip")
	w := zip.NewWriter(f)
	w.Encrypt("test.txt", "pass", zip.AES256Encryption)
	w.Close()
	fmt.Println("OK")
}
