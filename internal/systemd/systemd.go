package systemd

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

var (
	monServ = []string{
		"tomcat8.service",
		"rabbit-mq.service",
		"postgres.service",
		"mongod.service",
		"elasticsearch.service",
		"nginx.service",
	}
)

type Systemd struct {
	Data []map[string]string
}

func newRes(name, descr string) map[string]string {
	var r = make(map[string]string)
	n := strings.TrimSuffix(name, ".service")
	r["{#UNIT.NAME}"] = n
	r["{#UNIT.DESCRIPTION}"] = descr
	return r
}
func Discaver() error {
	serv := make(map[string]string)
	result := make(map[string][]map[string]string)
	var res []map[string]string
	o, err := exec.Command("/usr/bin/systemctl", "-t", "service", "-o", "json", "--no-page").Output()
	if err != nil {
		return err
	}
	system := Systemd{}
	err = json.Unmarshal([]byte(o), &system.Data)
	for _, service := range system.Data {
		serv[service["unit"]] = service["desciption"]
	}
	for _, v := range monServ {
		if r, ok := serv[v]; ok {
			res = append(res, newRes(v, r))
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
func Status(service string) error {
	o, err := exec.Command("/usr/bin/systemctl", "is-active", service).Output()
	if err != nil {
		return err
	}
	if strings.TrimSpace(string(o)) == "active" {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
	return nil
}
