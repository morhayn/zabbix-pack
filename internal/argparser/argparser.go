package argparser

import (
	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
	Res     string `cli:"r,res" usage:"name resurce for monitoring"`
	Service string `cli:"s,service" usage:"service for monitoring"`
	Name    string `cli:"n,name" usage:"name service"`
}

func Parser() {

}
