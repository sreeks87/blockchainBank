package downstream_test

import (
	"testing"

	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/downstream"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

// Testing positive scenario when API returns 200
func TestGetBlock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/get_block/BTC/000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		httpmock.NewStringResponder(200, `
	 {
		"status": "success",
		"data": {
		  "network": "BTC",
		  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		  "block_no": 200000,
		  "mining_difficulty": "2864140.507810974",
		  "time": 1348310759,
		  "confirmations": 509874,
		  "is_orphan": false,
		  "txs": [
			"dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"
			],
			"merkleroot": "a08f8101f50fd9c9b3e5252aff4c1c1bd668f878fffaf3d0dbddeb029c307e88",
			"previous_blockhash": "00000000000003a20def7a05a77361b9657ff954b2f2080e135ea6f5970da215",
			"next_blockhash": "00000000000002e3269b8a00caf315115297c626f954770e8398470d7f387e1c",
			"size": 247533
		  }
		}`),
	)

	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/tx/BTC/dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10",
		httpmock.NewStringResponder(200, `
		{
			"status": "success",
			"data": {
			  "network": "BTC",
			  "txid": "80efe43cf64a524d1417546a027786127ad87475f3af1c13b8f3719cd4268679",
			  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
			  "block_no": 200000,
			  "confirmations": 509875,
			  "time": 1348310759,
			  "size": 159,
			  "vsize": 159,
			  "version": 1,
			  "locktime": 0,
			  "sent_value": "50.00000000",
			  "fee": "0.0",
			  "inputs": [],
			  "outputs": [],
			  "tx_hex": "01000000017a38cfb70605d039c0528618b14d9f6a2a41fcaf407008cb32cc22aee78a2bdb000000004a493046022100d666783418029516503cfa10d7be7de6313a038fee3d035c589bed389e43fa42022100fae2c2e3648d53436ea2afc801f0b7930a7c58cd4186dd3d9a9e86f26499241001ffffffff0100f2052a010000001976a9140568015a9facccfd09d70d409b6fc1a5546cecc688ac00000000"
			},
			"code": 200,
			"message": ""
		  }`),
	)

	s := downstream.NewSochain("https://sochain.com/api/v2/")
	_, e := s.GetBlock("000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf", "BTC")
	assert.Equal(t, e, nil)
}

// Testing positive scenario when API returns 200
func TestGetTran(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/tx/BTC/000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		httpmock.NewStringResponder(200, `
		{
			"status": "success",
			"data": {
			  "network": "BTC",
			  "txid": "80efe43cf64a524d1417546a027786127ad87475f3af1c13b8f3719cd4268679",
			  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
			  "block_no": 200000,
			  "confirmations": 509875,
			  "time": 1348310759,
			  "size": 159,
			  "vsize": 159,
			  "version": 1,
			  "locktime": 0,
			  "sent_value": "50.00000000",
			  "fee": "0.0",
			  "inputs": [],
			  "outputs": [],
			  "tx_hex": "01000000017a38cfb70605d039c0528618b14d9f6a2a41fcaf407008cb32cc22aee78a2bdb000000004a493046022100d666783418029516503cfa10d7be7de6313a038fee3d035c589bed389e43fa42022100fae2c2e3648d53436ea2afc801f0b7930a7c58cd4186dd3d9a9e86f26499241001ffffffff0100f2052a010000001976a9140568015a9facccfd09d70d409b6fc1a5546cecc688ac00000000"
			},
			"code": 200,
			"message": ""
		  }`),
	)

	s := downstream.NewSochain("https://sochain.com/api/v2/")
	_, e := s.GetTransaction("000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf", "BTC")
	assert.Equal(t, e, nil)
}

// Testing error scenario when API returns 404
func TestGetBlock2(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/get_block/BTC/000000000000034a7dedef4a161fa073a90155f3a2fe6fc132e0ebf",
		httpmock.NewStringResponder(404, `
	 {
		"status": "success",
		"data": {
		  "network": "BTC",
		  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
		  "block_no": 200000,
		  "mining_difficulty": "2864140.507810974",
		  "time": 1348310759,
		  "confirmations": 509874,
		  "is_orphan": false,
		  "txs": [
			"dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10"
			],
			"merkleroot": "a08f8101f50fd9c9b3e5252aff4c1c1bd668f878fffaf3d0dbddeb029c307e88",
			"previous_blockhash": "00000000000003a20def7a05a77361b9657ff954b2f2080e135ea6f5970da215",
			"next_blockhash": "00000000000002e3269b8a00caf315115297c626f954770e8398470d7f387e1c",
			"size": 247533
		  }
		}`),
	)

	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/tx/BTC/dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10",
		httpmock.NewStringResponder(200, `
		{
			"status": "success",
			"data": {
			  "network": "BTC",
			  "txid": "80efe43cf64a524d1417546a027786127ad87475f3af1c13b8f3719cd4268679",
			  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
			  "block_no": 200000,
			  "confirmations": 509875,
			  "time": 1348310759,
			  "size": 159,
			  "vsize": 159,
			  "version": 1,
			  "locktime": 0,
			  "sent_value": "50.00000000",
			  "fee": "0.0",
			  "inputs": [],
			  "outputs": [],
			  "tx_hex": "01000000017a38cfb70605d039c0528618b14d9f6a2a41fcaf407008cb32cc22aee78a2bdb000000004a493046022100d666783418029516503cfa10d7be7de6313a038fee3d035c589bed389e43fa42022100fae2c2e3648d53436ea2afc801f0b7930a7c58cd4186dd3d9a9e86f26499241001ffffffff0100f2052a010000001976a9140568015a9facccfd09d70d409b6fc1a5546cecc688ac00000000"
			},
			"code": 200,
			"message": ""
		  }`),
	)

	s := downstream.NewSochain("https://sochain.com/api/v2/")
	_, e := s.GetBlock("000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf", "BTC")
	assert.NotEqual(t, e, nil)
}

// Testing error scenario when API returns 404
func TestGetTran2(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		"GET",
		"https://sochain.com/api/v2/tx/BTC/dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10",
		httpmock.NewStringResponder(404, `
		{
			"status": "success",
			"data": {
			  "network": "BTC",
			  "txid": "80efe43cf64a524d1417546a027786127ad87475f3af1c13b8f3719cd4268679",
			  "blockhash": "000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf",
			  "block_no": 200000,
			  "confirmations": 509875,
			  "time": 1348310759,
			  "size": 159,
			  "vsize": 159,
			  "version": 1,
			  "locktime": 0,
			  "sent_value": "50.00000000",
			  "fee": "0.0",
			  "inputs": [],
			  "outputs": [],
			  "tx_hex": "01000000017a38cfb70605d039c0528618b14d9f6a2a41fcaf407008cb32cc22aee78a2bdb000000004a493046022100d666783418029516503cfa10d7be7de6313a038fee3d035c589bed389e43fa42022100fae2c2e3648d53436ea2afc801f0b7930a7c58cd4186dd3d9a9e86f26499241001ffffffff0100f2052a010000001976a9140568015a9facccfd09d70d409b6fc1a5546cecc688ac00000000"
			},
			"code": 200,
			"message": ""
		  }`),
	)

	s := downstream.NewSochain("https://sochain.com/api/v2/")
	_, e := s.GetTransaction("000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf", "BTC")
	assert.NotEqual(t, e, nil)
}
