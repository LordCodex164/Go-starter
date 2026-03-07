package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("files", "example.txt")

	
	fmt.Println(filepath.Base(filePath))
	fmt.Println(filepath.Ext(filePath))

	dirtyDir := "./users/./files/..//.."
	cleanDir := filepath.Clean(dirtyDir)
	fmt.Println(cleanDir)
}