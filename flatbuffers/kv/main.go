package main

import (
	"time"

	"github.com/rustatian/GoPlayground/flatbuffers/kv/data"
	flatbuffers "github.com/google/flatbuffers/go"
)

// Item represents general storage item
type Item struct {
	// Key of item
	Key string
	// Value of item
	Value string
	// live until time provided by TTL in RFC 3339 format
	TTL string
}

type Data struct {
	Storage string   `json:"storage"`
	Keys    []string `json:"keys"`
	Timeout string   `json:"timeout"`
}

type SetData struct {
	Items    []Item   `json:"items"`
	Storages []string `json:"storages"`
}

func main() {
	builder := flatbuffers.NewBuilder(100)

	b := MakeItem(builder)

	ReadData(b)
}

func serializeItems(b *flatbuffers.Builder, item Item) flatbuffers.UOffsetT {
	key := b.CreateString(item.Key)
	val := b.CreateString(item.Value)
	ttl := b.CreateString(item.TTL)

	data.ItemStart(b)

	data.ItemAddKey(b, key)
	data.ItemAddValue(b, val)
	data.ItemAddTimeout(b, ttl)

	return data.ItemEnd(b)
}

func MakeItem(b *flatbuffers.Builder) []byte {
	b.Reset()

	itms := make([]Item, 0)
	itms = append(itms, Item{
		Key:   "key",
		Value: "value",
		TTL:   time.Now().Add(time.Second * 5).Format(time.RFC3339),
	})

	strg := make([]string, 0)
	strg = append(strg, "redis")

	///////////////////////////////////////////////////////////////////////////////

	offsets := make([]flatbuffers.UOffsetT, len(strg))

	for i := len(strg) - 1; i >= 0; i-- {
		offsets[i] = b.CreateString(strg[i])
	}

	data.SetDataStartStoragesVector(b, len(offsets))

	for i := len(offsets) - 1; i >= 0; i-- {
		b.PrependUOffsetT(offsets[i])
	}

	storagesOffset := b.EndVector(len(offsets))

	///////////////////////////////////////////////////////////////////////////////

	offsets = make([]flatbuffers.UOffsetT, len(itms))

	for i := len(itms) - 1; i >= 0; i-- {
		offsets[i] = serializeItems(b, itms[i])
	}

	data.SetDataStartItemsVector(b, len(offsets))

	for i := len(offsets) - 1; i >= 0; i-- {
		b.PrependUOffsetT(offsets[i])
	}

	itemsOffset := b.EndVector(len(offsets))

	///////////////////////////////////////////////////////////////////////////////

	data.SetDataStart(b)
	data.SetDataAddStorages(b, storagesOffset)
	data.SetDataAddItems(b, itemsOffset)

	finalOffset := data.SetDataEnd(b)

	b.Finish(finalOffset)

	return b.Bytes[b.Head():]
}

func ReadData(buf []byte) {
	itemRoot := data.GetRootAsSetData(buf, 0)

	it := &data.Item{}

	itemRoot.Items(it, 0)

	println(string(it.Key()))
	println(string(it.Value()))
	println(string(it.Timeout()))

	return
}
