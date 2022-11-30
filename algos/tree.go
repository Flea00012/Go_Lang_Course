package algos

import "fmt"

type Tree struct {
	LeftNode *Tree
	Value int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil{
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree){
	if tree != nil{
		fmt.Println("value: ", tree.Value)
		fmt.Println("tree node left")
		fmt.Println(tree.LeftNode)
	}
}
