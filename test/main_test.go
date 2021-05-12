package main

import (
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	tt := time.Now()
	test_btb(true)
	println(time.Now().Sub(tt).Nanoseconds())

	tt2 := time.Now()
	test_btb(true)
	println(time.Now().Sub(tt2).Nanoseconds())

	tt3 := time.Now()
	test_btb(true)
	println(time.Now().Sub(tt3).Nanoseconds())
}
