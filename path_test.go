package fs

import (
	"fmt"
	"path/filepath"
	"strings"
)

func ExampleSourceDir() {
	p := SourceDir()
	fmt.Println(strings.HasSuffix(filepath.ToSlash(p), `/fs`))
	// Output: true
}

func ExampleSourceFile() {
	p := SourceFile()
	fmt.Println(strings.HasSuffix(filepath.ToSlash(p), `/fs/path_test.go`))
	// Output: true
}
