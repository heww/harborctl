package command

import (
	"github.com/urfave/cli"

	"github.com/heww/harborctl/version"
)

// App ...
func App() *cli.App {
	app := cli.NewApp()
	app.Name = "harborctl"
	app.Usage = "command line utility"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server,s",
			Usage:  "server address",
			EnvVar: "HARBOR_SERVER",
		},
		cli.StringFlag{
			Name:   "username,u",
			Usage:  "server username",
			EnvVar: "HARBOR_USERNAME",
		},
		cli.StringFlag{
			Name:   "password,p",
			Usage:  "server password",
			EnvVar: "HARBOR_PASSWORD",
		},
		cli.BoolFlag{
			Name:   "insecure-skip-tls-verify",
			Usage:  "server insecure",
			EnvVar: "HARBOR_INSECURE_SKIP_TLS_VERIFY",
		},
	}
	app.Commands = []cli.Command{
		ProjectCommand(),
		RepositoryCommand(),
		TagCommand(),
	}

	app.Action = func(c *cli.Context) error {
		if showHelp(c) {
			cli.ShowAppHelp(c)
		}

		return nil
	}

	return app
}
