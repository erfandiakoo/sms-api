package handler

import (
	"net/http"
	"strings"

	"github.com/erfandiakoo/sms-api/models"
	"github.com/erfandiakoo/sms-api/modules"
	"github.com/erfandiakoo/sms-api/util"
	"github.com/gofiber/fiber/v2"
)

func SendSMS(c *fiber.Ctx) error {
	Input := new(models.SMSRequestModel)
	if err := c.BodyParser(Input); err != nil {
		return err
	}
	if Input.Receptor = util.CleanUpPhone(Input.Receptor); Input.Receptor == "" {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    500,
			Message: "شماره همراه اشتباه است.",
		})
	}

	//KaveNegar
	modules.KaveSendSMS(Input.Message, strings.Split(Input.Receptor, ""))
	//MeliPayamak
	modules.MeliSendSMS(Input.Receptor, Input.Message)
	//SMS IR

	return c.JSON(models.ResponseModel{
		Data:    nil,
		Code:    http.StatusOK,
		Message: "Sent succesfully!",
	})
}
