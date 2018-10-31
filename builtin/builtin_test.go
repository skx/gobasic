// builtin_test.go - Test-cases (trivial) for our registration-glue.

package builtin

import "testing"

//
// Noddy test that we can set/get a value.
//
func TestNoddy(t *testing.T) {

	// Create the holder
	b := New()

	// Register a function
	b.Register("steve", 1, nil)

	// Retrieve it.
	n, _ := b.Get("steve")
	if n != 1 {
		t.Errorf("Invalid argument count for dummy entry")
	}

	//
	// Now retrieve a missing built-in
	//
	foo, bar := b.Get("missing")
	if foo != 0 {
		t.Errorf("We found something unexpected on a missing entry!")
	}
	if bar != nil {
		t.Errorf("We found something unexpected on a missing entry!")
	}
}
