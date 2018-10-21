// The goserver utility is a simple HTTP server which allows you to
// interactively run BASIC scripts via your browser.
//
// The goservers purpose is to allow users to experiment with graphics,
// which it allows by the addition of several custom functions to the
// BASIC environment.
//
// The additions make it easy to change the colour of the pixels, draw
// points, circles, and view a rendered image containing the output.
//
// Graphing SIN and similar functions becomes very simple and natural.
package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// img holds the canvas we draw into.
var img *image.RGBA

// col holds our currently selected colour
var col color.RGBA

// Setup default color (black)
func init() {
	col = color.RGBA{0, 0, 0, 255}
}

// dotFunction is the golang implementation of the DOT primitive.
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
		c := color.RGBA{255, 255, 255, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	}

	// Draw the dot
	img.Set(int(x), int(y), col)

	return &object.NumberObject{Value: 0.0}, nil
}

// saveFunction is the golang implementation of the SAVE primitive,
// which is made available to BASIC.
// We save the image-canvas to the file `out.png`.
func saveFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {

	// If we have no image, create it.
	if img == nil {
		img = image.NewRGBA(image.Rect(0, 0, 600, 400))
		c := color.RGBA{255, 255, 255, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	}

	// Generate a temporary filename
	tmpfile, _ := ioutil.TempFile("", "goserver")

	// Now write out the image.
	f, _ := os.OpenFile(tmpfile.Name(), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)

	// And save the temporary filename in a variable
	env.SetVariable("file.name", &object.StringObject{Value: tmpfile.Name()})

	// Finally we can nuke the image
	img = nil

	return &object.NumberObject{Value: 0.0}, nil
}

// colorFunction allows drawing a circle upon our image.
func colorFunction(env eval.Interpreter, args []token.Token) (object.Object, error) {

	//
	// Get the args R, G, B values
	//
	r, _ := eval.TokenToFloat(env, args[0])
	// args1 is "COMMA"
	g, _ := eval.TokenToFloat(env, args[2])
	// args[3] is COMMA
	b, _ := eval.TokenToFloat(env, args[4])

	col = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
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
		c := color.RGBA{255, 255, 255, 255}
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	}

	// Now circle-magic happens.
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (int(r) * 2)

	for x > y {
		img.Set(x0+x, y0+y, col)
		img.Set(x0+y, y0+x, col)
		img.Set(x0-y, y0+x, col)
		img.Set(x0-x, y0+y, col)
		img.Set(x0-x, y0-y, col)
		img.Set(x0-y, y0-x, col)
		img.Set(x0+y, y0-x, col)
		img.Set(x0+x, y0-y, col)

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

//
// Runs the script the user submitted.
//
// Returns the base64-encoded version of the output image.
//
// Racy.
//
func runScript(code string) string {
	code += "\n9999999999 SAVE\n"
	t := tokenizer.New(code)

	e := eval.New(t)
	e.RegisterBuiltin("CIRCLE", 5, circleFunction)
	e.RegisterBuiltin("COLOR", 5, colorFunction)
	e.RegisterBuiltin("COLOUR", 5, colorFunction)
	e.RegisterBuiltin("PLOT", 3, plotFunction)
	e.RegisterBuiltin("SAVE", 0, saveFunction)

	err := e.Run()
	if err != nil {
		fmt.Printf("Error running code: %s\n", err.Error())
	}

	// Get the name of the file the SAVE function wrote to
	pathObj := e.GetVariable("file.name")
	path := pathObj.(*object.StringObject).Value

	// Read the file
	b, _ := ioutil.ReadFile(path)

	// BASE64 encode it.
	encoded := base64.StdEncoding.EncodeToString(b)

	// remove the temporary file
	os.Remove(path)

	// And return the value.
	return encoded
}

//
// Called via a HTTP-request.
//
// If GET serve `index.html`.
//
// If POST serve a PNG created by executing the user-submitted code.
//
func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		code := r.FormValue("code")
		out := runScript(code)
		fmt.Fprintf(w, "%s\n", out)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

//
// Entry-point.
//
func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Please open http://localhost:8080/ ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
