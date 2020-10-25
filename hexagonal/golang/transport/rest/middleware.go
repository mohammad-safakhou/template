package rest

import (
	"bizpooly/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
