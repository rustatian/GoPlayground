package main

type Adder interface{ Add(a, b int32) int32 }
type Subber interface{ Sub(a, b int32) int32 }
type Mather interface {
	Adder
	Subber
}

func main() {
	calc := Calculator{id: 6754}
	var m Mather = &calc
	m.Sub(10, 32)
}

type Calculator struct{ id int32 }

func (c *Calculator) Add(a, b int32) int32 { return a + b }
func (c *Calculator) Sub(a, b int32) int32 { return a - b }

//go.itab.*"".Calculator,"".Mather SRODATA dupok size=40
//0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
//0x0010 5e 33 ca c8 00 00 00 00 00 00 00 00 00 00 00 00  ^3..............
//0x0020 00 00 00 00 00 00 00 00                          ........
//rel 0+8 t=1 type."".Mather+0
//rel 8+8 t=1 type.*"".Calculator+0
//rel 24+8 t=1 "".(*Calculator).Add+0
//rel 32+8 t=1 "".(*Calculator).Sub+0
