package main

import (
	"github.com/spf13/afero"
	"math/rand"
	"strconv"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main() {
	var AppFs = afero.NewMemMapFs()
	afs := &afero.Afero{Fs: AppFs}

	//b := make([]byte, 1000*24)
	//for i := 0; i < 1000*24; i ++ {
	//	b[i] = byte(rand.Int())
	//}

	b := RandStringBytesMaskImprSrc(1000 * 24)

	m := map[int]afero.File{}

	for i := 0; i < 10000000; i++ {
		f, err := afs.Create("/tmp/file" + strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
		f.Write(b)
		m[i] = f
	}

	//for k, v := range m {
	//	fmt.Println(k)
	//	fmt.Println(v.Name())
	//}
}

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(src.Int63())
	}
	return b
}
