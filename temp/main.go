package main

func main() {
	u := []uint8{123, 34, 94, 94, 92, 117, 48, 48, 50, 54, 48, 50, 94, 49, 94, 50, 94, 50, 53, 92, 117, 48, 48, 50, 54, 48, 51, 94, 49, 94, 50, 94, 51, 92, 117, 48, 48, 50, 54, 48, 51, 94, 49, 94, 49, 94, 49, 48, 48, 48, 48, 34, 58, 49, 125}

	println(B2S(u))
}

func B2S(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}