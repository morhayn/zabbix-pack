package tomcat

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/morhayn/zabbix-pack/internal/getreq"
)

type TomcatWar struct {
	Name   string
	Status string
}

func TomcatParse(res *http.Response) ([]TomcatWar, error) {
	defer res.Body.Close()
	tomcat := []TomcatWar{}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []TomcatWar{}, err
	}
	lines := strings.Split(string(body), "\n")
	//First string in array "OK - Listed applications for virtual host localhost" its only status tomcat
	for _, line := range lines[1:] {
		arr_out := strings.Split(line, ":")
		if len(arr_out) > 1 {
			name_war := strings.TrimPrefix(strings.TrimSpace(arr_out[0]), "/")
			tomcat = append(tomcat, TomcatWar{
				Name:   name_war,
				Status: arr_out[1],
			})
		}
	}
	return tomcat, nil
}

func Discover(port, username, password string) error {
	result := make(map[string][]map[string]string)
	var res []map[string]string
	client, err := getreq.NewClient("http://127.0.0.1:8080", username, password)
	if err != nil {
		return err
	}
	client.SetTimeout(2 * time.Second)
	req, err := getreq.NewGETRequest(client, "manage/text/list")
	if err != nil {
		return err
	}
	response, err := getreq.ExecuteRequest(client, req)
	listTomcat, err := TomcatParse(response)
	if err != nil {
		return err
	}
	for _, warfile := range listTomcat {
		res = append(res, map[string]string{
			"{#WAR.NAME}":   warfile.Name,
			"{#WAR.STATUS}": warfile.Status,
		})
	}
	result["data"] = res
	out, err := json.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return nil
}
func Status() {}
