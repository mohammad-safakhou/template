package rest

import (
	authentication "backend-service/auth"
	"backend-service/transport/controllers"
	"github.com/labstack/echo/v4"
)

func (ac *RContext) RegisterRoutes(e *echo.Echo) {
	controllerContext := controllers.ControllerContext{ApplicationContext: &ac.ApplicationContext}
	authContext := authentication.AuthContext{ApplicationContext: &ac.ApplicationContext}
	v1 := e.Group("/v1")

	v1.GET("/hello", controllerContext.Hello)
	v1.POST("/auth/otp", controllerContext.OTP(), authContext.WithAuth())
	v1.POST("/auth/verify", controllerContext.VerifyOTP(), authContext.WithAuth())
}
