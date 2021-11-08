package fs

import (
	"fmt"
	"path/filepath"
)

func ExampleFilepathDir_linux() {
	fmt.Println(filepath.Dir(`/`))
	// Output:
	// /
}
