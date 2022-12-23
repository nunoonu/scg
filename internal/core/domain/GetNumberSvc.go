package domain

import (
	"net/http"
	"scg/errs"
	"strconv"
)

type GetNumberSvc interface {
	Execute(req GetNumberReq) (*GetNumberRes, error)
}

type GetNumberReq struct{}

type GetNumberRes struct {
	Numbers []Number `json:"numbers"`
}

type Number struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type Alphabet struct {
	Key   string
	Index int
}

type Strings []string

type Nums []*int

func (s Strings) PrepareModel() (Nums, []Alphabet) {

	var nums Nums
	var alphabets []Alphabet
	for i, v := range s {
		n, err := strconv.Atoi(v)
		if err != nil {
			nums = append(nums, nil)
			alphabets = append(alphabets, Alphabet{Key: v, Index: i})
		} else {
			nums = append(nums, &n)
		}
	}
	return nums, alphabets
}

func (n Nums) GetNext(i int) (*int, int) {

	index := i + 1
	var value *int
	for ; index < len(n); index++ {
		if n[index] != nil {
			value = n[index]
			break
		}
	}
	return value, index

}

func (n Nums) GetPrevious(i int) (*int, int) {

	index := i - 1
	var value *int
	for ; index >= 0; index-- {
		if n[index] != nil {
			value = n[index]
			break
		}
	}
	return value, index

}

func (n Nums) FillValue() Nums {

	for i, v := range n {
		if v == nil {
			preVal, preIdx := n.GetPrevious(i)
			nxVal, nxIdx := n.GetNext(i)

			if nxVal == nil {
				prePreVal, _ := n.GetPrevious(preIdx)
				it := *preVal + (*preVal - *prePreVal)
				n[i] = &it
			} else if preVal == nil {
				nxNxVal, _ := n.GetNext(nxIdx)
				it := *nxVal - (*nxNxVal - *nxVal)
				n[i] = &it
			} else {
				it := (*preVal + *nxVal) / (nxIdx - preIdx)
				n[i] = &it
			}

		}
	}
	return n

}

func (n Nums) Validate() error {

	for i := 1; i < len(n)-1; i++ {
		pre := *n[i] - *n[i-1]
		nxt := *n[i+1] - *n[i]

		if pre != nxt {
			return errs.New(http.StatusInternalServerError, "Input is not correct")
		}
	}
	return nil
}
