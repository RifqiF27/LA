package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)
func GenerateExcelReport(data [][]interface{}, headers []string, folder string, fileNamePrefix string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("data is empty, nothing to write")
	}

	// Tentukan nama file dan path untuk Excel
	fileName := fmt.Sprintf("%s/%s_%d.xlsx", folder, fileNamePrefix, time.Now().Unix())

	// Membuat file Excel baru
	f := excelize.NewFile()

	// Menambahkan sheet baru, memastikan sheet "Data Report" ada
	sheetName := "Data Report"
	index, _ := f.NewSheet(sheetName)

	// Menulis header ke sheet
	for colIndex, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+colIndex)) // Menulis header ke A1, B1, C1, dst.
		if err := f.SetCellValue(sheetName, cell, header); err != nil {
			return "", err
		}
	}

	// Menulis data ke dalam Excel
	for rowIndex, rowData := range data {
		for colIndex, fieldValue := range rowData {
			cell := fmt.Sprintf("%s%d", string('A'+colIndex), rowIndex+2) // Mulai dari baris 2
			if err := f.SetCellValue(sheetName, cell, fieldValue); err != nil {
				return "", err
			}
		}
	}

	// Menetapkan sheet aktif (sheet pertama yang akan dibuka)
	f.SetActiveSheet(index)

	// Simpan file Excel
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}

	log.Println("Excel file created:", fileName)
	return fileName, nil
}
