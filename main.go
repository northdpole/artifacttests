package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	m := md5.New()
	str := "test"
	m.Write([]byte(str))
	fmt.Printf("%x\n", m.Sum(nil))
}
