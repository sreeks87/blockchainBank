package delivery_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	delivery "github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/delivery/http"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/domain"
	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/service/mocks"
	"github.com/stretchr/testify/assert"
)

// Test bad request scenario- empty request
func TestGetBlockInfo(t *testing.T) {
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}

	req, err := http.NewRequest("POST", "/block", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetBlockInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}

// Test bad request scenario- empty request
func TestGetBlockInfo2(t *testing.T) {
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}

	req, err := http.NewRequest("POST", "/block", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetBlockInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}

// Testing success scenario when block data is fetched from external api
func TestGetBlockInfo3(t *testing.T) {
	r := []byte(`{"hash":"000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf","network_code":"BTC"}`)
	request := domain.Request{
		Hash:    "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		NetCode: "BTC",
	}
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}
	resp := &domain.Block{
		NetCode:      "BTC",
		BlockNum:     122,
		Timestamp:    1223,
		PrevBlock:    "1qwedqdwfwe",
		NextBlock:    "sfswf2fwvadfva",
		Size:         1,
		Transactions: nil,
	}
	mockSvc.On("GetBlockInfo", &request).Return(resp, nil)
	req, err := http.NewRequest("POST", "/block", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetBlockInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 200)

}

// Testing negative scenario when block data is not fetched from external api
func TestGetBlockInfo4(t *testing.T) {
	r := []byte(`{"hash":"000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf","network_code":"BTC"}`)
	request := domain.Request{
		Hash:    "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		NetCode: "BTC",
	}
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}
	resp := &domain.Block{
		NetCode:      "BTC",
		BlockNum:     122,
		Timestamp:    1223,
		PrevBlock:    "1qwedqdwfwe",
		NextBlock:    "sfswf2fwvadfva",
		Size:         1,
		Transactions: nil,
	}
	mockSvc.On("GetBlockInfo", &request).Return(resp, errors.New("test error"))
	req, err := http.NewRequest("POST", "/block", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetBlockInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}

// Testing scenario when tran data is not fetched from external api
func TestGetTranInfo(t *testing.T) {
	r := []byte(`{"hash":"000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf","network_code":"BTC"}`)
	request := domain.Request{
		Hash:    "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		NetCode: "BTC",
	}
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}
	resp := &domain.Transaction{
		TranId:    "1234",
		Fee:       "0",
		Timestamp: 1233,
		SentValue: "123",
	}
	mockSvc.On("GetTranInfo", &request).Return(resp, errors.New("test error"))
	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetTranInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}

// Testing success scenario when tran data is fetched from external api
func TestGetTranInfo2(t *testing.T) {
	r := []byte(`{"hash":"000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf","network_code":"BTC"}`)
	request := domain.Request{
		Hash:    "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		NetCode: "BTC",
	}
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}
	resp := &domain.Transaction{
		TranId:    "1234",
		Fee:       "0",
		Timestamp: 1233,
		SentValue: "123",
	}
	mockSvc.On("GetTranInfo", &request).Return(resp, nil)
	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(r))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetTranInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 200)

}

// Testing empty request scenario
func TestGetTranInfo3(t *testing.T) {
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}

	req, err := http.NewRequest("POST", "/transaction", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetTranInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}

// Testing empty request scenario
func TestGetTranInfo4(t *testing.T) {
	mockSvc := new(mocks.Service)
	c := delivery.BCInfoController{Svc: mockSvc}

	req, err := http.NewRequest("POST", "/transaction", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(c.GetTranInfo)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 400)

}
