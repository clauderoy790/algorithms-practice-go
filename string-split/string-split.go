package string_split

import (
	"errors"
	"fmt"
	"strings"
)

type StringSplit struct{}

//split it into 3 NON empty parts
// Length doesn't matter as long as > 1
// each part must contain the same number of letters a
// Find how many different way you can split

func (ht *StringSplit) Resolve() {
	s := "bbbbb"
	//s := "ababa"
	//s := "babaa"
	fmt.Println("string split problem!!")
	fmt.Printf("Solution is: %v\n", Solution(s))
	// should be 4
}

func Solution(S string) int {
	nbA := ACountPerPart(S)
	if nbA < 0 {
		return 0
	}
	count := 0
	runes := []rune(S)

	stack := Stack{}
	part := ""
	for i := 0; i < len(runes);i++ {
		part += string(runes[i])

		//stack size will always be 1 or 2
		aCnt := ACount(part)
		if stack.size() == 0 && (aCnt > nbA || len(part) == len(S)-1) {
			break
		} else if i == len(runes)-1 {
			if stack.size() == 2 {
				count++
			}
			v, _ := stack.pop()
			i,part = v.index,v.text
		} else if stack.size() < 2 && aCnt == nbA {
			stack.push(SentencePart{part,i})
			part = ""
		}
	}

	return count
}

func newSentencePart() SentencePart {
	return SentencePart{"",-1}
}

type Stack struct {
	 elements []SentencePart
}

func (s *Stack) push(p SentencePart) {
	s.elements = append(s.elements,p)
}

func (s *Stack) pop() (part SentencePart, err error) {
	if s.size() == 0 {
		err = errors.New("cannot pop an empty stack")
		return
	}
	part = s.elements[s.size()-1]
	s.elements = s.elements[:len(s.elements)-1]
	return
}

func (s *Stack) peek() (part SentencePart, err error) {
	if s.size() == 0 {
		err = errors.New("cannot peek an empty stack")
		return
	}
	part = s.elements[s.size()-1]
	return
}

func (s *Stack) size() int {
	return len(s.elements)
}

func (s *Stack) isEmpty() bool {
	return s.size() == 0
}

type SentencePart struct {
	text  string
	index int
}

func ACount(str string) int {
	return strings.Count(str,"a")
}

func ACountPerPart(str string) int {
	cnt := ACount(str)
	if cnt % 3 != 0 {
		return -1
	} else {
		return cnt / 3
	}
}
