package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "The path to the file")

	flag.Parse()

	if *filePath == "" {
		log.Fatal("You must specify a file path using the -file flag.")
	}

	ext := strings.ToLower(filepath.Ext(*filePath))

	videoExtensions := []string{".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm", ".mpeg", ".mpg"}

	isVideo := false

	for _, v := range videoExtensions {
		if ext == v {
			isVideo = true
			break
		}
	}

	if !isVideo {
		log.Fatal("Please provide a video file.")
	}

	fmt.Println("File Path:", *filePath)

}
