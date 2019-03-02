// The builtin package provides the ability to register our built-in functions.
//
// misc.go implements some misc. primitives.

package builtin

import (
	"fmt"

	"github.com/andydotxyz/gobasic/object"
)

// DUMP just displays the only argument it received.
func DUMP(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() == object.NUMBER {
		i := args[0].(*object.NumberObject).Value
		fmt.Printf("NUMBER: %f\n", i)
	}
	if args[0].Type() == object.STRING {
		s := args[0].(*object.StringObject).Value
		fmt.Printf("STRING: %s\n", s)
	}
	if args[0].Type() == object.ERROR {
		s := args[0].(*object.ErrorObject).Value
		fmt.Printf("Error: %s\n", s)
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: 0}
}

// PRINT handles displaying strings, integers, and errors.
func PRINT(env interface{}, args []object.Object) object.Object {
	for _, ent := range args {
		switch ent.Type() {
		case object.NUMBER:
			n := ent.(*object.NumberObject).Value
			if n == float64(int(n)) {
				fmt.Printf("%d", int(n))
			} else {
				fmt.Printf("%f", n)
			}
		case object.STRING:
			fmt.Printf("%s", ent.(*object.StringObject).Value)
		case object.ERROR:
			fmt.Printf("%s", ent.(*object.ErrorObject).Value)
		}
	}

	// Return the count of values we printed.
	return &object.NumberObject{Value: float64(len(args))}
}
