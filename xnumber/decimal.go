package xnumber

import (
	"fmt"
	"strconv"
)

func Decimal2(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func Decimal0(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.0f", value), 64)
	return value
}

func Decimal1(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", value), 64)
	return value
}
