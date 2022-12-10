package main

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func NewRequest(method string, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	viper.AddConfigPath("./helper")

	req.Header.Set("Cookie", GetConfigCookie())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bts, nil
}
