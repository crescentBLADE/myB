package rest

import (
	"encoding/json"

	"myb/client"
	"myb/util"
)

type Prices struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
type SpotService struct {
	client *client.BinanceClient
}

// 获取市场价格（例如 BTC/USDT）
func (s *SpotService) GetMarketPrice(symbol string) (Prices, error) {
	var data Prices
	params := map[string]string{
		"symbol": symbol,
	}
	response, err := s.client.RequestBinanceAPI("GET", "/api/v3/ticker/price", params)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		return data, err
	}
	// 将字节数组转换为字符串

	return data, nil
}

// 获取多个交易对市场价格（例如 BTC/USDT）
// TODO: Need Fix
func (s *SpotService) GetMarketPrices(symbols string) (data []Prices, err error) {

	params := map[string]string{
		"symbol": symbols,
	}
	response, err := s.client.RequestBinanceAPI("GET", "/api/v3/ticker/price", params)
	if err != nil {
		return data, err
	}
	util.PrintlnRawMessage(response)
	err = json.Unmarshal(response, &data)
	if err != nil {
		return data, err
	}

	return data, err
}

// 获取所有市场价格（例如 BTC/USDT）
func (s *SpotService) GetAllMarketPrices() (data []Prices, err error) {

	response, err := s.client.RequestBinanceAPI("GET", "/api/v3/ticker/price", nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
