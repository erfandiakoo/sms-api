package modules

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//request methods
func MediamakeRequestToSend(Data string) {

	dataAsByte := []byte(Data)
	response, err := http.Post("http://payamak-service.ir/SendService.svc?wsdl", "text/xml; charset=utf-8", bytes.NewBuffer(dataAsByte))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
