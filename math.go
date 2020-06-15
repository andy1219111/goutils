package goutils

import (
	"fmt"
	"math"
	"strconv"
)

//Round 四舍五入
func Round(number float64, decimal uint8) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f", math.Floor(number*100+0.5)/100), 64)
	return value
}
