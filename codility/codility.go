package codility

import (
	"fmt"
)

type Codility struct{}

func (cd *Codility) Resolve() {
	fmt.Println("test codility")
	ints := []int{2, 3, 3, 4}
	//ints := []int{5, 2, 5, 2}
	//fmt.Printf("the solution is: %v", Solution(ints, 8, 1))
	fmt.Printf("the solution is: %v", Solution(ints, 3, 1))
}

func Solution(A []int, L int, R int) (nbDisks int) {
	stacks := createStacks(L,R)

	check := true
	for check {
		check = false
		for i := 0; i < len(A); i++ {
			for _, stack := range stacks {
				if stack.canAdd(A[i]) {
					nbDisks++
					stack.Push(A[i])
					A, _ = removeAt(A, i)
					check = true
					i--
					break
				}
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

func createStacks(L, R int) []Stack {
	stackL := Stack{}
	stackL.Push(L)
	stackL.canAdd = func(newNb int) (canAdd bool) {
		if nb, ok := stackL.Peek(); ok {
			return newNb < nb
		}
		return false
	}

	stackR := Stack{}
	stackR.Push(R)
	stackR.canAdd = func(newNb int) (canAdd bool) {
		if nb, ok := stackR.Peek(); ok {
			return newNb > nb
		}
		return false
	}

	return []Stack{stackL, stackR}
}

func removeAt(sl []int, i int) ([]int, bool) {
	if i < 0 || i >= len(sl) {
		return sl, false
	}

	return append(sl[:i], sl[i+1:]...), true
}

type Stack struct {
	elements []int
	canAdd   func(nb int) bool
}

func (s *Stack) Push(nb int) {
	s.elements = append(s.elements, nb)
}

func (s *Stack) CanAdd(nb int) bool {
	return s.canAdd(nb)
}

func (s *Stack) Peek() (nb int, ok bool) {
	if len(s.elements) > 0 {
		ok = true
		nb = s.elements[len(s.elements)-1]
	}
	return
}
