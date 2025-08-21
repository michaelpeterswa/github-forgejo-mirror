package runner

import (
	"context"
	"fmt"

	"codeberg.org/mvdkleijn/forgejo-sdk/forgejo/v2"
	"github.com/google/go-github/v74/github"
	"github.com/urfave/cli/v3"
)

type Runner struct {
	cliCmd *cli.Command
}

func NewRunner(c *cli.Command) *Runner {
	return &Runner{
		cliCmd: c,
	}
}

func (r *Runner) Run(ctx context.Context) error {
	ghPat := r.cliCmd.String("gh-pat")
	org := r.cliCmd.String("org")
	forgejoURL := r.cliCmd.String("forgejo-url")
	forgejoToken := r.cliCmd.String("forgejo-token")

	// Initialize GitHub client
	githubClient := github.NewClient(nil).WithAuthToken(ghPat)

	// Initialize Forgejo client
	forgejoClient, err := forgejo.NewClient(forgejoURL, forgejo.SetToken(forgejoToken))
	if err != nil {
		return fmt.Errorf("failed to create Forgejo client: %w", err)
	}

	var allRepos []*github.Repository
	opts := &github.RepositoryListByUserOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	for {
		repos, resp, err := githubClient.Repositories.ListByUser(ctx, org, opts)
		if err != nil {
			return fmt.Errorf("failed to list repositories for %s: %w", org, err)
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	fmt.Printf("Found %d repositories for %s:\n", len(allRepos), org)

	// Create mirrors for each repository
	for _, repo := range allRepos {
		fmt.Printf("Creating mirror for %s...\n", repo.GetName())

		mirrorOpts := forgejo.MigrateRepoOption{
			RepoName:     repo.GetName(),
			CloneAddr:    repo.GetCloneURL(),
			Service:      forgejo.GitServiceGithub,
			AuthUsername: org,
			AuthToken:    ghPat,
			Mirror:       true,
			Private:      repo.GetPrivate(),
			Description:  repo.GetDescription(),
		}

		_, _, err := forgejoClient.MigrateRepo(mirrorOpts)
		if err != nil {
			fmt.Printf("  ❌ Failed to create mirror for %s: %v\n", repo.GetName(), err)
			continue
		}

		fmt.Printf("  ✅ Successfully created mirror for %s\n", repo.GetName())
	}

	return nil
}
