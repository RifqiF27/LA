package service

// import (
// 	"time"
// 	"voucher_system/models"
// )

// func (s *voucherService) ReportHistory(userID int, voucherCode string, transactionAmount float64, paymentMethod string, area string) {
// 	voucher, benefitValue, err := s.ValidateVoucher(userID, voucherCode, transactionAmount, 0, area, paymentMethod, time.Now())
// 	if err != nil {
// 		return
// 	}

// 	history := &models.History{
// 		UserID:            userID,
// 		VoucherID:         voucher.ID,
// 		TransactionAmount: transactionAmount,
// 		BenefitValue:      benefitValue,
// 		UsageDate:         time.Now(),
// 	}

// 	err = s.repo.History.CreateHistory(history)
// 	if err != nil {
// 		return
// 	}

// 	newQuota := voucher.Quota - 1
// 	if newQuota < 0 {
// 		return
// 	}

// 	err = s.repo.Voucher.UpdateVoucherQuota(voucher.ID, newQuota)
// 	if err != nil {
// 		return
// 	}

	
// }