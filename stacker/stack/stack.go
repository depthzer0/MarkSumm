package stack



type Stack []interface{}

func (stack *Stack) Push (x interface{}){
	*stack = append(*stack, x)
}

func (stack *Stack) Pop () interface{}{
	//*stack = append(*stack, x)
	return 1
}