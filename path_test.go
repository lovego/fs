package fs

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ExampleSourceDir() {
	p := SourceDir()
	fmt.Println(p)
	fmt.Println(strings.HasSuffix(filepath.ToSlash(p), `github.com/lovego/fs`))
	// Output: true
}

func ExampleSourceFile() {
	p := SourceFile()
	fmt.Println(p)
	fmt.Println(strings.HasSuffix(filepath.ToSlash(p), `github.com/lovego/fs/path_test.go`))
	// Output: true
}

func ExampleFilepathDir() {
	fmt.Println(filepath.Dir(`/`))
	// Output:
	// /
}
