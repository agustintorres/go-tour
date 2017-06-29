package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) { // in-order traversal
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func WalkTest() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i:= 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
}

func main() {
	WalkTest()
}
