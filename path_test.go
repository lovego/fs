package fs

import (
	"fmt"
	"strings"
)

func ExampleSourceDir() {
	fmt.Println(strings.TrimPrefix(SourceDir(), GetGoSrcPath()+"/"))
	// Output: github.com/lovego/fs
}

func ExampleSourceFile() {
	fmt.Println(strings.TrimPrefix(SourceFile(), GetGoSrcPath()+"/"))
	// Output: github.com/lovego/fs/path_test.go
}
