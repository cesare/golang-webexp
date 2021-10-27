package webexp

import (
	"database/sql"
	"webexp/internal/configs"

	_ "github.com/lib/pq"
)

type Context struct {
	Config   *configs.Config
	Database *sql.DB
}

func NewContext(config *configs.Config) (*Context, error) {
	db, err := sql.Open("postgres", config.Database.ConnectionString())
	if err != nil {
		return nil, err
	}

	context := Context{
		Config:   config,
		Database: db,
	}

	return &context, nil
}
