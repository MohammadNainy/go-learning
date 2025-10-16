package main

import (
	"fmt"

	"github.com/MohammadNainy/go-learning/Projects/Project1/filemanager"
	"github.com/MohammadNainy/go-learning/Projects/Project1/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("Projects/Project1/prices.txt", fmt.Sprintf("Projects/Project1/result_%.0f.txt", taxRate*100))
		priceJob := prices.NewTaxIncludedPricesJob(fm, taxRate)
		priceJob.Process()
	}
}
