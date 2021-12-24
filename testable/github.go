package testable

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	StargazersCount int `json:"stargazers_count"`
}

type RepositoriesAPI interface {
	GetRepos(username string) ([]Repo, error)
}

type MockGithub struct{}

func (m *MockGithub) GetRepos(username string) ([]Repo, error) {
	return []Repo{
		Repo{
			StargazersCount: 5,
		},
		Repo{
			StargazersCount: 2,
		},
	}, nil
}

type Github struct{}

func (g *Github) GetRepos(username string) ([]Repo, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))

	if err != nil {
		return nil, err
	}

	var repos []Repo
	if err := json.NewDecoder(res.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}
