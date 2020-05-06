# glinkedlist
glinkedlist is a GO implementation of a singularly linked list of strings

## Installation

    go get github.com/spydersweb/glinkedlist
    
## Import

    import linkedList "github.com/spydersweb/glinkedlist"

## Initialise

To generate a stack we create a variable of type linkedlist.Stack{}.  This sets the Head and Tail to nil.

    stack := linkedList.Stack{}

## Methods

Push(s string) -  Adds a string value to the stack.

    stack.Push(<value_to_add>)

Pop() - Removes the top of the stack and returns the popped Node.

    removedNode := stack.Pop()

Remove(s string) - Removes the Node with the matched string and returns it.

    stack.Remove(<value_to_remove>)

Iterate(fn func) - Iterates through the stack and passes each Node.data element to the passed callback if not nil. 

    stack.Iterate(func(node linkedList.Node){...})