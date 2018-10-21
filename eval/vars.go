// vars.go - Define an interface for getting/setting variables by name.
//
// NOTE: Names are assumed to be globally unique; there is no notion of scope.

package eval

import (
	"sync"

	"github.com/skx/gobasic/object"
)

// Variables holds our state
type Variables struct {
	// lock ensures we're thread-safe (ha!)
	lock sync.Mutex

	// data stores our data
	data map[string]object.Object
}

// NewVars handles a new variable-holder.
func NewVars() *Variables {
	return &Variables{lock: sync.Mutex{}, data: make(map[string]object.Object)}
}

// Set stores the given value against the specified name.
func (v *Variables) Set(name string, val object.Object) {
	v.lock.Lock()
	defer v.lock.Unlock()

	v.data[name] = val
}

// Get returns the value stored against the specified name.
func (v *Variables) Get(name string) object.Object {
	v.lock.Lock()
	defer v.lock.Unlock()
	return (v.data[name])
}
