package httpx

import (
	"github.com/kiem-toan/infrastructure/idx"
	"github.com/kiem-toan/infrastructure/jsonx"
)

type IDRequest struct {
	ID idx.ID `json:"id"`
}

func (m *IDRequest) String() string { return jsonx.MustMarshalToString(m) }
