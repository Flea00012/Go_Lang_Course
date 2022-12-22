package messenger

import (
	"fmt"
	"testing"
	"reflect"
)

type Block struct {
	name string
	stringPassing chan string
	numPassing chan int
	left *Block
	right *Block
}

func New() *Block {
	return new(Block)
}

func (a *Block) AddNode(b *Block) {
	if(a.left == nil){
		a.left = b
		return
	} else {
		a.right = b
		return
	}
	fmt.Printf("node not added")
}

type BlockList []*Block

func Test(t *testing.T){
	fmt.Println("start tests")
	n := New()
	m := New()
	n.AddNode(m)

	if !reflect.DeepEqual(m, n.left) {
		fmt.Println("node added is not the same as leaf node in parent")
	}
	
	
}

