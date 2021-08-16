package errorx

import (
	"fmt"
	"strings"
)

type Errorx struct {
	StatusCode int    `json:"code"`
	Err        error  `json:"error"`
	Msg        string `json:"msg"`
}

func New(statusCode int, err error, msg string, args ...interface{}) *Errorx {
	return &Errorx{
		StatusCode: statusCode,
		Err:        err,
		Msg:        fmt.Sprintf(msg, args...),
	}
}

func (e *Errorx) Error() string {
	return e.Msg
}

type Errors []error

func (errs Errors) ToError() error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}

func (errs Errors) Error() string {
	switch len(errs) {
	case 0:
		return ""
	case 1:
		e := errs[0]
		if e == nil {
			return "ok"
		}
		return e.Error()
	}

	var b strings.Builder
	for i, e := range errs {
		if i > 0 {
			b.WriteString("; ")
		}
		if e == nil {
			b.WriteString("ok")
		} else {
			b.WriteString(e.Error())
		}
	}
	return b.String()
}

func (errs Errors) NErrors() int {
	c := 0
	for _, err := range errs {
		if err != nil {
			c++
		}
	}
	return c
}

func (errs Errors) IsAll() bool {
	if len(errs) == 0 {
		return false
	}
	for _, err := range errs {
		if err == nil {
			return false
		}
	}
	return true
}

func (errs Errors) HasAny() bool {
	for _, err := range errs {
		if err != nil {
			return true
		}
	}
	return false
}

func (errs Errors) All() error {
	if errs.IsAll() {
		return errs
	}
	return nil
}

func (errs Errors) Any() error {
	if errs.HasAny() {
		return errs
	}
	return nil
}

func (errs Errors) Last() error {
	if len(errs) == 0 {
		return nil
	}
	return errs[len(errs)-1]
}
