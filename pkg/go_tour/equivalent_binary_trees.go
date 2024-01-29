package go_tour

import (
	"fmt"
	"strconv"

	"golang.org/x/tour/tree"
)

type EquivalentBinaryTrees struct{}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	walkRecursive(t, ch)
}

func walkRecursive(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkRecursive(t.Left, ch)

	ch <- t.Value

	walkRecursive(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var s1, s2 string
	for t1v := range ch1 {
		s1 += strconv.Itoa(t1v)
	}
	for t2v := range ch2 {
		s2 += strconv.Itoa(t2v)
	}

	return s1 == s2
}

func (e EquivalentBinaryTrees) Run() {
	ch := make(chan int, 10)
	t1 := tree.New(4)
	t2 := tree.New(4)

	go Walk(t1, ch)

	for t1v := range ch {
		fmt.Println(t1v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(t1, t2))
}
