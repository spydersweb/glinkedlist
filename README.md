# glinkedlist
glinkedlist is a GO implementation of a linked list of strings

## Import

import linkedList "github.com/spydersweb/glinkedlist"

## Initialise

To generate a stack we call the .Stack{} method which initialises the head and tail nodes of the list to nil

stack := linkedList.Stack{}

## Methods

Push(s string) -  Adds a string value to the stack

Pop() - Removes the top of the stack and returns the popped Node

Remove(s string) - Removes the Node with the matched string

Iterate(fn func) - Iterates through the stack and passes each Node.data element to the passed callback if not nil 