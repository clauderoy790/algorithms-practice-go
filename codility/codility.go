package codility

import (
	"fmt"
	"sort"
)

type Codility struct{}

func (cd *Codility) Resolve() {
	fmt.Println("test codility")
	ints := []int {2, 3, 3, 4}
	//ints := []int{5, 2, 5, 2}
	//fmt.Printf("the solution is: %v", Solution(ints, 8, 1))
	fmt.Printf("the solution is: %v", Solution(ints, 3, 1))
}

func Solution(A []int, L int, R int) (nbDisks int) {
	B := make([]int, len(A))
	C := make([]int, len(A))
	copy(B, A)
	copy(C, A)
	sort.Ints(C)
	sort.Sort(sort.Reverse(sort.IntSlice(B)))
	fmt.Println("reverse: ", B)
	fmt.Println("A: ", A)

	stackL := Stack{}
	stackR := Stack{}
	stackL.canAdd = func(newNb int) (canAdd bool) {
		if nb, ok := stackL.Peek(); ok {
			return newNb < nb
		}
		return false
	}

	stackR.canAdd = func(newNb int) (canAdd bool) {
		if nb, ok := stackR.Peek(); ok {
			return newNb > nb
		}
		return false
	}

	stackL.Push(L)
	stackR.Push(R)
	ok := false
	for _, nb := range C {
		if stackR.canAdd(nb) {
			A, ok = removeInt(A, nb)
			if ok {
				nbDisks++
				stackR.Push(nb)
			}
		}
	}

	for _, nb := range B {
		if stackL.canAdd(nb) {
			A, ok = removeInt(A, nb)
			if ok {
				nbDisks++
				stackL.Push(nb)
			}
		}
	}
	return

	// add left to left stack, add right to right stack
	// add function CanAdd(number) that checks if can add
	// Loop through the numbers array
	// check CanAdd
	//   if true, push on stack
	//		increment nbDisks
}

func removeInt(sl []int, nb int) ([]int, bool) {
	for i,v := range sl {
		if v == nb {
			return append(sl[:i], sl[i+1:]...), true
		}
	}
	return sl, false
}

type Stack struct {
	elements []int
	canAdd   func(nb int) bool
}

func (s *Stack) Push(nb int) {
	s.elements = append(s.elements, nb)
}

func (s *Stack) Pop(nb int, ok bool) {
	if len(s.elements) > 0 {
		ok = true
		nb = s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
	}
	return
}

func (s *Stack) CanAdd(nb int) bool {
	return s.canAdd(nb)
}

func (s *Stack) Count() int {
	return len(s.elements)
}

func (s *Stack) Peek() (nb int, ok bool) {
	if len(s.elements) > 0 {
		ok = true
		nb = s.elements[len(s.elements)-1]
	}
	return
}
