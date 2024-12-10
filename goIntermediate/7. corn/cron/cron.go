package cron

import (
	"fmt"
	"os"
	"voucher_system/service"
	"voucher_system/utils"

	"go.uber.org/zap"
)

type CronJob struct {
	voucherService service.Service
	log            *zap.Logger
}

func NewCronJob(voucherService service.Service, log *zap.Logger) *CronJob {
	return &CronJob{
		voucherService: voucherService,
		log:            log,
	}
}

func (c *CronJob) RunVoucherUsageJob() {
	c.log.Info("Starting voucher usage cron job")

	usedVouchers, err := c.voucherService.Voucher.GetUsedVouchers()
	if err != nil {
		c.log.Error("Failed to fetch used vouchers", zap.Error(err))
		return
	}

	c.log.Info("Fetched used vouchers", zap.Int("count", len(usedVouchers)))

	if len(usedVouchers) == 0 {
		c.log.Warn("No used vouchers to process")
	}


	var dataCsv [][]interface{}
	var dataExcel [][]interface{}
	for i, v := range usedVouchers {

			rowCsv := []interface{}{
			
				fmt.Sprintf("no: %d ", i+1),
				fmt.Sprintf("userID: %d ", v.UserID),
				fmt.Sprintf("voucherID: %d ", v.VoucherID),
				fmt.Sprintf("date_usage: %s ", v.UsageDate.Format("2006-01-02 15:04:05")),
				fmt.Sprintf("transaction_amount: %.2f ", v.TransactionAmount),
				fmt.Sprintf("benefit: %.2f", v.BenefitValue),
			}
			dataCsv = append(dataCsv, rowCsv)

			rowExcel := []interface{}{
				i + 1,
				v.UserID,
				v.VoucherID,
				v.UsageDate.Format("2006-01-02 15:04:05"),
				v.TransactionAmount,
				v.BenefitValue,
			}
			dataExcel = append(dataExcel, rowExcel)

		
	}

	// c.log.Info("Data to be written to CSV", zap.Int("dataCount", len(data)))
	
	headers := []string{"No", "UserID", "VoucherID", "Date Usage", "Transaction Amount", "Benefit"}
	// Prepare the folder for CSV reports
	folderCsv := "reports/csv"
	if err := os.MkdirAll(folderCsv, os.ModePerm); err != nil {
		c.log.Error("Failed to create folder for CSV", zap.Error(err))
		return
	}
	
	// Generate the CSV report
	filePathCsv, err := utils.GenerateCSVReport(dataCsv, folderCsv)
	if err != nil {
		c.log.Error("Failed to generate CSV report", zap.Error(err))
		return
	}
	c.log.Info("CSV report generated successfully", zap.String("filePath", filePathCsv))
	
	folderExcel := "reports/excel"
	if err := os.MkdirAll(folderExcel, os.ModePerm); err != nil {
		c.log.Error("Failed to create folder for excel", zap.Error(err))
		return
	}

	// Generate the CSV report
	filePathExcel, err := utils.GenerateExcelReport(dataExcel, headers, folderExcel, "voucher_report")
	if err != nil {
		c.log.Error("Failed to generate excel report", zap.Error(err))
		return
	}

	c.log.Info("Excel report generated successfully", zap.String("filePath", filePathExcel))
}

// func (c *CronJob) RunVoucherUsageJob() {
// 	c.log.Info("Starting voucher usage cron job")

// 	// Fetch all used vouchers
// 	usedVouchers, err := c.voucherService.Voucher.GetUsedVouchers()
// 	if err != nil {
// 		c.log.Error("Failed to fetch used vouchers", zap.Error(err))
// 		return
// 	}

// 	c.log.Info("Fetched used vouchers", zap.Int("count", len(usedVouchers)))

// 	if len(usedVouchers) == 0 {
//         c.log.Warn("No used vouchers to process")
//     }

// 	// Convert to interface{} for CSV generation
// 	data := make([]interface{}, len(usedVouchers))
// 	for i, v := range usedVouchers {
// 		row := []interface{}{
// 			v.UserID,
// 			v.VoucherID,
// 			v.UsageDate.Format("2006-01-02 15:04:05"),
// 			v.TransactionAmount,
// 			v.BenefitValue,
// 		}
// 		data[i] = row
// 	}

// 	c.log.Info("Data to be written to CSV", zap.Int("dataCount", len(data)))

// 	// Prepare the folder for CSV reports
// 	folder := "csv_reports"
// 	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
// 		c.log.Error("Failed to create folder for CSV", zap.Error(err))
// 		return
// 	}

// 	// Generate the CSV report
// 	filePath, err := utils.GenerateCSVReport(data, folder)
// 	if err != nil {
// 		c.log.Error("Failed to generate CSV report", zap.Error(err))
// 		return
// 	}

// 	c.log.Info("CSV report generated successfully", zap.String("filePath", filePath))
// }
