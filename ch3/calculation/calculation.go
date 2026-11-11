package calculation

import "errors"

func Add(params ...int) int {
	var sum int
	for _, val := range params {
		sum += val
	}
	return sum
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
