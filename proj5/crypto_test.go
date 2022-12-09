package proj5

import (
	"testing"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func TestCrypto(t *testing.T){
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	m := "the hash string for bytes"
	hash := sha256.Sum256([]byte(m))
	hashSign, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", hashSign)

	verified := ecdsa.VerifyASN1(&privKey.PublicKey, hash[:], hashSign)
	fmt.Println("signature verified: ", verified)
}