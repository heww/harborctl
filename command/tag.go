package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/urfave/cli"

	"github.com/heww/harborctl/pkg/harbor/client/project"
	"github.com/heww/harborctl/pkg/harbor/client/repository"
	"github.com/heww/harborctl/pkg/log"
)

func TagCommand() cli.Command {
	return cli.Command{
		Name:    "tag",
		Aliases: []string{"t"},
		Usage:   "manage tags",
		Subcommands: []cli.Command{
			TagListCommand(),
		},
		Action: func(c *cli.Context) error {
			if showHelp(c) {
				cli.ShowAppHelp(c)
			}

			return nil
		},
	}
}

func TagListCommand() cli.Command {
	return cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list tags",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name,n",
				Usage: "list tags of this name, eg library or library/photon, ignore when all flag set",
			},
			cli.BoolFlag{
				Name:  "all,a",
				Usage: "list all tags",
			},
		},
		Action: func(c *cli.Context) error {
			api, err := harbor(c)
			if err != nil {
				return err
			}

			all := c.Bool("all")

			name := c.String("name")
			if name == "" && !all {
				return errors.New("name required")
			}

			registry := hostname(c.GlobalString("server"))

			projectName, repositoryName := parseName(name)
			if all {
				projectName = ""
				repositoryName = ""
			}

			table := uitable.New()
			table.MaxColWidth = 80
			table.Wrap = true // wrap columns

			table.AddRow("REPOSITORY", "TAG", "DIGEST", "CREATED")

			ctx := context.TODO()

			projectParams := project.ListProjectsParams{}
			if !all {
				projectParams.Name = &projectName
			}

			for p := range projectsIterator(ctx, api, projectParams) {
				repositoryParams := repository.ListRepositoriesParams{
					ProjectID: p.ProjectID,
				}

				if repositoryName != "" {
					repositoryParams.Q = &repositoryName
				}

				for r := range repositoriesIterator(ctx, api, repositoryParams) {
					_, n := parseName(r.Name)

					resp, err := api.Repository.ListRepositoryTags(ctx, &repository.ListRepositoryTagsParams{
						ProjectName:    p.Name,
						RepositoryName: n,
					})

					if err != nil {
						log.G(ctx).WithError(err).Info("list repository tags failed")
						continue
					}

					for _, tag := range resp.Payload {
						table.AddRow(fmt.Sprintf("%s/%s", registry, r.Name), tag.Name, tag.Digest, tag.Created)
					}
				}
			}

			fmt.Println(table)

			return nil
		},
	}
}
