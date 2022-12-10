package main

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	configOnce   sync.Once
	configCookie string
)

func GetConfigCookie() string {
	configOnce.Do(func() {
		v := viper.New()
		v.SetConfigName("config")
		v.AddConfigPath("helper")
		v.SetConfigType("yaml")
		v.ReadInConfig()

		configCookie = v.GetString("Cookie")
	})
	return configCookie
}
