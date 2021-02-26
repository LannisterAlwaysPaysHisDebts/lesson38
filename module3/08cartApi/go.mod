module github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/08cartApi

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/05common v0.0.0-20210225173307-605f71bc8d38
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/micro v1.18.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/common v0.17.0
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
