# github-forgejo-mirror

A Go CLI tool to mirror GitHub repositories to Forgejo instances.

## Features

- Mirror all repositories from a GitHub organization or user to a Forgejo instance
- Preserves repository privacy settings and descriptions
- Handles pagination for organizations with many repositories
- Creates proper Git mirrors with automatic synchronization

## Usage

### Prerequisites

- Go 1.23.0 or later
- GitHub Personal Access Token with repository access
- Forgejo instance with API access token

### Installation

```bash
go build -o github-forgejo-mirror cmd/github-forgejo-mirror/main.go
```

### Running

```bash
./github-forgejo-mirror \
  --gh-pat "your-github-token" \
  --org "github-username-or-org" \
  --forgejo-url "https://your-forgejo-instance.com" \
  --forgejo-token "your-forgejo-token" \
  --forgejo-user "forgejo-username-or-org"
```

### Environment Variables

You can also set the following environment variables instead of using flags:

- `GH_PAT` - GitHub Personal Access Token
- `FORGEJO_URL` - Forgejo instance URL  
- `FORGEJO_TOKEN` - Forgejo access token
- `FORGEJO_USER` - Forgejo username or organization

### Command Line Flags

- `--gh-pat` - GitHub Personal Access Token (required)
- `--org` - GitHub organization or user name (required)
- `--forgejo-url` - Forgejo instance URL (required)
- `--forgejo-token` - Forgejo access token (required)
- `--forgejo-user` - Forgejo username or organization (required)

## Development

### Setup

```bash
make all
```

This will install commitlint and configure git hooks for conventional commits.

## License

See [LICENSE](LICENSE) file.
