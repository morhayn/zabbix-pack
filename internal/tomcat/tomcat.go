package tomcat

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/morhayn/zabbix-pack/internal/getreq"
)

type TomcatWar struct {
	Name   string
	Status string
}

func TomcatParse(url, username, password, port string) ([]TomcatWar, error) {
	tomcat := []TomcatWar{}
	client, err := getreq.NewClient("http://127.0.0.1:"+port, username, password)
	if err != nil {
		return tomcat, err
	}
	client.SetTimeout(2 * time.Second)
	req, err := getreq.NewGETRequest(client, "manager/text/list")
	if err != nil {
		return tomcat, err
	}
	res, err := getreq.ExecuteRequest(client, req)
	if err != nil {
		return tomcat, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return tomcat, err
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
	listTomcat, err := TomcatParse("manager/text/list", username, password, port)
	if err != nil {
		return err
	}
	for _, warfile := range listTomcat {
		if (warfile.Name != "manager") && (warfile.Name != "host-manager") {
			res = append(res, map[string]string{
				"{#WAR.NAME}":   warfile.Name,
				"{#WAR.STATUS}": warfile.Status,
			})
		}
	}
	result["data"] = res
	out, err := json.Marshal(result)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out)
	return nil
}
func Status(warname, port, username, password string) error {
	listTomcat, err := TomcatParse("manager/text/list", username, password, port)
	if err != nil {
		fmt.Print("0")
		return err
	}
	for _, warfile := range listTomcat {
		if warfile.Name == warname {
			if warfile.Status == "running" {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
			break
		}
	}
	return nil
}
