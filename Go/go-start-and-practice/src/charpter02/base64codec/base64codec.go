package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away from keyboard. https://golang.org"

	// encoding
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encoded)

	// decoding
	data, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
