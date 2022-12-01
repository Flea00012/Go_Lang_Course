package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

type Transaction struct {
	CCNum string `bson:"ccnum"`
	Date string `bson:"date"`
	Amount float32 `bson:"amount"`
	Cvv string `bson: "cvv"`
	Expiration string `bson: "exp"`

}

func main() {
	session, err := mgo.Dial("127.0.0.1")
}