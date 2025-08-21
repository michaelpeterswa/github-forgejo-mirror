package config

import "github.com/urfave/cli/v3"

const (
	GithubPAT = "gh-pat"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "gh-pat",
			Usage:    "GitHub Personal Access Token",
			Sources:  cli.EnvVars("GH_PAT"),
			Required: true,
		},
		&cli.StringFlag{
			Name:     "org",
			Usage:    "GitHub organization or user name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "forgejo-url",
			Usage:    "Forgejo instance URL",
			Sources:  cli.EnvVars("FORGEJO_URL"),
			Required: true,
		},
		&cli.StringFlag{
			Name:     "forgejo-token",
			Usage:    "Forgejo access token",
			Sources:  cli.EnvVars("FORGEJO_TOKEN"),
			Required: true,
		},
		&cli.StringFlag{
			Name:     "forgejo-user",
			Usage:    "Forgejo username or organization",
			Sources:  cli.EnvVars("FORGEJO_USER"),
			Required: true,
		},
	}
}
