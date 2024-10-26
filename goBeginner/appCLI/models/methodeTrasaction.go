package models

func (s *AccountService) GetTransactions(userID string) []Transaction {
    var transactions []Transaction
    for _, t := range s.Transactions {
        if t.FromUserID == userID || t.ToUserID == userID {
            transactions = append(transactions, t)
        }
    }
    return transactions
}