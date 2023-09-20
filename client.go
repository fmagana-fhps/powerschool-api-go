package ps

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	m "github.com/fmagana-fhps/powerschool-api-go/models"
)

type Client struct {
	ctx  context.Context
	opts *Options
}

type Options struct {
	BaseURL      string
	ClientId     string
	ClientSecret string
	Access       *m.AccessTokenResponse
	HTTPClient   *http.Client
}

func NewClient(host string, options *Options) (*Client, error) {
	if options.HTTPClient == nil {
		options.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}

	if host == "" {
		return nil, errors.New("host has not been provided but required")
	}

	options.BaseURL = fmt.Sprintf("https://%s", host)

	if options.ClientId == "" || options.ClientSecret == "" {
		return nil, errors.New("credentials have not been provided but required")
	}

	token, err := options.verifyAccessToken()
	if err != nil {
		return nil, errors.New("unable to retrieve access token; " + err.Error())
	}

	options.Access = token

	client := &Client{
		ctx:  context.Background(),
		opts: options,
	}

	return client, nil
}

func (o *Options) verifyAccessToken() (*m.AccessTokenResponse, error) {
	if !o.isTokenExpired() {
		return o.Access, nil
	}

	data := []byte(o.ClientId + ":" + o.ClientSecret)
	enc := base64.StdEncoding.EncodeToString(data)

	url := o.BaseURL + "/oauth/access_token/"
	reader := strings.NewReader(`grant_type=client_credentials`)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+enc)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	atr := o.Access
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, atr)
	if err != nil {
		return nil, err
	}

	i, err := strconv.Atoi(atr.LifeLength)
	if err != nil {
		return nil, err
	}
	atr.ExpiresAt = time.Now().Add(time.Second * time.Duration(i)).Unix()

	return atr, nil
}

func (o *Options) isTokenExpired() bool {
	return o.Access.ExpiresAt <= time.Now().Unix()
}

func (c *Client) request(method, url string, data io.Reader, respData interface{}) (*m.Response, error) {
	c.opts.verifyAccessToken()
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+c.opts.Access.AccessToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	resp, err := c.opts.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	br := &m.Response{}
	br.Header = resp.Header
	br.StatusCode = resp.StatusCode
	br.Body = respData

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return br, nil
	}

	return br, json.Unmarshal(body, &br.Body)
}

func (c *Client) do(method, endpoint string, reqData, respData interface{}) (*m.Response, error) {
	url := c.opts.BaseURL + endpoint

	if reqData != nil {
		b, err := json.Marshal(reqData)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(b))
		return c.request(method, url, strings.NewReader(string(b)), respData)
	}

	return c.request(method, url, nil, respData)
}

func (c *Client) get(endpoint string, respData interface{}) (*m.Response, error) {
	return c.do(http.MethodGet, endpoint, nil, respData)
}

func (c *Client) put(endpoint string, reqData, respData interface{}) (*m.Response, error) {
	return c.do(http.MethodPost, endpoint, reqData, respData)
}

func (c *Client) post(endpoint string, reqData, respData interface{}) (*m.Response, error) {
	return c.do(http.MethodPost, endpoint, reqData, respData)
}

func (c *Client) delete(endpoint string, reqData interface{}) (*m.Response, error) {
	return c.do(http.MethodDelete, endpoint, reqData, m.Response{})
}
