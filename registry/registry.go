package registry

import (
	"github.com/E-commerce-hapo/backend/core/config"
	"github.com/E-commerce-hapo/backend/pkg/database"
	"github.com/E-commerce-hapo/backend/pkg/event/dispatcher"
	"github.com/E-commerce-hapo/backend/thirdparty/email"
)

// Registry khởi tạo các client, config mà ứng dụng sử dụng...
// Dùng các client, config đã được khởi ở trên để khởi tạo các service Aggregate và Query của ứng dụng...
type Registry struct {
	Config      config.Config
	DB          *database.Database
	Dispatcher  *dispatcher.Dispatcher
	EmailClient *email.Client
}

// New ...
func New(c config.Config) (*Registry, error) {
	r := &Registry{
		Config:      c,
		DB:          database.New(c),
		Dispatcher:  dispatcher.NewDispatcher(),
		EmailClient: email.New(&email.SMTPConfig{}),
	}
	return r, nil
}
