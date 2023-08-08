package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

func writeCSV(content CSVContent, path string) (bool, error) {

	file, err := os.Create(path)
	if err != nil {
		return false, errors.New("unnable create csv file")
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(strings.Split(content.Header[0], ",")); err != nil {
		return false, errors.New("unnable write headers to csv file")
	}

	for _, row := range content.Content {
		if err := writer.Write(strings.Split(row, ",")); err != nil {
			return false, errors.New("unnable write headers to csv file")
		}
	}

	if err := writer.Error(); err != nil {
		return false, errors.New(fmt.Sprintf("error in writer, error: %v", err))
	}

	return true, nil
}
