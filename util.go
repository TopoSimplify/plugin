package main

import (
	b64 "encoding/base64"
	"log"
)

func usageHelp() {
	log.Println("usage : ./plugin base_64_aphanumeric_string")
}

func encode64(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func decode64(s string) string {
	sDec, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return string(sDec)
}
