package rest

import (
	"backend-service/config"
	"backend-service/transport"
	"backend-service/transport/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)


func StartRestServer(l *zap.SugaredLogger) {
	ac := transport.ApplicationContext{
		VConfig: nil,
		PsqlDb:  nil,
		Logger:  l,
	}
	vc, err := config.ViperConfig()
	if err != nil {
		panic(fmt.Sprintf("initializing viper failed %v", err))
	}
	ac.VConfig = vc

	dbContext := database.DBContext{ApplicationContext: &ac}
	dbContext.RegisterDatabases()

	e := echo.New()

	rc := RContext{ApplicationContext: ac}
	rc.RegisterMiddlewares(e)
	rc.RegisterRoutes(e)

	ac.Logger.Fatal(e.Start(":" + ac.VConfig.GetString("server.port")))
}
