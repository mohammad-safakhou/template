package rest

import (
	"git.siz-tel.com/charging/template/config"
	"git.siz-tel.com/charging/template/internal/repository/database"
	"git.siz-tel.com/charging/template/internal/services"
	"git.siz-tel.com/charging/template/utils"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	// ---------------- Dependency   ----------------
	db, err := utils.PostgresConnection(
		config.AppConfig.Databases.Postgres.Host,
		config.AppConfig.Databases.Postgres.Port,
		config.AppConfig.Databases.Postgres.User,
		config.AppConfig.Databases.Postgres.Pass,
		config.AppConfig.Databases.Postgres.DatabaseName,
		config.AppConfig.Databases.Postgres.SslMode,
		config.AppConfig.Databases.Postgres.MaxOpenConns,
		config.AppConfig.Databases.Postgres.MaxIdleConns,
		config.AppConfig.Databases.Postgres.Timeout)
	if err != nil {
		panic(err)
	}
	// ---------------- Repositories ----------------
	accountsRepository := database.NewAccountRepo(db, db)
	// ---------------- Services     ----------------
	accountService := services.NewAccountsService(accountsRepository)
	// ---------------- Controllers  ----------------
	accountControllers := accountController{accountsService: accountService}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	v1 := app.Group("/v1")

	v1.Post("/account", accountControllers.CreateAccount)
}
