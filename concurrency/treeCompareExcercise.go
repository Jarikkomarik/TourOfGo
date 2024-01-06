package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkDeep(t, ch)
	close(ch)

}

func walkDeep(t *tree.Tree, ch chan int) {
	if t.Left == nil && t.Right == nil {
		ch <- t.Value
		return
	}

	if t.Left != nil {
		walkDeep(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walkDeep(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, isWorking1 := <-ch1
		val2, isWorking2 := <-ch2

		if isWorking1 && isWorking2 { //if both active compare values
			if val1 != val2 {
				fmt.Println("values of trees differs")
				return false
			}
		} else if !isWorking1 && !isWorking2 { //if reached the end of both -> they are identical
			return true
		} else {
			return false
		}
	}

}

func main() {
	tree1, tree2 := tree.New(1), tree.New(1)

	fmt.Println(Same(tree1, tree2))

	fmt.Println(tree.New(3))

}
