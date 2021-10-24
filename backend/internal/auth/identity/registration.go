package identity

import (
	"database/sql"
	"webexp/internal/identities"
)

type IdentityRegistration struct {
	db *sql.DB
}

func NewIdentityRegistration(db *sql.DB) *IdentityRegistration {
	return &IdentityRegistration{db: db}
}

func (r *IdentityRegistration) Execute() (*identities.Identity, error) {
	return nil, nil
}
