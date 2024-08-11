package error_handling

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a / b, nil
}
