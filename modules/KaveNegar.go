package modules

import (
	"log"

	"github.com/kavenegar/kavenegar-go"
	"github.com/spf13/viper"
)

func KaveSendSMS(message string, receptor []string) {
	key := viper.GetString("kavenegar_api")
	api := kavenegar.New(key)
	sender := viper.GetString("kavenegar_sender")
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			log.Println(err.Error())
		case *kavenegar.HTTPError:
			log.Println(err.Error())
		default:
			log.Println(err.Error())
		}
	} else {
		for _, r := range res {
			log.Println("MessageID 	= ", r.MessageID)
			log.Println("Status    	= ", r.Status)
		}
	}
}
