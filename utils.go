package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
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

	informURL := fmt.Sprintf("https://sc.ftqq.com/%s", GetEvnWithDefaultVal("INFORM_URL", ""))
	url := fmt.Sprintf("%s?text=%s&desp=%s", informURL, text, desp)

	// fmt.Println(url)
	_, err := http.Get(url)
	return err
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
