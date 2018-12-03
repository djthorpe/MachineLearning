package util

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Error is an error message with potentially a line number
type Error struct {
	reason string
	line   int
}

// Value is a single value in the data
type Value struct {
	Str      string
	_Uint64  *uint64
	_Float64 *float64
}

// Table is the table of values with optional column headers
type Table struct {
	Columns []string
	colmap  map[string]int
	Rows    [][]*Value
}

var (
	ErrDuplicateColumn = &Error{reason: "Duplicate or invalid column name"}
	ErrDimensionError  = &Error{reason: "Too many values for row"}
	ErrOutOfRange      = &Error{reason: "Index out of range"}
	ErrNotFound        = &Error{reason: "Column Not Found"}
)

// NewTable creates a new table with specified columns
func NewTable(columns ...string) (*Table, error) {
	this := new(Table)
	if err := this.SetColumns(columns...); err != nil {
		return nil, err
	}
	return this, nil
}

// SetColumns sets the columns for the table
func (this *Table) SetColumns(columns ...string) error {
	this.Columns = make([]string, 0, len(columns))
	this.colmap = make(map[string]int, len(columns))
	if err := this.AppendColumns(columns...); err != nil {
		return err
	}
	return nil
}

// NumberOfColumns returns the number of columns for the
// table
func (this *Table) NumberOfColumns() int {
	return len(this.Columns)
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

// AppendStringRow appends a row of string values onto the table
// and will return an error if the length of the string exceeds
// the number of columns. If you set treat_empty_as_nil to true
// then any string value which is only whitespace or of zero length
// is treated as nil
func (this *Table) AppendStringRow(values []string, treat_empty_as_nil bool) error {
	if len(values) > len(this.Columns) {
		return ErrDimensionError
	}
	// Create a row of values
	row := make([]*Value, len(this.Columns))
	for i := 0; i < len(values); i++ {
		if i >= len(values) {
			continue
		} else if treat_empty_as_nil && (values[i] == "" || strings.TrimSpace(values[i]) == "") {
			continue
		} else {
			row[i] = &Value{Str: values[i]}
		}
	}

	// Append row
	if this.Rows == nil {
		this.Rows = make([][]*Value, 0, 1)
	}
	this.Rows = append(this.Rows, row)

	// Return success
	return nil
}

// StringRow returns a row as an array of string values for row index n. If
// any values are nil then the nil_string is used
func (this *Table) StringRow(n int, nil_string string) ([]string, error) {
	if n < 0 || n >= len(this.Rows) {
		return nil, ErrOutOfRange
	}
	values := this.Rows[n]
	row := make([]string, len(this.Columns))
	for i := range row {
		if i >= len(values) || values[i] == nil {
			row[i] = nil_string
		} else {
			row[i] = values[i].Str
		}
	}
	return row, nil
}

// StringColumn returns all values in a specific named column, c. If
// any values are nil then the nil_string is used
func (this *Table) StringColumn(c string, nil_string string) ([]string, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]string, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_string
			} else {
				column[i] = values[n].Str
			}
		}
		return column, nil
	}
}

// FloatColumn returns all values in a specific named column, c as float64 values. If
// any values are nil then the nil_value is used (usually 0.0). If any value cannot be
// converted to a float, then an error is returned
func (this *Table) FloatColumn(c string, nil_value float64) ([]float64, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]float64, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_value
			} else {
				if float, err := values[n].Float64(); err != nil {
					return nil, err
				} else {
					column[i] = float
				}
			}
		}
		return column, nil
	}
}

// UintColumn returns all values in a specific named column, c as uint values. If
// any values are nil then the nil_value is used (usually 0). If any value cannot be
// converted to a uint, then an error is returned
func (this *Table) UintColumn(c string, nil_value uint) ([]uint, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]uint, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_value
			} else {
				if value, err := values[n].Uint64(); err != nil {
					return nil, err
				} else {
					column[i] = uint(value)
				}
			}
		}
		return column, nil
	}
}

// ReadCSV reads data from a CSV file. Sometimes there are comments
// and a header line within the file
func (this *Table) ReadCSV(filename string, skip_header, skip_comments, treat_empty_as_nil bool) error {
	if f, err := os.Open(filename); err != nil {
		return err
	} else {
		defer f.Close()
		if rows, err := csv.NewReader(f).ReadAll(); err != nil {
			return err
		} else {
			is_header := !skip_header
			for _, row := range rows {
				if len(row) == 0 && skip_comments {
					continue
				} else if skip_comments && strings.TrimSpace(row[0]) == "" {
					continue
				} else if skip_comments && (strings.HasPrefix(row[0], "#") || strings.HasPrefix(row[0], "//")) {
					continue
				} else if is_header {
					// Set the columns from the header, over-writing the
					// existing columns
					if err := this.SetColumns(row...); err != nil {
						return err
					}
					is_header = false
				} else if err := this.AppendStringRow(row, treat_empty_as_nil); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Stringify
func (this *Value) String() string {
	return this.Str
}

func (this *Error) Error() string {
	if this.line == 0 {
		return this.reason
	} else {
		return fmt.Sprintf("%v @ line %v", this.reason, this.line)
	}
}

func (this *Error) Line(i int) {
	this.line = i
}

func (this *Value) Float64() (float64, error) {
	if this._Float64 != nil {
		return *this._Float64, nil
	} else if f, err := float64conv(this.Str); err != nil {
		return math.NaN(), err
	} else {
		this._Float64 = &f
		return f, nil
	}
}

func (this *Value) Uint64() (uint64, error) {
	if this._Uint64 != nil {
		return *this._Uint64, nil
	} else if v, err := strconv.ParseUint(this.Str, 10, 64); err != nil {
		return 0, err
	} else {
		this._Uint64 = &v
		return v, nil
	}
}

func float64conv(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Output table as ASCII table
func (this *Table) String() string {
	buf := new(bytes.Buffer)
	table := tablewriter.NewWriter(buf)
	table.SetHeader(this.Columns)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for i := range this.Rows {
		if row, err := this.StringRow(i, "<nil>"); err != nil {
			buf.WriteString(fmt.Sprintf("%v\n", err))
		} else {
			table.Append(row)
		}
	}
	table.Render()
	return buf.String()
}
