package server

import (
	"log"

	"github.com/arcesti/projeto_go/internal/entity"
	"github.com/arcesti/projeto_go/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StrServer(port string) {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	user := entity.User{ID:1, Nome:"Arcesti", Email:"arcesti15@gmail.com", Idade: 20};

	app.Post("/createUser", func (c *fiber.Ctx) error {
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"erro": "Erro ao processar o corpo da requisição",
			})
		}
		_,err := service.CreateUser(user)
		if ; err!= nil {

		}
	})

	app.Listen(":"+port)
	log.Println("Servidor escutando em: localhost:" + port)
}