package main

import (
	"time"

	"github.com/ValeryPiashchynski/GoPlayground/fluuuut/items"
	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(100)

	b := MakeItem(builder, []byte("key"), []byte("value"), []byte(time.Now().Format(time.RFC3339)))

	key, value, ttl := ReadData(b)

	println(key)
	println(value)
	println(ttl)
}

func MakeItem(b *flatbuffers.Builder, key []byte, value []byte, ttl []byte) []byte {
	b.Reset()

	keyOffset := b.CreateByteString(key)
	valueOffset := b.CreateByteString(value)
	ttlOffset := b.CreateByteString(ttl)

	items.ItemStart(b)
	items.ItemAddKey(b, keyOffset)
	items.ItemAddValue(b, valueOffset)
	items.ItemAddTTL(b, ttlOffset)

	itemOffset := items.ItemEnd(b)

	b.Finish(itemOffset)

	return b.Bytes[b.Head():]
}

func ReadData(buf []byte) (key, value, ttl string) {
	itemRoot := items.GetRootAsItem(buf,0)


	key = string(itemRoot.Key())
	value = string(itemRoot.Value())
	ttl = string(itemRoot.TTL())
	return
}