package proj3

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

type ChatRoom struct {
    roomName string
    Message chan *ChatMessage
}

type ChatMessage struct {
    Message string
    SenderID string
    SenderNick string
}

type ChatServer int

func GetSignKeys() []byte {
	privkey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256([]byte("passwordforaccess"))
	signature , err := ecdsa.SignASN1(rand.Reader, privkey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %v", signature)
	valid := ecdsa.VerifyASN1(&privkey.PublicKey, hash[:], signature)
	fmt.Println("signature verified: ", valid)

	return signature
}