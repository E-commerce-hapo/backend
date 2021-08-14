package httpx

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/kiem-toan/infrastructure/auth"

	"github.com/gin-gonic/gin"
	"github.com/kiem-toan/infrastructure/errorx"
)

type Gin struct {
	C  *gin.Context
	SS *auth.SessionInfo
}

func (g *Gin) ParseRequest(p interface{}) error {
	r := g.C.Request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&p)
	if err != io.EOF {
		return err
	}
	return nil
}

func (g *Gin) Response(status int, response interface{}) {
	g.C.JSON(status, response)
}

func (g *Gin) ResponseError(err error) {
	if _err, ok := err.(*errorx.Errorx); ok {
		g.Response(_err.StatusCode, _err)
	} else {
		pp.Println(err)
		g.Response(http.StatusInternalServerError, err)
	}
}
