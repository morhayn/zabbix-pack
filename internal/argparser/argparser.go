package argparser

import (
	"fmt"
	"strings"

	"github.com/mkideal/cli"
	"github.com/morhayn/zabbix-pack/internal/rabbitmq"
	"github.com/morhayn/zabbix-pack/internal/systemd"
)

type ArgT struct {
	cli.Helper
	Res   string `cli:"r,res" usage:"name resurce for monitoring"`
	Name  string `cli:"n,name" usage:"name service"`
	Vhost string `cli:"v,vhost" usage:"rabbitmq vhost"`
}

func Parser(arg *ArgT) error {
	argSpl := strings.Split(arg.Res, ".")
	if len(argSpl) < 2 {
		return fmt.Errorf("argSpl < 3")
	}
	switch argSpl[0] {
	case "systemd":
		if argSpl[1] == "discover" {
			systemd.Discaver()
		}
		if argSpl[1] == "status" {
			systemd.Status(arg.Name)
		}
	case "rabbitmq":
		if argSpl[1] == "discover" {
			rabbitmq.Discaver()
		}
		if argSpl[1] == "lenmessage" {
			rabbitmq.LenMessage(arg.Name, arg.Vhost)
		}
		if argSpl[1] == "redeliver" {
			rabbitmq.RedeliverMessage(arg.Name, arg.Vhost)
		}
		if argSpl[1] == "activeconsume" {
			rabbitmq.ActiveConsumer(arg.Name, arg.Vhost)
		}
	}
	return nil
}
