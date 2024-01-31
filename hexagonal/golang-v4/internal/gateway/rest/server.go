package rest

import (
	"errors"
	"net/http"
	"os"
	"os/signal"

	"git.siz-tel.com/charging/template/config"
	"git.siz-tel.com/charging/template/logger"
	"git.siz-tel.com/charging/template/utils"
	"github.com/gofiber/fiber/v2"
)

func Start() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if errors.Is(err, fiber.ErrBadRequest) {
				return ctx.Status(http.StatusBadRequest).JSON(utils.StandardHttpResponse{
					Message: utils.NotValidData,
					Status:  http.StatusBadRequest,
					Data:    err.Error(),
				})
			} else if errors.Is(err, fiber.ErrInternalServerError) {
				return ctx.Status(http.StatusInternalServerError).JSON(utils.StandardHttpResponse{
					Message: utils.ProblemInSystem,
					Status:  http.StatusInternalServerError,
					Data:    err.Error(),
				})
			} else if err != nil {
				return ctx.Status(http.StatusInternalServerError).JSON(utils.StandardHttpResponse{
					Message: utils.ProblemInSystem,
					Status:  http.StatusInternalServerError,
					Data:    err.Error(),
				})
			}

			return ctx.Status(http.StatusOK).JSON(utils.StandardHttpResponse{
				Message: utils.Ok,
				Status:  http.StatusOK,
				Data:    err.Error(),
			})
		},
	})

	registerMiddlewares(app)
	registerRoutes(app)

	go func() {
		if err := app.Listen(":" + config.AppConfig.General.Listen); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Logger.Fatal("shutting down server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of X seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := app.ShutdownWithTimeout(config.AppConfig.General.ShutdownTimeout); err != nil {
		logger.Logger.Fatal(err)
	}
}
