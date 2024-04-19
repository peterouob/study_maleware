package rsa_key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func GenerateKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return errors.New("Cannot generate private key")
	}
	privateBitsKey := x509.MarshalPKCS1PrivateKey(privateKey)
	privateBitsPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBitsKey,
	})
	err = WriteToFile("private_key.pem", string(privateBitsPem))
	if err != nil {
		return errors.New("cannot write into private_key.pem")
	}
	publicKey := privateKey.PublicKey
	publicBitsKey := x509.MarshalPKCS1PublicKey(&publicKey)
	publicKeyBitsPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBitsKey,
	})
	err = WriteToFile("public_key.pem", string(publicKeyBitsPem))
	if err != nil {
		return errors.New("cannot write into public_key.pem")
	}
	return nil
}

func WriteToFile(path string, content string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()
	if _, err = f.Write([]byte(content)); err != nil {
		return errors.New("Cannot write into the file " + err.Error())
	}
	return nil
}
