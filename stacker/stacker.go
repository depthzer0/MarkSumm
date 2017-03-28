package main

import(
	"fmt"
	"github.com/depthzer0/MarkSumm/stacker/stack"
)

func main()  {
	
	var newStack stack.Stack
	newStack.Push(123)
	newStack.Push("asd")

	fmt.Println(newStack)
}