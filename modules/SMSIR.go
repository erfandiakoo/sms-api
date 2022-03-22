package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SMSIRmakeRequest(jsonData map[string]string, op string) {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("http://RestfulSms.com/api/"+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
func SMSIRSendSMS(username string, password string, to string, from string, text string, isFlash bool) {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	MelimakeRequest(jsonData, "MessageSend")
}
