package main

import "golang.org/x/tour/tree"

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	cht1 := make(chan int)
	cht2 := make(chan int)
	go Walk(t1, cht1)
	go Walk(t2, cht2)

	for {
		v1, ok1 := <-cht1
		v2, ok2 := <-cht2
		switch {
		case !ok1, !ok2:
			return ok1 == ok2
		case v1 != v2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
