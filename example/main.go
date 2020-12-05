package main

import (
	"fmt"
	"qrgo/encode"
)

func main() {
	length := encode.Counter("Hello")
	fmt.Println(length)
}
