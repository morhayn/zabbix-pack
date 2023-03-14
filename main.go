package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/mkideal/cli"
)

// systemctl -t service -o json  --no-pager

type argT struct {
	cli.Helper
	Res     string `cli:"r,res" usage:"name resurce for monitoring"`
	Service string `cli:"s,service" usage:"service for monitoring"`
	Name    string `cli:"n,name" usage:"name service"`
}
type Systemd struct {
	Data []map[string]string
	// Unit        string `json:"unit"`
	// Load        string `json:"load"`
	// Active      string `json:"active"`
	// Sub         string `json:"sub"`
	// Description string `json:"description"`
	// }
}

func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fmt.Println(argv, argv.Help, argv.Helper)
		fmt.Println(argv.Res)
		fmt.Println(argv.Service)
		out, err := exec.Command("/usr/bin/systemctl", "-t", "service", "-o", "json", "--no-page").Output()
		exec.Command("/usr/bin/systemctl", "is-active", service)
		if err != nil {
			fmt.Println(err)
		}
		system := Systemd{}
		err = json.Unmarshal([]byte(out), &system.Data)
		for _, service := range system.Data {
			fmt.Println(service["unit"])
		}
		return nil
	})
}
