// Package requests wrap net/http like requests did for python
// it is easy to use, but not very efficient
package requests

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dyweb/gommon/errors"
)

const (
	ShadowSocksLocal = "127.0.0.1:1080"
	ContentJSON      = "application/json"
)

var defaultClient = &Client{h: NewDefaultClient(), content: ContentJSON}

type Client struct {
	h       *http.Client
	content string
}

func NewClient(options ...func(h *http.Client)) *Client {
	c := &Client{h: NewDefaultClient(), content: ContentJSON}
	for _, option := range options {
		option(c.h)
	}
	return c
}

func (c *Client) makeRequest(method string, url string, body io.Reader) (*Response, error) {
	if c.h == nil {
		return nil, errors.New("http client is not initialized")
	}
	var res *http.Response
	var err error
	switch method {
	case http.MethodGet:
		res, err = c.h.Get(url)
	case http.MethodPost:
		res, err = c.h.Post(url, c.content, body)
	}
	response := &Response{}
	if err != nil {
		return response, errors.Wrap(err, "error making request")
	}
	defer res.Body.Close()
	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, errors.Wrap(err, "error reading body")
	}
	response.Res = res
	response.StatusCode = res.StatusCode
	response.Data = resContent
	return response, nil
}

// TODO: this defaultish wrapping methods should be generated by gommon generator
func Post(url string, data io.Reader) (*Response, error) {
	return defaultClient.Post(url, data)
}

func (c *Client) Post(url string, data io.Reader) (*Response, error) {
	return c.makeRequest(http.MethodPost, url, data)
}

func PostString(url string, data string) (*Response, error) {
	return defaultClient.PostString(url, data)
}

func (c *Client) PostString(url string, data string) (*Response, error) {
	return c.makeRequest(http.MethodPost, url, ioutil.NopCloser(strings.NewReader(data)))
}

func PostBytes(url string, data []byte) (*Response, error) {
	return defaultClient.PostBytes(url, data)
}

func (c *Client) PostBytes(url string, data []byte) (*Response, error) {
	return c.makeRequest(http.MethodPost, url, ioutil.NopCloser(bytes.NewReader(data)))
}

func Get(url string) (*Response, error) {
	return defaultClient.Get(url)
}

func (c *Client) Get(url string) (*Response, error) {
	return c.makeRequest(http.MethodGet, url, nil)
}

func GetJSON(url string, data interface{}) error {
	return defaultClient.GetJSON(url, data)
}

func (c *Client) GetJSON(url string, data interface{}) error {
	res, err := c.Get(url)
	if err != nil {
		return errors.Wrap(err, "error getting response")
	}
	err = res.JSON(data)
	if err != nil {
		return errors.Wrap(err, "error parsing json")
	}
	return nil
}

func GetJSONStringMap(url string) (map[string]string, error) {
	return defaultClient.GetJSONStringMap(url)
}

func (c *Client) GetJSONStringMap(url string) (map[string]string, error) {
	var data map[string]string
	res, err := c.Get(url)
	if err != nil {
		return data, errors.Wrap(err, "error getting response")
	}
	data, err = res.JSONStringMap()
	if err != nil {
		return data, errors.Wrap(err, "error parsing json")
	}
	return data, nil
}
