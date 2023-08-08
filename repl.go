package main

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strings"
)

func readCSV(path string) (CSVContent, error) {
	file, err := os.Open(path)
	if err != nil {
		return CSVContent{[]string{""}, []string{""}}, errors.New("unnable open csv file")
	}

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	var header []string
	var body []string
	var cntr int

	for scan.Scan() {
		if cntr == 0 {
			header = append(header, scan.Text())
			cntr++
			continue
		}

		body = append(body, scan.Text())
		cntr++
	}

	file.Close()
	return CSVContent{header, body}, nil
}

func convertReflect(value reflect.Value) ([]string, error) {
	if value.Kind() == reflect.Slice {
		i := value.Interface()
		if convert, ok := i.([]string); ok {
			return convert, nil
		} else {
			return []string{""}, errors.New("unable convert to []string")
		}

	} else {
		return []string{""}, errors.New("not reflect.Slice")
	}
}

func replace(content CSVContent, field string, toReplace string) (CSVContent, error) {

	var indx int
	var isFind bool
	var cHeader []string
	var cContent []string

	rcontent := reflect.ValueOf(content)

	for i := 0; i < rcontent.NumField(); i++ {
		if i == 0 {
			data, err := convertReflect(rcontent.Field(i))
			if err != nil {
				return CSVContent{[]string{""}, []string{""}}, errors.New("unable convert to []string type header of csv")
			}

			headers := strings.Split(data[0], ",")
			for i, v := range headers {
				if v == field {
					isFind = true
					indx = i + 1
				}
			}

			cHeader = append(cHeader, strings.Join(headers[:], ","))
			continue
		}

		if !isFind || indx == 0 {
			return CSVContent{[]string{""}, []string{""}}, errors.New("incoming value in csv not finded")
		}

		data, err := convertReflect(rcontent.Field(i))
		if err != nil {
			return CSVContent{[]string{""}, []string{""}}, errors.New("unable convert to []string type body of csv")
		}

		for _, lines := range data {
			var line []string
			sv := strings.Split(lines, ",")
			for i, v := range sv {
				if i == indx-1 {
					line = append(line, toReplace)
					continue
				}
				line = append(line, v)
			}
			cContent = append(cContent, strings.Join(line[:], ","))
		}
	}

	return CSVContent{cHeader, cContent}, nil
}
