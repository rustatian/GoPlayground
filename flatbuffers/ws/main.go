package main

import (
	"fmt"

	"github.com/rustatian/GoPlayground/flatbuffers/ws/message"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Message struct {
	// Command (join, leave, headers)
	Command string `json:"command"`

	// Broker (redis, memory)
	Broker string `json:"broker"`

	// Topic message been pushed into.
	Topics []string `json:"topic"`

	// Payload to be broadcasted
	Payload []byte `json:"payload"`
}

func main() {
	builder := flatbuffers.NewBuilder(100)

	messages := make([]Message, 0, 1)
	messages = append(messages, Message{
		Command: "join",
		Broker:  "memory",
		Topics:  []string{"aaaaaaaa", "bbbbbbbb"},
		Payload: []byte("hello, flatc"),
	}, Message{
		Command: "joinfff",
		Broker:  "masdfasdf",
		Topics:  []string{"aaaaaaaa", "bbbbbbbb"},
		Payload: []byte("hello, flatc"),
	},
	)
	bb := msgs(builder, messages)

	readmsgs(bb)
}

func serializeMsg(b *flatbuffers.Builder, msg Message) flatbuffers.UOffsetT {
	cmdOff := b.CreateString(msg.Command)
	brokerOff := b.CreateString(msg.Broker)

	offsets := make([]flatbuffers.UOffsetT, len(msg.Topics))
	for j := len(msg.Topics) - 1; j >= 0; j-- {
		offsets[j] = b.CreateString(msg.Topics[j])
	}

	message.MessageStartTopicsVector(b, len(offsets))

	for j := len(offsets) - 1; j >= 0; j-- {
		b.PrependUOffsetT(offsets[j])
	}

	tOff := b.EndVector(len(offsets))
	pOff := b.CreateByteVector(msg.Payload)

	message.MessageStart(b)

	message.MessageAddCommand(b, cmdOff)
	message.MessageAddBroker(b, brokerOff)
	message.MessageAddTopics(b, tOff)
	message.MessageAddPayload(b, pOff)

	return message.MessageEnd(b)
}

func msgs(b *flatbuffers.Builder, msgs []Message) []byte {
	b.Reset()

	mOff := make([]flatbuffers.UOffsetT, len(msgs))

	for i := len(msgs) - 1; i >= 0; i-- {
		mOff[i] = serializeMsg(b, msgs[i])
	}

	message.MessagesStartMessagesVector(b, len(mOff))

	for i := len(mOff) - 1; i >= 0; i-- {
		b.PrependUOffsetT(mOff[i])
	}

	msgsOff := b.EndVector(len(msgs))

	message.MessagesStart(b)
	message.MessagesAddMessages(b, msgsOff)
	fOff := message.MessagesEnd(b)
	b.Finish(fOff)

	return b.Bytes[b.Head():]
}

func readmsgs(buf []byte) {
	root := message.GetRootAsMessages(buf, 0)

	tmpMsg := &message.Message{}
	for i := 0; i < root.MessagesLength(); i++ {
		root.Messages(tmpMsg, i)

		l := tmpMsg.TopicsLength()

		fmt.Println(string(tmpMsg.Command()))
		fmt.Println(string(tmpMsg.Broker()))

		for i := 0; i < l; i++ {
			fmt.Println(string(tmpMsg.Topics(i)))
		}

		for i := 0; i < tmpMsg.PayloadLength(); i++ {
			fmt.Println(string(byte(tmpMsg.Payload(i))))
		}
	}
}
