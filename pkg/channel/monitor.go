package channel

import (
	"sync"

	"github.com/linclaus/stock-exportor/pkg/cache"
	"github.com/linclaus/stock-exportor/pkg/model"
)

var lock sync.RWMutex

func Init() {
	go monitorStockAdd(cache.StockMap.StockAddChan)
}

// Add custom monitor handler
func monitorStockAdd(stock chan model.Stock) {
	for s := range stock {
		handlerStock(s)
	}
}

func handlerStock(stock model.Stock) {}
