package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 获取 AK 和 SK
func GetCredentials() (string, string) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		log.Fatal("API Key or Secret Key is not set in environment variables.")
	}

	return apiKey, secretKey
}
func PrintlnRawMessage(r []byte) {
	// 使用 RawMessage 来直接获取原始数据
	var raw json.RawMessage
	err := json.Unmarshal([]byte(r), &raw)
	if err != nil {
		return
	}
	responseStr := string(r)
	fmt.Println("This is Raw", responseStr)
}
