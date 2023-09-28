package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func CloneRepo(url, directory string) error {
	cmd := exec.Command("git", "clone", url, directory)
	err := cmd.Run()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return fmt.Errorf("Error au clonage du repo: %w", err)
	}
	return nil
}
