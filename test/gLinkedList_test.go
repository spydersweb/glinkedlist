package main

import (
	ll "glinkedlist"
	"testing"
)


func TestLinkedListHeadandTail(t *testing.T) {

	l := ll.Stack{}

	seedData := seedLinkedList(&l, []string{})

	// Test the Head matches the first seedData item
	if l.Head.Data != seedData[0] {
		t.Errorf("Error in linked list Head data, expected %s, got %s",
			seedData[0],
			l.Head.Data,
		)
	}

	// Test the Tail matches the last seedData item
	if l.Tail.Data != seedData[len(seedData) - 1] {
		t.Errorf("Error in linked list Head data, expected %s, got %s",
			seedData[0],
			l.Tail.Data,
		)
	}

}

func TestIterationDataMatchesSeedData(t *testing.T) {
	l := ll.Stack{}

	seedData := seedLinkedList(&l, []string{})

	ptr := l.Head
	counter := 0
	for {
		node := *ptr
		if node.Pointer == nil {
			break
		}

		if node.Data != seedData[counter] {
			t.Errorf("Iterating through list found mismatched linkedList and seed value data.  Exp: %s, Got: %s",
				seedData[counter],
				node.Data,
			)
		}

		ptr = node.Pointer
		counter++
	}
}

func TestListCountMatchesSeed(t *testing.T) {
	l := ll.Stack{}

	seedData := seedLinkedList(&l, []string{})

	// Test that the seedData count matches the linkedList count
	if l.Count != len(seedData) {
		t.Errorf("List count does not match seedData count. Exp: %d, Got: %d",
			len(seedData),
			l.Count,
		)
	}
}

func TestListPopping(t *testing.T){
	l := ll.Stack{}

	seedData := seedLinkedList(&l, []string{})

	lastIndex := len(seedData) - 1

	node := l.Pop()
	if node.Data != seedData[lastIndex] {
		t.Errorf("Popped node Data doesn't match last item of seed data. Exp: %s, Got: %s",
			seedData[lastIndex],
			node.Data,
		)
	}

	// Check the Count of the linkedList
	seedDataAfterPopping := seedData[0:2]
	lastIndex = len(seedDataAfterPopping) - 1

	if len(seedDataAfterPopping) != l.Count {
		t.Errorf("LinkedList Count does not match popped list length. Exp: %d, Got: %d",
			len(seedDataAfterPopping),
			l.Count,
		)
	}

	// Check the Tail Data matches the new slice data
	if l.Tail.Data != seedDataAfterPopping[lastIndex] {
		t.Errorf("linkedList Tail should have been updated to the new value. Exp: %s, Got: %s",
			seedDataAfterPopping[lastIndex],
			l.Tail.Data,
		)
	}

}

func TestRemoveListItem(t *testing.T) {
	l := ll.Stack{}

	removeValue := "Barbara"
	seedData := seedLinkedList(&l, []string{"John", "Jack", "Barbara", "Tony", "Harry"})

	var removedData = []string{}
	removedData = append(seedData[0:2], seedData[3:]...)

	removedNode := l.Remove(removeValue)

	testRemovedNode := ll.Node{removeValue, nil}

	// Test the removed Node has a nil pointer reference
	if removedNode.Pointer != nil {
		t.Errorf("Removed node pointer ref should be nil but got %p", removedNode.Pointer)
	}

	// Test the linkedList count now matches the removedData count
	if len(removedData) != l.Count {
		t.Errorf("New list count doesn't match the removed list count.  Exp: %d, Got: %d",
			len(removedData),
			l.Count,
		)
	}

	// Test the removed Node
	if removedNode != testRemovedNode {
		t.Errorf("Removed Node doesn't match dummy Removed Node. Exp: %v, Got: %v",
			testRemovedNode,
			removedNode,
		)
	}
}

func TestRemoveHeadUsingRemove(t *testing.T) {
	l := ll.Stack{}

	removeValue := "John"
	_ = seedLinkedList(&l, []string{"John", "Jack", "Barbara", "Tony", "Harry"})
	expectedNewHead := l.Head.Pointer

	l.Remove(removeValue)

	if l.Head != expectedNewHead {
		t.Errorf("Head Pointer doesn't match expected Head Pointer. Exp: %p, Got: %p",
			expectedNewHead,
			l.Head,
		)
	}
}

func seedLinkedList(ll *ll.Stack, values []string) (seed []string){

	seed = []string{"harry", "sharon", "graham"}
	if len(values) != 0 {
		seed = values
	}

	for _,n := range seed {
		ll.Push(n)
	}

	return
}