package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"
	"time"
)

func HmacSignRequest(r *http.Request, secret string) {
	query := r.URL.Query()
	ts := time.Now().UnixMilli()
	query.Set("timestamp", string(ts))

	// 参数排序
	params := make([]string, 0, len(query))
	for k := range query {
		params = append(params, k+"="+query.Get(k))
	}
	sort.Strings(params)
	queryString := strings.Join(params, "&")

	// 生成签名
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))

	query.Set("signature", signature)
	r.URL.RawQuery = query.Encode()
}
