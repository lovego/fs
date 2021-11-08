package fs

import (
	"fmt"
	"path/filepath"
)

func ExampleFilepathDir_windows() {
	fmt.Println(filepath.Dir(`C:\`))
	// Output:
	// C:\
}
