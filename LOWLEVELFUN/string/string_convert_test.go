package main

import (
	"reflect"
	"testing"
	"unsafe"
)

// BenchmarkToStringUnsafe-12    	566317729	         1.91 ns/op	       0 B/op	       0 allocs/op
// inline BenchmarkToStringUnsafe-12    	1000000000	         0.295 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToStringUnsafe(b *testing.B) {
	testPayload := []byte("falsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtoj")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res := AsString(testPayload)
		_ = res
	}
}

// BenchmarkToStringSafe-12    	28584489	        39.1 ns/op	     112 B/op	       1 allocs/op
// inline BenchmarkToStringSafe-12    	28926276	        46.6 ns/op	     128 B/op	       1 allocs/op
func BenchmarkToStringSafe(b *testing.B) {
	testPayload := []byte("falsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtoj")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res := toStringNotFun(testPayload)
		_ = res
	}
}

// BenchmarkToStringUnsafe-12    	566317729	         1.91 ns/op	       0 B/op	       0 allocs/op
// inline BenchmarkToStringUnsafe-12    	1000000000	         0.295 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToStringWrong(b *testing.B) {
	testPayload := []byte("falsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtojfalsflasjlifjwpoihejfoiwejow{}{}{}{}jelfjasjfhwaopiehjtopwhtgohrgouahsgkljasdlfjasl;fjals;jdflkndgouwhetopwqhjtoj")
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		res := AsString(testPayload)
		_ = res
	}
}

func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func toStringNotFun(data []byte) string {
	return string(data)
}

// AsBytes returns a slice that refers to the data backing the string s.
func AsBytes(s string) []byte {
	// get the pointer to the data of the string

	p := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)

	var b []byte
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Data = uintptr(p)
	// we need to set the cap and len for the string to byte convert
	// because string is shorter than []bytes
	hdr.Cap = len(s)
	hdr.Len = len(s)

	return b
}

// AsString returns a string that refers to the data backing the slice s.
func AsString(b []byte) string {
	p := unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)

	var s string
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(p)
	hdr.Len = len(b)

	return s
}
