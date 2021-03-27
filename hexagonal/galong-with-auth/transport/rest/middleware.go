package rest

import (
	"backend-service/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func (ac *RContext) RegisterMiddlewares(e *echo.Echo) {

	// TODO change logger to custom logger
	e.Use(middleware.Logger())
	// TODO change recover to custom logger
	e.Use(middleware.Recover())
	t := logger.LoggerUtils{SugaredLogger: ac.Logger}
	e.HTTPErrorHandler = t.CustomZapHttpErrorHandler

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}
