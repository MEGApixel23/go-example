package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sort"
	"sync"
)

var numberOfNodes int = 10
var wg sync.WaitGroup

func Walk(t *tree.Tree, ch chan int) {
	defer wg.Done()

	if len(ch) == 0 {
		wg.Add(1)
	}

	ch <- t.Value

	if t.Left != nil {
		wg.Add(1)
		go Walk(t.Left, ch)
	}

	if t.Right != nil {
		wg.Add(1)
		go Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	var slice1, slice2 []int

	ch1 := make(chan int, numberOfNodes)
	ch2 := make(chan int, numberOfNodes)

	Walk(t1, ch1)
	Walk(t2, ch2)

	wg.Wait()

	close(ch1)
	close(ch2)

	for i := range ch1 {
		slice1 = append(slice1, i)
	}
	sort.Ints(slice1)

	for i := range ch2 {
		slice2 = append(slice2, i)
	}
	sort.Ints(slice2)

	for i, _ := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}

	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
