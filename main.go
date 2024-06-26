package test

import (
	"embed"
	"io/fs"
)

//go:embed somedir
var f embed.FS

func FS() fs.FS {
	return f
}
