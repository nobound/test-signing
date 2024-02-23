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
	"log/slog"
	"os"
)

// Read the RSA private key from the PEM file
// (PEM blocks of type "RSA PRIVATE KEY").
func getSigningKey(keyFile string) (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	content, err := os.ReadFile(keyFile)
	if err != nil {
		return key, err
	}

	block, _ := pem.Decode(content)
	if block.Type == "RSA PRIVATE KEY" {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else if block.Type == "PRIVATE KEY" {
		// this is for PEM blocks of type "PRIVATE KEY".
		parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tmpKey, ok := parsedKey.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("unable to parse private key, unexpected private key type: %T", tmpKey)
		}
		key = tmpKey
	} else {
		return nil, fmt.Errorf("unexpected PEM block type: %v", block.Type)
	}
	return key, err
}

// Test RSA Signing
func testRSASigning() {
	message := []byte("message to be signed")
	hashed := sha256.Sum256(message)

	keyFile := "./rsa_private_key.pem"
	privateKey, err := getSigningKey(keyFile)
	if err != nil {
		slog.Error("Failed to get private key for signing: %v", err)
		return
	}

	signing := func(key *rsa.PrivateKey, hashed []byte) ([]byte, error) {
		var signature []byte
		var err error
		signature, err = rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
		// signature, err = rsa.SignPSS(rand.Reader, key, crypto.SHA256, hashed, nil)
		if err != nil {
			return nil, err
		}
		return signature, nil
	}

	signature, err := signing(privateKey, hashed[:])
	if err != nil {
		slog.Error("Failed to signing: %v", err)
		return
	}

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer encoder.Close()
	encoder.Write(signature)
}

func main() {
	testRSASigning()
}
