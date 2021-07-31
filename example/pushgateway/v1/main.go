package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)


var pusher *push.Pusher
var gauge prometheus.Gauge

func init(){
	// 自定义指标
	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "",
		Subsystem: "push",
		Name:      "some",
		Help:      "some push", // 只是描述
		ConstLabels: prometheus.Labels{
		"label": "value",
		},
	})

	// TODO 安装poshgateway 默认是9091端口
	// 连接pushgateway
	pusher = push.New("http://localhost:9091", "some_job")
}

func main()  {
	// 注册指标
	pusher.Collector(gauge)
	// 增加value计数
	gauge.Inc()
	// push到pushgateway
	err := pusher.Push()
	fmt.Println(err)
}
