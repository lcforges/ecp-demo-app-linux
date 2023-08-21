package main

import (
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
	ciphertext, err := key.Encrypt(byteSlice)
	if err != nil {
		t.Errorf("Encrypt: got %v, want nil err", err)
		return
	}
	t.Logf("Encrypted: %v", ciphertext)
}
