package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
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

	output480p := strings.TrimSuffix(*filePath, ext) + "_480p" + ext
	output720p := strings.TrimSuffix(*filePath, ext) + "_720p" + ext
	outputAudio := strings.TrimSuffix(*filePath, ext) + "_audio.mp3"

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		runFFmpegCommand(*filePath, "-vf", "scale=-2:480", "-c:a", "copy", output480p)
		fmt.Println("Compressed to 480p saved to:", output480p)
	}()

	go func() {
		defer wg.Done()
		runFFmpegCommand(*filePath, "-vf", "scale=-2:720", "-c:a", "copy", output720p)
		fmt.Println("Compressed to 720p saved to:", output720p)

	}()

	go func() {
		defer wg.Done()
		runFFmpegCommand(*filePath, "-vn", "-c:a", "libmp3lame", "-q:a", "4", outputAudio)
		fmt.Println("Audio extracted and saved to:", outputAudio)
	}()

}

func runFFmpegCommand(inputFile string, args ...string) {
	cmdArgs := append([]string{"-i", inputFile}, args...)
	cmd := exec.Command("ffmpeg", cmdArgs...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed: %v\nOutput: %s\nError: %s", err, out.String(), stderr.String())
	}
}
