package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	QUOTES_URL      = "https://xueqiu.com/"
	QUOTES_JSON_URL = "https://stock.xueqiu.com/v5/stock/quote.json?symbol=%s&extend=detail"
)

//获取任意股票对应的价格，该方法只适合于A股股票，并且只是从HTML代码中提取价格
func GetPrice(code string) (float64, error) {
	url := fmt.Sprintf("%s/S/%s", QUOTES_URL, _getSymbolCode(code))
	res, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return 0.0, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return 0.0, errors.New("request failed")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
		return 0.0, err
	}

	ret := 0.0
	sel := doc.Find(".stock-current").First()
	if sel != nil {
		price := sel.Find("strong").Text()
		price = strings.ReplaceAll(price, "¥", "")
		price = strings.ReplaceAll(price, "HK$", "")

		f, err := strconv.ParseFloat(price, 64)
		if err == nil {
			ret = f
		}
	}

	return ret, nil
}

//访问雪球主页获取cookie，用于后续授权请求
func _getCookies() []*http.Cookie {
	res, err := http.Get(QUOTES_URL)
	if err != nil {
		log.Println(err)
		return nil
	}
	return res.Cookies()
}

//用随机访问获取的cookies，请求给定股票对应的JSON数据
func GetQuoteData(code string, cookies []*http.Cookie) *Data {

	url := fmt.Sprintf(QUOTES_JSON_URL, _getSymbolCode(code))

	client := &http.Client{Timeout: time.Second * 30}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	req.Header.Add("Host", "xueqiu.com")
	req.Header.Add("Referer", "https://xueqiu.com/")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body. ", err)
		return nil
	}

	retData := &struct {
		Data `json:"data"`
	}{}

	err = json.Unmarshal(body, retData)
	// fmt.Println(string(body))
	if err != nil {
		log.Printf("invalid data: %s", string(body))
		return nil
	}
	return &retData.Data

}

// 将普通股票代码转换成雪球可以识别的代码
func _getSymbolCode(code string) string {
	if len(code) > 5 && code != "SH000001" {
		firstChar := code[:1]
		if firstChar == "5" || firstChar == "6" || firstChar == "9" {
			return fmt.Sprintf("SH%s", code)
		} else {
			return fmt.Sprintf("SZ%s", code)
		}
	} else {
		return code
	}
}
