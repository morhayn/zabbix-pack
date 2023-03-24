package rabbitmq

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/morhayn/zabbix-pack/internal/getreq"
)

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

// Discover return all queue from rabbitmq-server
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

// LenMessage return count message in queue
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

// RedeliverMessage return count redeliver messages
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

// ActiveConsumer return count consumer for queue
func ActiveConsumer(port, queue, vhost, user, pass string) error {
	client, err := makeRabbitMQClient("http://127.0.0.1:"+port, user, pass, 2*time.Second)
	if err != nil {
		return err
	}
	q, err := getQueue(client, vhost, queue)
	if err != nil {
		return err
	}
	fmt.Println(q.Consumers)
	return nil
}
