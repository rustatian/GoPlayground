package main

import "encoding/json"

type B struct {
	FOO string `json:"foo"`
	FAA string `json:"faa"`
}

type A struct {
	FEE string `json:"fee"`
	B   `json:"b"`
}

func main() {
	str := `{
  "fee":"ffff",
  "b":{
    "foo": "asdfasdf",
    "faa": "fasdfasldfasdf"
  }
}`

	var a A

	json.Unmarshal([]byte(str), &a)

	println("fsadfas")
}
