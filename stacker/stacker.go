package main

import(
	"fmt"
	"github.com/depthzer0/MarkSumm/stacker/stack"
)

func main()  {
	
	var newStack stack.Stack
	newStack.Push(123)
	newStack.Push("asd")
	newStack.Push(-888)
	newStack.Push([]string{"asf", "2r23r23r", "asdqc1122121"})

	fmt.Println(newStack)
}