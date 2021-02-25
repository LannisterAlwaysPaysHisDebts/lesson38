# jaeger实现链路追踪

> jaeger文档： https://www.jaegertracing.io/docs/1.22/

启动jaeger
```
docker run -d --rm -p 6831:6831/udp -p 16686:16686 cap1573/jaeger
```

jaeger tracer
```
t, io, err := common.NewTracer(serviceName, tracerAddr)
if err != nil {
    log.Fatal(err)
}
defer io.Close()
opentracing.SetGlobalTracer(t)
```



