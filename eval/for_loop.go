// for_loop.go - Handles the state required for for-loops
//
// A for-loop looks like this:
//
//    FOR i=START to END [STEP OFFSET]
//      ..
//    NEXT i
//
// The variable in the FOR-loop is unique.
//

package eval

import "sync"

// ForLoop is the structure used to record a for-loop
type ForLoop struct {
	// variable the loop refers to
	id string

	// offset of the start of the loop-body
	offset int

	// start is the initial value of the variable at the start of the loop
	start float64

	// end is the terminating value of the variable
	end float64

	// increment is how much to step by
	step float64

	// is the loop finnished?
	finished bool
}

// Loops is the structure which holds ForLoop entries
type Loops struct {
	// lock ensures we're thread-safe (ha!)
	lock sync.Mutex

	// data stores our data
	data map[string]ForLoop
}

// NewLoops creates a new for-loop holder
func NewLoops() *Loops {
	return &Loops{lock: sync.Mutex{}, data: make(map[string]ForLoop)}
}

// Add stores a reference to a for-loop in our map
func (l *Loops) Add(x ForLoop) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.data[x.id] = x
}

// Get returns a reference to a for-loop.
func (l *Loops) Get(id string) ForLoop {
	l.lock.Lock()
	defer l.lock.Unlock()

	return (l.data[id])
}

// Remove removes a reference to a for-loop.
func (l *Loops) Remove(id string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	delete(l.data, id)
}

// Empty returns true if we have no open FOR loop-references.
func (l *Loops) Empty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	return (len(l.data) == 0)
}
