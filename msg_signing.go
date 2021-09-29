package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func getSigningKey() *rsa.PrivateKey {
	content, err := ioutil.ReadFile("./rsa_private_key.pem") // the file is inside the local directory
	if err != nil {
		fmt.Println(err)
		return nil
	}
	block, _ := pem.Decode(content)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func testRSASigning() {
	message := []byte("message to be signed")
	hashed := sha256.Sum256(message)
	privateKey := getSigningKey()
	if privateKey == nil {
		return
	}

	var signature []byte
	var err error

	if true {
		// this one matches the output from the Openssl RSA_Sign()
		signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	} else {
		signature, err = rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer encoder.Close()

	encoder.Write(signature)
}

func main() {
	testRSASigning()
}
