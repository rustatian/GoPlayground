package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/auth/jwt"
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.
type contextKey struct {
	name string
}

var (
	// JwtKey is key for jwt in context or headers, name of the key is compatible with go-kit
	JwtKey    = &contextKey{"JWTToken"}
	// ZipkinKey is key for zipkin context
	ZipkinKey = &contextKey{"zipkin"}
)

// String return string representation of the key
func (ck *contextKey) String() string {
	switch ck {
	case JwtKey:
		return "JWTToken"
	case ZipkinKey:
		return "zipkin"
	default:
		return "unrecognized"
	}
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, JwtKey.String(), "17")

	fmt.Println(ctx.Value(JwtKey.String()))
	tt := string(jwt.JWTTokenContextKey)

	fmt.Println(ctx.Value(tt))

	//aa := JwtKey.String()
	//fmt.Println(aa)
	//
	//fmt.Println(ctx.Value(JwtKey.String()))


	
}
