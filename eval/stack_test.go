// stack_test.go - Simple test-cases for our stack

package eval

import "testing"

// TestEmpty: Test that the Empty() function works as expected.
func TestEmpty(t *testing.T) {
	s := NewStack()

	if !s.Empty() {
		t.Errorf("New stack is not empty!")
	}

	s.Push(33)

	if s.Empty() {
		t.Errorf("Despite storing a value the stack is still empty!")
	}
}

// TestEmptyPop: Test that pop'ing from an empty stack fails.
func TestEmptyPop(t *testing.T) {
	s := NewStack()

	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expected an error popping from an empty stack!")
	}
}

// TestPushPop: Test that we can store/retrieve as we expect.
func TestPushPop(t *testing.T) {
	s := NewStack()

	s.Push(33)

	out, err := s.Pop()
	if err != nil {
		t.Errorf("We shouldn't get an error popping from our stack")
	}
	if out != 33 {
		t.Errorf("We retrieved a value from our stack, but it was wrong")
	}
}
