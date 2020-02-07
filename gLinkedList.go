package glinkedlist

import "fmt"

//Stack tracks the linked list head and tail
type Stack struct {
	Head *Node
	Tail *Node
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
	n := Node{value, nil}

	// On first initialisation the head will always be nil
	if s.Head == nil {
		s.Head = &n
	} else {
		// update the head of the stack to the newly generated node pointer
		n.Pointer = s.Head
		s.Head = &n
	}
}

// Remove takes a string value and removes the node with that value
// re-assigning the node pointers before and after the node being removed
func (s *Stack) Remove(value string) Node {

	testNode := *s.Head

	var returnNode Node = Node{}
	var prevNode Node = Node{}

	for {
		nextNode := *testNode.Pointer

		fmt.Println("Previous node:", prevNode)
		// break the loop if we have hit the tail
		if testNode.Pointer == nil {
			break
		}

		// first iteration, basically the head of the stack
		// then we are removing the head of the stack only, popping
		if testNode.Data == value && testNode == *s.Head {
			returnNode = s.Pop()
			break

		} else if testNode.Data == value {

			// Reassign prevNode's pointer to the testNode.Pointer element
			// and return testNode
			if nextNode.Pointer != nil {
				*prevNode.Pointer = *nextNode.Pointer
			}

			returnNode = testNode
			break
		}

		// Assign the prevNode to the currently tested node
		prevNode = testNode

		// get the next list item
		testNode = *testNode.Pointer

	}

	return returnNode
}

//Iterate loops through the stack of linked nodes
func (s *Stack) Iterate(f func(n Node)) {

	node := *s.Head

	for {
		if f != nil {
			f(node)
		}

		if node.Pointer == nil {
			break
		}
		node = *node.Pointer
	}
}
