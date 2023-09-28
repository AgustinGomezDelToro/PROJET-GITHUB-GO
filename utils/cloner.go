package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func CloneRepo(url, directory string) error {
	cmd := exec.Command("git", "clone", url, directory)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return fmt.Errorf("Error au clonage du repo: %w", err)
	}
	return nil
}

func UpdateRepo(directory string) error {
	log.Printf("Mise à jour du dépôt %s", directory)
	cmd := exec.Command("git", "pull", "origin")
	cmd.Dir = directory

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return fmt.Errorf("Error lors de la mise à jour du dépot : %w \n", err)
	}
	log.Printf("Dépôt %s mis à jour avec succès sur UpdateRepo\n", directory)
	return nil
}

func FetchRepo(directory string) error {
	log.Printf("Récup des références du dépôt %s sur FetchRepo\n", directory)
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = directory

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return fmt.Errorf("Erreur lors de la récupération des références : %w", err)
	}
	log.Printf("Références du dépôt %s récupérées avec succès\n", directory)
	return nil
}
