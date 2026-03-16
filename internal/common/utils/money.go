package utils

import (
	"fmt"
	"strings"
)

func MoneyFmt(m float64) string {
	var format string
	if m > 100000000 {
		format = fmt.Sprintf("%.2f亿", m/100000000)
	} else if m > 10000 {
		format = fmt.Sprintf("%.2f万", m/10000)
	} else {
		format = fmt.Sprintf("%.2f", m)
	}
	return strings.Replace(format, ".00", "", -1)
}
