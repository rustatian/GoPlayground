package main

type ExchangeType string

const (
	Direct  ExchangeType = "direct"
	Fanout  ExchangeType = "fanout"
	Topic   ExchangeType = "topic"
	Headers ExchangeType = "headers"
)

func main() {
	//isValidExchangeType := map[string]bool{"direct": true, "fanout": true, "topic": true, "headers": true}

	foo()

}

func foo(e ExchangeType) {

}

type queue struct {
	exchType ExchangeType
}
