package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var build = "0"

func main() {
	app := cli.NewApp()
	app.Name = "git-with-ssh plugin"
	app.Usage = "git-with-ssh plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ssh_private_key",
			Usage:  "SSH private key for git",
			EnvVar: "PLUGIN_SSH_PRIVATE_KEY",
		},
		cli.StringSliceFlag{
			Name:   "ssh_hosts",
			Usage:  "SSH hosts by git with ssh (disable StrictHostKeyChecking)",
			EnvVar: "PLUGIN_SSH_HOSTS",
		},
		cli.StringSliceFlag{
			Name:   "commands",
			Usage:  "shell commands",
			EnvVar: "PLUGIN_COMMANDS",
		},
		cli.StringFlag{
			Name:   "home",
			Usage:  "home directory for ssh",
			EnvVar: "PLUGIN_HOME",
			Value:  "/root",
		},

		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Home:     c.String("home"),
		SSHKey:   strings.Replace(c.String("ssh_private_key"), `\n`, "\n", -1),
		Hosts:    c.StringSlice("ssh_hosts"),
		Commands: c.StringSlice("commands"),
	}

	return plugin.Exec()
}
