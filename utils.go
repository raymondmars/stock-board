package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func GetEvnWithDefaultVal(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	} else {
		return defaultVal
	}
}

func SendInform(name string, price float64) error {
	text := fmt.Sprintf("%s 到达目标价通知：）", name)
	desp := fmt.Sprintf("%s 已经到达目标价：￥%.2f", name, price)
	desp = url.QueryEscape(desp)
	text = url.QueryEscape(text)

	url := fmt.Sprintf("%s?text=%s&desp=%s", "https://sc.ftqq.com/SCU19880Te116691c07d63925173ee3175f92533d5a55b93258cfd.send", text, desp)

	// fmt.Println(url)
	_, err := http.Get(url)
	return err
}
