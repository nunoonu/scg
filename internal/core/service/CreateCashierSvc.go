package service

import "scg/internal/core/domain"

type createCashierSvc struct{}

func NewCreateCashierSvc() domain.CreateCashierSvc {
	return &createCashierSvc{}
}

func (g createCashierSvc) Execute(req domain.CreateCashierReq) (*domain.CreateCashierRes, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	cashierLimits := getCashierLimit()

	res, err := req.Calculate(cashierLimits)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func getCashierLimit() []domain.Limit {

	return []domain.Limit{{1000, 10}, {500, 20}, {100, 15}, {50, 20}, {20, 30}, {10, 20}, {5, 20}, {1, 20}, {0.25, 50}}

}
