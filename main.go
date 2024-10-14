package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	// "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func RootDir() string {
	var _, b, _, _ = runtime.Caller(0)
	return filepath.Dir(b)

}

func main() {
	println("Hello, World!")
	var root = RootDir()
	file_csapp := filepath.Join(root, "examples", "csapp.pdf")
	println(file_csapp)

	dummyDir := filepath.Join(root, "dummy")
	println(dummyDir)

	outDir := filepath.Join(dummyDir)

	selectedPages := []string{"1"}
	if err := api.ExtractFontsFile(
		file_csapp,
		outDir,
		selectedPages,
		nil); err != nil {
		log.Fatalf("Extract Image Error %s: %v\n", file_csapp, err)
	}
}
