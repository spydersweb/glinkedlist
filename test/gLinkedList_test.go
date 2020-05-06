package main

import (
	ll "glinkedlist"
	"testing"
)

var (
	seedData = []string{"harry", "sharon", "graham"}
)

func TestLinkedListHeadandTail(t *testing.T) {

	l := ll.Stack{}

	seedLinkedList(&l, seedData)

	if l.Head.Data != seedData[0] {
		t.Errorf("Error in linked list Head data, expected %s, got %s", seedData[0], l.Head.Data)
	}

	if l.Tail.Data != seedData[len(seedData) - 1] {
		t.Errorf("Error in linked list Head data, expected %s, got %s", seedData[0], l.Tail.Data)
	}

}

func TestIterationDataMatchesSeedData(t *testing.T) {
	l := ll.Stack{}

	seedLinkedList(&l, seedData)

	ptr := l.Head
	counter := 0
	for {
		node := *ptr
		if node.Pointer == nil {
			break
		}

		if node.Data != seedData[counter] {
			t.Errorf("Iterating through list found mismatched linkedList and seed value data.  Exp: %s, Got: %s", seedData[counter], node.Data)
		}

		ptr = node.Pointer
		counter++
	}
}

func TestListCountMatchesSeed(t *testing.T) {
	l := ll.Stack{}

	seedLinkedList(&l, seedData)

	if l.Count != len(seedData) {
		t.Errorf("List count does not match seedData count. Exp: %d, Got: %d", len(seedData), l.Count)
	}
}

func TestListPopping(t *testing.T){
	l := ll.Stack{}

	seedLinkedList(&l, seedData)

	lastIndex := len(seedData) - 1

	node := l.Pop()
	if node.Data != seedData[lastIndex] {
		t.Errorf("Popped node Data doesn't match last item of seed data. Exp: %s, Got: %s", seedData[lastIndex], node.Data)
	}

	// Check the Count of the linkedList
	seedDataAfterPopping := seedData[0:2]
	lastIndex = len(seedDataAfterPopping) - 1

	if len(seedDataAfterPopping) != l.Count {
		t.Errorf("LinkedList Count does not match popped list length. Exp: %d, Got: %d", len(seedDataAfterPopping), l.Count)
	}

	// Check the Tail Data matches the new slice data
	if l.Tail.Data != seedDataAfterPopping[lastIndex] {
		t.Errorf("linkedList Tail should have been updated to the new value. Exp: %s, Got: %s",seedDataAfterPopping[lastIndex], l.Tail.Data )
	}

}

func seedLinkedList(ll *ll.Stack, values []string) {
	for _,n := range values {
		ll.Push(n)
	}
}