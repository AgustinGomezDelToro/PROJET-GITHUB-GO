package main

import (
	"PROJET-GIT-GO/controllers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error au deployement du fichier .env")
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

	log.Fatal(app.Listen(":3000"))
}
