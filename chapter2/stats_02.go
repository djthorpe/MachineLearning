// Usage:
//  go run stats_02.go iris.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

///////////////////////////////////////////////////////////////////////////////

func ParseFloat(string_value string) (float64, error) {
	// Check for empty value
	if len(strings.TrimSpace(string_value)) == 0 {
		return 0.0, ErrEmpty
	}
	// Convert to integer
	if float_value, err := strconv.ParseFloat(string_value, 64); err != nil {
		return 0.0, err
	} else {
		return float64(float_value), nil
	}
}

func AnalysePetalLength(rows [][]string) error {
	petal_length := make([]float64, 0)
	for line_number, row := range rows {
		if len(row) != 5 {
			return fmt.Errorf("Line %v: Expected 5 values", line_number+1)
		}
		// Skip header
		if *flagSkipHeader && line_number == 0 {
			continue
		}
		// Retrieve sepal_length
		if petal_length_value, err := ParseFloat(row[2]); err != nil {
			return err
		} else {
			petal_length = append(petal_length, petal_length_value)
		}
	}

	// Calculate the Max of the variable.
	minVal := floats.Min(petal_length)
	maxVal := floats.Max(petal_length)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(petal_length, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(petal_length, nil)

	// Sort the values.
	inds := make([]int, len(petal_length))
	floats.Argsort(petal_length, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, petal_length, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, petal_length, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, petal_length, nil)

	fmt.Printf("Petal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)

	return nil
}

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}
	if f, err := os.Open(flag.Arg(0)); err != nil {
		log.Println(err)
		return -1
	} else {
		defer f.Close()
		r := csv.NewReader(f)
		if rows, err := r.ReadAll(); err != nil {
			log.Println(err)
			return -1
		} else if err := AnalysePetalLength(rows); err != nil {
			log.Println(err)
			return -1
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
