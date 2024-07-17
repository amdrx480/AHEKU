package utils

import (
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// FormatRupiah mengonversi bilangan menjadi format mata uang Rupiah Indonesia
func FormatRupiah(amount interface{}, decimalPlaces int, decimalSeparator, thousandSeparator string) string {
	var formattedAmount string
	switch v := amount.(type) {
	case int:
		formattedAmount = strconv.Itoa(v)
	case float64:
		formattedAmount = strconv.FormatFloat(v, 'f', decimalPlaces, 64)
	default:
		return "Invalid amount"
	}

	parts := strings.Split(formattedAmount, ".")
	integerPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	// Format integer part with thousand separator
	integerPartWithSeparator := addThousandSeparator(integerPart, thousandSeparator)

	if decimalPlaces > 0 {
		// Add leading zeros to decimal part if needed
		for len(decimalPart) < decimalPlaces {
			decimalPart += "0"
		}
		return "Rp " + integerPartWithSeparator + decimalSeparator + decimalPart[:decimalPlaces]
	}
	return "Rp " + integerPartWithSeparator
}

func addThousandSeparator(value, separator string) string {
	n := len(value)
	if n <= 3 {
		return value
	}
	remainder := n % 3
	if remainder == 0 {
		remainder = 3
	}
	return value[:remainder] + separator + addThousandSeparator(value[remainder:], separator)
}

//

// SetColAlignRight mengatur rata kanan untuk kolom tertentu di lembar Excel dan menambahkan border
func SetColAlignRight(f *excelize.File, sheetName, col string) error {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}

	// Menggunakan gaya untuk border
	rightAlignedStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		return err
	}

	// Loop through rows and set right-aligned style with border
	for r := range rows {
		cell := col + strconv.Itoa(r+1)
		if err := f.SetCellStyle(sheetName, cell, cell, rightAlignedStyle); err != nil {
			return err
		}
	}

	return nil
}

// getRightAlignedStyle mengembalikan style yang mengatur rata kanan
func getRightAlignedStyle(f *excelize.File) (styleID int) {
	styleID, _ = f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
		},
	})
	return styleID
}
