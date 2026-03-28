package archiver

import (
	"os"
	"path/filepath"
	"testing"
)

func TestZipUnzip(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create a test file
	testFileName := "test_hello.txt"
	testFilePath := filepath.Join(tempDir, testFileName)
	err := os.WriteFile(testFilePath, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	zipDest := filepath.Join(tempDir, "test.zip")

	// Test Zip with Password
	opts := ZipOptions{
		Recursive:        false,
		Move:             false,
		CompressionLevel: 8,
		Password:         "secret",
	}
	
	// Temporarily change working dir to avoid absolute paths in zip
	originalWD, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalWD)

	err = ZipFiles("test.zip", []string{testFileName}, opts)
	if err != nil {
		t.Fatalf("ZipFiles failed: %v", err)
	}

	// Test Unzip
	extractDir := "extract"
	unzipOpts := UnzipOptions{
		DestDir:   extractDir,
		Overwrite: 1,
		Password:  "secret",
	}
	
	err = UnzipFile("test.zip", &unzipOpts)
	if err != nil {
		t.Fatalf("UnzipFile failed: %v", err)
	}

	// Verify extraction
	extractedContent, err := os.ReadFile(filepath.Join(extractDir, testFileName))
	if err != nil {
		t.Fatalf("Failed to read extracted file: %v", err)
	}

	if string(extractedContent) != "hello world" {
		t.Errorf("Extracted content mismatch. Expected 'hello world', got '%s'", string(extractedContent))
	}
}
