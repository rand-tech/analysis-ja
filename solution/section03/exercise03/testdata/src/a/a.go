package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.Atoi("100") // want "ignore error"
	fmt.Println(n)
}
