package metrics

import "math"

func Accuracy(yTrue, yPred []int) float64 {
	if len(yTrue) != len(yPred) || len(yTrue) == 0 {
		return 0
	}
	correct := 0
	for i := range yTrue {
		if yTrue[i] == yPred[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(yTrue))
}

func Recall(yTrue, yPred []int) float64 {
	tp, fn := 0, 0
	for i := range yTrue {
		if yTrue[i] == 1 {
			if yPred[i] == 1 {
				tp++
			} else {
				fn++
			}
		}
	}
	if tp+fn == 0 {
		return 0
	}
	return float64(tp) / float64(tp+fn)
}

// Precision = TP / (TP + FP)
func Precision(yTrue, yPred []int) float64 {
	tp, fp := 0, 0
	for i := range yTrue {
		if yPred[i] == 1 {
			if yTrue[i] == 1 {
				tp++
			} else {
				fp++
			}
		}
	}
	if tp+fp == 0 {
		return 0
	}
	return float64(tp) / float64(tp+fp)
}

// F1 Score = 2 * (Precision * Recall) / (Precision + Recall)
func F1Score(yTrue, yPred []int) float64 {
	p := Precision(yTrue, yPred)
	r := Recall(yTrue, yPred)
	if p+r == 0 {
		return 0
	}
	return 2 * (p * r) / (p + r)
}

// Mean Absolute Error
func MAE(yTrue, yPred []float64) float64 {
	if len(yTrue) != len(yPred) || len(yTrue) == 0 {
		return 0
	}
	sum := 0.0
	for i := range yTrue {
		sum += math.Abs(yTrue[i] - yPred[i])
	}
	return sum / float64(len(yTrue))
}

// Root Mean Squared Error
func RMSE(yTrue, yPred []float64) float64 {
	if len(yTrue) != len(yPred) || len(yTrue) == 0 {
		return 0
	}
	sum := 0.0
	for i := range yTrue {
		diff := yTrue[i] - yPred[i]
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(yTrue)))
}
