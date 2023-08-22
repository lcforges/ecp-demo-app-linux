package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"log"

	"github.com/googleapis/enterprise-certificate-proxy/client"
)

const (
	configFilePath = "linux_config.json"
)

func main() {
	key, err := client.Cred(configFilePath)
	if err != nil {
		log.Printf("Cred: got %v, want nil err", err)
		return
	}
	mode := flag.String("mode", "", "the flag for action: -e for encryption or -d for decryption")
	msg := flag.String("msg", "", "the message to encrypt or decrypt")
	flag.Parse()

	if *mode == "" || *msg == "" {
		log.Printf("Usage: go run main.go -mode <-e|-d> -msg <msg>")
		return
	}

	if *mode != "-e" && *mode != "-d" {
		log.Printf("Invalid mode: %v", *mode)
		return
	}

	if *mode == "-e" {
		byteSlice := []byte(*msg)
		ciphertext, err := key.Encrypt(byteSlice)
		if err != nil {
			log.Printf("Encrypt: got %v, want nil err", err)
			return
		}
		cipherString := base64.StdEncoding.EncodeToString(ciphertext)
		log.Printf("Encrypted message: %v", cipherString)
	} else {
		str, err := base64.StdEncoding.DecodeString(*msg)
		if err != nil {
			log.Printf("Base64: got %v, want nil err", err)
			return
		}
		plaintext, err := key.Decrypt([]byte(str))
		if err != nil {
			log.Printf("Decrypt: got %v, want nil err", err)
			return
		}
		plaintext = bytes.Trim(plaintext, "\x00")
		plainString := string(plaintext)
		log.Printf("Decrypted message: %v", plainString)
	}
}
