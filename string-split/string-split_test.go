package string_split

import (
	"fmt"
	"testing"
)

func TestACount(t *testing.T) {
	tests := map[string]int {
		"aaaa":4,
		"bfdsfdsf":0,
		"salut":1,
		"allolesamis":2,
		"aaloloiua":3,
		"lalalalala":5,
	}
	for k, v := range tests {
		if c := ACount(k); c != v {
			t.Errorf("ACount returned %v A for string %v but expected value is %v", c,k, v)
		}
	}
}

func TestACountPerPart(t *testing.T) {
	tests := map[string]int {
		"aaaa":-1,
		"bfdsfdsf":0,
		"salut":-1,
		"allolesamisa":1,
		"aalooiua":1,
		"lalalalalaa":2,
	}
	for k, v := range tests {
		if c := ACountPerPart(k); c != v {
			t.Errorf("ACountPerPart returned %v for string %v but expected value is %v", c,k, v)
		}
	}
}

func TestSentencePartNew(t *testing.T) {
	s := newSentencePart()
	if len(s.text) > 0 {
		t.Errorf("new sentence part's text should be empty, but has text %v",s.text)
	}

	if s.index != -1 {
		t.Errorf("new sentence part should have start index -1, but has %v",s.index)
	}
}

func TestStackSize(t *testing.T) {
	s := Stack{}
	if s.size() != 0 {
		t.Errorf("empty stack's size should be 0, got: %v",s.size())
	}

	s.push(SentencePart{})
	if s.size() != 1 {
		t.Errorf("After inserting an element, stack's size should be 1, got: %v",s.size())
	}

	s.push(SentencePart{})
	s.push(SentencePart{})
	s.push(SentencePart{})
	s.push(SentencePart{})
	if s.size() != 5 {
		t.Errorf("After inserting 5 elements, stack's size should be 5, got: %v",s.size())
	}

	s.pop()
	s.pop()

	if s.size() != 3 {
		t.Errorf("After inserting 5 elements and popping 2, stack's size should be 3, got: %v",s.size())
	}
}

func TestSentencePartSetValues(t *testing.T) {
	s := newSentencePart()
	s.index = 2
	s.text = "salut"
	if s.index != 2 {
		t.Errorf("Wrong value after setting index, expected 2, but got %v",s.index)
	}

	if s.text != "salut" {
		t.Errorf("Wrong value after setting text, expected salut, but got %v",s.text)
	}
}

func TestStackNewEmpty(t *testing.T) {
	s := Stack{}
	if !s.isEmpty() {
		t.Errorf("Freshly created stack should be empty")
	}

	if s.size() != 0 {
		t.Errorf("Freshly created stack's size should be 0")
	}
}

func TestStackPushOne(t *testing.T) {
	s := Stack{}
	st := newSentencePart()
	s.push(st)
	if s.isEmpty() {
		t.Errorf("Stack should not be empty after inserting an element")
	}

	if s.size() != 1 {
		t.Errorf("Stack's size should be 1 after inserting an element, but got %v",s.size())
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := Stack{}

	v,e := s.pop()
	if e == nil{
		t.Errorf("Popping an empty stack should return an error, got value %v",v)
	}
	s.push(SentencePart{"salut",2})
	s.push(SentencePart{"coucou",2})
	s.push(SentencePart{"bonjour",2})
	s.pop()
	s.pop()
	s.pop()

	v,e = s.pop()
	if e == nil{
		t.Errorf("Popping an empty stack should return an error, got value %v",v)
	}
}

func TestStackOnePushOnePop(t *testing.T) {
	s := Stack{}

	s.push(SentencePart{"hello",0})

	v, e := s.pop()

	if e != nil {
		t.Errorf("Got error when popping a stack with one element: %v",e)
	}

	if v.text != "hello" {
		t.Errorf("Wrong text, expected 'hello' but got: %v",v.text)
	}

	if v.index != 0 {
		t.Errorf("Wrong index, expected 0 but got: %v",v.index)
	}
}

func TestStackThreePushTwoPop(t *testing.T) {
	s := Stack{}

	v,e := s.pop()
	if e == nil{
		t.Errorf("Popping an empty stack should return an error, got value %v",v)
	}
	s.push(SentencePart{"salut",1})
	s.push(SentencePart{"coucou",2})
	s.push(SentencePart{"bonjour",3})
	s.pop()
	v, e = s.pop()

	if v.text != "coucou" {
		t.Errorf("Wrong text, expected 'coucou' but got: %v",v.text)
	}

	if v.index != 2 {
		t.Errorf("Wrong index, expected 2 but got: %v",v.index)
	}
}

func TestStackPeekEmpty(t *testing.T) {
	s := Stack{}
	v,e := s.peek()
	if e == nil{
		t.Errorf("Peeking an empty stack should return an error, got value %v",v)
	}
}

func TestStackPushOnePeek(t *testing.T) {
	s := Stack{}

	s.push(SentencePart{"salut",1})
	v,_ := s.peek()

	if s.size() != 1 {
		t.Errorf("Wrong size after peeking, expected 1 but got: %v",s.size())
	}

	if v.text != "salut" {
		t.Errorf("Wrong text, expected 'salut' but got: %v",v.text)
	}

	if v.index != 1 {
		t.Errorf("Wrong index, expected 1 but got: %v",v.index)
	}
}

func TestSolution(t *testing.T) {
	tests := map[string]int {
		"babaa":2,
		"ababa":4,
		"aba":0,
		//"bbbbb":6,
	}

	for k, v := range tests {
		testName := fmt.Sprintf("%v is expected to return %v", k, v)
		t.Run(testName, func(t *testing.T) {
			if s := Solution(k); s != v {
				t.Errorf("got %v, want %v", s, v)
			}
		})
	}
}