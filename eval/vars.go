// vars.go - Hold variables

package eval

import "sync"

// Variables holds our state
type Variables struct {
	// lock ensures we're thread-safe (ha!)
	lock sync.Mutex

	// data stores our data
	data map[string]interface{}
}

// NewVars handles a new variable-holder.
func NewVars() *Variables {
	return &Variables{lock: sync.Mutex{}, data: make(map[string]interface{})}
}

// Set stores the given value against the specified name.
func (v *Variables) Set(name string, val interface{}) {
	v.lock.Lock()
	defer v.lock.Unlock()

	v.data[name] = val
}

// Get returns the value stored against the specified name.
func (v *Variables) Get(name string) interface{} {
	v.lock.Lock()
	defer v.lock.Unlock()
	return (v.data[name])
}
