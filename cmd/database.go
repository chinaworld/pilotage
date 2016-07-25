package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/containerops/pilotage/models"
)

var CmdDatabase = cli.Command{
	Name:        "database",
	Usage:       "database utils for backend database",
	Description: "Pilotage run base SQL database like MySQL, database command provide some utils of migrate, backup, config and so on.",
	Action:      runDatabase,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "action",
			Usage: "Action，[sync/backup/restore]",
		},
	},
}

func runDatabase(c *cli.Context) error {
	if len(c.String("action")) > 0 {
		action := c.String("action")

		switch action {
		case "sync":
			if err := models.Sync(); err != nil {
				fmt.Println("Init database struct error, ", err.Error())
			}
			break
		default:
			break
		}
	}

	return nil
}
