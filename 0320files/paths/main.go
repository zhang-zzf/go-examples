package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

var p = fmt.Println

func main() {
	aPath := filepath.Join("/a/b/", "c.txt")
	p(aPath)
	p(filepath.Dir(aPath))
	p(filepath.Base(aPath))
	p(filepath.IsAbs(aPath))
	ext := filepath.Ext(aPath)
	p(ext)
	p(strings.TrimSuffix(filepath.Base(aPath), ext))

}
