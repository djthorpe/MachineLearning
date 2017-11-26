// Usage:
//  go run csv_reader.go data.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

///////////////////////////////////////////////////////////////////////////////

func ParseInteger(string_value string) (int, error) {
	// Check for empty value
	if len(strings.TrimSpace(string_value)) == 0 {
		return 0, ErrEmpty
	}
	// Convert to integer
	if int_value, err := strconv.ParseInt(string_value, 10, 32); err != nil {
		return 0, err
	} else {
		return int(int_value), nil
	}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func AnalyseData(rows [][]string) error {
	var maxFilesRemaining int
	for line_number, row := range rows {
		if len(row) != 5 {
			return fmt.Errorf("Line %v: Expected 5 values", line_number+1)
		}
		// Skip header
		if *flagSkipHeader && line_number == 0 {
			continue
		}
		// Parse the second value
		filesRemaining, err := ParseInteger(row[1])
		if err == ErrEmpty {
			continue
		} else if err != nil {
			return fmt.Errorf("Line %v: %v", line_number+1, err)
		}
		maxFilesRemaining = MaxInt(filesRemaining, maxFilesRemaining)
	}

	fmt.Printf("maxFilesRemaining=%v\n", maxFilesRemaining)

	return nil
}

func RunMain() int {
	if f, err := os.Open("data.csv"); err != nil {
		log.Println(err)
		return -1
	} else {
		defer f.Close()
		r := csv.NewReader(f)
		if rows, err := r.ReadAll(); err != nil {
			log.Println(err)
			return -1
		} else if err := AnalyseData(rows); err != nil {
			log.Println(err)
			return -1
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	os.Exit(RunMain())
}
