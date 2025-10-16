package main

import (
	"fmt"

	"github.com/MohammadNainy/go-learning/Projects/Project1/cmdmanager"
	// "github.com/MohammadNainy/go-learning/Projects/Project1/filemanager"
	"github.com/MohammadNainy/go-learning/Projects/Project1/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("Projects/Project1/prices.txt", fmt.Sprintf("Projects/Project1/result_%.0f.txt", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPricesJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}

}
