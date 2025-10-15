package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricesJob) LoadData() {
	file, err := os.Open("Projects/Project1/prices.txt")
	if err != nil {
		fmt.Println("An error occurred: ", err)
		return
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("Reading the file content failed: ", err)
		return
	}

	prices := make([]float64, len(lines))

	for lineindex, line := range lines {
		val, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Converting price to float failed: ", err)
			return
		}

		prices[lineindex] = val
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPricesJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}
	job.TaxIncludedPrices = result
	fmt.Println(result)
}

func NewTaxIncludedPricesJob(taxrate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		TaxRate: taxrate,
	}
}
