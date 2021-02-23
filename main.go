package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var ConfigTargets = GetEvnWithDefaultVal("CONFIG_TARGETS", "600030,50.5|300750,200|TSLA,600")

const MaxInformCount = 5

func main() {

	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	w := new(Watcher)
	interval, _ := cfg.Section("conf").Key("interval").Int64()
	codes := strings.Split(cfg.Section("conf").Key("codes").String(), ",")

	w.Run(codes, interval)

	// if ConfigTargets != "" {
	// 	list := strings.Split(ConfigTargets, "|")
	// 	if len(list) > 0 {
	// 		dataList := []TargetData{}
	// 		for _, item := range list {
	// 			arr := strings.Split(item, ",")
	// 			if len(arr) > 1 {
	// 				td := TargetData{}
	// 				td.Code = arr[0]
	// 				f, _ := strconv.ParseFloat(arr[1], 64)
	// 				td.TargetPrice = f
	// 				dataList = append(dataList, td)
	// 			} else {
	// 				log.Panicf("target data is wrong: %s", item)
	// 			}
	// 		}
	// 		if len(dataList) > 0 {
	// 			informCounter := map[string]int{}
	// 			ticker := time.NewTicker(15 * time.Minute)
	// 			for {
	// 				select {
	// 				case <-ticker.C:
	// 					hour := time.Now().Hour()
	// 					if hour >= 9 && hour <= 16 {
	// 						cookies := _getCookies()
	// 						for _, data := range dataList {
	// 							if informCounter[data.Code] >= MaxInformCount {
	// 								continue
	// 							}
	// 							rt := GetQuoteData(data.Code, cookies)
	// 							if rt != nil && rt.Market.Status != "休市" && rt.Market.Status != "已收盘" {
	// 								if rt.Quote.Current > data.TargetPrice {
	// 									err := SendInform(rt.Quote.Name, rt.Quote.Current)
	// 									if err == nil {
	// 										informCounter[data.Code]++
	// 										log.Printf("%s success touch ￥%.2f", rt.Quote.Name, rt.Quote.Current)
	// 									} else {
	// 										log.Printf("send inform failed: %v", err)
	// 									}
	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	} else {
	// 		log.Panicf("config is wrong: %s", ConfigTargets)
	// 	}
	// }
}
