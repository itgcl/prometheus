package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// 自定义指标
var counter prometheus.Counter

/* 另外三种指标类型
	var gauge = prometheus.NewGauge(prometheus.GaugeOpts{})
	var summary = prometheus.NewSummary(prometheus.SummaryOpts{})
	var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{})
*/

func init(){
	// 设置options
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "",
		Subsystem: "http",
		Name:      "some",
		Help:      "some http", // 只是描述
		ConstLabels: prometheus.Labels{
			"label": "value",
		},
	})
	// 相同的options只需要注册一次, 重复注册会panic
	// 注册有多种方式, 这里是其中一种
	prometheus.MustRegister(counter)
}

func main() {
	// 暴露接口 prometheus定时获取指标
	http.Handle("/metrics", promhttp.Handler())

	// 自定义指标计数
	http.HandleFunc("/some", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		// 增加自定义value值
		//counter.Add(10)
		w.Write([]byte("pass"))
		return
	})

	http.ListenAndServe(":8080", nil)
}
