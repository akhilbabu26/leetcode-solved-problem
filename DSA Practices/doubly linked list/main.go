package main

import "fmt"

type Node struct{
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct{
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) InsertHead(data int){
	newNode := &Node{data:data}

	if dll.head == nil{
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

func (dll *DoublyLinkedList) InsertEnd(data int){
	newNode := &Node{data: data}

	if dll.head == nil{
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}

func (dll *DoublyLinkedList) PrintForward(){
	temp := dll.head

	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (dll *DoublyLinkedList) PrintBackwords(){
	temp := dll.tail

	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.prev
	}
}

func main(){
	dll := DoublyLinkedList{}

	dll.InsertHead(1)
	dll.InsertHead(3)
	dll.InsertHead(9)
	dll.InsertHead(12)
	dll.InsertHead(15)

	dll.PrintForward()
	dll.PrintBackwords()
}