package fs

import (
	"fmt"
	"os"
	"strings"
)

func ExampleSourceDir() {
	printPath(SourceDir())
	// Output: github.com/lovego/fs
}

func ExampleSourceFile() {
	printPath(SourceFile())
	// Output: github.com/lovego/fs/path_test.go
}

func printPath(p string) {
	fmt.Println(p)
	sep := string(os.PathSeparator)
	p = strings.TrimPrefix(p, GetGoSrcPath()+sep)
	if os.PathSeparator != '/' {
		p = strings.Replace(p, sep, "/", -1)
	}
	fmt.Println(p)
}
