package argparser

import (
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

type ArgT struct {
	cli.Helper
	Res   string `cli:"r,res" usage:"name resurce for monitoring"`
	Name  string `cli:"n,name" usage:"name service"`
	Vhost string `cli:"v,vhost" usage:"rabbitmq vhost"`
	User  string `cli:"u,user" usage:"user name"`
	Pass  string `cli:"p,pass" usage:"password"`
	Port  string `cli:"port" usage:"tomcat port"`
}

func Parser(arg *ArgT) error {
	// argSpl := strings.Split(arg.Res, ".")
	// if len(argSpl) < 2 {
	// return fmt.Errorf("argSpl < 3")
	// }
	switch arg.Res {
	case "systemd.discover":
		systemd.Discaver()
	case "systemd.status":
		systemd.Status(arg.Name)
	case "rabbitmq.discover":
		rabbitmq.Discaver(arg.User, arg.Pass)
	case "rabbitmq.lenmessage":
		rabbitmq.LenMessage(arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "rabbitmq.redeliver":
		rabbitmq.RedeliverMessage(arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "rabbitmq.activeconsume":
		rabbitmq.ActiveConsumer(arg.Name, arg.Vhost, arg.User, arg.Pass)
	case "tomcat.discover":
		tomcat.Discover(arg.Port, arg.User, arg.Pass)
	case "tomcat.status":
		tomcat.Status(arg.Name, arg.Port, arg.User, arg.Pass)
	}
	return nil
}
