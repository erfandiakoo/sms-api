package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/erfandiakoo/sms-api/config"
	"github.com/spf13/viper"
)

func SMSIRRequest(jsonData map[string]string, method string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(config.SMSIRUrl+method, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

func SMSIRSendSMS(message string, mobile string) {
	jsonData := map[string]string{
		"Messages":                 message,
		"MobileNumbers":            mobile,
		"LineNumber":               "1000091001919",
		"SendDateTime":             time.Now().UTC().String(),
		"CanContinueInCaseOfError": strconv.FormatBool(false),
	}
	SMSIRRequest(jsonData, "MessageSend")
}

func SMSIRGetToken() {
	jsonData := map[string]string{
		"UserApiKey": viper.GetString("kavenegar_api"),
		"SecretKey":  viper.GetString("kavenegar_api"),
	}
	SMSIRRequest(jsonData, "Token")
}
