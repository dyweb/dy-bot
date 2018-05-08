package gh

import (
	"context"
	"sync"

	"github.com/dyweb/gommon/util/logutil"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	gitHubClient *Client
	log          = logutil.NewPackageLogger()
	once         sync.Once
)

func InitGitHubClient(owner, repo, accessToken string) {
	once.Do(func() {
		gitHubClient = NewClient(owner, repo, accessToken)
	})
}

func GetGitHubClient() *Client {
	if gitHubClient == nil {
		log.Error("GitHubClient is not initialized")
	}
	return gitHubClient
}

// Client refers to a client which wishes to connect to specific repository of a user.
type Client struct {
	sync.Mutex
	*github.Client
	owner string
	repo  string
}

// NewClient constructs a new instance of Client.
func NewClient(owner, repo, token string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Client{
		Client: github.NewClient(tc),
		owner:  owner,
		repo:   repo,
	}
}

// Owner returns owner of client.
func (c *Client) Owner() string {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.owner
}

// Repo returns repo name of client.
func (c *Client) Repo() string {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.repo
}
