package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// AuthResponse -
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

// NewClient -
func NewClient(ctx context.Context, host, username, password *string) (Client, error) {
	c := client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	data := []byte(fmt.Sprintf("%s:%s", *username, *password))
	token := base64.StdEncoding.EncodeToString(data)

	// get token
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v2/token", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	body, _, err := c.doRequest(ctx, req)

	// parse response body
	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		panic(string(body))
		return nil, err
	}

	c.Token = fmt.Sprintf("Bearer %s", ar.AccessToken)
	return &c, nil
}

func (c *client) doRequest(ctx context.Context, req *http.Request) ([]byte, *http.Response, error) {
	if len(c.Token) > 0 {
		req.Header.Set("Authorization", c.Token)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	return body, res, err
}
