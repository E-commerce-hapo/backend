package httpx

import "github.com/kiem-toan/pkg/jsonx"

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
