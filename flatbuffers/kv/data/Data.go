// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Data struct {
	_tab flatbuffers.Table
}

func GetRootAsData(buf []byte, offset flatbuffers.UOffsetT) *Data {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Data{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Data) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Data) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Data) Storage() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Data) Keys(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Data) KeysLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Data) Timeout() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func DataStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func DataAddStorage(builder *flatbuffers.Builder, Storage flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(Storage), 0)
}
func DataAddKeys(builder *flatbuffers.Builder, Keys flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(Keys), 0)
}
func DataStartKeysVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DataAddTimeout(builder *flatbuffers.Builder, Timeout flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Timeout), 0)
}
func DataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
