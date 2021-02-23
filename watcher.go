package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type Watcher struct {
}

func (w *Watcher) Run(symbols []string, interval int64) {
	if len(symbols) == 0 {
		fmt.Println("please input stock symbols")
		return
	}
	if len(symbols) > 20 {
		fmt.Println("stock quantity can't exceed 20")
		return
	}
	var listArray [20]*Data

	if interval < 20 {
		interval = 20
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			hour := time.Now().Hour()
			fmt.Println(hour)
			if hour >= 0 {
				fmt.Println("ok")
				cookies := _getCookies()
				var wg sync.WaitGroup
				// var dataChannel = make(chan *Data, len(symbols))
				for i := 0; i < len(symbols); i++ {
					wg.Add(1)
					go func(scode string, index int) {
						defer wg.Done()
						rt := GetQuoteData(scode, cookies)
						listArray[index] = rt
						// dataChannel <- rt

					}(symbols[i], i)
				}
				wg.Wait()
				// close(dataChannel)
				data := [][]string{}
				green := color.New(color.FgGreen).SprintFunc()
				red := color.New(color.FgRed).SprintFunc()
				yellow := color.New(color.FgYellow).SprintFunc()
				cyan := color.New(color.FgCyan).SprintFunc()

				// for item := range dataChannel {
				for i := 0; i < len(listArray); i++ {
					if listArray[i] == nil {
						break
					}
					item := listArray[i]

					status := cyan(item.Status)
					if item.Status == "盘前交易" {
						status = yellow(item.Status)
					}

					if item.Change > 0 {
						data = append(data, []string{fmt.Sprintf("%d", i+1), item.Name, status, red(fmt.Sprintf("↑ +%.2f%% (%.2f)", item.Percent, item.Current))})
					} else {
						if item.Change < 0 {
							data = append(data, []string{fmt.Sprintf("%d", i+1), item.Name, status, green(fmt.Sprintf("↓ %.2f%% (%.2f)", item.Percent, item.Current))})
						} else {
							data = append(data, []string{fmt.Sprintf("%d", i+1), item.Name, status, fmt.Sprintf("%.2f", item.Current)})
						}
					}

				}
				ClearScreen()
				fmt.Println(time.Now().Format("2006-01-02 15:04"))
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"No", "Name", "", ""})
				table.SetAutoWrapText(false)
				table.SetAutoFormatHeaders(true)
				table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
				table.SetAlignment(tablewriter.ALIGN_LEFT)
				table.SetCenterSeparator("")
				table.SetColumnSeparator("")
				table.SetRowSeparator("")
				table.SetHeaderLine(false)
				table.SetBorder(false)
				table.SetTablePadding("\t") // pad with tabs
				table.SetNoWhiteSpace(true)
				table.AppendBulk(data) // Add Bulk Data
				// table.SetTablePadding(" ")
				table.Render()

				// fmt.Printf("%v => %s\n", time.Now().Format("2006-01-02 15:04"), strings.Join(list, " , "))
				// for i := 0; i < len(list); i++ {
				// 	fmt.Println(list[i])
				// }
			}
		}
	}
}
