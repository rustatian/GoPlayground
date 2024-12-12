package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// Assuming you have the video data stored in a []byte called videoData
	//videoData, err := os.ReadFile("./tmp/video.mp4")
	//if err != nil {
	//	panic(err)
	//}

	// Create an FFmpeg command to extract audio
	cmd := exec.Command("ffmpeg",
		"-i", "./tmp/video.mp4", // Read input from stdin
		"-vn",          // disable video
		"-c:a", "flac", // codec to use: flac
		"-b:a", "320k", // bitrate
		"-map", "0:a", // map audio stream
		"-ar", "44100",
		"-f", "flac",
		"pipe:1", // Write output to stdout
	)

	// Set the video data as the input to FFmpeg's stdin

	// Create pipes to read the output and error streams from FFmpeg
	outputPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Stderr = os.Stderr

	// Start the FFmpeg process
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./tmp/audio.wav")
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Start a goroutine to wait for the FFmpeg process to finish
	go func() {
		defer wg.Done()
		// Start a goroutine to read the output stream
		reader := bufio.NewReader(outputPipe)
		chunkSize := 20 * 1024 // 20KB
		chunk := make([]byte, chunkSize)

		for {
			n, err := reader.Read(chunk)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}

			_, err2 := f.Write(chunk[:n])
			if err2 != nil {
				panic(err2)
			}

			fmt.Printf("Received audio chunk of size: %d bytes\n", n)
		}
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println("FFmpeg process error:", err)
	}

	wg.Wait()
}
