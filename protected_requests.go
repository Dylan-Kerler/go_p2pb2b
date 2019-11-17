package p2pb2b

import (
	"encoding/json"
	"net/http"
	"time"
)

type currencyBalanceResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type CurrencyBalanceJsonRes struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Result  currencyBalanceResult `json:"result"`
}

type CurrencyBalanceJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Currency   string `json:"currency"`
}

type CurrencyBalanceParams struct {
	Currency string /// e.g "ETH"
}

func (clt *Client) CurrencyBalance(opts CurrencyBalanceParams) (*CurrencyBalanceJsonRes, error) {
	endpoint := "/api/v1/account/balance"
	postBody := CurrencyBalanceJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().Unix(),
		Currency:   opts.Currency,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CurrencyBalanceJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type BalancesJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
}

type balancesResult struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

type BalancesJsonRes struct {
	Success bool                      `json:"success"`
	Message string                    `json:"message"`
	Result  map[string]balancesResult `json:"result"`
}

func (clt *Client) Balances() (*BalancesJsonRes, error) {
	endpoint := "/api/v1/account/balances"
	postBody := BalancesJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().Unix(),
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes BalancesJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}

type CreateOrderParams struct {
	Market string `json:"market"`
	Side   string `json:"side"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

type CreateOrderJsonBody struct {
	RequestUrl string `json:"request"`
	Nonce      int64  `json:"nonce"`
	Market     string `json:"market"`
	Side       string `json:"side"`
	Amount     string `json:"amount"`
	Price      string `json:"price"`
}

type CreateOrderJsonRes struct {
}

func (clt *Client) CreateOrder(opts CreateOrderParams) (*CreateOrderJsonRes, error) {
	endpoint := "/api/v1/order/new"
	postBody := CreateOrderJsonBody{
		RequestUrl: endpoint,
		Nonce:      time.Now().Unix(),
		Market:     opts.Market,
		Side:       opts.Side,
		Amount:     opts.Amount,
		Price:      opts.Price,
	}

	res, err := clt.AuthAPIRequest(postBody, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	var jsonRes CreateOrderJsonRes
	if err := json.Unmarshal(res, &jsonRes); err != nil {
		return nil, err
	}

	return &jsonRes, nil
}
