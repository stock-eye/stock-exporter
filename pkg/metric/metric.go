package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	StockCurrentMetricCountVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_current_gauge",
		Help: "stock current value",
	}, []string{"code", "name"})

	StockIncreaseMetricCountVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_increase_gauge",
		Help: "stock increase value",
	}, []string{"code", "name"})

	StockTradeVolmeMetricTotalVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_trade_volume_total",
		Help: "stock trade value total value",
	}, []string{"code", "name"})
	StockLastBuy3MetricVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_last_buy_3_gauge",
		Help: "stock of sum last buy 5 value",
	}, []string{"code", "name"})
	StockLastSell3MetricVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_last_sell_3_gauge",
		Help: "stock of sum last sell 5 value",
	}, []string{"code", "name"})
	StockWeibiMetricVec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "stock_weibi_gauge",
		Help: "stock of weibi",
	}, []string{"code", "name"})

	StockIncreaseHistogramVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "stock_increase_histogram",
			Help:    "Histogram of stock increase.",
			Buckets: prometheus.LinearBuckets(-10.5, 1, 22),
		},
		[]string{},
	)
)

//Init metric
func init() {
	RegistMetric()
}

func RegistMetric() {
	prometheus.MustRegister(StockCurrentMetricCountVec)
	prometheus.MustRegister(StockIncreaseMetricCountVec)
	prometheus.MustRegister(StockTradeVolmeMetricTotalVec)
	prometheus.MustRegister(StockLastBuy3MetricVec)
	prometheus.MustRegister(StockLastSell3MetricVec)
	prometheus.MustRegister(StockWeibiMetricVec)
	prometheus.MustRegister(StockIncreaseHistogramVec)
}

func UnRegistMetric() {
	prometheus.Unregister(StockCurrentMetricCountVec)
	prometheus.Unregister(StockIncreaseMetricCountVec)
	prometheus.Unregister(StockTradeVolmeMetricTotalVec)
	prometheus.Unregister(StockLastBuy3MetricVec)
	prometheus.Unregister(StockLastSell3MetricVec)
	prometheus.Unregister(StockWeibiMetricVec)
	prometheus.Unregister(StockIncreaseHistogramVec)
}
