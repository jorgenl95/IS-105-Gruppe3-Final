package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename := os.Args[1]
	file, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}
	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileSize := file.Size()
	fileSizeB := fileSize
	fileSizeKB := fileSizeB / 1024
	fileSizeMB := float32(fileSizeKB / 1024)
	fileSizeGB := float32(fileSizeMB / 1024)

	fileMode := file.Mode()
	isDir := fileMode.IsDir()
	isReg := fileMode.IsRegular()
	permissionBits := fileMode.Perm()
	isAppendOnly := fileMode&os.ModeAppend != 0
	isDevice := false
	isUnixChar := false
	isUnixBloc := false
	if fileMode&os.ModeDevice != 0 {
		isDevice = true
		if fileMode&os.ModeCharDevice != 0 {
			isUnixChar = true
		} else {
			isUnixBloc = true
		}
	}
	isSymbolicLink := fileMode&os.ModeSymlink != 0

	fmt.Printf("Information about %s:\n", absPath)
	fmt.Printf("Size: %d bytes, %d KB, %.2f MB, %f GB\n", fileSizeB, fileSizeKB, fileSizeMB, fileSizeGB)
	fmt.Println("IS Dir ", isDir)
	fmt.Println("Is Reg ", isReg)
	fmt.Println("Unix permission bits: ", permissionBits)
	fmt.Println("Is append only: ", isAppendOnly)
	fmt.Println("Is device file: ", isDevice)
	fmt.Println("Is Unix character device: ", isUnixChar)
	fmt.Println("Is Unix block device", isUnixBloc)
	fmt.Println("Is symbolic link", isSymbolicLink)
}