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
	// Get a reference to the Tail node
	nodePopped := *s.Tail
	ptr := s.Head

	// Loop through the linkedList to find the preceding pointer reference
	for {
		node := *ptr
		if node.Pointer == s.Tail {
			s.Tail = &node
			break
		}
		ptr = node.Pointer
	}

	// Reduce the counter by 1
	s.Count--
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
func (s *Stack) Remove(value string) (returnNode Node) {

	// Start at the head
	ptr := s.Head
	node := *ptr
	var previousNodePtr *Node

	for {
		if node.Data == value {

			if ptr == s.Head {
				// Am I removing the Head
				s.Head = node.Pointer
			} else if ptr == s.Tail {
				// Am I removing the Tail
				s.Tail = previousNodePtr
			} else {
				// Am I removing anything else
				if previousNodePtr != nil {
					previousNodePtr.Pointer = node.Pointer
				}
			}

			// Remove any pointer reference
			returnNode = node
			returnNode.Pointer = nil

			s.Count--
			return
		}

		// Break out of the loop if don't find the string
		if node.Pointer == nil {
			return
		}

		// Assign the previous node to the current node
		previousNodePtr = ptr

		// Get the next node in the list
		if node.Pointer != nil {
			ptr = node.Pointer
			node = *ptr
		}
	}
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
		} else {
			node = *node.Pointer
		}
	}
}
