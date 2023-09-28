package utils

import (
	"PROJET-GIT-GO/models"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func WriteCSV(repos []models.Repository, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error a la creation du dossier ", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"ID", "Name", "Full Name", "Clone URL", "Last Updated"})
	if err != nil {
		log.Println("Error a l'ecriture du dossier :", err)
		return err
	}

	for _, repo := range repos {
		err = writer.Write([]string{
			strconv.Itoa(repo.ID),
			repo.Name,
			repo.FullName,
			repo.CloneURL,
			repo.LastUpdated,
		})
		if err != nil {
			log.Println("Error a l'ecriture du dossier :", err)
			return err
		}
	}
	return nil
}
