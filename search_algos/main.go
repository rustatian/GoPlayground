package main

import (
	"crypto/rand"
	"sort"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const searchStrin = "fsadnfjaslfjnawejf"

var targetSlice []string

func init() {
	targetSlice = make([]string, 0, 100000001)
	for i := 0; i < 10; i++ {
		targetSlice = append(targetSlice, string(RandASCIIBytes(18)))
	}

	targetSlice = append(targetSlice, searchStrin)

}

func RandASCIIBytes(n int) []byte {
	output := make([]byte, n)     // We will take n bytes, one byte for each character of output.
	randomness := make([]byte, n) // read all random
	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}
	l := len(letterBytes)
	// fill output
	for pos := range output {
		// get random item
		random := uint8(randomness[pos]) // random % 64
		randomPos := random % uint8(l)   // put into output
		output[pos] = letterBytes[randomPos]
	}
	return output
}

func main() {


}

//go:noinline
func naive_search(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//go:noinline
func sorted_search(s []string, e string) bool {
	i := sort.SearchStrings(s, e)
	if i <= len(s) {
		return true
	}
	return false
}
