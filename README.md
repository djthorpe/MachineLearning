# MachineLearning

Examples from "Machine Learning with Go" by Daniel Whitenack

## Chapter 3

The data file called `time_series.csv` has two columns, one the
predicated value and one the observed value. These are continuously
changing sets of data (floating point). You can calculate three values
from this data set:

  * Mean-squared error
  * Mean Absolute error
  * R-squared

In order to compute:

```
  go run chapter3/mean_squared_error.go chapter3/time_series.csv
```