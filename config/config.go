package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Running in debug mode")
	}
}
