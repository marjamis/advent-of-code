package structures

// Stack is a stack type
type Stack struct {
	value    interface{}
	previous *Stack
	next     *Stack
}

func (s *Stack) findLastNode() *Stack {
	currentNode := s

	for true {
		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}

	return currentNode
}

// Push will push a value onto the top of a given stack
func (s *Stack) Push(value interface{}) {
	if s.value == nil {
		s.value = value
		return
	}
	currentNode := s.findLastNode()

	currentNode.next = &Stack{
		value:    value,
		previous: currentNode,
		next:     nil,
	}
}

// Pop will take the top value off the given stack and return it
func (s *Stack) Pop() (value interface{}) {
	currentNode := s.findLastNode()

	value = currentNode.value

	if currentNode.previous != nil {
		currentNode.previous.next = nil
	} else {
		s.value = nil
	}

	return
}

// CreateStack creates an new empty Stack
func CreateStack() *Stack {
	return &Stack{}
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return (s.value == nil) && (s.previous == nil) && (s.next == nil)
}
