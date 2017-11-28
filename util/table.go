package util

import (
	"fmt"
)

// Value is a single value in the data
type Value struct {
	String string
}

// Table is the table of values with optional column headers
type Table struct {
	Columns []string
	colmap  map[string]int
	Rows    [][]*Value
}

var (
	ErrDuplicateColumn = fmt.Errorf("Duplicate or invalid column name")
)

// NewTable creates a new table with specified columns
func NewTable(columns ...string) (*Table, error) {
	this := new(Table)
	this.Columns = make([]string, 0, len(columns))
	this.colmap = make(map[string]int, len(columns))
	if err := this.AppendColumns(columns...); err != nil {
		return nil, err
	}
	return this, nil
}

// AppendColumns appends columns onto the table
func (this *Table) AppendColumns(columns ...string) error {
	// Update columns and colmap
	for i, column := range columns {
		if _, exists := this.colmap[column]; exists {
			return ErrDuplicateColumn
		}
		this.colmap[column] = i
		this.Columns = append(this.Columns, column)
	}
	return nil
}

// AddRowString appends a row of string values onto the table
// and will return an error if the length of the string exceeds
// the number of columns
func (this *Table) AddRowString(values []string) error {
	// TODO
	return nil
}

// ReadCSV reads data from a CSV file. Sometimes there are comments
// and a header line within the file
func (this *Table) ReadCSV(filename string, skip_header, skip_comments bool) error {
	// TODO
	return nil
}
