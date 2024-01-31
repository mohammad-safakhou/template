package rest

import (
	"git.siz-tel.com/charging/template/internal/services"
	"git.siz-tel.com/charging/template/internal/services/service_models"
	"git.siz-tel.com/charging/template/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type accountController struct {
	accountsService services.Accounts
}

type createAccountsRequest struct {
	IMSI string `json:"IMSI"`
}

func (c *accountController) CreateAccount(ctx *fiber.Ctx) error {
	var req createAccountsRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.StandardHttpResponse{
			Message: utils.NotValidData,
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
		})
	}

	err = c.accountsService.Save(ctx.Context(), service_models.Account{
		IMSI: req.IMSI,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.StandardHttpResponse{
			Message: utils.ProblemInSystem,
			Status:  http.StatusInternalServerError,
			Data:    err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(utils.StandardHttpResponse{
		Message: utils.Ok,
		Status:  http.StatusCreated,
		Data:    nil,
	})
}
