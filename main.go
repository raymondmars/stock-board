package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

func main() {
	// prase arguments
	group := flag.String("g", "default", "# read which group config")
	flag.Parse()

	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	w := new(Board)

	interval, _ := cfg.Section(*group).Key("interval").Int64()
	language := cfg.Section(*group).Key("language").String()
	codes := strings.Split(cfg.Section(*group).Key("codes").String(), ",")
	w.Run(codes, interval, language)
}
