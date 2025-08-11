package linear

// Package linear implements linear regression algorithms.
// It provides functionalities for fitting linear models to datasets,
// making predictions, and evaluating model performance.
import (
	"fmt"
)

// LinearRegression represents a linear regression model.
type LinearRegression struct {
	Weights []float64 // Coefficients of the linear model
}

// Fit trains the linear regression model using the provided features and target values.
func (lr *LinearRegression) Fit(features [][]float64, target []float64) error {
	if len(features) == 0 || len(target) == 0 {
		return fmt.Errorf("features and target must not be empty")
	}
	if len(features) != len(target) {
		return fmt.Errorf("number of features must match number of target values")
	}

	numFeatures := len(features[0])
	lr.Weights = make([]float64, numFeatures)

	// Simple implementation of gradient descent or normal equation can be added here.
	// For now, we will just initialize weights to zero.
	for i := range lr.Weights {
		lr.Weights[i] = 0.0
	}

	return nil
}

// Predict makes predictions using the trained linear regression model.
func (lr *LinearRegression) Predict(features [][]float64) ([]float64, error) {
	if len(lr.Weights) == 0 {
		return nil, fmt.Errorf("model is not trained yet")
	}
	if len(features) == 0 {
		return nil, fmt.Errorf("features must not be empty")
	}

	predictions := make([]float64, len(features))
	for i, feature := range features {
		if len(feature) != len(lr.Weights) {
			return nil, fmt.Errorf("feature length does not match model weights length")
		}
		var prediction float64
		for j, weight := range lr.Weights {
			prediction += weight * feature[j]
		}
		predictions[i] = prediction
	}

	return predictions, nil
}

// Evaluate calculates the mean squared error of the model predictions against the target values.
func (lr *LinearRegression) Evaluate(features [][]float64, target []float64) (float64, error) {
	predictions, err := lr.Predict(features)
	if err != nil {
		return 0.0, err
	}

	if len(predictions) != len(target) {
		return 0.0, fmt.Errorf("number of predictions must match number of target values")
	}

	var mse float64
	for i := range predictions {
		diff := predictions[i] - target[i]
		mse += diff * diff
	}
	mse /= float64(len(predictions))

	return mse, nil
}

// String returns a string representation of the linear regression model.
func (lr *LinearRegression) String() string {
	return fmt.Sprintf("LinearRegression(Weights: %v)", lr.Weights)
}

// NewLinearRegression creates a new instance of LinearRegression.
func NewLinearRegression() *LinearRegression {
	return &LinearRegression{
		Weights: []float64{},
	}
}
