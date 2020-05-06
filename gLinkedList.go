package glinkedlist

import "fmt"

//Stack tracks the linked list head and tail
type Stack struct {
	Head *Node
	Tail *Node
	Count int
	Debug bool
}

// Node is a linked list item
type Node struct {
	Data    string
	Pointer *Node
}

// Pop takes the top node from the linked list and returns
func (s *Stack) Pop() Node {
	// Get a reference to the head node
	nodePopped := *s.Head

	// Set the head of the stack to its next pointer
	*s.Head = *nodePopped.Pointer
	return nodePopped
}

// Push generates a new node for the linked list and assigns a pointer
// reference if necessary
func (s *Stack) Push(value string) {

	// Create a new node
	newNode := Node{value, nil}

	if s.Head == nil {
		// First iteration set the head to the newNode
		s.Head = &newNode
	} else if s.Count == 1 {
		// Second iteration we set the head nodes pointer to the newNode
		s.Head.Pointer = &newNode
	} else {
		// Subsequent iterations we set the tail nodes pointer to the newNode
		s.Tail.Pointer = &newNode
	}

	// Set the tail to the newNode
	s.Tail = &newNode

	// Increment the counter
	s.Count++

	// Output debug if needed
	if s.Debug {
		fmt.Printf("Pushing:\nPointer to n: %p\t%v\n", &newNode, newNode)
		fmt.Printf("Current Stack: %v\n", s)
	}
}

// Remove takes a string value and removes the node with that value
// re-assigning the node pointers before and after the node being removed
func (s *Stack) Remove(value string) Node {

	currentNode := *s.Head

	var returnNode Node

	for {
		prevNode := currentNode
		nextNode := *currentNode.Pointer

		// break the loop if we have hit the tail
		if currentNode.Pointer == nil {
			break
		}

		// first iteration, basically the head of the stack
		// then we are removing the head of the stack only, popping
		if currentNode.Data == value && currentNode == *s.Head {
			returnNode = s.Pop()
			break

		} else if currentNode.Data == value {

			// Reassign prevNode's pointer to the testNode.Pointer element
			// and return testNode
			if nextNode.Pointer != nil {
				prevNode.Pointer = &nextNode
			} else {
				prevNode.Pointer = nil
			}

			returnNode = currentNode
			break
		}

		// get the next Node in the stack
		currentNode = *currentNode.Pointer

	}

	return returnNode
}

//Iterate loops through the stack of linked nodes
func (s *Stack) Iterate(f func(n *Node)) {

	node := *s.Head

	for {
		if f != nil {
			f(&node)
		}

		if node.Pointer == nil {
			break
		}
		node = *node.Pointer
	}
}
