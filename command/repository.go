package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/urfave/cli"

	"github.com/heww/harborctl/pkg/harbor/client/project"
	"github.com/heww/harborctl/pkg/harbor/client/repository"
)

func RepositoryCommand() cli.Command {
	return cli.Command{
		Name:    "repository",
		Aliases: []string{"repo"},
		Usage:   "manage repositories",
		Subcommands: []cli.Command{
			RepositoryListCommand(),
		},
		Action: func(c *cli.Context) error {
			if showHelp(c) {
				cli.ShowAppHelp(c)
			}

			return nil
		},
	}
}

func RepositoryListCommand() cli.Command {
	return cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list repositories",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "project,p",
				Usage: "list repositories of this name, eg library, ignore when all flag set",
			},
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "list all repositories",
			},
		},
		Action: func(c *cli.Context) error {
			api, err := harbor(c)
			if err != nil {
				return err
			}

			all := c.Bool("all")

			projectName := c.String("project")
			if projectName == "" && !all {
				return errors.New("project required")
			}

			if all {
				projectName = ""
			}

			table := uitable.New()
			table.MaxColWidth = 80
			table.Wrap = true // wrap columns

			table.AddRow("REPOSITORY", "TAGS", "PULLS", "STARS", "CREATED")

			ctx := context.TODO()

			projectParams := project.ListProjectsParams{}
			if !all {
				projectParams.Name = &projectName
			}

			for p := range projectsIterator(ctx, api, projectParams) {
				params := repository.ListRepositoriesParams{
					ProjectID: p.ProjectID,
				}

				for r := range repositoriesIterator(ctx, api, params) {
					table.AddRow(r.Name, r.TagsCount, r.PullCount, r.StarCount, r.CreationTime)
				}
			}

			fmt.Println(table)

			return nil
		},
	}
}
