package command

import (
	"context"
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/urfave/cli"

	"github.com/heww/harborctl/pkg/harbor/client/project"
)

func ProjectCommand() cli.Command {
	return cli.Command{
		Name:    "project",
		Aliases: []string{"p"},
		Usage:   "manage projects",
		Subcommands: []cli.Command{
			ProjectListCommand(),
		},
		Action: func(c *cli.Context) error {
			if showHelp(c) {
				cli.ShowAppHelp(c)
			}

			return nil
		},
	}
}

func ProjectListCommand() cli.Command {
	return cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list projects",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name,n",
				Usage: "filter projects by, eg library",
			},
		},
		Action: func(c *cli.Context) error {
			api, err := harbor(c)
			if err != nil {
				return err
			}

			name := c.String("name")

			table := uitable.New()
			table.MaxColWidth = 80
			table.Wrap = true // wrap columns

			table.AddRow("NAME", "REPOSITORIES", "CHARTS", "CREATED")

			ctx := context.TODO()

			projectParams := project.ListProjectsParams{}
			if name != "" {
				projectParams.Name = &name
			}

			for p := range projectsIterator(ctx, api, projectParams) {
				table.AddRow(p.Name, p.RepoCount, p.ChartCount, p.CreationTime)
			}

			fmt.Println(table)

			return nil
		},
	}
}
