package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// Image
var img *image.RGBA

func peekFunction(env eval.Interpreter, args []token.Token) (float64, error) {
	fmt.Printf("PEEK called with %v\n", args[0])
	return 0, nil
}
func pokeFunction(env eval.Interpreter, args []token.Token) (float64, error) {
	fmt.Printf("POKE called.\n")
	for i, e := range args {
		fmt.Printf(" Arg %d -> %v\n", i, e)
	}
	return 0, nil
}

// Draw a DOT at a given X,Y coordinate
func dotFunction(env eval.Interpreter, args []token.Token) (float64, error) {

	x := 0
	y := 0

	// Get the args
	//
	if args[0].Type == token.INT {
		i, err := strconv.ParseFloat(args[0].Literal, 64)
		if err != nil {
			return 0, err
		}

		x = int(i)
	}
	if args[0].Type == token.IDENT {
		// Get.
		x = int(env.GetVariable(args[0].Literal))
	}

	// y
	if args[2].Type == token.INT {
		i, err := strconv.ParseFloat(args[2].Literal, 64)
		if err != nil {
			return 0, err
		}

		y = int(i)
	}
	if args[2].Type == token.IDENT {
		// Get.
		y = int(env.GetVariable(args[2].Literal))
	}

	// If we have no image, create it.
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, 600, 400))
		black := color.RGBA{0, 0, 0, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
	}

	// Draw the dot
	img.Set(x, y, color.RGBA{255, 0, 0, 255})

	return 0, nil
}

// Save our generated image to out.png
func saveFunction(env eval.Interpreter, args []token.Token) (float64, error) {

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	return 0, nil
}

func main() {

	//
	// Ensure we seed a random-number source
	//
	// This is required such that RND() returns suitable
	// values that change.
	//
	rand.Seed(time.Now().UnixNano())

	//
	// This is the program we're going to execute
	//
	prog := `
 10 PRINT "HELLO, I AM EMBEDDED BASIC\n"
 20 LET S = S + PI
 30 POKE 23659 , 0
 40 PEEK 30
 50 PRINT "I'M NOW CREATING AN IMAGE!!!!\n"
 55 REM 640 should be enough for anybody
 60 FOR I = 1 TO 640
 70 let x = RND 600
 80 let y = RND 400
 90 DOT x, y
100 NEXT I
110 SAVE
120 PRINT "OPEN 'out.png' TO VIEW YOUR IMAGE!\n"
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
	e.RegisterBuiltin("PEEK", 1, peekFunction)
	e.RegisterBuiltin("POKE", 3, pokeFunction)
	e.RegisterBuiltin("DOT", 3, dotFunction)
	e.RegisterBuiltin("SAVE", 0, saveFunction)

	//
	// Set an initial value to the variable "S".
	//
	e.SetVariable("S", 3)

	//
	// Run the code.
	//
	err := e.Run()
	if err != nil {
		fmt.Printf("Error running program: %s\n", err.Error())
	}

	//
	// The value of the variable is now different
	//
	fmt.Printf("Output value is %v\n", e.GetVariable("S"))
}
