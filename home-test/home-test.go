package home_test

import (
	"fmt"
	"strings"
)

type HomeTest struct{}

//split it into 3 NON empty parts
// Length doesn't matter as long as > 1
// each part must contain the same number of letters a
// Find how many different way you can split

func (ht *HomeTest) Resolve() {
	s := "ababa"
	//s := "babaa"
	fmt.Println("hello home test!!!!")
	fmt.Printf("Solution is: %v\n", Solution(s))
	// should be 2
}

var sh = SentenceHelper{}

func Solution(S string) int {
	//1. count number of a
	//2. count max number of a which would be len (S) / 3
	//3. make part struct containing p1,p2,p3
	//4. loop from 1 until <= max number (step 2)
	//5. Iterate through characters.
	//6. when finding that number of a is desired number, check if any current part are equal to this current string
	//6.1 when finding a a, check if number of remaining a will be enough to split it into other parts
	//7. if no, add this as part one, if yes,keep adding to the part and perform the same check
	//8. When completing string to the end, add Part object to list
	//9. Reiterate from beginning UNTIL your part one contains too many a and there isn't enough other a to make part 2 and 3
	info := getProblemInfo(S)
	fmt.Println("info: ", info)
	for i := 1; i <= info.MaxAPerPart; i++ {
		fmt.Println("counting pars for A: ", i)
		for j := 0; j < 3; j++ {
			processPart(i, S)
		}
	}

	return sh.Count()
}

func processPart(nbA int, str string) {
	morePossibilities := true //What is no one part is valid?
	//for morePossibilities {

	for morePossibilities {
		fmt.Println("more possib", morePossibilities)
		sentence := newSentence()
		currentPart := ""
		partId := 0

		for i, c := range str {
			currentPart += string(c)
			//if you have enough A in the part and it's not already contained in the did one, add it
			// AND if there are enough remaining A
			// Or if you're at the end of the string
			if (ACount(currentPart) == nbA && !sh.containsSentence(partId, currentPart, sentence) &&
				sh.EnoughRemainingA(str, i, nbA, partId)) || i == len(str)-1 {
				sentence.Set(partId, currentPart)
				if currentPart == str {
					fmt.Println("DONE2")
					morePossibilities = false
				}
				fmt.Printf("setting part %v to %v\n", partId, currentPart)
				currentPart = ""
				partId++
			} else if ACount(sentence.parts[0]) > nbA {
				fmt.Printf("DONE part 0: %v, nb: %v\n",sentence.parts[0],nbA)
				morePossibilities = false
				break
			}
		}
		//a ba ba --
		//ab a ba --
		//a bab a
		//ab ab a
		fmt.Println("try sentence: ",sentence)
		if sentence.isValid(nbA) {
			sh.Add(sentence)
			fmt.Printf("adding sentence: %v\n", &sentence)
		}
	}

}

func ACount(str string) int {
	return strings.Count(str, "a")
}

func newSentence() Sentence {
	s := Sentence{}
	s.parts = make([]string, 3)
	return s
}

type Sentence struct {
	parts []string
}

func (s *Sentence) isValid(aCount int) bool {
	p1ACount := ACount(s.parts[0])
	p2ACount := ACount(s.parts[1])
	p3ACount := ACount(s.parts[2])
	return p1ACount == aCount && p1ACount == p2ACount && p2ACount == p3ACount
}

func (s *Sentence) Set(id int, part string) {
	s.parts[id] = part
}

type SentenceHelper struct {
	sentences []Sentence
}

func (s *SentenceHelper) Add(sentence Sentence) {
	s.sentences = append(s.sentences, sentence)
}

func (s *SentenceHelper) Count() int {
	return len(s.sentences)
}

func (s *SentenceHelper) containsSentence(partId int, currentPart string, sentence Sentence) bool {
	sentence.Set(partId, currentPart)
	contained := true

	for i, part := range sentence.parts {
		if contained && i <= partId {
			contained = contained && s.containPart(i, part)
		}
	}

	return contained
}

func (s *SentenceHelper) containPart(id int, str string) bool {
	for _, sentence := range s.sentences {
		if sentence.parts[id] == str {
			return true
		}
	}
	return false
}

func (s *SentenceHelper) EnoughRemainingA(str string, ind, nbA, partId int) bool {
	remaining := str[(ind + 1):]
	remainingCount := ACount(remaining)
	remainingParts := 3 - (partId + 1)
	//fmt.Println("str: ",str)
	//fmt.Printf("remaining: %v, reaminingCount: %v, remainingParts: %v, current part: %v, criss: %v\n",remaining,remainingCount,remainingParts,partId,str[(ind+1):])
	//fmt.Println("remainingCount: ",remainingCount)
	//fmt.Println("needed count: ",(remainingParts*nbA))
	//fmt.Println("NB A:",nbA)

	return remainingCount == remainingParts*nbA
}

func getProblemInfo(s string) ProblemInfo {
	aCount := ACount(s)
	return ProblemInfo{aCount, aCount / 3}
}

type ProblemInfo struct {
	TotalA, MaxAPerPart int
}
