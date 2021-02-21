package common

import (
	"github.com/uber/jaeger-client-go"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	return cfg.NewTracer()
}
