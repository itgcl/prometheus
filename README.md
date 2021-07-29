# prometheus
go prometheus


//todo readme样式待完善

prometheus监控指标分为两种
1.prometheus pull, 通过服务端暴露http接口, prometheus定时查询接口指标。
2.服务端主动push, 将指标信息push到一个中间网关(pushgateway), prometheus定时去中间网关拉取数据。 
