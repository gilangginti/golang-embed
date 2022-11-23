package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed Example.txt
var version string

func TestVersion(t *testing.T) {
	fmt.Println(version)
}

//go:embed Logo.png
var logo []byte

func TestByteSlice(t *testing.T) {
	err := os.WriteFile("logo_next2.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed Example.txt
//go:embed Example2.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("Example.txt")
	b, _ := files.ReadFile("Example2.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content", string(content))
		}
	}
}
