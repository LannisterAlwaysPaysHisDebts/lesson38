# cart / cartApi

熔断、限流、负载均衡

## 熔断
熔断器添加在客户端； 使用hystrix，import: `"github.com/afex/hystrix-go/hystrix"`。启动：
```
hystrixStreamHandler := hystrix.NewStreamHandler()
hystrixStreamHandler.Start()

// 注册熔断器的监听端口
go func() {
    err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", " "), hystrixStreamHandler)
    if err != nil {
        log.Error(err)
    }
}()
```

编写熔断方法:
```
type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		fmt.Println(e)
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{}
	}
}
```

在go-micro里添加wrapper
```
micro.NewService(
    micro.WrapClient(NewClientHystrixWrapper()),
)
```

运行dashboard
```
docker pull cap1573/hystrix-dashboard
docker run -d -p 9002:9002 cap1573/hystrix-dashboard
```
访问dashboard： http://192.168.199.198:9002/hystrix

## 限流
限流代码在*服务端*添加；使用go-micro的ratelimit插件, import: `ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"`
在micro里添加:
```
QPS := 100
micro.NewService(
    micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS))
)
```


## 负载均衡
负载均衡写在*客户端*,
import: `"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"`

在go-micro里添加wrapper
```
micro.NewService(
    micro.WrapClient(roundrobin.NewClientWrapper()),
)
```



