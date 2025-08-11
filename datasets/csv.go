package datasets

import (
	"encoding/csv"
	"os"
	"strconv"
)

// LoadCSV loads a CSV file and returns the data as a slice of slices of strings.
func LoadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SaveCSV saves a slice of slices of strings to a CSV file.
func SaveCSV(filePath string, data [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

// ConvertToFloat converts a slice of strings to a slice of float64.
func ConvertToFloat(data []string) ([]float64, error) {
	floatData := make([]float64, len(data))
	for i, v := range data {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		floatData[i] = value
	}
	return floatData, nil
}

// ConvertToInt converts a slice of strings to a slice of int.
func ConvertToInt(data []string) ([]int, error) {
	intData := make([]int, len(data))
	for i, v := range data {
		value, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		intData[i] = value
	}
	return intData, nil
}

type CSVData struct {
	Headers []string
	Rows    [][]string
}

func NewCSVData(headers []string, rows [][]string) *CSVData {
	return &CSVData{
		Headers: headers,
		Rows:    rows,
	}
}
