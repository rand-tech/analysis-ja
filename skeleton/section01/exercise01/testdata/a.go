package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	X [2]string
	Y string
}

func main() {
	t := T{
		X: [...]string{"A", "B"},
		Y: "C",
	}

	xp := uintptr(unsafe.Pointer(&t.X))
	yp := (*string)(unsafe.Pointer(xp + unsafe.Sizeof("")*2))
	fmt.Println(*yp)
}
