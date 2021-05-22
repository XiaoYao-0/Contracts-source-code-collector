package domain

import (
	"errors"
)

type Contract struct {
	Number       int
	Address      string
	Name         string
	Json         string
	Sols         []*Sol
	Error        error
	ContractRepo ContractRepo
}

type Sol struct {
	Name string
	Src  string
}

type ContractRepo interface {
	Save(*Contract) error
	Collect(*Contract) error
}

func (contract *Contract) Save() {
	if contract == nil {
		contract.Error = errors.New("nil contract")
		return
	}
	if contract.Error != nil {
		return
	}
	err := contract.ContractRepo.Save(contract)
	if err != nil {
		contract.Error = errors.New("storage failure")
	}
	return
}

func (contract *Contract) Collect() {
	if contract == nil {
		contract.Error = errors.New("nil contract")
		return
	}
	if contract.Error != nil {
		return
	}
	err := contract.ContractRepo.Collect(contract)
	if err != nil {
		contract.Error = errors.New("collect failure")
	}
	return
}
