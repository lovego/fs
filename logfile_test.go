package fs

import "fmt"

func ExampleNewLogFile() {
	f, err := NewLogFile("./a.txt")
	fmt.Println(f.path, f.File != nil, err)

	// Output: ./a.txt true <nil>
}
