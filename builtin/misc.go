// The builtin package provides the ability to register our built-in functions.
//
// misc.go implements some misc. primitives.

package builtin

import (
	"bufio"
	"fmt"
	"os"

	"github.com/skx/gobasic/object"
)

// DUMP just displays the only argument it received.
func DUMP(env Environment, args []object.Object) object.Object {
	var out *bufio.Writer
	if env == nil {
		out = bufio.NewWriter(os.Stdout)
	} else {
		out = env.StdOutput()
	}

	// Get the (float) argument.
	if args[0].Type() == object.NUMBER {
		i := args[0].(*object.NumberObject).Value
		out.WriteString(fmt.Sprintf("NUMBER: %f\n", i))
	}
	if args[0].Type() == object.STRING {
		s := args[0].(*object.StringObject).Value
		out.WriteString(fmt.Sprintf("STRING: %s\n", s))
	}
	if args[0].Type() == object.ERROR {
		s := args[0].(*object.ErrorObject).Value
		out.WriteString(fmt.Sprintf("Error: %s\n", s))
	}
	out.Flush()

	// Otherwise return as-is.
	return &object.NumberObject{Value: 0}
}

// PRINT handles displaying strings, integers, and errors.
func PRINT(env Environment, args []object.Object) object.Object {
	var out *bufio.Writer
	if env == nil {
		out = bufio.NewWriter(os.Stdout)
	} else {
		out = env.StdOutput()
	}
	for _, ent := range args {
		switch ent.Type() {
		case object.NUMBER:
			n := ent.(*object.NumberObject).Value
			if n == float64(int(n)) {
				out.WriteString(fmt.Sprintf("%d", int(n)))
			} else {
				out.WriteString(fmt.Sprintf("%f", n))
			}
		case object.STRING:
			out.WriteString(ent.(*object.StringObject).Value)
		case object.ERROR:
			out.WriteString(ent.(*object.ErrorObject).Value)
		}
	}
	if env != nil {
		out.WriteString(env.LineEnding())
	}
	out.Flush()

	// Return the count of values we printed.
	return &object.NumberObject{Value: float64(len(args))}
}
