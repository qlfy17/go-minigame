package minigame

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	libVersion     = "0.1.0"
	defaultBaseURL = "https://api.weixin.qq.com"
	userAgent      = "go-minigame/" + libVersion
)

// Client 负责与微信小游戏API进行通信
type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	AccessToken *AccessTokenService
}

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:    client,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}

	c.AccessToken = &AccessTokenService{client: c}

	return c
}

func (c *Client) newRequest(method, apiURL string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	url_ := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url_.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) Get(apiURL string, body interface{}) (*http.Request, error) {
	return c.newRequest("get", apiURL, body)
}

func (c *Client) Post(apiURL string, body interface{}) (*http.Request, error) {
	return c.newRequest("post", apiURL, body)
}

func (c *Client) Put(apiURL string, body interface{}) (*http.Request, error) {
	return c.newRequest("put", apiURL, body)
}

func (c *Client) Delete(apiURL string, body interface{}) (*http.Request, error) {
	return c.newRequest("delete", apiURL, body)
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
