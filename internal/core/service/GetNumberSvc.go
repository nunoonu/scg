package service

import (
	"scg/internal/core/domain"
)

type getNumberSvc struct{}

func NewGetNumberSvc() domain.GetNumberSvc {
	return &getNumberSvc{}
}

func (c getNumberSvc) Execute(req domain.GetNumberReq) (*domain.GetNumberRes, error) {

	_ = req

	nums, alphabets := getNumbers().PrepareModel()
	if err := nums.FillValue().Validate(); err != nil {
		return nil, err
	}

	return buildRes(nums, alphabets), nil

}

func buildRes(nums domain.Nums, alphabets []domain.Alphabet) *domain.GetNumberRes {

	var numbers []domain.Number
	for _, v := range alphabets {

		numbers = append(numbers, domain.Number{Key: v.Key, Value: *nums[v.Index]})

	}
	return &domain.GetNumberRes{Numbers: numbers}

}

func getNumbers() domain.Strings {
	return []string{"1", "x", "8", "17", "y", "z", "78", "113"}
	//return []string{"a", "5", "10", "b", "20", "c", "30", "d", "e"}
}
