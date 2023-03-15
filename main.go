package main

import (
	"github.com/mkideal/cli"
	"github.com/morhayn/zabbix-pack/internal/argparser"
)

// systemctl -t service -o json  --no-pager

// Unit        string `json:"unit"`
// Load        string `json:"load"`
// Active      string `json:"active"`
// Sub         string `json:"sub"`
// Description string `json:"description"`
// }
// }

func main() {
	cli.Run(new(argparser.ArgT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argparser.ArgT)
		argparser.Parser(argv)
		return nil
	})
}
