package httpx

import (
	"github.com/E-commerce-hapo/backend/pkg/idx"
	"github.com/E-commerce-hapo/backend/pkg/jsonx"
)

type IDRequest struct {
	ID idx.ID `json:"id"`
}

func (m *IDRequest) String() string { return jsonx.MustMarshalToString(m) }
