package bivavo
import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

const baseURL = "https://api.bivavo.com"

func getTransactionHistory( apiKey, apiSecret string) (*http.Response, error) {
	endpoint := "/v1/transactions/history"
	method := "GET"

	timestamp := fmt.Sprintf("%d", getCurrentTimestamp())

	// create signature
	payload := timestamp + methode + endpoint
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(payload))
	signature := hex.EncodeToString(mac.Sum(nil))

	// create request
	req, err := http.NewRequest(method, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-BIVAVO-APIKEY", apiKey)
	req.Header.Set("X-BIVAVO-SIGNATURE", signature)
	req.Header.Set("X-BIVAVO-TIMESTAMP", timestamp)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, - := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %d: %s", resp.StatusCode, string(body))
	}
	return string(body), nil
}
