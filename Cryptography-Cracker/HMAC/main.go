package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var key = []byte("some random key")

func checkMAC(message, recvMAC []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	calcMAC := mac.Sum(nil)

	return hmac.Equal(calcMAC, recvMAC)
}

func main() {
	message := []byte("The red eagle flies at 10:00") //Example of message
	mac, _ := hex.DecodeString("69d2c7b6fbbfcaeb72a3172f4662601d1f16acfb46339639ac8c10c8da64631d") //Example of HexString
	if checkMAC(message, mac) {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("NOT EQUAL")
	}
}