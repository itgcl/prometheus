# prometheus

> prometheus是由go语言开发的一个开源的监控系统。



#### 安装包

```go
go get github.com/prometheus/client_golang/prometheus
```



本文只是记录prometheus pull 和 push 方式。

至于如何安装prometheus，各种指标类型区别等等...请查看 [prometheus官方文档](https://prometheus.io/docs/introduction/overview/)



#### 架构图

![20180906141552432](C:\Users\dell\Desktop\20180906141552432.png)

#### 指标计数使用方式

prometheus监控指标分为两种

1. 主动pull，通过服务端暴露http接口，prometheus定时查询接口获取指标数据。

   ```go
   import (
   	"github.com/prometheus/client_golang/prometheus"
   	"github.com/prometheus/client_golang/prometheus/promhttp"
   	"net/http"
   )
   
   func main() {
       // 暴露接口 prometheus定时获取指标
   	http.Handle("/metrics", promhttp.Handler())
   }
   ```

   > 示例代码： [暴露http接口方式](https://github.com/itgcl/prometheus/example/http)

   

2. 服务端主动push，将指标信息push到一个中间网关**(pushgateway)**， prometheus定时去中间网关拉取数据。

   ```go
   import (
   	"fmt"
   	"github.com/prometheus/client_golang/prometheus"
   	"github.com/prometheus/client_golang/prometheus/push"
   )
   // 连接pushgateway
   pusher = push.New("http://localhost:9091", "some_job")
   
   // TODO counter 是自定义的指标类型 详情见实例代码
   // 注册指标和http方式不同
   pusher.Collector(counter)
   // 指标计数
   counter.Inc() // or counter.Add(1)
   // 提交指标信息
   err := pusher.Push()
   ```

   > pushgateway需要安装的，这里跳过。
   >
   > 示例代码： [push方式](https://github.com/itgcl/prometheus/example/pushgateway/v1)



#### 特殊情况

如果注册相同的指标信息会报错，例如:

```go
duplicate metrics collector registration attempted
```

业务上可能需要动态注册指标信息，就会导致可能会出现注册相同的指标信息问题。

###### 解决方式

```go
// 按不同维度划分的同一个东西, 只注册label的key, 通过动态增加value
gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
   Namespace: "",
   Subsystem: "push",
   Name:      "some",
   Help:      "some push", // 只是描述
}, []string{"status", "fail_kind"})

// 对value是pass递增2次
gauge.WithLabelValues("pass", "").Inc()
gauge.WithLabelValues("pass", "").Inc()

// 对value是fail递增
gauge.WithLabelValues("fail", "db_exec_error").Inc()

```

示例代码： [解决注册相同指标信息问题](https://github.com/itgcl/prometheus/example/pushgateway/v2)