package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/erfandiakoo/sms-api/config"
	"github.com/spf13/viper"
)

func SMSIRmakeRequest(jsonData map[string]string, op string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post(config.SMSIRUrl+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func SMSIRSendSMS(messages string, mobile string) {

	jsonData := map[string]string{
		"Messages":                 messages,
		"MobileNumbers":            mobile,
		"LineNumber":               viper.GetString("kavenegar_api"),
		"SendDateTime":             time.Now().String(),
		"CanContinueInCaseOfError": "false",
	}
	MelimakeRequest(jsonData, "MessageSend")
}
