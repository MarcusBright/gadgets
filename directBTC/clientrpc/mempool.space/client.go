package mempoolspace

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	url        string
	httpClient *http.Client
}

// ThrottledTransport Rate Limited HTTP Client
type ThrottledTransport struct {
	roundTripperWrap http.RoundTripper
	ratelimiter      *rate.Limiter
}

func (c *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := c.ratelimiter.Wait(r.Context()) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	return c.roundTripperWrap.RoundTrip(r)
}

// NewThrottledTransport wraps transportWrap with a rate limitter
// examle usage:
// client := http.DefaultClient
// client.Transport = NewThrottledTransport(10*time.Seconds, 60, http.DefaultTransport) allows 60 requests every 10 seconds
func NewThrottledTransport(limitPeriod time.Duration, requestCount int, transportWrap http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		roundTripperWrap: transportWrap,
		ratelimiter:      rate.NewLimiter(rate.Every(limitPeriod), requestCount),
	}
}

func NewClient(url string, requestCount, periodSecond int) *Client {
	return &Client{
		url: url,
		httpClient: &http.Client{
			Transport: NewThrottledTransport(time.Duration(periodSecond)*time.Second, requestCount, http.DefaultTransport),
		},
	}
}

func (c *Client) GetAddressNewTransactions(address string, seenTxid string) ([]AddressTransaction, error) {
	url := fmt.Sprintf("%s/api/address/%s/txs", c.url, address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var txs []AddressTransaction
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()

	for {
		// fmt.Println("url:", req.URL.String())
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}
		var respBody []AddressTransaction
		err = json.Unmarshal(body, &respBody)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response body: %w", err)
		}
		for _, tx := range respBody {
			if tx.Txid != seenTxid {
				txs = append(txs, tx)
			} else {
				return txs, nil
			}
		}
		if len(respBody) < 10 {
			break
		}
		params.Set("after_txid", respBody[len(respBody)-1].Txid)
		req.URL.RawQuery = params.Encode()
	}
	return txs, nil
}

func (c *Client) GetTx(txHash string) (AddressTransaction, error) {

	url := fmt.Sprintf("%s/api/tx/%s", c.url, txHash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return AddressTransaction{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AddressTransaction{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return AddressTransaction{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AddressTransaction{}, fmt.Errorf("error reading response body: %w", err)
	}
	var tx AddressTransaction
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return AddressTransaction{}, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return tx, nil
}

func (c *Client) GetLatestBlockNumber() (uint64, error) {
	url := fmt.Sprintf("%s/api/blocks/tip/height", c.url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}
	var height uint64
	err = json.Unmarshal(body, &height)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return height, nil
}

func (c *Client) GetBlockHash(height uint64) (string, error) {
	url := fmt.Sprintf("%s/api/block-height/%d", c.url, height)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}
	return string(body), nil
}

func (c *Client) GetBlockTxId(hash string) ([]string, error) {
	url := fmt.Sprintf("%s/api/block/%s/txids", c.url, hash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	var txids []string
	err = json.Unmarshal(body, &txids)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return txids, nil
}
