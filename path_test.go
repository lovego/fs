package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExampleSourceDir() {
	p := SourceDir()
	fmt.Println(filepath.ToSlash(strings.TrimPrefix(p, GetGoSrcPath()+string(os.PathSeparator))))
	// Output: github.com/lovego/fs
}

func ExampleSourceFile() {
	p := SourceFile()
	fmt.Println(filepath.ToSlash(strings.TrimPrefix(p, GetGoSrcPath()+string(os.PathSeparator))))
	// Output: github.com/lovego/fs/path_test.go
}
