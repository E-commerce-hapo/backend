package middleware

import (
	"fmt"
	"net/http"

	"github.com/kiem-toan/pkg/httpx"
	log2 "github.com/kiem-toan/pkg/log"
	"github.com/openzipkin/zipkin-go"
)

func APILoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := httpx.NewLoggingResponseWriter(w)
		span := zipkin.SpanFromContext(r.Context())
		defer func() {
			log2.Info(fmt.Sprintf("API infomation: %v [%v]", r.RequestURI, r.Method), span, map[string]interface{}{
				"method": r.Method,
				"path":   r.RequestURI,
				"status": lrw.StatusCode,
			})
		}()
		next.ServeHTTP(lrw, r)
	})
}
