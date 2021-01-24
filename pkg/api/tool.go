package api

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/linclaus/stock-exportor/pkg/model"
	"github.com/spf13/viper"
)

func StoreCodes(codeSet *model.CodeSet) {
	codes := codeSet.List()
	file, err := os.OpenFile(viper.GetString("filePath")+viper.GetString("codeFileName"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	write := bufio.NewWriter(file)
	write.WriteString(strings.Join(codes, ","))
	write.Flush()
}

func StoreStocks(sm *model.StockMap) {
	file, err := os.OpenFile(viper.GetString("filePath")+viper.GetString("stockFileName"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprint(sm.List()))
	write.WriteString("\n")
	write.Flush()
}
