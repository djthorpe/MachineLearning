// Usage:
//  go get -u gonum.org/v1/gonum/...
//  go run matrix_03.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gonum.org/v1/gonum/mat"
)

///////////////////////////////////////////////////////////////////////////////
// The core packages of the gonum suite are written in pure Go with
// some assembly. Installation is done using go get:
//
//   go get -u gonum.org/v1/gonum/...
//

// RunMain runs the main program. Demonstrates adding, multiplying, powers,
// applying function to elements
func RunMain() int {
	// Create a new matrix a.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	fmt.Printf("a = %v\n\n", mat.Formatted(a, mat.Prefix("    ")))

	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
