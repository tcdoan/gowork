package main

import (
	"math/rand"
	"time"
)

type tree struct {
	value       int
	left, right *tree
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}

	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var x []int
	for i := 0; i < 20; i++ {
		x = append(x, r.Intn(1000))
	}

}
