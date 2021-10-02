package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	// encode
	girlfriend := "cyw"
	fmt.Println("before encode:", []byte(girlfriend))
	debyte := base64Encode([]byte(girlfriend))
	fmt.Println("after encode:", debyte)

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println("encode, error")
	}
	fmt.Println("after decode:", enbyte)

	if string(enbyte) != girlfriend {
		fmt.Println("error girlfriend")
	}

	fmt.Println(string(girlfriend))
}
