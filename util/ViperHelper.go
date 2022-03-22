package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/erfandiakoo/sms-api/models"
	"github.com/spf13/viper"
)

func InitViper() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigName("config")
			viper.AddConfigPath(".")
			viper.SetConfigType(".ini")
			viper.AutomaticEnv()
			viper.WriteConfig()
		} else {
			log.Fatalln("Error in Config file")
		}
	}
	ReadViper()
}

func ReadViper() {
	jsonFile, err := os.Open("conf.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cfg models.ConfigModel
	json.Unmarshal(byteValue, &cfg)

	viper.Set("MPUsername", cfg.MelipayamakUsername)
	viper.Set("MPPassword", cfg.MelipayamakPasswd)
	viper.Set("MPNum", cfg.MelipayamakNum)
	viper.Set("KaveNegarSender", cfg.KaveNegarSender)
	viper.Set("KaveNegarApi", cfg.KaveNegarApi)

	viper.WriteConfig()
}
