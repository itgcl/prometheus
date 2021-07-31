package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

// TODO 已解决相同指标信息正常记录

var pusher *push.Pusher
var gauge *prometheus.GaugeVec

func init(){
	// 正常来说一个服务指标信息基础描述和label的key可以预先定义保证不变, 一直变的是label的value, 对不同的value进行递增
	// 例如 status的值可能是pass或fail。pass不存在就创建 存在就递增
	/*
		{
			"service_name": "push_some",
			"labels": {
				"status":"pass", // or "fail",
				"fail_kind": "db_exec_error", // or "redis_exec_error", "request_params_error" ....
			}
		}
	*/

	// 按不同维度划分的同一个东西
	gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "",
		Subsystem: "push",
		Name:      "some",
		Help:      "some push", // 只是描述
	}, []string{"status", "fail_kind"})
	// TODO 安装poshgateway 默认是9091端口
	// 连接pushgateway
	pusher = push.New("http://localhost:9091", "some_job")
}

func main()  {
	// 注册指标 此时只有label的key没有value，动态对不同的value递增
	pusher.Collector(gauge)
	// 对value是pass递增2次
	gauge.WithLabelValues("pass", "").Inc()
	gauge.WithLabelValues("pass", "").Inc()
	// 对value是fail递增
	gauge.WithLabelValues("fail", "db_exec_error").Inc()

	gauge.WithLabelValues("fail", "redis_exec_error").Inc()
	// push到pushgateway
	err := pusher.Push()
	fmt.Println(err)
}



