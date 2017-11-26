// Usage:
//  go run stats_01.go iris.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

func AnalyseSepalLength(rows [][]string) error {
	sepal_length := make([]float64, 0)
	for line_number, row := range rows {
		if len(row) != 5 {
			return fmt.Errorf("Line %v: Expected 5 values", line_number+1)
		}
		// Skip header
		if *flagSkipHeader && line_number == 0 {
			continue
		}
		// Retrieve sepal_length
		if sepal_length_value, err := ParseFloat(row[0]); err != nil {
			return err
		} else {
			sepal_length = append(sepal_length, sepal_length_value)
		}
	}

	modeVal, modeCount := stat.Mode(sepal_length, nil)
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", stat.Mean(sepal_length, nil))
	fmt.Printf("Mode value & count: %0.2f, %f\n", modeVal, modeCount)

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
		} else if err := AnalyseSepalLength(rows); err != nil {
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
