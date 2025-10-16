package prices

import (
	"fmt"

	"github.com/MohammadNainy/go-learning/Projects/Project1/conversion"
	"github.com/MohammadNainy/go-learning/Projects/Project1/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_price"`
	TaxIncludedPrices map[string]string       `json:"tex_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrices := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrices)
	}
	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)

}

func NewTaxIncludedPricesJob(iom iomanager.IOManager, taxrate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: iom,
		TaxRate:   taxrate,
	}
}
