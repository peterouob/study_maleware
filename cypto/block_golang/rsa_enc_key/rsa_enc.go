package rsaenckey

import (
	"bgf/rsa_key"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

var public_key = Fileload("../rsa_key/public_key.pem")

func Fileload(path string) []byte {
	privatefile, err := os.Open(path)
	defer privatefile.Close()
	if errors.Is(err, os.ErrNotExist) {
		if err = rsa_key.GenerateKey(); err != nil {
			log.Println("Couldnot load file" + err.Error())
		}
	}
	privateKey := make([]byte, 2048)
	num, err := privatefile.Read(privateKey)
	if err != nil {
		log.Println("could not read privateKey")
	}
	return privateKey[:num]
}

// 加密函數
func RSA_ENC_KEY(origin []byte) ([]byte, error) {
	block, _ := pem.Decode(public_key)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origin) //解密函數
}

func RSAEncrypt(key string) string { //主RSAEncrypt
	var data []byte
	var err error
	data, err = RSA_ENC_KEY([]byte(key))
	if err != nil {
		log.Println("Error " + err.Error())
	}
	fmt.Println("加密 :", base64.StdEncoding.EncodeToString(data))
	return base64.StdEncoding.EncodeToString(data)
}
