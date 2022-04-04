package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	w := new(Board)
	sectionValue := "default"
	interval, _ := cfg.Section(sectionValue).Key("interval").Int64()
	language := cfg.Section(sectionValue).Key("language").String()
	codes := strings.Split(cfg.Section(sectionValue).Key("codes").String(), ",")
	w.Run(codes, interval, language)
}
