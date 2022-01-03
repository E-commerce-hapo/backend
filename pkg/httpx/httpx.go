package httpx

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/E-commerce-hapo/backend/pkg/errorx"
)

func ParseRequest(r *http.Request, p interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != io.EOF {
		return err
	}
	return nil
}

func WriteError(ctx context.Context, w http.ResponseWriter, err error) {
	errInterface := errorx.ToErrorInterface(err)
	statusCode := errInterface.GetCode()
	jsonErr := errorx.ToJSONError(errInterface)
	errBody, err := json.Marshal(&jsonErr)
	if err != nil {
		errBody = []byte("{\"type\": \"internal\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}
	w.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	w.Header().Set("Content-Length", strconv.Itoa(len(errBody)))
	w.WriteHeader(statusCode) // set HTTP status code and send response

	_, writeErr := w.Write(errBody)
	if writeErr != nil {
		_ = writeErr
	}
}

func WriteReponse(ctx context.Context, w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Can not marshal response`))
		return
	}
	w.WriteHeader(status)
	w.Write(response)
}
