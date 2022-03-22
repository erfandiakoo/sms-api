package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/erfandiakoo/sms-api/models"
	"github.com/erfandiakoo/sms-api/util"
	"github.com/gofiber/fiber/v2"
)

func Config(c *fiber.Ctx) error {
	Input := new(models.ConfigModel)
	if err := c.BodyParser(Input); err != nil {
		log.Fatal(err)
	}

	file, _ := json.MarshalIndent(&Input, "", " ")
	_ = ioutil.WriteFile("conf.json", file, 0644)

	util.ReadViper()

	return c.JSON(&models.ResponseModel{
		Data:    nil,
		Message: "Ok",
		Code:    200,
	})
}
