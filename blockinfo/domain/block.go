package domain

type Request struct {
	Hash    string `json:"hash"`
	NetCode string `json:"network_code"`
}

type Block struct {
	NetCode      string        `json:"net_code"`
	BlockNum     int           `json:"block_num"`
	Timestamp    int           `json:"timestamp"`
	PrevBlock    string        `json:"previous"`
	NextBlock    string        `json:"next"`
	Size         int           `json:"size"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	TranId    string `json:"txid"`
	Fee       string `json:"fee"`
	Timestamp int    `json:"timestamp"`
	SentValue string `json:"sent_value"`
}

type Service interface {
	GetBlockInfo(*Request) (*Block, error)
	GetTranInfo(*Request) (*Transaction, error)
}
