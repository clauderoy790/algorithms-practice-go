package linked_list

import (
	"fmt"
	"strconv"
)

func LinkedListEx() {
	lst := LinkedList{}
	lst.AddMany([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(lst.String())
	//reverseLinkedList2(&lst)
	reverseLinkedList(&lst)
	fmt.Println("reversed list: ", lst.String())
}

func reverseLinkedLis2(lst *LinkedList) {
	var slice []*Node

	if lst.head == nil || lst.head.next == nil {
		return
	}

	//Append all nodes in slice
	current := lst.head
	for ; current != nil; {
		slice = append(slice, current)
		current = current.next
	}

	//Remove current head next node
	if lst.head.next != nil {
		lst.head.next = nil
	}

	//Change head
	lst.head = slice[len(slice)-1]

	//Loop in reverse order to set next node
	for i := len(slice) - 1; i > 0; i-- {
		slice[i].next = slice[i-1]
	}
}

func reverseLinkedList(lst *LinkedList) {
	var prev, temp *Node
	current := lst.head

	for ;current != nil; {
		temp = current.next
		current.next = prev
		prev = current
		current = temp
	}

	lst.head = prev

}

type LinkedList struct {
	head *Node
}

func (lst *LinkedList) AddMany(values []int) {
	for i := 0; i < len(values); i++ {
		lst.Add(values[i])
	}
}

func (lst *LinkedList) Add(val int) {
	if lst.head == nil {
		lst.head = &Node{val, nil}
	} else {
		current := lst.head
		for ; lst.head != nil && current.next != nil; {
			current = current.next
		}
		if lst.head != nil && current.next == nil {
			newNode := Node{val, nil}
			current.next = &newNode
		}
	}
}

func (lst *LinkedList) String() string {
	str := ""
	current := lst.head
	for ; current != nil; {
		str += strconv.Itoa(current.val)
		if current.next != nil {
			str += ", "
		} else {
			str += "."
		}
		current = current.next
	}

	return str
}

type Node struct {
	val  int
	next *Node
}
