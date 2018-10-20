// stack.go holds a simple stack which can hold ints.
//
// (This is used solely for the GOSUB/RETURN handler.)

package eval

import (
	"errors"
	"sync"
)

// Stack holds the stack-data, protected by a mutex
type Stack struct {
	lock sync.Mutex
	s    []int
}

// NewStack returns a new stack (for holding integers)
func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]int, 0)}
}

// Push adds a new item to our stack.
func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

// Pop returns an item from our stack.
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

// Empty returns `true` if our stack is empty.
func (s *Stack) Empty() bool {

	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	return (l == 0)
}
