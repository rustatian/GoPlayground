package main

import (
	"unsafe"
)

type struct1 struct {
	uuid string
}

type uuidb [16]byte

type struct2 struct {
	uuid uuidb
}

func main() {
	s := struct1{
		uuid: "c599f965-e518-4f8c-98b9-7c53c10fcaiiiiiiiiiiiijrjpewjrtopiwje[prijtoeirjtoiiiiiiiiiiig'k;jkfgjks;dfgl;'kds;fkiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiic1",
	}
	s2 := struct2{
		uuid: [16]byte{1, 2, 3, 4, 5, 6, 7, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	println(int(unsafe.Sizeof(s)) + int(len(s.uuid)))
	println(unsafe.Sizeof(s2))
}
