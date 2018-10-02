package main

import (
	"bytes"
	cr "crypto/rand"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/goware/urlx"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	Base              = 10
	rateLimitInterval = 10 * time.Millisecond

	exitFailure = 1

	//RANDOM
	letters       = "qwertyuioplkjhgfdsazxcvbnmPOIUYTREWQASDFGHJKLMNBVCXZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	//RANDOM

	nilStr = "nilStr"

	defaultTestDuration  = 10 * time.Second
	defaultNumberOfConns = uint64(125)
	defaultTimeout       = 2 * time.Second

	panicZeroRate         = "rate can't be zero"
	panicNegativeAdjustTo = "adjustTo can't be negative or zero"
)

var src = rand.NewSource(time.Now().UnixNano())

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
	Gl := NewConcurrentLevel(1)
	for i := 0; i < 1; i-- {
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

	var fs = afero.NewMemMapFs()
	afs := &afero.Afero{Fs: fs}

	//f, err := os.Open("/Users/0xdev/Projects/HomeProjects/GoPlayground/testing_scripts/file.txt")
	f, err := afs.Create("/tmp/testFile.txt")
	if err != nil {
		panic(err)
	}

	boundary := randomBoundary()

	client := &http.Client{}
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	writer.SetBoundary(boundary)
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

	b := randStringBytes(1024 * 1024)
	f.Write(b)

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	req := &http.Request{
		Method: http.MethodPost,
		Body:   ioutil.NopCloser(&buf),
		GetBody: func() (io.ReadCloser, error) {
			r := bytes.NewReader(buf.Bytes())
			return ioutil.NopCloser(r), nil
		},
	}

	req.Header = http.Header{}
	req.URL, err = urlx.Parse("http://192.168.101.25:5685/v1/documents?company_uuid=33333333-3333-3333-3333-333333333333&uid_uuid=11111111-1111-1111-1111-111111111111")

	//req, err := http.NewRequest(http.MethodPost, "http://192.168.101.25:5685/v1/documents?company_uuid=33333333-3333-3333-3333-333333333333&uid_uuid=11111111-1111-1111-1111-111111111111", &buf)
	//req.Header.Set("Content-Type", "multipart/form-data; boundary=d54d4ce7498780c2cf76f30dcdf7881c41ae7a0435eaf20ffd1a8324976b")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWRfdXVpZCI6IjExMTExMTExLTExMTEtMTExMS0xMTExLTExMTExMTExMTExMSIsImNvbXBhbnlfdXVpZCI6IjMzMzMzMzMzLTMzMzMtMzMzMy0zMzMzLTMzMzMzMzMzMzMzMyJ9.4ZFpGRWFfUGDDr8rKk_T30bC5W248Av7D3S8JB1Zyvg")
	req.Close = true

	//ssss := writer.Boundary()
	//print(ssss)
	//hdr := writer.FormDataContentType()

	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
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

// generates random string with specified len
func randStringBytes(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(cr.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}
