package domain

type Downstream interface {
	GetBlock(string, string) (*Block, error)
	GetTransaction(string, string) (*Transaction, error)
}

type BlockResp struct {
	Status    string `json:"status"`
	DataFrame BData  `json:"data"`
}

type BData struct {
	Network       string   `json:"network"`
	BlockHash     string   `json:"blockhash"`
	BlockNum      int      `json:"block_no"`
	MiningDiff    string   `json:"mining_difficulty"`
	Timestamp     int      `json:"time"`
	Confirmations int      `json:"confirmations"`
	IsOrphan      bool     `json:"is_orphan"`
	Transactions  []string `json:"txs"`
	MerkleRoot    string   `json:"merkleroot"`
	Previous      string   `json:"previous_blockhash"`
	Next          string   `json:"next_blockhash"`
	Size          int      `json:"size"`
}

type TranResp struct {
	Status    string `json:"status"`
	DataFrame TData  `json:"data"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type TData struct {
	Network       string                   `json:"network"`
	TxID          string                   `json:"txid"`
	BlockHash     string                   `json:"blockhash"`
	BlockNum      int                      `json:"block_no"`
	Confirmations int                      `json:"confirmations"`
	Timestamp     int                      `json:"time"`
	Size          int                      `json:"size"`
	Vsize         int                      `json:"vsize"`
	Version       int                      `json:"version"`
	Locktime      int                      `json:"locktime"`
	SentValue     string                   `json:"sent_value"`
	Fee           string                   `json:"fee"`
	Inputs        []map[string]interface{} `json:"inputs"`
	Outputs       []map[string]interface{} `json:"outputs"`
	TXHex         string                   `json:"tx_hex"`
}
