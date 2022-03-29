package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/erfandiakoo/sms-api/config"
	"github.com/spf13/viper"
)

func MeliPayamakRequest(jsonData map[string]string, method string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(config.MeliPayamakUrl+method, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func MeliSendSMS(to string, text string) {

	jsonData := map[string]string{
		"username": viper.GetString("kavenegar_api"),
		"password": viper.GetString("kavenegar_api"),
		"to":       to,
		"from":     viper.GetString("kavenegar_api"),
		"text":     text,
		"isFlash":  strconv.FormatBool(false),
	}
	MeliPayamakRequest(jsonData, "SendSMS")
}
func MeliGetCredit(username string, password string) {
	jsonData := map[string]string{
		"UserName": viper.GetString("kavenegar_api"),
		"PassWord": viper.GetString("kavenegar_api"),
	}
	MeliPayamakRequest(jsonData, "GetCredit")
}
