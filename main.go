package test

import (
	"encoding/json"
	"net/url"

	"github.com/valyala/fasthttp"
)

type Client struct {
	Secret  string
	Sitekey string
}

type Response struct {
	Success     bool   `json:"success"`
	Credit      bool   `json:"credit"`
	Hostname    string `json:"hostname"`
	ChallengeTs string `json:"challenge_ts"`
}

func New(secret_key, site_key string) *Client {
	return &Client{
		Secret:  secret_key,
		Sitekey: site_key,
	}
}

func (c *Client) Verify(token string) bool {
	resp := fasthttp.AcquireResponse()
	req := fasthttp.AcquireRequest()

	req.SetRequestURI("https://hcaptcha.com/siteverify")

	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/x-www-form-urlencoded")

	req.SetBodyString(url.Values{
		"sitekey":  []string{c.Sitekey},
		"secret":   []string{c.Secret},
		"response": []string{token},
	}.Encode())

	defer fasthttp.ReleaseResponse(resp)
	defer fasthttp.ReleaseRequest(req)

	err := fasthttp.Do(req, resp)

	if err != nil {
		return false
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return false
	}

	var result Response

	if err := json.Unmarshal([]byte(string(resp.Body())), &result); err != nil {
		return false
	}

	if result.Success {
		return true
	}

	return false
}
