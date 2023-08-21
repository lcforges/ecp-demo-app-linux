package main

import (
	"bytes"
	"testing"

	"github.com/googleapis/enterprise-certificate-proxy/client"
)

const (
	linuxConfigFilePath = "linux_config.json"
)
func TestEncrypt(t *testing.T) {
	key, err := client.Cred(linuxConfigFilePath)
	if err != nil {
		t.Errorf("Cred: got %v, want nil err", err)
		return
	}
	msg := "This is my secret message"
	byteSlice := []byte(msg)
	_, err = key.Encrypt(byteSlice)
	if err != nil {
		t.Errorf("Encrypt: got %v, want nil err", err)
		return
	}
}

func TestDecrypt(t *testing.T) {
	key, err := client.Cred(linuxConfigFilePath)
	if err != nil {
		t.Errorf("Cred: got %v, want nil err", err)
		return
	}
	msg := "This is my secret message"
	byteSlice := []byte(msg)
	ciphertext, err := key.Encrypt(byteSlice)
	if err != nil {
		t.Errorf("Encrypt: got %v, want nil err", err)
		return
	}
	decrypted, err := key.Decrypt(ciphertext)
	if err != nil {
		t.Errorf("Decrypt: got %v want nil err", err)
		return
	}
	decrypted = bytes.Trim(decrypted, "\x00")
	if string(decrypted) != msg {
		t.Errorf("Decrypt error expected: %v, got: %v", msg, string(decrypted))
	}

}
