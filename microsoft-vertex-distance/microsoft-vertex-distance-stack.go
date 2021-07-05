package microsoft_vertex_distance

import "fmt"

func ResolveStack() {
	stack := Stack{}

	for _,point := range input {
		dist := distance(point,vertex)
		if stack.len() == index {
			if dist < stack.peek().distance {
				stack.pop()
				stack.push(stackNode{point,dist})
			}
		} else {
			stack.push(stackNode{point,dist})
		}
	}

	node := stack.pop()
	fmt.Println(fmt.Sprintf("the %v closest is : %v",index,node.point))
}

type stackNode struct {
	point Point
	distance float64
}

type Stack struct {
	elements []stackNode
}


func (s *Stack) len() int {
	return len(s.elements)
}

func (s *Stack) peek() stackNode {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) push(node stackNode) {
	s.elements = append(s.elements,node)
}

func (s *Stack) pop() stackNode {
	node := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return node
}

