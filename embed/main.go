// The embed script demonstrates embedding the BASIC interpreter
// into a custom program of your own - along with extending the
// BASIC interpreter to add your own custom functions.
//
// This example demonstrates several things:
//
// Setting a variable from golang which will be visible to BASIC.
//
// Defining custom functions (CIRCLE, DOT, PEEK, POKE, SAVE).
//
// Retrieving the contents of BASIC values back to golang.
//
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// img holds a canvas.
//
// The BASIC program embedded in this program will draw upon an image
// this is the actual image they draw upon.
//
var img *image.RGBA

// peekFunction is the golang implementation of the PEEK primitive,
// which is made available to BASIC.
// We just log that we've been invoked here.
func peekFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {
	fmt.Printf("PEEK called with %v\n", args[0])
	return &object.NumberObject{Value: 0.0}, nil
}

// pokeFunction is the golang implementation of the PEEK primitive,
// which is made available to BASIC.
// We just log that we've been invoked here, along with the (three) args.
func pokeFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {
	fmt.Printf("POKE called.\n")
	for i, e := range args {
		fmt.Printf(" Arg %d -> %v\n", i, e)
	}
	return &object.NumberObject{Value: 0.0}, nil
}

// circleFunction allows drawing a circle upon our image.
func circleFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {

	//
	// Get the args X, Y, & radius
	//
	xx, _ := eval.TokenToFloat(env, args[0])
	// args1 is "COMMA"
	yy, _ := eval.TokenToFloat(env, args[2])
	// args[3] is COMMA
	rr, _ := eval.TokenToFloat(env, args[4])

	//
	// They need to be ints.
	//
	x0 := int(xx)
	y0 := int(yy)
	r := int(rr)

	// If we have no image, create it.
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, 600, 400))
		black := color.RGBA{0, 0, 0, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	}

	// Create the colour
	c := color.RGBA{0, 255, 0, 255}

	// Now circle-magic happens.
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (int(r) * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}

	// All done.
	return &object.NumberObject{Value: 0.0}, nil
}

// plotFunction is the golang implementation of the PLOT primitive.
//
// It is invoked with three arguments (NUMBER COMMA NUMBER) and sets
// the corresponding pixel in our canvas to be Red.
func plotFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {

	//
	// Get the args: X, Y
	//
	x, _ := eval.TokenToFloat(env, args[0])
	// args1 is "COMMA"
	y, _ := eval.TokenToFloat(env, args[2])

	// If we have no image, create it.
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, 600, 400))
		black := color.RGBA{0, 0, 0, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	}

	// Draw the dot
	img.Set(int(x), int(y), color.RGBA{255, 0, 0, 255})

	return &object.NumberObject{Value: 0.0}, nil
}

// saveFunction is the golang implementation of the SAVE primitive,
// which is made available to BASIC.
// We save the image-canvas to the file `out.png`.
func saveFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {

	// If we have no image, create it.
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, 600, 400))
		black := color.RGBA{0, 0, 0, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	}

	// Now write out the image.
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	return &object.NumberObject{Value: 0.0}, nil
}

func main() {

	//
	// This is the program we're going to execute
	//
	prog := `
 10 PRINT "HELLO, I AM EMBEDDED BASIC IN YOUR GOLANG!\n"
 20 LET S = S + PI
 30 POKE 23659 , 0
 40 PEEK 30
 50 PRINT "\n" "I'M NOW CREATING AN IMAGE!!!!\n"
 60 REM
 70 REM Draw 100 random pixels
 80 REM
 90 FOR I = 1 TO 100
100  LET x = RND 600
110  LET y = RND 400
120  PLOT x, y
130 NEXT I
140 REM
150 REM Draw a random number of circles
160 REM
170 LET R = RND 30
180 IF R < 2 THEN LET R=2
190 PRINT "\tWe will draw", R, "random circles upon the image\n"
200 FOR I = 1 TO R
210  LET x = RND 600
220  LET y = RND 400
230  LET r = RND 100
240  CIRCLE x, y, r
250 NEXT I
260 SAVE
270 PRINT "\tOPEN 'out.png' TO VIEW YOUR IMAGE!\n"
`

	//
	// Load the program
	//
	t := tokenizer.New(prog)

	//
	// Create an interpreter
	//
	e := eval.New(t)

	//
	// Register some  functions.
	//
	e.RegisterBuiltin("CIRCLE", 5, circleFunction)
	e.RegisterBuiltin("DOT", 3, plotFunction)
	e.RegisterBuiltin("PEEK", 1, peekFunction)
	e.RegisterBuiltin("PLOT", 3, plotFunction)
	e.RegisterBuiltin("POKE", 3, pokeFunction)
	e.RegisterBuiltin("SAVE", 0, saveFunction)

	//
	// Set an initial value to the variable "S".
	//
	e.SetVariable("S", &object.NumberObject{Value: 3})

	//
	// Run the code.
	//
	err := e.Run()
	if err != nil {
		fmt.Printf("Error running program: %s\n", err.Error())
	}

	fmt.Printf("\n\n")

	//
	// The value of the variable is now different
	//
	result := e.GetVariable("S")
	if result.Type() == object.NUMBER {
		fmt.Printf("After calling BASIC 'S' is a number '%f'\n",
			result.(*object.NumberObject).Value)
	}
	if result.Type() == object.STRING {
		fmt.Printf("After calling BASIC 'S' is a string '%s'\n",
			result.(*object.StringObject).Value)
	}
}
