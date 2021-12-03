package zipkin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kiem-toan/core/config"

	"os"

	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	host = "localhost"
)

func NewTracer() (*zipkin.Tracer, error) {
	var endpointURL = config.GetAppConfig().Zipkin.URL + "/api/v2/spans"
	// The reporter sends traces to zipkin server
	reporter := reporterhttp.NewReporter(endpointURL)
	hostPort := fmt.Sprintf("%v:%v", host, config.GetAppConfig().ServerPort)
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
	// We add the instrumented transport to the defaultClient
	// that comes with the zipkin-go library
	http.DefaultClient.Transport, err = zipkinhttp.NewTransport(
		t,
		zipkinhttp.TransportTrace(true),
	)
	if err != nil {
		log.Fatal(err, nil, nil)
	}
	return t, err
}
