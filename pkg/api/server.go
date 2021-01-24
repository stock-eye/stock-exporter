package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linclaus/stock-exportor/pkg/cache"
	"github.com/linclaus/stock-exportor/pkg/metric"
	"github.com/linclaus/stock-exportor/pkg/model"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	e     *gin.Engine
	Codes *model.CodeSet
}

// Get Server
func New(cs *model.CodeSet) *Server {
	e := gin.Default()
	s := Server{
		e:     e,
		Codes: cs,
	}

	e.GET("/metrics", s.metricHandler(), gin.WrapH(promhttp.Handler()))
	e.GET("/stocks", s.GetStocks)
	e.GET("/stock/:code", s.GetStockByCode)
	e.POST("/stock/:code", s.AddStockByCode)
	e.DELETE("/stock/:code", s.DeleteStockByCode)
	return &s
}

// Custom Handler
func (s Server) metricHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before metric handler")
		for _, value := range cache.StockMap.List() {
			metric.StockCurrentMetricCountVec.WithLabelValues(value.Code, value.Name).Set(value.Current)
			metric.StockIncreaseMetricCountVec.WithLabelValues(value.Code, value.Name).Set(value.Increase)
			metric.StockTradeVolmeMetricTotalVec.WithLabelValues(value.Code, value.Name).Set(value.TradeVolume)
			metric.StockLastBuy3MetricVec.WithLabelValues(value.Code, value.Name).Set(value.Last3BuySum)
			metric.StockLastSell3MetricVec.WithLabelValues(value.Code, value.Name).Set(value.Last3SellSum)
			metric.StockWeibiMetricVec.WithLabelValues(value.Code, value.Name).Set(value.Weibi)
			metric.StockIncreaseHistogramVec.WithLabelValues().Observe(value.Increase)
		}
		c.Next()
		log.Println("after metric handler")
	}

}

// Custom Inteceptor
func (s Server) handleFuncInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before handlerFunc")
		h(w, r)
		log.Println("after handlerFunc")
	}
}

func (s Server) Start(address string) {
	log.Println("Starting listener on", address)
	log.Fatal(s.e.Run(address))
}
