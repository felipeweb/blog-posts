package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func generatePublicKey(privateKey *rsa.PrivateKey) (err error) {
	publicKey := privateKey.PublicKey
	pub, _ := ssh.NewPublicKey(&publicKey)
	fmt.Println("PUBLIC KEY:", string(ssh.MarshalAuthorizedKey(pub)))
	return
}

func generatePrivateKey(bits int) (privateKey *rsa.PrivateKey, err error) {
	privateKey, _ = rsa.GenerateKey(rand.Reader, bits)
	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := pem.EncodeToMemory(&privateKeyBlock)
	fmt.Println("PRIVATE KEY:", string(privateKeyPem))
	return
}

func main() {
	privateKey, err := generatePrivateKey(4096)
	if err != nil {
		panic(err)
	}
	err = generatePublicKey(privateKey)
	if err != nil {
		panic(err)
	}
}
