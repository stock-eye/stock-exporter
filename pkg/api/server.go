package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linclaus/stock-exportor/pkg/cache"
	"github.com/linclaus/stock-exportor/pkg/metric"
	"github.com/linclaus/stock-exportor/pkg/model"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	r     *mux.Router
	Codes *model.CodeSet
}

// Get Server
func New(cs *model.CodeSet) *Server {
	r := mux.NewRouter()
	s := Server{
		r:     r,
		Codes: cs,
	}
	r.Handle("/metrics", s.metricHandler(promhttp.Handler()))
	r.HandleFunc("/stocks", s.GetStocks).Methods("GET")
	r.HandleFunc("/stock-operator/{code}", s.GetStockByCode).Methods("GET")
	r.HandleFunc("/stock-operator/{code}", s.AddStockByCode).Methods("POST")
	r.HandleFunc("/stock-operator/{code}", s.DeleteStockByCode).Methods("DELETE")
	return &s
}

// Custom Handler
func (s Server) metricHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		next.ServeHTTP(w, r)
		log.Println("after metric handler")
	})

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
	log.Fatal(http.ListenAndServe(address, s.r))
}
