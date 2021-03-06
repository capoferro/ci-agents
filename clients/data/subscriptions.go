package data

import (
	"context"

	"github.com/tinyci/ci-agents/errors"
	"github.com/tinyci/ci-agents/grpc/services/data"
	"github.com/tinyci/ci-agents/model"
)

// ListSubscriptions lists the subscriptions that the user has selected.
func (c *Client) ListSubscriptions(name string) (model.RepositoryList, *errors.Error) {
	rl, err := c.client.ListSubscriptions(context.Background(), &data.Name{Name: name})
	if err != nil {
		return nil, errors.New(err)
	}

	return makeRepoList(rl)
}

// AddSubscription adds a subscription for the user.
func (c *Client) AddSubscription(name, repo string) *errors.Error {
	_, err := c.client.AddSubscription(context.Background(), &data.RepoUserSelection{RepoName: repo, Username: name})
	if err != nil {
		return errors.New(err)
	}

	return nil
}

// DeleteSubscription removes a subscription for the user.
func (c *Client) DeleteSubscription(name, repo string) *errors.Error {
	// sigh.. these names.
	_, err := c.client.RemoveSubscription(context.Background(), &data.RepoUserSelection{RepoName: repo, Username: name})
	if err != nil {
		return errors.New(err)
	}

	return nil
}
