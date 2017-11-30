// Usage:
//  go get -u gonum.org/v1/gonum/stat
//  go run chisquare_02.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	// Define the observed frequencies
	observed := []float64{
		260.0, // This number is the number of observed with no regular exercise.
		135.0, // This number is the number of observed with sporatic exercise.
		105.0, // This number is the number of observed with regular exercise.
	}
	// Define the total observed.
	totalObserved := 500.0

	// Calculate the expected frequencies (again assuming the null Hypothesis).
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Output the test statistic to standard out.
	fmt.Println("Observed", observed)
	fmt.Println("Expected", expected)
	fmt.Println("chiSquare", chiSquare)

	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
