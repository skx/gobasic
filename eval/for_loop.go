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

// ForLoop is the structure used to record a for-loop
type ForLoop struct {
	// variable the loop refers to
	id string

	// offset of the start of the loop-body
	offset int

	// start is the initial value of the variable at the start of the loop
	start int

	// end is the terminating value of the variable
	end int

	// increment is how much to step by
	step int

	// is the loop finnished?
	finished bool
}

// FORS holds all known/current for-loops.
var FORS map[string]ForLoop

// AddForLoop stores a reference to a for-loop in our map
func AddForLoop(x ForLoop) {
	if FORS == nil {
		FORS = make(map[string]ForLoop)
	}

	FORS[x.id] = x
}

// GetForLoop returns a reference to a for-loop.
func GetForLoop(id string) ForLoop {
	return (FORS[id])
}

// RemoveForLoop removes a reference to a for-loop.
func RemoveForLoop(id string) {
	delete(FORS, id)
}
