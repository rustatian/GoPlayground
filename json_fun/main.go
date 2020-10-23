package main

import (
	"encoding/json"
	"unicode/utf8"
)

func main() {
	rs := "GreyFlannelHeather Mol&;"

	var i string
	data, err := json.Marshal(&rs)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &i)
	if err != nil {
		panic(err)
	}
	_ = data
	_ =i

	g := string(data)
	_ = g

	bs := make([]byte, len(rs)*utf8.UTFMax)
	n := 0
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	bs = bs[:n]
}

