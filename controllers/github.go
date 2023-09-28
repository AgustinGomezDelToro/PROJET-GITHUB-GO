package controllers

import (
	"PROJET-GIT-GO/models"
	"PROJET-GIT-GO/services"
	"PROJET-GIT-GO/utils"
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	directories := make(chan string, len(repos))

	for _, repo := range repos {
		wg.Add(1)

		go func(r models.Repository) {
			defer wg.Done()

			fmt.Printf("\nClonage du dépôt : %s, URL : %s\n", r.Name, r.CloneURL)

			t, err := time.Parse(time.RFC3339, r.LastUpdated)
			if err != nil {
				fmt.Printf("La date de la dernière mise à jour est vide pour le dépôt : %s\n", r.Name)
				return
			}
			r.LastUpdatedTime = t

			directory := "./clonedRepos/" + r.Name

			err = utils.CloneRepo(r.CloneURL, directory)
			if err != nil {
				fmt.Printf("Erreur lors du clonage du dépôt %s : %v\n", r.Name, err)
				return
			}

			err = utils.UpdateRepo(directory)
			if err != nil {
				fmt.Printf("Erreur lors de la mise à jour du dépôt %s : %v\n", r.Name, err)
				return
			}

			err = utils.FetchRepo(directory)
			if err != nil {
				fmt.Printf("Erreur lors de la récupération des références du dépôt %s : %v\n", r.Name, err)
				return
			}

			directories <- directory
		}(repo)
	}

	wg.Wait()
	close(directories)

	var dirSlice []string
	for dir := range directories {
		dirSlice = append(dirSlice, dir)
	}

	err = utils.WriteCSV(repos, "repositories.csv")
	if err != nil {
		return err
	}

	zipName := "./zipRepo/ReposEnZip.zip"
	err = utils.ZipRepos(dirSlice, zipName)
	if err != nil {
		fmt.Printf("\nErreur lors de la création du fichier zip pour les dépôts : %v\n", err)
		return err
	}
	fmt.Printf("\nTous les dépôts ont été compressés avec succès en tant que %s\n", zipName)

	return nil
}
