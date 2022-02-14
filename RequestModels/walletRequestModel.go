package RequestModels

import "github.com/thedevsaddam/govalidator"

type Transaction struct {
	Amount govalidator.Float64 `json:"amount"`
}

type Transaction2 struct {
	Amount govalidator.Float64 `json:"amount"`
}
