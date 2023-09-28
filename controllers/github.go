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

	var directories []string

	cloneCount := 0
	for _, repo := range repos {
		if cloneCount >= 5 {
			break
		}

		fmt.Printf("\nClonage du dépôt : %s, URL : %s\n", repo.Name, repo.CloneURL)

		t, err := time.Parse(time.RFC3339, repo.LastUpdated)
		if err != nil {
			fmt.Printf("La date de la dernière mise à jour est vide pour le dépôt : %s\n", repo.Name)
			continue
		}
		repo.LastUpdatedTime = t

		directory := "./clonedRepos/" + repo.Name
		directories = append(directories, directory)

		err = utils.CloneRepo(repo.CloneURL, directory)
		if err != nil {
			fmt.Printf("Erreur lors du clonage du dépôt %s : %v\n", repo.Name, err)
			continue
		}

		err = utils.UpdateRepo(directory)
		if err != nil {
			fmt.Printf("Erreur lors de la mise à jour du dépôt %s : %v\n", repo.Name, err)
			continue
		}

		err = utils.FetchRepo(directory)
		if err != nil {
			fmt.Printf("Erreur lors de la récupération des références du dépôt %s : %v\n", repo.Name, err)
			continue
		}

		cloneCount++
	}

	err = utils.WriteCSV(repos, "repositories.csv")
	if err != nil {
		return err
	}

	zipName := "./zipRepo/ReposEnZip.zip"
	err = utils.ZipRepos(directories, zipName)
	if err != nil {
		fmt.Printf("\nErreur lors de la création du fichier zip pour les dépôts : %v\n", err)
		return err
	}
	fmt.Printf("\nTous les dépôts ont été compressés avec succès en tant que %s\n", zipName)

	return nil
}
