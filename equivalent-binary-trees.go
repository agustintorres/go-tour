package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) { // in-order traversal
    WalkRecursive(t, ch)
    close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) { // in-order traversal
    if t == nil {
        return
    }
    WalkRecursive(t.Left, ch)
    ch <- t.Value
    WalkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
	
    for {
        v1,ok1 := <- ch1
        v2,ok2 := <- ch2

        if v1 != v2 || ok1 != ok2 {
            return false
        }

        if !ok1 {
            break
        }
    }
    return true
}

func WalkTest() {
    ch := make(chan int, 10)
    go Walk(tree.New(1), ch)
    for i := range ch {
	fmt.Println(i)
    }
}

func main() {
    WalkTest()
    result := Same(tree.New(1), tree.New(1)) 
    fmt.Println(result)
	
    result = Same(tree.New(1), tree.New(2)) 
    fmt.Println(result)
}
