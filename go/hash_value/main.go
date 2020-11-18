package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5_1 := md5.Sum([]byte("foo bar"))
	md5_2 := md5.Sum([]byte("foo bar baz"))
	md5_3 := md5.Sum([]byte("foo bar baz qux"))

	fmt.Println("md5: " + hex.EncodeToString(md5_1[:]))
	fmt.Println("md5: " + hex.EncodeToString(md5_2[:]))
	fmt.Println("md5: " + hex.EncodeToString(md5_3[:]))
}
