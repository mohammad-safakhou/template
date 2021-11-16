package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"template/adapters/repository"
	"template/core/usecase"
)


func (rc *restControllers) Hello(ctx echo.Context) error {
	helloRepository := repository.NewHelloRepository(rc.PsqlDb)
	helloService := usecase.NewHelloService(helloRepository)
	helloService.SayHello()
	return ctx.JSON(http.StatusOK, "hello just worked awesome")
}
