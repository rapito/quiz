// Credit for this collection
// goes to https://gist.github.com/bemasher/1777766
package collections

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next *Element
}

func NewStack() *Stack {
	return new(Stack)
}

// Tells whether the stack's empty or not
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
