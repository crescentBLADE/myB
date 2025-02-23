package rest

import (
	"fmt"
	"myb/client"
	"myb/util"
	"testing"
)

func TestGetMarketPrice(t *testing.T) {
	ak, sk := util.GetCredentials()
	c := client.NewBinanceClient(ak, sk)
	s := SpotService{client: c}
	p, err := s.GetMarketPrice("BTCUSDT")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("This is one p", p)
	// 调用 SDK 函数
}

func TestGetMarketPrices(t *testing.T) {
	ak, sk := util.GetCredentials()
	c := client.NewBinanceClient(ak, sk)
	s := SpotService{client: c}
	p, err := s.GetMarketPrices("BTCUSDT,ETHUSDT")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("This is  two p", p)
	// 调用 SDK 函数
}

func TestGetAllMarketPrice(t *testing.T) {
	ak, sk := util.GetCredentials()
	c := client.NewBinanceClient(ak, sk)
	s := SpotService{client: c}
	p, err := s.GetAllMarketPrices()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("This is all p", p)
	// 调用 SDK 函数
}
