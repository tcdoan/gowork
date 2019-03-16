package main

import "fmt"

type myset struct {
	words []uint64
}

func (s *myset) Add(x uint64) {
	word, bits := x/64, x%64
	for word >= uint64(len(s.words)) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bits)
}

func (s myset) Contains(x uint64) bool {
	word, bits := x/64, x%64
	if word < uint64(len(s.words)) {
		return (s.words[word] & (1 << bits)) != 0
	}
	return false
}

func main() {
	s1 := myset{}
	s1.Add(9)
	s1.Add(8)
	s1.Add(42)
	s1.Add(12564788999)
	fmt.Println(s1.Contains(9))
	fmt.Println(s1.Contains(11))
	fmt.Println(s1.Contains(42))
	fmt.Println(s1.Contains(12564788999))
	fmt.Println(s1.Contains(12564788999 + 1))
}
