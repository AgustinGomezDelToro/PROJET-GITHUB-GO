package controllers

import (
	"PROJET-GIT-GO/models"
	"PROJET-GIT-GO/services"
	"PROJET-GIT-GO/utils"
	"fmt"
	"time"
)

type ByLastUpdated []models.Repository

func (a ByLastUpdated) Len() int           { return len(a) }
func (a ByLastUpdated) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastUpdated) Less(i, j int) bool { return a[i].LastUpdatedTime.After(a[j].LastUpdatedTime) }

func GetAndCloneRepositories(user string, token string) error {
	repos, err := services.GetRepositories(user, token)
	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Printf("Clonage du dépôt : %s, URL : %s\n", repo.Name, repo.CloneURL)

		t, err := time.Parse(time.RFC3339, repo.LastUpdated)
		if err != nil {
			fmt.Printf("La date de la dernière mise à jour est vide pour le dépôt : %s\n", repo.Name)
			continue
		}
		repo.LastUpdatedTime = t

		err = utils.CloneRepo(repo.CloneURL, "./clonedRepos/"+repo.Name)
		if err != nil {
			fmt.Printf("Erreur lors du clonage du dépôt %s : %v\n", repo.Name, err)
			continue
		}
	}

	err = utils.WriteCSV(repos, "repositories.csv")
	if err != nil {
		return err
	}

	return nil
}
