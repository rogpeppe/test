package test

import (
	"embed"
	"io/fs"
)

//go:embed all:somedir
var f embed.FS

func FS() fs.FS {
	return f
}
