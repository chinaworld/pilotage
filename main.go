package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/containerops/pilotage/cmd"
	"github.com/containerops/pilotage/setting"
)

func init() {
	//
}

func main() {
	app := cli.NewApp()

	app.Name = setting.AppName
	app.Usage = setting.Usage
	app.Version = setting.Version
	app.Author = setting.Author
	app.Email = setting.Email

	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdDatabase,
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
