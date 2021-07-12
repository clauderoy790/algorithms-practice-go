package string_split

//import (
//	"fmt"
//	"strings"
//)
//
//var sh = SentenceHelper{}
//var str = ""
//
//func Solution(S string) int {
//	//1. count number of a
//	//3. make part struct containing p1,p2,p3
//	//5. Iterate through characters.
//	//6. when finding that number of a is desired number, check if any current part are equal to this current string
//	//7. if no, add this as part one, if yes,keep adding to the part and perform the same check
//	//8. When completing string to the end, add Part object to list
//	//9. Reiterate from beginning UNTIL your part one contains too many a and there isn't enough other a to make part 2 and 3
//
//	//a ba ba
//	//ab a ba
//	//
//	//missing
//	//a bab a
//	//ab ab a
//	str = S
//	info := getProblemInfo(str)
//	if info.TotalA % 3 != 0 {
//		return 0
//	}
//	fmt.Println("info: ", info)
//	for i := 1; i <= info.MaxAPerPart; i++ {
//		fmt.Println("counting pars for A: ", i)
//		for j := 0; j < 3; j++ {
//			processPart(i)
//		}
//	}
//
//	return sh.Count()
//}
//
//func processPart(nbA int) {
//	morePossibilities := true //What is no one part is valid?
//	//for morePossibilities {
//
//	for morePossibilities {
//		fmt.Println("more possib", morePossibilities)
//		sentence := newSentence()
//		currentPart := ""
//		partId := 0
//
//		for i, c := range str {
//			currentPart += string(c)
//			//if you have enough A in the part and it's not already contained in the did one, add it
//			// AND if there are enough remaining A
//			// Or if you're at the end of the string
//			if (ACount(currentPart) == nbA && !sh.containsSentence(currentPart, sentence) &&
//				sh.EnoughRemainingA(sentence,currentPart, nbA)) || i == len(str)-1 {
//				sentence.Set(partId, currentPart)
//				if currentPart == str {
//					fmt.Println("DONE2")
//					morePossibilities = false
//				}
//				fmt.Printf("setting part %v to %v\n", partId, currentPart)
//				currentPart = ""
//				partId++
//			} else if ACount(sentence.parts[0]) > nbA {
//				fmt.Printf("DONE part 0: %v, nb: %v\n",sentence.parts[0],nbA)
//				morePossibilities = false
//				break
//			}
//		}
//		//a ba ba --
//		//ab a ba --
//		//a bab a
//		//ab ab a
//		fmt.Println("try sentence: ",sentence)
//		if sentence.isValid(nbA) {
//			sh.Add(sentence)
//			fmt.Printf("adding sentence: %v\n", &sentence)
//		}
//	}
//
//}
//
//func ACount(str string) int {
//	return strings.Count(str, "a")
//}
//
//func newSentence() Sentence {
//	s := Sentence{}
//	s.parts = make([]string, 3)
//	return s
//}
//
//type Sentence struct {
//	parts []string
//}
//
//func (s *Sentence) isValid(aCount int) bool {
//	p1ACount := ACount(s.parts[0])
//	p2ACount := ACount(s.parts[1])
//	p3ACount := ACount(s.parts[2])
//	return p1ACount == aCount && p1ACount == p2ACount && p2ACount == p3ACount
//}
//
//func (s *Sentence) nbEmptyPart() int {
//	count := 0
//	for _, part := range s.parts {
//		if len(part) == 0 {
//			count++
//		}
//	}
//	return count
//}
//
//func (s *Sentence) Set(id int, part string) {
//	s.parts[id] = part
//}
//
//func (s *Sentence) addPart(part string) {
//	id := -1;
//
//	for i,v := range s.parts {
//		if len(v) == 0 {
//			id = i
//		}
//	}
//
//	if id != -1 {
//		s.parts[id] = part
//	}
//
//}
//
//func (s *Sentence) String() string {
//	str := ""
//	for _, part := range s.parts{
//		str += part
//	}
//	return str
//}
//
//type SentenceHelper struct {
//	sentences []Sentence
//}
//
//func (s *SentenceHelper) Add(sentence Sentence) {
//	s.sentences = append(s.sentences, sentence)
//}
//
//func (s *SentenceHelper) Count() int {
//	return len(s.sentences)
//}
//
//func (s *SentenceHelper) containsSentence(currentPart string, sentence Sentence) bool {
//	sentence.addPart(currentPart)
//	for _, sentence := range s.sentences {
//		contained := true
//		atLeastOnePart := false
//		for i, part := range sentence.parts {
//			if len(part) > 0 {
//				atLeastOnePart = true
//				contained = contained && sentence.parts[i] == part
//			}
//			if atLeastOnePart && contained {
//				return true
//			}
//		}
//	}
//
//	return false
//}
//
//func (s *SentenceHelper) containPart(id int, str string) bool {
//	return false
//}
//
//func (s *SentenceHelper) EnoughRemainingA(sentence Sentence,currentPart string, nbA int) bool {
//	sentence.addPart(currentPart)
//	remaining := str[len(sentence.String())+1:]
//	remainingCount := ACount(remaining)
//
//	return remainingCount == sentence.nbEmptyPart()*nbA
//}
//
//func getProblemInfo(s string) ProblemInfo {
//	aCount := ACount(s)
//	return ProblemInfo{aCount, aCount / 3}
//}
//
//type ProblemInfo struct {
//	TotalA, MaxAPerPart int
//}