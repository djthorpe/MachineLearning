// Usage:
//  go get -u gonum.org/v1/gonum/...
//  go run matrix.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

///////////////////////////////////////////////////////////////////////////////
// The core packages of the gonum suite are written in pure Go with
// some assembly. Installation is done using go get:
//
//   go get -u gonum.org/v1/gonum/...
//

// RunMain runs the main program
func RunMain() int {

	// Create a flat representation of our matrix.
	components := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix (the first argument is the number of
	// rows and the second argument is the number of columns).
	a := mat.NewDense(2, 2, components)

	// As a sanity check, output the matrix to standard out.
	fmt.Printf(" %v\n", mat.Formatted(a, mat.Prefix(" ")))

	// Get a single value from the matrix.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)

	// Get the values in a specific column.
	col := mat.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	// Get the values in a kspecific row.
	row := mat.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	// Modify a single element.
	a.Set(0, 1, 11.2)

	// Modify an entire row.
	a.SetRow(0, []float64{14.3, -4.2})

	// Modify an entire column.
	a.SetCol(0, []float64{1.7, -0.3})

	// As a sanity check, output the matrix to standard out.
	fmt.Printf(" %v\n", mat.Formatted(a, mat.Prefix(" ")))

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
