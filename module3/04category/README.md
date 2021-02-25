# category分类 micro service
consul实现配置中心与服务注册中心

## 配置中心
> package：
> github.com/micro/go-micro/v2/config
> "github.com/micro/go-plugins/config/source/consul/v2"

见`./common/config.go`;代码写完之后记得在consul里添加对应的配置； 


## 服务注册中心
> package:
> github.com/micro/go-plugins/registry/consul/v2
```
// 注册中心
consul.NewRegistry(func(options *registry.Options) {
    options.Addrs = []string{
        consulHost + ":8500",
    }
})
micro.NewService(micro.Registry(consulRegister))
```


## consul服务配置
### server端
拉取consul容器: `docker pull consul`；

#### 容器执行
server1: 
```
docker run -d --net=host -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}' \
 consul agent -server -bind=192.168.199.198 -bootstrap-expect=2 -ui -client=0.0.0.0 -node=192.168.199.198
```
docker参数`--net=host`代表网络是共享主机的network；

consul参数`-bind`代表consul监听的地址；
`-server`代表该consul服务是服务端；
`-ui`: 启动图形服务

server2:
```
docker run -d --net=host -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}' \
 consul agent -server -bind=192.168.199.197 -bootstrap-expect=2 -ui -retry-join=192.168.199.198 -node=192.168.199.197
```
`-retry-join`: 启动时要加入的另一个代理的地址

查看状态：
```
docker exec -t <containerId> consul members
```


