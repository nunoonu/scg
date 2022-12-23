package domain

import (
	"net/http"
	"scg/errs"
)

type CreateCashierSvc interface {
	Execute(req CreateCashierReq) (*CreateCashierRes, error)
}

type CreateCashierReq struct {
	ProductPrice float64 `json:"productPrice" binding:"required"`
	Pay          float64 `json:"pay" binding:"required"`
}

type CreateCashierRes struct {
	ChangeMoneys changeMoneys `json:"changeMoneys"`
}

type changeMoney struct {
	Amount int     `json:"amount"`
	Value  float64 `json:"value"`
}

type Limit struct {
	Value  float64 `json:"value"`
	Amount int     `json:"amount"`
}

type changeMoneys []changeMoney

func (r CreateCashierReq) Validate() error {

	if r.ProductPrice > r.Pay {
		return errs.New(http.StatusBadRequest, "product price is less than pay money")
	}
	return nil
}

func (r CreateCashierReq) Calculate(limits []Limit) (*CreateCashierRes, error) {

	remaining := r.Pay - r.ProductPrice

	var chMoneys changeMoneys
	for _, v := range limits {

		value := v.Value
		amountLimit := v.Amount

		amount := int(remaining / value)
		if amount > amountLimit {
			amount = amountLimit
		}

		chMoneys.append(amount, v.Value)

		remaining = remaining - (float64(amount) * value)
		if remaining == 0 {
			break
		}
	}

	if remaining != 0 {
		return nil, errs.New(http.StatusInternalServerError, "out of money to return")
	}

	return &CreateCashierRes{ChangeMoneys: chMoneys}, nil
}

func (ch *changeMoneys) append(amount int, value float64) {
	*ch = append(*ch, changeMoney{amount, value})
}
