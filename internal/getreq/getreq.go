package getreq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Client for interaction with RabbitMQ HTTP API.
type Client struct {
	// URI of a RabbitMQ node to use, not including the path, e.g. http://127.0.0.1:15672.
	Endpoint string
	// Username to use. This RabbitMQ user must have the "management" tag.
	Username string
	// Password to use.
	Password  string
	Host      string
	Transport http.RoundTripper
	Timeout   time.Duration
}

// ErrorResponse represents an error reported by an API response.
type ErrorResponse struct {
	StatusCode int
	Message    string `json:"error"`
	Reason     string `json:"reason"`
}

func (rme ErrorResponse) Error() string {
	return fmt.Sprintf("Error %d (%s): %s", rme.StatusCode, rme.Message, rme.Reason)
}

// NewClient instantiates a client.
func NewClient(uri string, username string, password string) (me *Client, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	me = &Client{
		Endpoint: uri,
		Host:     u.Host,
		Username: username,
		Password: password,
	}
	return me, nil
}

// SetTimeout changes the HTTP timeout that the Client will use.
// By default there is no timeout.
func (c *Client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

func NewGETRequest(client *Client, path string) (*http.Request, error) {
	s := client.Endpoint + "/api/" + path
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.SetBasicAuth(client.Username, client.Password)

	return req, err
}

func ExecuteRequest(client *Client, req *http.Request) (resp *http.Response, err error) {
	httpc := &http.Client{
		Timeout: client.Timeout,
	}
	if client.Transport != nil {
		httpc.Transport = client.Transport
	}
	resp, err = httpc.Do(req)
	if err != nil {
		return nil, err
	}
	if err = parseResponseErrors(resp); err != nil {
		if resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	return resp, err
}

func ExecuteAndParseRequest(client *Client, req *http.Request, rec interface{}) (err error) {
	res, err := ExecuteRequest(client, req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&rec); err != nil {
		return err
	}
	return nil
}

func parseResponseErrors(res *http.Response) (err error) {
	if res.StatusCode == http.StatusUnauthorized {
		return errors.New("Error: API responded with a 401 Unauthorized")
	}
	// handle a "404 Not Found" response for a DELETE request as success.
	if res.Request.Method == http.MethodDelete && res.StatusCode == http.StatusNotFound {
		return nil
	}
	if res.StatusCode >= http.StatusBadRequest {
		rme := ErrorResponse{}
		if err = json.NewDecoder(res.Body).Decode(&rme); err != nil {
			rme.Message = fmt.Sprintf("Error %d from RabbitMQ: %s", res.StatusCode, err)
		}
		rme.StatusCode = res.StatusCode
		return rme
	}
	return nil
}
