package argparser

import (
	"github.com/mkideal/cli"
	"github.com/morhayn/zabbix-pack/internal/systemd"
)

type ArgT struct {
	cli.Helper
	Res     string `cli:"r,res" usage:"name resurce for monitoring"`
	Service string `cli:"s,service" usage:"service for monitoring"`
	Name    string `cli:"n,name" usage:"name service"`
}

func Parser(arg *ArgT) {
	if arg.Res == "systemd.discover" {
		systemd.Discaver()
	}
	if arg.Res == "systemd.status" {
		systemd.Status(arg.Name)
	}
}
