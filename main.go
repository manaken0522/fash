package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strings"
)

func main() {
	extension := []string{"png", "jpeg", "mov", "webp", "gif", "mp4", "webm", "jpg"}
	files, _ := ioutil.ReadDir("./")
	for _, file := range files {
		if !file.IsDir() {
			if slices.Contains(extension, strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]) {
				file_bytes, _ := ioutil.ReadFile(file.Name())
				hash_bytes := md5.Sum(file_bytes)
				hash := hex.EncodeToString(hash_bytes[:16])
				fmt.Println(file.Name(), fmt.Sprintf("%s.%s", hash, strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]))
				os.Rename(file.Name(), fmt.Sprintf("%s.%s", hash, strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]))
			}
		}
	}
}
