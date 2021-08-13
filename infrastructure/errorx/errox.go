package errorx

type Errorx struct {
	StatusCode int    `json:"code"`
	Err        error  `json:"error"`
	Msg        string `json:"msg"`
}

func New(statusCode int, err error, msg string) *Errorx {
	return &Errorx{
		StatusCode: statusCode,
		Err:        err,
		Msg:        msg,
	}
}

func (e *Errorx) Error() string {
	return e.Msg
}
