package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (cc *ControllerContext) Hello(ctx echo.Context) error {
	cc.Logger.Info("hey yo this happened")
	return ctx.JSON(http.StatusOK, "hello mate")
}
