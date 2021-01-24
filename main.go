package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/linclaus/stock-exportor/pkg/cache"
	"github.com/spf13/viper"

	"github.com/linclaus/stock-exportor/pkg/api"
	"github.com/linclaus/stock-exportor/pkg/channel"
	"github.com/linclaus/stock-exportor/pkg/metric"
	"github.com/linclaus/stock-exportor/pkg/model"

	"github.com/linclaus/stock-exportor/pkg/util"
)

var cfgFile string

func main() {
	fmt.Println("Stock Started")
	// Set Args
	flag.StringVar(&cfgFile, "file", "", "The config file.")
	flag.Parse()
	initConfig()
	// Get os signal
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	codes := initCodeSet()
	s := api.New(codes)
	//Init metrics
	channel.Init()

	util.SetByCodes(s.Codes.List(), cache.StockMap)
	//Start server
	go s.Start(":" + viper.GetString("port"))
	//Set stock message scrape interval
	ticker := time.NewTicker(time.Second * time.Duration(viper.GetInt("interval")))
	scrapeFlag := true
	lastTime := util.GetStockTime("sh000001")
	for {
		select {
		case <-ticker.C:
			if viper.GetBool("alwaysScrape") {
				fmt.Printf("Update %d Metrics\n", len(s.Codes.List()))
				util.SetByCodes(s.Codes.List(), cache.StockMap)
				break
			}

			currentTime := util.GetStockTime("sh000001")
			now := time.Now()
			todayMinutes := now.Hour()*60 + now.Minute()
			fmt.Printf("hour: %d, minute: %d, todayMinutes: %d\n", now.Hour(), now.Minute(), todayMinutes)
			fmt.Printf("lastTime: %s, currentTime: %s\n", lastTime, currentTime)
			if (todayMinutes > 9*60+30-1 && todayMinutes < 15*60+5) && !(strings.HasPrefix(lastTime, "15") && strings.HasPrefix(currentTime, "15")) {
				if !scrapeFlag {
					fmt.Println("Stock is opened")
					metric.RegistMetric()
					scrapeFlag = true
				}
				fmt.Printf("Update %d Metrics\n", len(s.Codes.List()))
				util.SetByCodes(s.Codes.List(), cache.StockMap)
			} else {
				if scrapeFlag {
					fmt.Println("Stock is closed")
					metric.UnRegistMetric()
					scrapeFlag = false
				}
			}
			lastTime = currentTime
		case signal := <-c:
			//Save stock message before exit
			fmt.Println("退出信号", signal)
			fmt.Println(strings.Join(s.Codes.List(), ","))
			fmt.Println(cache.StockMap)
			api.StoreStocks(cache.StockMap)
			return
		}
	}
}

func initCodeSet() *model.CodeSet {
	// Initial the monitor stock code
	file, err := os.OpenFile(viper.GetString("filePath")+viper.GetString("codeFileName"), os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	buf := bufio.NewReader(file)

	codeStr := ""
	for {
		codeByte, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		codeStr = codeStr + string(codeByte)
	}

	file.Close()
	codes := strings.Split(codeStr, ",")
	cs := model.NewCodeSet()
	if len(codes) == 0 || codes[0] == "" {
		codes = []string{"sh000001", "sz399001"}
		cs.AddSet(codes)
		api.StoreCodes(cs)
		return cs
	}
	cs.AddSet(codes)
	return cs
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.AddConfigPath("/etc/")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
