package metrics

import "testing"

func TestAccuracy(t *testing.T) {
	yTrue := []int{1, 0, 1, 1, 0}
	yPred := []int{1, 0, 1, 0, 0}
	got := Accuracy(yTrue, yPred)
	want := 0.8
	if got != want {
		t.Errorf("Accuracy() = %v; want %v", got, want)
	}
}

func TestPrecisionRecallF1(t *testing.T) {
	yTrue := []int{1, 0, 1, 1, 0}
	yPred := []int{1, 0, 1, 0, 0}

	if Precision(yTrue, yPred) != 1.0 {
		t.Errorf("Precision failed")
	}
	if Recall(yTrue, yPred) != 0.6666666666666666 {
		t.Errorf("Recall failed")
	}
	if F1Score(yTrue, yPred) != 0.8 {
		t.Errorf("F1 Score failed")
	}
}

func TestMAERMSE(t *testing.T) {
	yTrue := []float64{2.5, 0.0, 2.1, 7.8}
	yPred := []float64{3.0, -0.1, 2.1, 7.8}

	if MAE(yTrue, yPred) != 0.15 {
		t.Errorf("MAE failed")
	}
	if RMSE(yTrue, yPred) != 0.22360679774997896 {
		t.Errorf("RMSE failed")
	}
}
