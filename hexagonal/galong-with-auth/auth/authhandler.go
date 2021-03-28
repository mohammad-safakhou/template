package authentication

import (
	"backend-service/utils"
	"github.com/labstack/echo/v4"
)

func (ac *AuthContext) WithAuth() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			t := ctx.Request().Header.Get("Authorization")
			if t == "" {
				return ctx.JSON(401, utils.StandardHttpErrorResponse{Message: "you are not authorized by server"})
			}
			t = t[len("Bearer "):]
			token, err := ac.validateToken(t)
			if err != nil {
				return ctx.JSON(401, utils.StandardHttpErrorResponse{Message: "you are not authorized by server"})
			}
			ident, err := ac.newIdentity(token)
			if err != nil {
				return ctx.JSON(401, utils.StandardHttpErrorResponse{Message: "you are not authorized by server"})
			}
			cancan := false
			// TODO its for test
			{
				for _, role := range ident.Roles {
					if role.Name == "test" {
						ctx.Set("identity", ident)
						return h(ctx)
					}
				}
			}

			for _, per := range ident.Permissions {
				if per.Route.String == ctx.Request().RequestURI && per.Method.String == ctx.Request().Method {
					cancan = true
					break
				}
			}
			if !cancan {
				return ctx.JSON(403, utils.StandardHttpErrorResponse{Message: "you dont have access to this route"})
			}
			ctx.Set("identity", ident)
			return h(ctx)
		}
	}
}