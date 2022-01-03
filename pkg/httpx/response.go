package httpx

import (
	"net/http"

	"github.com/E-commerce-hapo/backend/pkg/jsonx"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type DeletedResponse struct {
	Deleted int `json:"deleted"`
}

func (m *DeletedResponse) String() string { return jsonx.MustMarshalToString(m) }

type UpdatedResponse struct {
	Updated int `json:"updated"`
}

func (m *UpdatedResponse) String() string { return jsonx.MustMarshalToString(m) }

type CreatedResponse struct {
	Created int `json:"created"`
}

func (m *CreatedResponse) String() string { return jsonx.MustMarshalToString(m) }
