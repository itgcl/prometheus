package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)


var pusher *push.Pusher
func init(){
	// TODO 安装poshgateway 默认是9091端口
	// 连接pushgateway
	pusher = push.New("http://localhost:9091", "some_job")
}

func main()  {
	gaugeA := SomeA()
	gaugeB := SomeB()
	// 注册指标
	pusher.Collector(gaugeA)
	pusher.Collector(gaugeB)
	// 增加value计数
	gaugeA.Inc()
	gaugeB.Inc()
	// push到pushgateway
	err := pusher.Push()
	fmt.Println(err)
}

func SomeA() prometheus.Gauge {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "",
		Subsystem: "push",
		Name:      "some",
		Help:      "some push", // 只是描述
		ConstLabels: prometheus.Labels{
			"label": "value",
		},
	})
	return gauge
}

func SomeB() prometheus.Gauge {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "",
		Subsystem: "push",
		Name:      "some",
		Help:      "some push", // 只是描述
		ConstLabels: prometheus.Labels{
			"label": "value",
		},
	})
	return gauge
}
