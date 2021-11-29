package errorx

import (
	"fmt"
	"strings"
)

// JSONError sẽ là cấu trúc Error API trả về cho client
type JSONError struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

func ToJSONError(errInterface ErrorInterface) *JSONError {
	return &JSONError{
		Code: fmt.Sprint(errInterface.GetCode()),
		Msg:  errInterface.Msg(),
		Meta: errInterface.MetaMap(),
	}
}

func (e *JSONError) Error() (s string) {
	if len(e.Meta) == 0 {
		return e.Msg
	}
	b := strings.Builder{}
	b.WriteString(e.Msg)
	b.WriteString(" (")
	for _, v := range e.Meta {
		b.WriteString(v)
		break
	}
	b.WriteString(")")
	return b.String()
}
