package main

import (
	"encoding/json"
	"fmt"
	"os"
	"unsafe"
)

func stringMemSize(s string) uintptr {
	// Sizeof the string header + the memory for its content.
	return unsafe.Sizeof(s) + uintptr(len(s))
}

func main() {
	m := make(map[string]string)
	m["oneeeeee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["two"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"
	m["oneeeee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["oneeee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["onee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["oneee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["oneeedeee"] = "foooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooofoooooooooooooooooooooooooooo"
	m["twof"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"
	m["twgo"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"
	m["twso"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"
	m["tweo"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"
	m["twbo"] = "bbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbarbbbbbbbbbbbbbbbbbbbbbbbbbbbar"

	mapSize := unsafe.Sizeof(m)

	for key, value := range m {
		keySize := uintptr(len(key))
		valueSize := uintptr(len(value))
		mapSize += keySize + valueSize
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("/home/valery/Downloads/foo", b, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Estimated size of map in bytes: %d\n", mapSize)
}
