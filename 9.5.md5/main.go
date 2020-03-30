package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "123")
	fmt.Printf("%x\n", h.Sum(nil))

	h = md5.New()
	io.WriteString(h, "123")
	io.WriteString(h, "snail")
	fmt.Printf("%x\n", h.Sum(nil))
}
