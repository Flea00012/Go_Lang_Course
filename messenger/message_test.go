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
	} else if (a.right == nil) {
		a.right = b
		return
	}
	if (a.left.left == nil){
		a.left.AddNode(b)
	}
	a.right.AddNode(b)
}

func CreateBlockList() []*Block{
	return make([]*Block, 0)
}

type BlockList []*Block

func TestNodeInsertion(t *testing.T){
	fmt.Println("start tests")
	n := New()
	m := New()
	n.AddNode(m)
	o := New()
	n.AddNode(o)
	p := New()
	n.AddNode(p)
	Q := New()
	n.AddNode(Q)


	if !reflect.DeepEqual(m, n.left) && !reflect.DeepEqual(o, n.right) {
		fmt.Println("node added is not the same as leaf node in parent")
	}
	if !reflect.DeepEqual(p, n.left.left) && !reflect.DeepEqual(Q, n.right){
		fmt.Println("recursive insertion failed")
	}
}

func TestSending(t *testing.T) {
	a := &Block{
		name: "blocky",
		stringPassing: make(chan string),
		numPassing: make(chan int),
		left: &Block{},
		right: &Block{},
	}
	a.numPassing <- 1   
	x := <- a.numPassing         
	
	if !reflect.DeepEqual(1, x){
		fmt.Println("number passed between chans not arrived")
	}
	close(a.numPassing)

}

