package services

import (
	"PROJET-GIT-GO/models"
	"encoding/json"
	"net/http"
)

func GetRepositories(user string, token string) ([]models.Repository, error) {
	url := "https://api.github.com/users/" + user + "/repos"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Add("Authorization", "token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []models.Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}
