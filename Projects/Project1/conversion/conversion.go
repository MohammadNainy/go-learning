package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFloat(str []string) ([]float64, error) {

	flt := make([]float64, len(str))

	for IndexStr, ValStr := range str {

		val, err := strconv.ParseFloat(ValStr, 64)

		if err != nil {
			fmt.Println("Converting price to float failed: ", err)
			return nil, errors.New("Failed to convert string to float! ")
		}

		flt[IndexStr] = val
	}
	return flt, nil
}
