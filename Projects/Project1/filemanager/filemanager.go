package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	Inputfile  string
	Outputfile string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.Inputfile)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, errors.New("reading the file content failed")
	}
	return lines, nil

}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.Outputfile)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed convert data to JSON")
	}
	return nil
}

func New(input, output string) FileManager {
	return FileManager{
		Inputfile: input ,
		Outputfile: output,
	}
}