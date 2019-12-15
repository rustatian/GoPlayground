package main

import "unsafe"

const maxVarintBytes = 10
//923472
func main() {
	b := EncodeVarintC(128)
	println(unsafe.Sizeof(b))
	d := DecodeVarint(b)
	println(d)
}

// 11100001011101010000
//
func EncodeVarint(x uint64) []byte {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		tmp3 := x & 0x7F
		tmp := uint8(tmp3)
		tmp2 := 0x80 | tmp
		println(unsafe.Sizeof(tmp2))
		buf[n] = tmp2
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

func DecodeVarint(buf []byte) (x uint64) {
	i := 0

	if buf[i] < 0x80 {
		return uint64(buf[i])
	}
	var b uint64
	// we already checked the first byte
	x = uint64(buf[i]) - 0x80
	i++

	b = uint64(buf[i])
	i++
	x = x + b<<7
	if b&0x80 == 0 {
		goto done
	}
	x = x - 0x80<<7

	b = uint64(buf[i])
	i++
	x += b << 14
	if b&0x80 == 0 {
		goto done
	}
	x = x - 0x80<<14

	b = uint64(buf[i])
	i++
	x = x + b<<21
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 21

	b = uint64(buf[i])
	i++
	x += b << 28
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 28

	b = uint64(buf[i])
	i++
	x += b << 35
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 35

	b = uint64(buf[i])
	i++
	x += b << 42
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 42

	b = uint64(buf[i])
	i++
	x += b << 49
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 49

	b = uint64(buf[i])
	i++
	x += b << 56
	if b&0x80 == 0 {
		goto done
	}
	x -= 0x80 << 56

	b = uint64(buf[i])
	i++
	x += b << 63
	if b&0x80 == 0 {
		goto done
	}

	return 0

done:
	return x
}

func EncodeVarintC(x uint64) []byte {
	var buf [maxVarintBytes]byte
	i := 0
	if x&0x80 == 0 {
		buf[i] = uint8(x)
		return buf[0:i+1]
	}

	for x > 127 {
		tmp := uint8(x&0x7F) // get first 7 bytes
		tmp2 := tmp | 0x80
		buf[i] = tmp2
		x = x >> 7
		i++
	}
	buf[i] = uint8(x)
	return buf[0:i+1]

}
