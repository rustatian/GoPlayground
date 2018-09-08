package main

import (
	"bytes"
	"fmt"
	tm "github.com/buger/goterm"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

type GoroutineLimiter struct {
	sync.WaitGroup
	currentRunning      int64
	concurrentLevel     int
	routinesDistributor chan int
}

func NewConcurrentLevel(numGoroutines int) *GoroutineLimiter {
	gl := &GoroutineLimiter{
		concurrentLevel:     numGoroutines,
		routinesDistributor: make(chan int, numGoroutines),
	}
	//gl.Add(numGoroutines)

	// add maximum values
	for i := 0; i < gl.concurrentLevel; i++ {
		gl.routinesDistributor <- i
	}

	return gl
}

func (gl *GoroutineLimiter) Start(fn func()) {
	//gl.Done()
	getAndFree := <-gl.routinesDistributor
	atomic.AddInt64(&gl.currentRunning, 1)
	go func() {
		defer func() {
			//gl.Done()
			gl.routinesDistributor <- getAndFree
			atomic.AddInt64(&gl.currentRunning, -1)
		}()

		fn()
		//gl.Add(1)
	}()
}

func (gl *GoroutineLimiter) GetNumberOfRunningGoroutines() int64 {
	return gl.currentRunning
}

func main() {
	Gl := NewConcurrentLevel(2000)
	for i := 0; i < 10000; i-- {
		Gl.Start(func() {
			doPostRequest("file.txt")
		})

		tm.MoveCursor(1, 1)
		tm.Printf("Number of running Goroutines: %d", Gl.GetNumberOfRunningGoroutines())
		tm.Flush()
	}
}

func doPostRequest(fn string) {
	i := rand.Intn(10000000000)
	params := make(map[string]string)
	params["filename"] = "file_" + strconv.Itoa(i) + ".txt"

	f, err := os.Open("/Users/0xdev/Projects/HomeProjects/GoPlayground/testing_scripts/file.txt")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", fn)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, f)
	if err != nil {
		panic(err)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://192.168.101.96:5685/v1/documents", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwOi8vYXV0aC5pbnR1cm4uZG9ja2VyOjgwODIvYXV0aC9hY2Nlc3NfdG9rZW4iLCJpYXQiOjE1MzQ5NjM5ODEsImV4cCI6MTUzNzU5MTk4MSwibmJmIjoxNTM0OTYzOTgxLCJqdGkiOiJVR09TemptUkppcjRiTXlyIiwic3ViIjoxNTQ3LCJ1aWRfdXVpZCI6IjExMTExMTExLTExMTEtMTExMS0xMTExLTExMTExMTExMTExMSIsImNvbXBhbnlfdXVpZCI6IjMzMzMzMzMzLTMzMzMtMzMzMy0zMzMzLTMzMzMzMzMzMzMzMyJ9.-91UgS0AWUxWvpn115Tb5NAfDEc9EWaG4nGPgursVgM")
	req.Close = true

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println(err)
	} else {
		defer req.Body.Close()
		defer resp.Body.Close()
		resp.Close = true

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		//
		fmt.Printf("%s\n", string(body))
	}
}
