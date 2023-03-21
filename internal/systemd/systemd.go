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
		"tomcat1.service",
		"rabbitmq-server.service",
		"postgresql.service",
		"postgresql@9.6-main.service",
		"mongod.service",
		"elasticsearch.service",
		"nginx.service",
		"haproxy.service",
		"kafka.service",
		"zookeeper.service",
		"hazelcast.service",
	}
)

func newRes(name, descr string) map[string]string {
	var r = make(map[string]string)
	n := strings.TrimSuffix(name, ".service")
	r["{#UNIT.NAME}"] = n
	r["{#UNIT.DESCRIPTION}"] = descr
	return r
}
func Discover() error {
	// serv := make(map[string]string)
	result := make(map[string][]map[string]string)
	var res []map[string]string
	o, err := exec.Command("/bin/systemctl", "-t", "service", "--state", "active", "--no-legend", "--no-page").Output()
	if err != nil {
		fmt.Println(err)
		return err
	}
	l := strings.Split(string(o), "\n")
	system := make(map[string]string)
	for _, str := range l {
		field := (strings.Fields(str))
		if len(field) > 3 {
			system[field[0]] = field[3]
		}
	}
	for _, v := range monServ {
		if r, ok := system[v]; ok {
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
	o, err := exec.Command("/bin/systemctl", "is-active", service).Output()
	if err != nil {
		fmt.Print("0")
		return err
	}
	if strings.TrimSpace(string(o)) == "active" {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
	return nil
}
