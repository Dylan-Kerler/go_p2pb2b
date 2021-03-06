package p2pb2b

import (
	"context"
	"flag"
	"testing"
)

func TestMarkets(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.get_markets()
	if err != nil {
		t.Errorf("Markets() get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("Markets() get request success")
}

func TestTickers()(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.get_tickers()
	if err != nil {
		t.Errorf("Tickers() get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("Tickers() get request success, expected %+v\n, got %+v\n", res, res)
}

func TestTicker(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.Get_ticker(Get_ticker_params{Symbol: "ETH_BTC"})
	if err != nil {
		t.Errorf("Ticker() get request failed, expected %v, got %v", nil, err)
		return
	}

	t.Logf("Ticker() get request success")

}

func TestOrderBook(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.OrderBook(OrderBook_params{Market: "ETH_BTC", Side: "buy", Offset: "0", Limit: "100"})
	if err != nil {
		t.Errorf("OrderBook() get request failed, expected Err=%v, got %Err=v\n", nil, err)
	}

	t.Logf("OrderBook() get request success")
}

func TestHistory(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.History(History_params{Market: "ETH_BTC", Last_id: "1",  Limit: "100"})
	if err != nil {
		t.Errorf("History() get request failed, expected %v, got %v\n", nil, err)
	}

	if !res.Success {
		t.Errorf("History() get request failed, expected res.Success to be true,instead got %+v\n", res)
	}

	if len(res.Result) == 0 {
		t.Errorf("History() get request failed, expected res.Result length to be greater than 1, instead got %+v\n", res)
		return
	}

	if i := res.Result[0]; i.Id == 0 || i.Type == "" || i.Time == 0.0 || i.Amount == "" || i.Price == "" {
		t.Errorf("History() get request failed, expected res.Result[0] to have values in all fields, instead got %+v\n", res)
		return
	}

	t.Logf("History() get request success")
}

func TestDepth(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.Depth(Depth_params{Market: "ETH_BTC", Limit: "100"})
	if err != nil {
		t.Errorf("Depth get request failed, expected %v, got %v\n", nil, err)
	}

	if !res.Success {
		t.Errorf("Depth() get request failed, expected res.Success to be true,instead got %+v\n", res)
		return
	}

	if len(res.Result.Asks) < 1 || len(res.Result.Bids) < 1 {
		t.Errorf("Depth() get request failed, expected to get some Asks or Bids, instead got %+v\n", res)
		return
	}

	if len(res.Result.Asks[0]) != 2 || len(res.Result.Bids[0]) != 2 {
		t.Errorf("Depth() get request failed. res.Result.Bids or res.Result.Asks has malformed data, expected [][2]string, instead got %v+\n", res)
	}
	t.Logf("Depth get request success")
}

func TestProducts(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.Products()
	if err != nil {
		t.Errorf("Products() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("Products() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	if len(res.Result) < 1 {
		t.Errorf("Products() get request failed, expected res.Result length to have data, instead got %+v\n", res)
		return
	}

	t.Logf("Products() get request success")
}

func TestSymbols(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", "", ctx)
	res, err := client.Symbols()
	if err != nil {
		t.Errorf("Symbols() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("Symbols() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	if len(res.Result) < 1 {
		t.Errorf("Symbols() get request failed, expected res.Result length to have data, instead got %+v\n", res)
		return
	}

	t.Logf("Symbol() get request success")
}

var APISecret *string = flag.String("secret", "", "Pass in api secret to test protected requests")
var APIKey *string = flag.String("key", "", "Pass in api key to test protected requests")

func TestCurrencyBalance(t *testing.T) {
	fmt.Printf("key: %+v\n secret: %+v", *APIKey, *APISecret)
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.CurrencyBalance(CurrencyBalanceParams{Currency: "ETH"})
	if err != nil {
		t.Errorf("CurrencyBalance() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("CurrencyBalance() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("CurrencyBalance() get request success")
}


func TestBalances(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.Balances()
	if err != nil {
		t.Errorf("Balances() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("Balances() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("Balances() get request success")

}

var OrderIdTmp int64

func TestCreateOrder(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.CreateOrder(CreateOrderParams{
		Market: "ETH_BTC",
		Side:   "sell",
		Amount: "0.03",
		Price:  "0.3",
	})

	if err != nil {
		t.Errorf("CreateOrder() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	OrderIdTmp = res.Result.OrderId
	if !res.Success {
		t.Errorf("CreateOrder() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("CreateOrder() get request success")
}

func TestCancelOrder(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.CancelOrder(CancelOrderParams{
		Market:  "ETH_BTC",
		OrderId: OrderIdTmp,
	})
	if err != nil {
		t.Errorf("CancelOrder() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}
	
	if !res.Success {
		t.Errorf("CancelOrder() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("CancelOrder() get request success")

}

func TestGetOrders(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.GetOrders(GetOrdersParams{
		Market: "ETH_BTC",
		Offset: 0,
		Limit:  100,
	})

	if err != nil {
		t.Errorf("GetOrders() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}
	
	if !res.Success {
		t.Errorf("GetOrders() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("GetOrders() get request success")

}

/// TODO: Still testing this
func TestGetOrder(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.GetOrder(GetOrderParams{
		Offset:  0,
		OrderID: OrderIdTmp,
		Limit:   100,
	})
	if err != nil {
		t.Errorf("GetOrder() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}
	
	if !res.Success {
		t.Errorf("GetOrder() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("GetOrder() get request success")

}

func TestOrderHistory(t *testing.T) {
	ctx := context.Background()
	client := NewClient("https://api.p2pb2b.io", *APIKey, *APISecret, ctx)
	res, err := client.OrderHistory(OrderHistoryParams{
		Offset: 0,
		Limit:  100,
	})
	if err != nil {
		t.Errorf("OrderHistory() get request failed, expected Error=%v, instead got Error=%v\n", nil, err)
		return
	}

	if !res.Success {
		t.Errorf("OrderHistory() get request failed, expected res.Success to be true, instead got %+v\n", res)
		return
	}

	t.Logf("OrderHistory() get request success")
}
