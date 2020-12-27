package main

import (
	"context"
	"fmt"
)

//TODO change the package name

const (
	// Authorization is a key used for transport headers
	Authorization = "authorization"
	// CorrelationID is a key used for transport headers and context
	CorrelationID = "correlation-id"
	// jwt is a key used for context
	jwt = "jwt-token"
	// zipkin is a key used for context
	zipkin = "zipkin"
	// Routing is a key used for context
	Routing = "routing"
	// timeout
	timeout = "timeout"
	// unrecognized is a key used in case if there is no match for existing keys
	unrecognized = "unrecognized"
	// itemsCount is a key used in grpc to send number of products which client should expect from server
	itemsCount = "items-count"
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

var (
	// JwtKey is key for jwt in context or headers, name of the key is compatible with go-kit
	JwtKey = &contextKey{jwt}
	// ZipkinKey is key for zipkin context
	ZipkinKey = &contextKey{zipkin}
	// CorrelationIDKey is key for inter-service communication
	CorrelationIDKey = &contextKey{CorrelationID}
	// TimeoutKey is a key for product decode timeout
	TimeoutKey = &contextKey{timeout}
	// ItemsCount is a grpc key
	ItemsCount = &contextKey{itemsCount}
)

// String return string representation of the key
func (ck *contextKey) String() string {
	switch ck {
	case JwtKey:
		return jwt
	case ZipkinKey:
		return zipkin
	case CorrelationIDKey:
		return CorrelationID
	case TimeoutKey:
		return timeout
	case ItemsCount:
		return itemsCount
	default:
		return unrecognized
	}
}

// in order to pass linters we can't use primitive types like string
var contextKeys = []*contextKey{
	JwtKey, ZipkinKey, CorrelationIDKey, TimeoutKey, ItemsCount,
}

// CopyContext copies microservice related keys from old context to new
// exclude http/golang related data (timeouts, dedlines etc..)
/*
The types are:
ctx: old context

Returns
	New background context

ATTENTION: DO NOT FORGET TO ADD NEW KEYS TO THE SLICE ABOVE (contextKeys)
*/
func CopyContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	for i := 0; i < len(contextKeys); i++ {
		if value := ctx.Value(contextKeys[i].String()); value != nil {
			newCtx = context.WithValue(newCtx, interface{}(contextKeys[i].String()), value)
		}
	}
	return newCtx
}

func main() {
	ctx := context.Background()
	ctx2 := context.WithValue(ctx, contextKeys[1].String(), "1234")

	ctx3 := CopyContext(ctx2)
	fmt.Println(ctx3)
}

type sss struct {
}
