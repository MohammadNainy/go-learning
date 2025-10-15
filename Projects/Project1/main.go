package main

import (
	"github.com/MohammadNainy/go-learning/Projects/Project1/prices"	
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPricesJob(taxRate)
		priceJob.Process()
	}
}
