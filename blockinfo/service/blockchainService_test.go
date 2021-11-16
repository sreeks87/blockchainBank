package service_test

import (
	"errors"
	"testing"

	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/domain"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/downstream/mocks"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/service"
	"github.com/stretchr/testify/assert"
)

// Postive scenario when the request is valid and no error returned
func TestGetBlockInfo(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "BTC",
	}
	d.On("GetBlock", "12w1wefafdv", "BTC").Return(nil, nil)
	_, e := s.GetBlockInfo(&r)
	assert.Equal(t, e, nil)
}

// Postive scenario when the request is valid and no error returned
func TestGetBlockInfo2(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "BTC",
	}
	d.On("GetBlock", "12w1wefafdv", "BTC").Return(nil, errors.New("test error"))
	_, e := s.GetBlockInfo(&r)
	assert.NotEqual(t, e, nil)
}

// Negative scenario validation fails
func TestGetBlockInfoValErr(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "XRP",
	}
	d.On("GetBlock", "12w1wefafdv", "BTC").Return(nil, errors.New("test error"))
	_, e := s.GetBlockInfo(&r)
	assert.NotEqual(t, e, nil)
}

// Negative scenario when tran validation fails
func TestGetTranInfoValErr(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "XRP",
	}
	d.On("GetTransaction", "12w1wefafdv", "BTC").Return(nil, nil)
	_, e := s.GetTranInfo(&r)
	assert.NotEqual(t, e, nil)
}

// Positive scenario when tran info is returned and no errors
func TestGetTranInfo(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "BTC",
	}
	d.On("GetTransaction", "12w1wefafdv", "BTC").Return(nil, nil)
	_, e := s.GetTranInfo(&r)
	assert.Equal(t, e, nil)
}

// Positive scenario when tran info is returned and no errors
func TestGetTranInfo2(t *testing.T) {
	d := new(mocks.Downstream)
	s := service.NewBlockAPI(d)
	r := domain.Request{
		Hash:    "12w1wefafdv",
		NetCode: "BTC",
	}
	d.On("GetTransaction", "12w1wefafdv", "BTC").Return(nil, errors.New("test error"))
	_, e := s.GetTranInfo(&r)
	assert.NotEqual(t, e, nil)
}
