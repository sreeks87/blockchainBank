package downstream

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"sync"
	"time"

	"github.com/NuriCareers/Sreekanth-Cheriyanath-coding-challenge/blockinfo/domain"
)

type sochain struct {
	BlockAPI *url.URL
	TranAPI  *url.URL
}

const (
	BLOCK_API = "/get_block/"
	TRAN_API  = "/tx/"
)

// NewSochain returns a new object of downstream - sochain implementation
func NewSochain(baseURL string) domain.Downstream {
	log.SetPrefix("NewSochain :")
	log.Println("base url :", baseURL)
	urlBlock, _ := getFullURL(baseURL, BLOCK_API)
	urlTran, _ := getFullURL(baseURL, TRAN_API)
	log.Println("urlBlock ", urlBlock, "urlTran ", urlTran)
	return &sochain{
		BlockAPI: urlBlock,
		TranAPI:  urlTran,
	}

}

// GetBlock function responsible for the actual GET of the Block data
// input : request
// output : block,error
func (s *sochain) GetBlock(hash string, code string) (*domain.Block, error) {
	log.SetPrefix("GetBlock :")

	body, e := makeGETCall(s.BlockAPI, hash, code)
	if e != nil {
		return nil, e
	}
	var blockResp domain.BlockResp
	e = json.Unmarshal(body, &blockResp)
	if e != nil {
		return nil, e
	}
	var wg sync.WaitGroup
	tranChannel := make(chan domain.Transaction)
	// if the length of transactions is less than 10, then use that
	max := 10
	if len(blockResp.DataFrame.Transactions) < 10 {
		max = len(blockResp.DataFrame.Transactions)
	}
	for i := 0; i < max; i++ {
		// concurrently get the trasaction details
		tranId := blockResp.DataFrame.Transactions[i]
		log.Println("fetching tran details for ", tranId)
		wg.Add(1)
		go s.concurrentTranDetails(s.TranAPI, code, tranId, &wg, tranChannel)
	}
	var transactions []domain.Transaction
	go checkChan(&wg, tranChannel)
	// for every item returned by concurrent execution, append it to the final lsist
	for tran := range tranChannel {
		transactions = append(transactions, tran)
	}
	return mappedBlockResp(&blockResp, transactions)
}

// GetTransaction function responsible for the actual GET of the Tran data
// input : request
// output : Transaction,error
func (s *sochain) GetTransaction(hash string, code string) (*domain.Transaction, error) {
	log.SetPrefix("GetTransaction :")
	url := s.TranAPI
	log.Println("url", url, "hash ", hash, "code ", code)
	body, e := makeGETCall(s.TranAPI, hash, code)
	if e != nil {
		log.Println("get failed ", e)
		return nil, e
	}
	var tranResp domain.TranResp
	e = json.Unmarshal(body, &tranResp)
	if e != nil {
		log.Println("unmarshal failed ", e)
		return nil, e
	}
	return mappedTranResp(&tranResp)
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 20 * time.Second,
	}
}

// mappedBlockResp is responsible for mapping the unmarshalled response to the actual Block response
func mappedBlockResp(resp *domain.BlockResp, transactions []domain.Transaction) (*domain.Block, error) {
	return &domain.Block{
		NetCode:      resp.DataFrame.Network,
		BlockNum:     resp.DataFrame.BlockNum,
		Timestamp:    resp.DataFrame.Timestamp,
		PrevBlock:    resp.DataFrame.Previous,
		NextBlock:    resp.DataFrame.Next,
		Transactions: transactions,
		Size:         resp.DataFrame.Size,
	}, nil
}

// mappedBlockResp is responsible for mapping the unmarshalled response to the actual Transaction response
func mappedTranResp(resp *domain.TranResp) (*domain.Transaction, error) {
	return &domain.Transaction{
		TranId:    resp.DataFrame.TxID,
		Fee:       resp.DataFrame.Fee,
		Timestamp: resp.DataFrame.Timestamp,
		SentValue: resp.DataFrame.SentValue,
	}, nil
}

// Actual Rest call for GETs
func makeGETCall(url *url.URL, hash string, code string) ([]byte, error) {
	log.SetPrefix("makeGETCall ")
	client := newHTTPClient()
	fullURL := url.String() + "/" + code + "/" + hash
	log.Println("full url ", fullURL)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		log.Println("new request failed ", err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("http call failed ", err, " for url ", fullURL)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("sochain returned error ", resp.StatusCode, " for url ", fullURL)
		return nil, errors.New("downstream returned error ")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response ", err)
		return nil, err
	}
	return body, nil
}

func getFullURL(baseURL string, pathURL string) (*url.URL, error) {
	log.SetPrefix("getFullURL")
	base, e := url.Parse(baseURL)
	if e != nil {
		log.Println("error occured ", e.Error())
		return nil, e
	}
	log.Println("base url :", base)
	base.Path = path.Join(base.Path, pathURL)
	return base, nil
}

func (s *sochain) concurrentTranDetails(url *url.URL, code string, tranId string, wg *sync.WaitGroup, out chan<- domain.Transaction) {
	defer wg.Done()
	var tranDetails *domain.Transaction
	tranDetails, e := s.GetTransaction(tranId, code)
	if e != nil {
		log.Println("could not get transactions for ", tranId)
		tranDetails = &domain.Transaction{
			TranId: tranId,
		}
	}
	out <- *tranDetails
}

// checkChan will monitor the dirrent concurrently running funcions and the channel usage
func checkChan(wg *sync.WaitGroup, ch chan domain.Transaction) {
	wg.Wait()
	close(ch)
}
