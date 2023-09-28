package main

import (
	"log"
	"os"
	"path/filepath"

	"PROJET-GIT-GO/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	app := fiber.New()

	app.Get("/api/repos/:user", func(c *fiber.Ctx) error {
		user := c.Params("user")
		token := os.Getenv("GITHUB_TOKEN")

		err := controllers.GetAndCloneRepositories(user, token)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString("Repo clonéé, ZIP creé et CSV created ! Regardez les resultat sur votre console !")
	})

	app.Get("/api/download", func(c *fiber.Ctx) error {

		zipPath := "./zipRepo/ReposEnZip.zip"

		c.Type("zip")
		c.Append("Content-Disposition", "attachment; filename="+filepath.Base(zipPath))

		return c.SendFile(zipPath)
	})

	log.Fatal(app.Listen(":3000"))
}
