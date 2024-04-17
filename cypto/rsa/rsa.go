package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Cannot generate key")
	}
	privateBitsKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBitsPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBitsKey,
	})
	WriteToFile("private_key.pem", string(privateBitsPem))

	publicKey := privateKey.PublicKey
	publicBitsKey := x509.MarshalPKCS1PublicKey(&publicKey)
	publicKeyBitsPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBitsKey,
	})
	WriteToFile("public_key.pem", string(publicKeyBitsPem))
}

func WriteToFile(path string, content string) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	if _, err = f.Write([]byte(content)); err != nil {
		fmt.Println("Cannot write into the file " + err.Error())
	}
}
