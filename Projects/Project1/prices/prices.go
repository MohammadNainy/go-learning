package prices

import (
	"fmt"

	"github.com/MohammadNainy/go-learning/Projects/Project1/conversion"
	"github.com/MohammadNainy/go-learning/Projects/Project1/filemanager"
)

type TaxIncludedPricesJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_price"`
	TaxIncludedPrices map[string]string       `json:"tex_included_prices"`
}

func (job *TaxIncludedPricesJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
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

	job.IOManager.WriteResult(job)

}

func NewTaxIncludedPricesJob(fm filemanager.FileManager, taxrate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager: fm,
		TaxRate:   taxrate,
	}
}
