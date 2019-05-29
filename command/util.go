package command

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	rtclient "github.com/go-openapi/runtime/client"
	"github.com/urfave/cli"

	"github.com/heww/harborctl/pkg/harbor/client"
	"github.com/heww/harborctl/pkg/harbor/client/project"
	"github.com/heww/harborctl/pkg/harbor/client/repository"
	"github.com/heww/harborctl/pkg/harbor/models"
	"github.com/heww/harborctl/pkg/log"
	"github.com/heww/harborctl/pkg/util"
)

func showHelp(c *cli.Context) bool {
	return c.Command.FullName() == ""
}

func harbor(c *cli.Context) (*client.Harbor, error) {
	server := c.GlobalString("server")
	if server == "" {
		return nil, errors.New("server required")
	}
	username := c.GlobalString("username")
	if username == "" {
		return nil, errors.New("username required")
	}
	password := c.GlobalString("password")
	if password == "" {
		return nil, errors.New("password required")
	}

	if strings.HasSuffix(server, "/") {
		server = strings.TrimSuffix(server, "/")
	}

	_url, _ := url.Parse(server + client.DefaultBasePath)

	config := client.Config{
		URL:       _url,
		Transport: http.DefaultTransport,
		AuthInfo:  rtclient.BasicAuth(username, password),
	}

	insecureSkipTlsVerify := c.GlobalBool("insecure-skip-tls-verify")
	if insecureSkipTlsVerify {
		config.Transport = util.NewInsecureTransport()
	}

	return client.New(config), nil
}

func projectsIterator(ctx context.Context, h *client.Harbor, params project.ListProjectsParams) <-chan *models.Project {
	ch := make(chan *models.Project)

	go func() {

		page := int32(1)
		size := int32(10)

		for {
			params.Page = &page
			params.PageSize = &size

			resp, err := h.Project.ListProjects(ctx, &params)
			if err != nil {
				log.G(ctx).WithError(err).Info("harborctl: List projects failed")
				break
			}

			for _, p := range resp.Payload {
				ch <- p
			}

			if len(resp.Payload) < int(size) {
				break
			}

			page = page + 1
		}

		close(ch)
	}()

	return ch
}

func repositoriesIterator(ctx context.Context, h *client.Harbor, params repository.ListRepositoriesParams) <-chan *models.Repository {
	ch := make(chan *models.Repository)

	go func() {

		page := int32(1)
		size := int32(10)

		for {
			params.Page = &page
			params.PageSize = &size

			resp, err := h.Repository.ListRepositories(ctx, &params)
			if err != nil {
				log.G(ctx).WithError(err).Info("Harbor adapter: List repositories failed")
				break
			}

			for _, r := range resp.Payload {
				ch <- r
			}

			if len(resp.Payload) < int(size) {
				break
			}

			page = page + 1
		}

		close(ch)
	}()

	return ch
}

func hostname(server string) string {
	u, err := url.Parse(server)
	if err != nil {
		return ""
	}

	return u.Hostname()
}

func parseName(name string) (string, string) {
	parts := strings.SplitN(name, "/", 2)

	if len(parts) == 1 {
		return parts[0], ""
	} else if len(parts) == 2 {
		return parts[0], parts[1]
	} else {
		return "", ""
	}
}
