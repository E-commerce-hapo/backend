package zipkin

import (
	"fmt"

	"github.com/kiem-toan/pkg/config"

	"os"

	"github.com/openzipkin/zipkin-go"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	host = "localhost"
)

func NewTracer() (*zipkin.Tracer, error) {
	var endpointURL = "config.GetAppConfig().Zipkin.URL" + "/api/v2/spans"
	// The reporter sends traces to zipkin server
	reporter := reporterhttp.NewReporter(endpointURL)
	hostPort := fmt.Sprintf("%v:%v", host, config.Port)
	localEndpoint, err := zipkin.NewEndpoint(os.Getenv("APPLICATION_NAME"), hostPort)
	if err != nil {
		return nil, err
	}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}
	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
	)
	if err != nil {
		return nil, err
	}
	return t, err
}
