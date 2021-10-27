package webexp

import (
	"database/sql"
	"webexp/internal/configs"

	_ "github.com/lib/pq"
)

type Context struct {
	config   *configs.Config
	database *sql.DB
}

func NewContext(config *configs.Config) (*Context, error) {
	db, err := sql.Open("postgres", config.Database.ConnectionString())
	if err != nil {
		return nil, err
	}

	context := Context{
		config:   config,
		database: db,
	}

	return &context, nil
}

func (context *Context) Config() *configs.Config {
	return context.config
}

func (context *Context) Database() *sql.DB {
	return context.database
}
