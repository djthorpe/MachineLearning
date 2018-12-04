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

Alternatively, let's say the observed and the predicted sets of data
are categories. The categories in `labeled.csv` are 0, 1 and 2. Then
you can see the number of true positives and false positives (that is,
where the observed category is the same as the predicted category)
as follows:

```
  go run chapter3/category_accuracy.go chapter3/labeled.csv
```

Let's say now we are testing against the '0' category:

  * If the predicted value and observed value are both 0, then this is a true positive (TP)
  * If the predicted value is 0 but the observed value is not 0, then this is a false positive (FP)
  * If the predicted value is not 0 but observed is 0 this is a false negative (FN)
  * Finally, if the predicted value is not 0 and the observed vaoue is not 0, this is a true negative (TN)

Then we can create some metrics:

  * Accuracy: The ratio of true predictions vs false predictions: (TP+TN)/(FP+FN+TP+TN)
  * Precision: The ratio of true predictions over all predictions: TP/(TP+FP)
  * Recall

When evaluating data, you can create training and testing sets. See how this works with
the following command, which subsamples one set of data into two distinct sets:

```
  go run chapter3/subsample.go chapter3/time_series.csv
```

