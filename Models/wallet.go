package Models

import (
	"quik-assessment/Config"
)

type Wallet struct {
	ID            int64         `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Balance       float64       `json:"balance"`
	HashedBalance string        `json:"-"`
	Transactions  []Transaction `gorm:"foreignKey:WalletId" json:"transactions,omitempty"`
}

func (b *Wallet) TableName() string {
	return "wallet"
}

//GetUserByID ... Fetch only one user by Id
func GetWalletBalanceByID(wallet *Wallet, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(wallet).Error; err != nil {
		return err
	}
	return nil
}
func UpdateWallet(wallet *Wallet) error {
	err := Config.DB.Model(&wallet).Updates(wallet)
	if err != nil {
		return err.Error
	}
	return nil
}
