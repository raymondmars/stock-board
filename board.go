package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type Board struct {
	Cookies []*http.Cookie
}

func (w *Board) Run(symbols []string, interval int64, language string) {
	if len(symbols) == 0 {
		fmt.Println("please input stock symbols")
		return
	}
	if len(symbols) > 20 {
		fmt.Println("stock quantity can't exceed 20")
		return
	}
	if language == "" {
		language = "zh"
	}

	var listArray [20]*Data

	if interval < 20 {
		interval = 20
	}
	w.Cookies = _getCookies()

	w.render(symbols, &listArray, language)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-ticker.C:
			w.render(symbols, &listArray, language)
		}
	}
}

func (w *Board) render(symbols []string, listArray *[20]*Data, language string) {
	if w.Cookies == nil || len(w.Cookies) == 0 {
		panic("Err: need to load cookies first")
	}
	// cookies := _getCookies()
	var wg sync.WaitGroup
	for i := 0; i < len(symbols); i++ {
		wg.Add(1)
		go func(scode string, index int) {
			defer wg.Done()
			rt := GetQuoteData(scode, w.Cookies)
			listArray[index] = rt

		}(symbols[i], i)
	}
	wg.Wait()

	data := [][]string{}
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	for i := 0; i < len(listArray); i++ {
		if listArray[i] == nil {
			break
		}
		item := listArray[i]

		status := item.ZhStatusDescription
		if language == "en" {
			status = GetEnStatusDescription(item.Status)
		}
		if item.Status == OpenBefore {
			status = yellow(status)
		} else {
			status = cyan(status)
		}

		displayName := item.Name
		if language == "en" {
			displayName = item.Code
		}
		if item.Change > 0 {
			data = append(data, []string{fmt.Sprintf("%d", i+1), displayName, item.Exchange, status, red(fmt.Sprintf("↑ +%.2f%% (%.3f)", item.Percent, item.Current))})
		} else {
			if item.Change < 0 {
				data = append(data, []string{fmt.Sprintf("%d", i+1), displayName, item.Exchange, status, green(fmt.Sprintf("↓ %.2f%% (%.3f)", item.Percent, item.Current))})
			} else {
				data = append(data, []string{fmt.Sprintf("%d", i+1), displayName, item.Exchange, status, fmt.Sprintf("%.3f", item.Current)})
			}
		}

	}
	ClearScreen()
	fmt.Println(time.Now().Format("2006-01-02 15:04"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Name", "ExG", "", ""})
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

}
