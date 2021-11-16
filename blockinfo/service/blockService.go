package service

import (
	"errors"

	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/domain"
)

type blockService struct {
	downAPI domain.Downstream
}

func NewBlockAPI(d domain.Downstream) domain.Service {
	return &blockService{
		downAPI: d,
	}
}

// GetBlockInfo function ios responible for calling the downstream in our case sochain
// to get the block details
func (s *blockService) GetBlockInfo(r *domain.Request) (*domain.Block, error) {
	b, e := Validate(r)
	if b {
		return s.downAPI.GetBlock(r.Hash, r.NetCode)
	}
	return nil, e
}

// GetTranInfo function ios responible for calling the downstream in our case sochain
// to get the tran details
func (s *blockService) GetTranInfo(r *domain.Request) (*domain.Transaction, error) {
	b, e := Validate(r)
	if b {
		return s.downAPI.GetTransaction(r.Hash, r.NetCode)
	}
	return nil, e
}

func Validate(r *domain.Request) (bool, error) {
	// Validate if the network code is one of BTC/LTC/DOGE
	coins := map[string]bool{
		"BTC":  true,
		"LTC":  true,
		"DOGE": true,
	}

	if !coins[r.NetCode] {
		return false, errors.New("invalid network code")
	}
	return true, nil
}
