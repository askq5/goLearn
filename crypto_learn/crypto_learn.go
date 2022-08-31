package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

func Struct2String(obj interface{}) string {
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}

func main() {
	s := "sha1 this string"

	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()
	type data struct {
		Email string   `json:"email"`
		Auth  []string `json:"auth"`
	}
	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	data1 := data{Email: "huangming.666@bytedance.com", Auth: []string{"all"}}
	s1 := Struct2String(data1)
	s2 := "1631868560210"
	s1 = ""
	s2 = ""
	s3 := "dfed44eb-b1df-4231-83f4-b3754b57d0db"
	h.Write([]byte(s1 + s2 + s3))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
	fmt.Println(fmt.Sprintf("%x", bs))
}
