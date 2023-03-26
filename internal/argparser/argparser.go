package argparser

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/morhayn/zabbix-pack/internal/rabbitmq"
	"github.com/morhayn/zabbix-pack/internal/systemd"
	"github.com/morhayn/zabbix-pack/internal/tomcat"
)

// var (
// module = map[string]interface{}{
// "systemd.discover": interface{}
// }
// )
var (
	checkFun func(...string) error
)
// ArgT input parametres module
type ArgT struct {
	cli.Helper
	Res   string `cli:"r,res" usage:"name resurce for monitoring"`
	Name  string `cli:"n,name" usage:"name service"`
	Vhost string `cli:"v,vhost" usage:"rabbitmq vhost"`
	User  string `cli:"u,user" usage:"user name"`
	Pass  string `cli:"p,pass" usage:"password"`
	Port  string `cli:"port" usage:"tomcat port"`
}

// Parser call function for collect couners  
func Parser(arg *ArgT) error {
	// argSpl := strings.Split(arg.Res, ".")
	// if len(argSpl) < 2 {
	// return fmt.Errorf("argSpl < 3")
	// }
	var err error
	switch arg.Res {
	case "systemd.discover":
		err = systemd.Discover()
	case "systemd.status":
		err = systemd.Status(arg.Name)
	case "rabbitmq.discover":
		err = rabbitmq.Discover(arg.Port, arg.User, arg.Pass)
	case "rabbitmq.lenmessage":
		err = rabbitmq.LenMessage(arg.Port, arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "rabbitmq.redeliver":
		err = rabbitmq.RedeliverMessage(arg.Port, arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "rabbitmq.activeconsume":
		err = rabbitmq.ActiveConsumer(arg.Port, arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "tomcat.discover":
		err = tomcat.Discover(arg.Port, arg.User, arg.Pass)
	case "tomcat.status":
		err = tomcat.Status(arg.Name, arg.Port, arg.User, arg.Pass)
	}
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
