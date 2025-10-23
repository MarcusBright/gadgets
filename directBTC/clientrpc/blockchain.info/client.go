package blockchaininfo

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

func (c *Client) GetAddressNewTransactions(address string, seenTxid string) ([]Transaction, error) {
	url := fmt.Sprintf("%s/rawaddr/%s", c.url, address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var txs []Transaction
	params := req.URL.Query()
	params.Add("limit", fmt.Sprintf("%d", 30))
	params.Add("offset", fmt.Sprintf("%d", 0))
	req.URL.RawQuery = params.Encode()

	for {
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
		var respBody AddressTransactions
		err = json.Unmarshal(body, &respBody)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling response body: %w", err)
		}
		for _, tx := range respBody.Txs {
			if tx.Hash != seenTxid {
				txs = append(txs, tx)
			} else {
				return txs, nil
			}
		}
		if len(respBody.Txs) < 30 {
			break
		}
		params.Set("offset", fmt.Sprintf("%d", len(txs)))
		req.URL.RawQuery = params.Encode()
	}
	return txs, nil
}
