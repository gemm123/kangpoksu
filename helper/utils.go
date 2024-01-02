package helper

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func RandomNumberOrder(totalOrder int) int {
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(90) + 10
	totalOrder = totalOrder / 100
	totalOrder = totalOrder*100 + randNumber

	return totalOrder
}

func FormatRupiah(angka float64) string {
	numberString := fmt.Sprintf("%.0f", angka)
	split := strings.Split(numberString, ",")
	sisa := len(split[0]) % 3
	rupiah := split[0][:sisa]
	ribuan := regexp.MustCompile(`\d{3}`).FindAllString(split[0][sisa:], -1)

	// Add a dot if the input is already in thousands
	if ribuan != nil {
		separator := ""
		if sisa != 0 {
			separator = "."
		}
		rupiah += separator + strings.Join(ribuan, ".")
	}

	if len(split) > 1 {
		rupiah += "," + split[1]
	}

	if rupiah == "" {
		return ""
	}

	return "Rp. " + rupiah
}
