package main

import (
	"flag"
	"log"

	"github.com/googleapis/enterprise-certificate-proxy/client"
)

func main() {
	key, err := client.Cred("")
	if err != nil {
		log.Printf("Cred: got %v, want nil err", err)
	}
	mode := flag.String("mode", "", "the flag for action: -e for encryption or -d for decryption")
	msg := flag.String("msg", "", "the message to encrypt or decrypt")
	password := flag.String("password", "", "the password to use encryption or decryption")
	flag.Parse()

	if *mode == "" || *msg == "" || *password == "" {
		log.Printf("Usage: go run main.go -mode <-e|-d> -msg <msg> -password <password>")
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
		log.Printf("Encrypted message: %v", ciphertext)
	} else {
		byteSlice := []byte(*msg)
		plaintext, err := key.Decrypt(byteSlice)
		if err != nil {
			log.Printf("Decrypt: got %v, want nil err", err)
			return
		}
		log.Printf("Decrypted message: %v", plaintext)
	}
}
