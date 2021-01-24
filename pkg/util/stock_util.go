package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/linclaus/stock-exportor/pkg/model"
	"golang.org/x/text/encoding/simplifiedchinese"
)

const (
	// SINA_URL = "http://hq.sinajs.cn/"
	SINA_URL = "http://123.125.104.104/"
)

func GetStockTime(code string) string {
	stockStr, ok := getStocksInfoFromSina([]string{code})
	if !ok {
		return ""
	} else {
		rst := strings.Split(stockStr, "=\"")
		values := strings.Split(rst[1], ",")
		currentTime := values[31]
		return currentTime
	}
}

func StockOpened() bool {
	stockCurrents := []string{}
	for i := 0; i < 5; i++ {
		stockStr, ok := getStocksInfoFromSina([]string{"sh000001"})
		if !ok {
			return false
		} else {
			rst := strings.Split(stockStr, "=\"")
			values := strings.Split(rst[1], ",")
			current := values[3]
			stockCurrents = append(stockCurrents, current)
			if stockCurrents[0] != current {
				return true
			}
			time.Sleep(10 * time.Second)
		}
	}
	return false
}

func getStocksInfoFromSina(codes []string) (string, bool) {
	resp, err := http.Get(SINA_URL + "list=" + strings.Join(codes, ","))
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(body)
	results := string(decodeBytes)
	ok := true
	if resp.StatusCode != 200 {
		ok = false
	}
	return results, ok
}

func StockExixts(code string) bool {
	results, _ := getStocksInfoFromSina([]string{code})
	return resultValid(results)
}

func resultValid(stockStr string) bool {
	return len(stockStr) > 30
}

// Parse stocks message and set Stock to stockMap
func ParseAndPutStock(stocksStr string, sm *model.StockMap) {
	for _, stockStr := range strings.Split(stocksStr, ";") {
		if stockStr != "" && stockStr != "\n" && resultValid(stockStr) {
			rst := strings.Split(stockStr, "=\"")
			l := len(rst[0])
			code := rst[0][l-8 : l]
			values := strings.Split(rst[1], ",")
			time := values[31]
			current, _ := strconv.ParseFloat(values[3], 32)
			currentTime := values[31]
			close, _ := strconv.ParseFloat(values[2], 32)
			increase := (current - close) / close * 100
			tradeVolume, _ := strconv.ParseFloat(values[8], 32)
			lastBuy1, _ := strconv.ParseFloat(values[10], 32)
			lastBuy2, _ := strconv.ParseFloat(values[12], 32)
			lastBuy3, _ := strconv.ParseFloat(values[14], 32)
			last3BuySum := (lastBuy1 + lastBuy2 + lastBuy3)
			lastSell1, _ := strconv.ParseFloat(values[20], 32)
			lastSell2, _ := strconv.ParseFloat(values[22], 32)
			lastSell3, _ := strconv.ParseFloat(values[24], 32)
			last3SellSum := (lastSell1 + lastSell2 + lastSell3)
			weibi := (last3BuySum - last3SellSum) / (last3BuySum + last3SellSum)
			if last3BuySum+last3SellSum == 0 {
				weibi = 0
			}

			if s, ok := sm.Get(code); !ok {
				open, _ := strconv.ParseFloat(values[1], 32)
				highest, _ := strconv.ParseFloat(values[4], 32)
				lowest, _ := strconv.ParseFloat(values[5], 32)
				s := model.Stock{
					Name:         values[0],
					Code:         code,
					Open:         float64(open),
					Close:        float64(close),
					Current:      float64(current),
					Highest:      float64(highest),
					Lowest:       float64(lowest),
					Increase:     float64(increase),
					TradeVolume:  float64(tradeVolume) / 100,
					Last3BuySum:  float64(last3BuySum),
					Last3SellSum: float64(last3SellSum),
					Weibi:        float64(weibi),
					CurrentTime:  currentTime,
					History: map[string]float64{
						currentTime: float64(current),
					},
				}
				sm.Add(code, s)
				// fmt.Printf("InitStock:%s\n", s)
			} else {
				s.Increase = float64(increase)
				s.TradeVolume = float64(tradeVolume)
				s.Last3BuySum = float64(last3BuySum)
				s.Last3SellSum = float64(last3SellSum)
				s.Weibi = float64(weibi)
				s.Current = float64(current)
				if s.IsRecord {
					s.History[time] = float64(current) / 100
				}
				s.CurrentTime = currentTime
				sm.Add(code, s)
				// fmt.Printf("UpdateStock:%s\n", s)
			}
		}
	}
}

//Set stockMap by codes
func SetByCodes(codes []string, sm *model.StockMap) {
	codeLists := SplitStringSlice(codes, 600)
	var wg sync.WaitGroup
	wg.Add(len(codeLists))
	for _, codeList := range codeLists {
		go func(codes []string, sm *model.StockMap) {
			results, ok := getStocksInfoFromSina(codes)
			if ok {
				ParseAndPutStock(results, sm)
			}
			wg.Done()
		}(codeList, sm)
	}
	wg.Wait()
}
