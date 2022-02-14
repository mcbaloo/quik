package Models

type Transaction struct {
	ID       int64 `gorm:"primary_key, AUTO_INCREMENT"`
	Amount   float64
	Type     string
	WalletID int64 `gorm:"column:wallet_id"`
	Wallet   Wallet
}

func (b *Transaction) TableName() string {
	return "transaction"
}
