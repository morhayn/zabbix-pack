package rabbitmq

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/morhayn/zabbix-pack/internal/getreq"
)

// NewTLSClient instantiates a client with a transport; it is up to the developer to make that layer secure.
// func NewTLSClient(uri string, username string, password string, transport http.RoundTripper) (me *Client, err error) {
// u, err := url.Parse(uri)
// if err != nil {
// return nil, err
// }

// me = &Client{
// Endpoint:  uri,
// host:      u.Host,
// Username:  username,
// Password:  password,
// transport: transport,
// }

// return me, nil
// }

// SetTransport changes the Transport Layer that the Client will use.
// func (c *Client) SetTransport(transport http.RoundTripper) {
// c.transport = transport
// }

// func newGETRequestWithParameters(client *Client, path string, qs url.Values) (*http.Request, error) {
// s := client.Endpoint + "/api/" + path + "?" + qs.Encode()

// req, err := http.NewRequest("GET", s, nil)
// if err != nil {
// return nil, err
// }

// req.Close = true
// req.SetBasicAuth(client.Username, client.Password)

// return req, err
// }

// func newRequestWithBody(client *Client, method string, path string, body []byte) (*http.Request, error) {
// s := client.Endpoint + "/api/" + path

// req, err := http.NewRequest(method, s, bytes.NewReader(body))
// if err != nil {
// return nil, err
// }

// req.Close = true
// req.SetBasicAuth(client.Username, client.Password)

// req.Header.Add("Content-Type", "application/json")

// return req, err
// }

// ListQueues lists all queues in the cluster. This only includes queues in the
// virtual hosts accessible to the user.
func listQueues(c *getreq.Client) (rec []QueueInfo, err error) {
	req, err := getreq.NewGETRequest(c, "api/queues")
	if err != nil {
		return []QueueInfo{}, err
	}
	if err = getreq.ExecuteAndParseRequest(c, req, &rec); err != nil {
		return []QueueInfo{}, err
	}
	return rec, nil
}

// GetQueue returns information about a queue.
func getQueue(c *getreq.Client, vhost, queue string) (rec *QueueInfo, err error) {
	req, err := getreq.NewGETRequest(c, "api/queues/"+url.PathEscape(vhost)+"/"+url.PathEscape(queue))
	if err != nil {
		return nil, err
	}
	if err = getreq.ExecuteAndParseRequest(c, req, &rec); err != nil {
		return nil, err
	}
	return rec, nil
}
func makeRabbitMQClient(dsn string, username string, password string, timeout time.Duration) (*getreq.Client, error) {
	var (
		rmqc *getreq.Client
		err  error
	)
	rmqc, err = getreq.NewClient(dsn, username, password)
	if err != nil {
		return nil, err
	}
	rmqc.SetTimeout(timeout)
	return rmqc, nil
}
func newQueue(q QueueInfo) map[string]string {
	var res = make(map[string]string)
	res["{#QUEUENAME}"] = q.Name
	res["{#VHOST}"] = q.Vhost
	return res
}
func Discover(port, user, pass string) error {
	result := make(map[string][]map[string]string)
	var res []map[string]string
	client, err := makeRabbitMQClient("http://127.0.0.1:"+port, user, pass, 2*time.Second)
	if err != nil {
		return err
	}
	listQ, err := listQueues(client)
	if err != nil {
		return err
	}
	for _, queue := range listQ {
		res = append(res, newQueue(queue))
	}
	result["data"] = res
	out, err := json.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return nil
}
func LenMessage(port, queue, vhost, user, pass string) error {
	client, err := makeRabbitMQClient("http://127.0.0.1:"+port, user, pass, 2*time.Second)
	if err != nil {
		return err
	}
	q, err := getQueue(client, vhost, queue)
	if err != nil {
		return err
	}
	fmt.Println(q.Messages)
	return nil
}
func RedeliverMessage(port, queue, vhost, user, pass string) error {
	client, err := makeRabbitMQClient("http://127.0.0.1:"+port, user, pass, 2*time.Second)
	if err != nil {
		return err
	}
	q, err := getQueue(client, vhost, queue)
	if err != nil {
		return err
	}
	fmt.Println(q.MessageStats.Redeliver)
	return nil
}
func ActiveConsumer(port, queue, vhost, user, pass string) error {
	client, err := makeRabbitMQClient("http://127.0.0.1:"+port, user, pass, 2*time.Second)
	if err != nil {
		return err
	}
	q, err := getQueue(client, vhost, queue)
	if err != nil {
		return err
	}
	fmt.Println(q.MessageStats.Redeliver)
	return nil
}
