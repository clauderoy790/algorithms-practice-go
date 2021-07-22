package string_split

import (
	"fmt"
	"strings"
)

type StringSplit struct{}

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
