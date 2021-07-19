package httpx

import (
	"encoding/json"
	"github.com/kiem-toan/infrastructure/errorx"
	"io"
	"net/http"
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

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Can not marshal response`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func RespondError(w http.ResponseWriter, err error) {
	if _err, ok := err.(*errorx.Errorx); ok {
		RespondJSON(w, _err.StatusCode, err)
	} else {
		RespondJSON(w, http.StatusInternalServerError, err)
	}

}
