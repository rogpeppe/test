//go:build never

package main

import (
	"fmt"
	"io/fs"

	"github.com/rogpeppe/test"
)

func main() {
	fs.WalkDir(test.FS(), ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
}
