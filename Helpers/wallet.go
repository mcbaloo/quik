package Helpers

import (
	"errors"

	"github.com/shopspring/decimal"
)

func ComputeCredit(currentBalance decimal.Decimal, creditAmount decimal.Decimal) decimal.Decimal {
	newBalance := currentBalance.Add(creditAmount)
	return newBalance
}
func CheckTransactionAmount(creditAmount decimal.Decimal) error {
	if creditAmount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return errors.New("amount cannot be less than or equals zero(0)")
	}
	return nil
}

func ComputeBalanceAfterWithdrawal(currentBalance decimal.Decimal, debitAmount decimal.Decimal)error{
	diff := currentBalance.Sub(debitAmount)
	if diff.LessThan(decimal.NewFromInt(0)){
		return errors.New("wallet balance cannot be less than zero")
	}
	return nil
}