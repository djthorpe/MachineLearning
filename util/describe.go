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
	that.AppendStringRow(this.describe_type(), true)
	r1, r2, r3 := this.describe_samples()
	that.AppendStringRow(r1, true)
	that.AppendStringRow(r2, true)
	that.AppendStringRow(r3, true)
	return that, nil
}

func (this *Table) describe_type() []string {
	types_str := make([]string, len(this.Columns)+1)
	types_str[0] = "type"
	for i, c := range this.Columns {
		if t, err := this.TypeForColumn(c); err == nil {
			types_str[i+1] = t
		}
	}
	return types_str
}

func (this *Table) describe_samples() ([]string, []string, []string) {
	samples_str := make([]string, len(this.Columns)+1)
	sum_str := make([]string, len(this.Columns)+1)
	mean_str := make([]string, len(this.Columns)+1)
	samples_str[0] = "samples"
	sum_str[0] = "sum"
	mean_str[0] = "mean"
	for column := range this.Columns {
		var sum float64
		var cells, samples uint
		for row := range this.Rows {
			if value := this.Rows[row][column]; value != nil {
				cells++
				if v, err := value.Float64(); err == nil {
					sum += v
					samples++
				}
			}
		}
		if cells > 0 {
			samples_str[column+1] = fmt.Sprint(cells)
		}
		if samples > 0 {
			sum_str[column+1] = fmt.Sprintf("%.2f", sum)
			mean_str[column+1] = fmt.Sprintf("%.2f", sum/float64(samples))
		}

	}
	return samples_str, sum_str, mean_str
}
