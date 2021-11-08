package fs

import (
	"fmt"
	"path/filepath"
)

func ExampleFilepathDir() {
	fmt.Println(filepath.Dir(`C:\`))
	// Output:
	// C:\\
}
