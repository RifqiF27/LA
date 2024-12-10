package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	// "reflect"
	"time"
)

func GenerateCSVReport(data [][]interface{}, folder string) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("data is empty, nothing to write")
	}

	fileName := fmt.Sprintf("%s/voucher_report_%d.csv", folder, time.Now().Unix())
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// if err := writer.Write(headers); err != nil {
	// 	return "", err
	// }

	for _, row := range data {
		var data []string
		for _, cell := range row {
			// Mengonversi setiap elemen interface{} menjadi string
			data = append(data, fmt.Sprintf("%v", cell))
		}
		if err := writer.Write(data); err != nil {
			return "", err
		}
	}

	// info, err := os.Stat(fileName)
	// if err != nil {
	// 	return "", err
	// }
	// log.Printf("CSV file created, size: %d bytes", info.Size())

	return fileName, nil
}

// func GenerateCSVReport(data []interface{}, folder string) (string, error) {
// 	// log.Println(data, "<<<<<<<<")
// 	if len(data) == 0 {
// 		fmt.Println("No data to write to CSV")
// 		return "", fmt.Errorf("data is empty, nothing to write")
// 	}
// 	// Tentukan nama file dan path
// 	fileName := fmt.Sprintf("%s/voucher_report_%d.csv", folder, time.Now().Unix())
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	writer := csv.NewWriter(file)
// 	defer writer.Flush()

// 	if len(data) == 0 {
// 		return "", fmt.Errorf("data is empty, nothing to write")
// 	}

// 	// Menulis header CSV berdasarkan field struktur data pertama
// 	if len(data) > 0 {
// 		val := reflect.ValueOf(data[0])
// 		if val.Kind() == reflect.Struct {
// 			// Ambil field names untuk header
// 			var headers []string
// 			for i := 0; i < val.NumField(); i++ {
// 				headers = append(headers, val.Type().Field(i).Name)
// 			}
// 			if err := writer.Write(headers); err != nil {
// 				return "", err
// 			}
// 		}
// 	}

// 	// Menulis data transaksi voucher ke dalam CSV
// 	for _, item := range data {
// 		val := reflect.ValueOf(item)
// 		var row []string
// 		if val.Kind() == reflect.Struct {
// 			for i := 0; i < val.NumField(); i++ {
// 				fieldValue := fmt.Sprintf("%v", val.Field(i).Interface())

// 				// Log setiap field untuk melihat nilai yang diterima
// 				// log.Println("Field Name:", val.Type().Field(i).Name, "Field Value:", fieldValue)
// 				row = append(row, fieldValue)
// 			}
// 			// log.Println("row", row, "<<<<<<<<<")
// 			if err := writer.Write(row); err != nil {
// 				// log.Fatal("Error writing row:", err)
// 				return "", err
// 			}
// 		}
// 	}
// 	// info, err := os.Stat(fileName)
// 	// if err != nil {
// 	// 	log.Fatal("Error checking file size:", err)
// 	// }
// 	// log.Println("CSV file created, size:", info.Size(), "bytes")

// 	return fileName, nil
// }
