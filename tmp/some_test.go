package main

import (
	"testing"

	"github.com/rustatian/GoPlayground/tmp/test_package"
)

func TestOne(t *testing.T) {
	p := new(test_package.Plugin)
	p.Init()
}

func TestTwo(t *testing.T) {
	p := new(test_package.Plugin)
	p.Init()
}
