package prices

import "fmt"

type TaxIncludedPricesJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPricesJob) Process() {
	result := make(map[string]float64)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f",price)] = price * (1 + job.TaxRate)
	}
	job.TaxIncludedPrices = result
	fmt.Println(result)
}

func NewTaxIncludedPricesJob(taxrate float64) *TaxIncludedPricesJob {
	prices := []float64{10, 20, 30}

	return &TaxIncludedPricesJob{
		InputPrices: prices,
		TaxRate:     taxrate,
	}
}
