// builtin-support.go - Helpers for registering "built-in" functions

package eval

import (
	"sync"

	"github.com/skx/gobasic/token"
)

// BuiltinSig is the signature of a builtin-function.
//
// Each built-in will receive an array of tokens, and will return a
// result/error which will be made available to the BASIC caller.
type BuiltinSig func(env Variables, args []token.Token) (int, error)

// Builtins holds our state.
type Builtins struct {
	// lock holds a mutex to prevent corruption.
	lock sync.Mutex

	// arg_registry holds the number of arguments the given
	// name requires.
	arg_registry map[string]int

	// fn_registry holds a reference to the golang function which
	// implements the builtin.
	fn_registry map[string]BuiltinSig
}

// NewBuiltins returns a new helper/holder for builtin functions.
func NewBuiltins() *Builtins {
	t := &Builtins{}
	t.arg_registry = make(map[string]int)
	t.fn_registry = make(map[string]BuiltinSig)

	return t
}

// Register records a built-in function.
// The three arguments are:
//  NAME  - The thing that the BASIC program will call
//  nARGS - The number of arguments (tokens) it requires.
//  FT    - The function which provides the implementation.
func (b *Builtins) Register(name string, nArgs int, ft BuiltinSig) {
	b.lock.Lock()
	defer b.lock.Unlock()

	// Register our arg-count and function
	b.arg_registry[name] = nArgs
	b.fn_registry[name] = ft
}

// Exists tests if the given name exists as a built-in function.
func (b *Builtins) Exists(name string) bool {
	b.lock.Lock()
	defer b.lock.Unlock()

	// Does it exist?
	return (b.fn_registry[name] != nil)
}

// Get the values associated with the given built-in.
func (b *Builtins) Get(name string) (int, BuiltinSig) {
	b.lock.Lock()
	defer b.lock.Unlock()

	return b.arg_registry[name], b.fn_registry[name]
}
