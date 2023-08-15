package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Add(val string) {

	newNode := &Node{
		data: val,
		next: nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	return
}

func (l *LinkedList) Print() {
	currentNode := l.head

	for currentNode != nil {

		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}

}

func main() {

	Ll := LinkedList{}
	Ll.Add("3")
	Ll.Add("4")
	Ll.Add("5")
	Ll.Print()

}
