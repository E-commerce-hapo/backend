package httpx

import (
	"github.com/kiem-toan/pkg/idx"
	"github.com/kiem-toan/pkg/jsonx"
)

type IDRequest struct {
	ID idx.ID `json:"id"`
}

func (m *IDRequest) String() string { return jsonx.MustMarshalToString(m) }
