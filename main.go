package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
)

func main() {
	folder_path := flag.String("p", "./", "folder path")
	flag.Parse()

	extensions := []string{".png", ".jpeg", ".mov", ".webp", ".gif", ".mp4", ".webm", ".jpg"}
	files, _ := os.ReadDir(*folder_path)
	for _, file := range files {
		if !file.IsDir() {
			file_path := path.Join(*folder_path, file.Name())
			file_extension := filepath.Ext(file_path)
			if slices.Contains(extensions, file_extension) {
				file_bytes, _ := os.ReadFile(file_path)
				hash_bytes := md5.Sum(file_bytes)
				hash := hex.EncodeToString(hash_bytes[:16])
				fmt.Println(file_path, path.Join(*folder_path, fmt.Sprintf("%s%s", hash, file_extension)))
				os.Rename(file_path, path.Join(*folder_path, fmt.Sprintf("%s%s", hash, file_extension)))
			}
		}
	}
}
