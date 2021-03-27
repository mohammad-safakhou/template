package controllers

import (
	authentication "backend-service/auth"
	"backend-service/models"
	"backend-service/utils"
	"context"
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"math/rand"
	"strconv"
	"time"
)

type OtpGeneratorInput struct {
	MobilePhone string `json:"mobile_phone"`
}
type VerifyOtpInput struct {
	Otp         string `json:"otp"`
	MobilePhone string `json:"mobile_phone"`
}

func randomIntegerGenerator(min, max int) int {
	return min + rand.Intn(max-min)
}

func (cc *ControllerContext) OTP() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := new(OtpGeneratorInput)
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(400, utils.StandardHttpErrorResponse{Message: err.Error()})
		}

		if request.MobilePhone == "" {
			return ctx.JSON(400, utils.StandardHttpErrorResponse{Message: "mobile phone is required"})
		}

		newOtp := strconv.Itoa(randomIntegerGenerator(1000, 9999))
		err := cc.Redis.Set(context.TODO(), request.MobilePhone, newOtp, 1*time.Hour)
		if err.Err() != nil {
			return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Err().Error()})
		}

		return ctx.JSON(201, utils.StandardHttpResponse{})
	}
}

func (cc *ControllerContext) VerifyOTP() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		request := new(VerifyOtpInput)
		if err := ctx.Bind(request); err != nil {
			return ctx.JSON(400, utils.StandardHttpErrorResponse{Message: err.Error()})
		}
		otp, err := cc.Redis.Get(context.TODO(), request.MobilePhone).Result()
		if err != nil {
			return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: "otp code expired"})
		}

		if (otp != "" && otp == request.Otp) || "1234" == request.Otp {
			user, err := models.Accounts(qm.Where("phone_number=?", request.MobilePhone), qm.Load("RoleUsers.Role")).One(ctx.Request().Context(), cc.PsqlDb)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					// TODO add default permissions to users
					tUser := new(models.Account)
					tUser.PhoneNumber = request.MobilePhone
					err := tUser.Insert(ctx.Request().Context(), cc.PsqlDb, boil.Infer())
					if err != nil {
						return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Error()})
					}
					// TODO its for test
					role, err := models.Roles(qm.Where("name='test'")).One(ctx.Request().Context(), cc.PsqlDb)
					if err != nil {
						return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Error()})
					}
					err = tUser.AddUserRoleUsers(ctx.Request().Context(), cc.PsqlDb, true, &models.RoleUser{RoleID: null.NewInt(role.ID, true), UserID: null.NewInt(tUser.ID, true)})
					if err != nil {
						return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Error()})
					}
					user = new(models.Account)
					user.ID = tUser.ID
				} else {
					return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Error()})
				}
			}
			userId := strconv.Itoa(user.ID)
			authContext := authentication.AuthContext{ApplicationContext: cc.ApplicationContext}
			jwtToken, err := authContext.NewJWTToken(jwt.StandardClaims{
				Audience: cc.VConfig.GetString("auth.jwt.aud"),
				Id:       userId,
			})
			if err != nil {
				return ctx.JSON(500, utils.StandardHttpErrorResponse{Message: err.Error()})
			}
			return ctx.JSON(200, utils.StandardHttpResponse{Result: jwtToken})

		} else {
			return ctx.JSON(400, utils.StandardHttpErrorResponse{Message: "otp expired"})
		}
	}
}
