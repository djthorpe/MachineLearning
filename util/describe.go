package util

import "fmt"

func (this *Table) Describe() (*Table, error) {
	// Create a new table with the same columns and one additional column at the start
	that := new(Table)
	if err := that.SetColumns("[parameter]"); err != nil {
		return nil, err
	}
	if err := that.AppendColumns(this.Columns...); err != nil {
		return nil, err
	}
	// Compute the number of samples for each column
	that.AppendStringRow(this.describe_samples(), true)
	that.AppendStringRow(this.describe_type(), true)
	return that, nil
}

func (this *Table) describe_samples() []string {
	samples := make([]uint, len(this.Columns))
	for _, row := range this.Rows {
		for column := range row {
			if row[column] != nil {
				samples[column]++
			}
		}
	}
	samples_str := make([]string, len(samples)+1)
	samples_str[0] = "samples"
	for i := range samples {
		samples_str[i+1] = fmt.Sprint(samples[i])
	}
	return samples_str
}

func (this *Table) describe_type() []string {
	types_str := make([]string, len(this.Columns)+1)
	return types_str
}
