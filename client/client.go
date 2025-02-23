package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"myb/logger"
)

const (
	BaseURL = "https://api.binance.com"
	TestUrl = " https://testnet.binance.vision/api"
)

var standby = [5]string{"https://api-gcp.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
	"https://api4.binance.com"}

type BinanceClient struct {
	APIKey     string
	SecretKey  string
	HTTPClient *http.Client
	BaseURL    string
	standby    []string
}

func NewBinanceClient(apiKey, secretKey string) *BinanceClient {

	return &BinanceClient{
		APIKey:    apiKey,
		SecretKey: secretKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		BaseURL: BaseURL,
		standby: standby[:],
	}
}

func (c *BinanceClient) SwitchBaseUrl() error {
	r, err := c.RequestBinanceAPI("GET", "/api/v3/ping", nil)
	if err != nil || r != nil {
		logger.LogInfo("Biannce Api Cant Connect,Try Another Url")
		for _, v := range c.standby {
			r, err := c.RequestBinanceAPI("GET", "/api/v3/ping", nil)
			if r == nil && err == nil {
				c.BaseURL = v
				return nil
			}

		}
		logger.LogError("No Avaliable Url, Stop the Server Now" + err.Error())
		return errors.New("No Avaliable Url, Stop the Server Now, err is : %s" + err.Error())
	}
	return nil
}

// 请求 Binance API 的函数
func (c *BinanceClient) RequestBinanceAPI(method string, endpoint string, params map[string]string) ([]byte, error) {

	// 拼接查询字符串
	queryString := "?"
	for key, value := range params {
		queryString += fmt.Sprintf("%s=%s&", key, value)
	}
	queryString = queryString[:len(queryString)-1]

	req, err := http.NewRequest(method, c.BaseURL+endpoint+queryString, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-MBX-APIKEY", c.APIKey)
	//------------
	proxyURL, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return nil, err
	}

	// 创建自定义的 HTTP Transport，使用 SOCKS5 代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	//------------

	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
