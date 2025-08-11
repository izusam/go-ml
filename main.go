package main

import (
	"fmt"

	"github.com/izusam/go-ml/metrics"
)

func main() {

	yTrue := []int{1, 0, 1, 1, 0}
	yPred := []int{1, 0, 0, 1, 1}

	fmt.Println("Accuracy:", metrics.Accuracy(yTrue, yPred))
	fmt.Println("Recall:", metrics.Recall(yTrue, yPred))
	fmt.Println("Recall:", metrics.Recall(yTrue, yPred))
	fmt.Println("F1 Score:", metrics.F1Score(yTrue, yPred))

	yTrueF := []float64{2.5, 0.0, 2.1, 7.8}
	yPredF := []float64{3.0, -0.1, 2.1, 7.8}

	fmt.Println("MAE:", metrics.MAE(yTrueF, yPredF))
	fmt.Println("RMSE:", metrics.RMSE(yTrueF, yPredF))
}
