package string_split

import "errors"

func NewSentencePart() SentencePart {
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