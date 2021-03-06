package registry

import (
	"github.com/pressly/goose/v3"
)

// MigrateMySQL ...
func (r *Registry) MigratePostgres() error {
	var err error
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	db, err := r.DB.Db.DB()
	return goose.Up(db, "./script/database_script")
}
