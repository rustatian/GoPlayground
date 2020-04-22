package main

import (
	"time"

	"github.com/ValeryPiashchynski/GoPlayground/fluuuut/items"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(100)

	b := MakeItem(builder, []byte("key"), []byte("value"), []byte(time.Now().Format(time.RFC3339)))

	key, value, ttl, d := ReadData(b)

	println(key)
	println(value)
	println(ttl)
	println(d)
}

func MakeItem(b *flatbuffers.Builder, key []byte, value []byte, ttl []byte) []byte {
	b.Reset()

	m := make([]string, 0)
	m = append(m, "asasdfasdfd")
	m = append(m, "dasdfasdfasdfasdaaaaaaaaaaaaad")
	m = append(m, "dasdd")

	offsets := make([]flatbuffers.UOffsetT, len(m))

	for i := len(m) - 1; i >= 0; i-- {
		offsets[i] = b.CreateString(m[i])
	}

	items.ItemStartDataVector(b, len(offsets))

	for i := len(offsets) - 1; i >= 0; i-- {
		b.PrependUOffsetT(offsets[i])
	}

	x := b.EndVector(len(offsets))

	keyOffset := b.CreateByteString(key)
	valueOffset := b.CreateByteString(value)
	ttlOffset := b.CreateByteString(ttl)

	items.ItemStart(b)
	items.ItemAddData(b, x)
	items.ItemAddKey(b, keyOffset)
	items.ItemAddValue(b, valueOffset)
	items.ItemAddTTL(b, ttlOffset)

	itemOffset := items.ItemEnd(b)

	b.Finish(itemOffset)

	return b.Bytes[b.Head():]
}

func ReadData(buf []byte) (key, value, ttl, data string) {
	itemRoot := items.GetRootAsItem(buf, 0)

	key = string(itemRoot.Key())
	value = string(itemRoot.Value())
	ttl = string(itemRoot.TTL())


	l := itemRoot.DataLength()
	for i := 0; i < l; i++ {
		println(string(itemRoot.Data(i)))
	}


	return
}
