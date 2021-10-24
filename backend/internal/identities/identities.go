package identities

import (
	"database/sql"
	"time"
)

type Identity struct {
	Id                 string
	PrividerIdentifier string
	Alive              bool
	RegisterdAt        time.Time
}

type IdentityRepository struct {
	db *sql.DB
}

func NewIdentityRepository(db *sql.DB) *IdentityRepository {
	return &IdentityRepository{db: db}
}

func (r *IdentityRepository) Find(id string) (*Identity, error) {
	return nil, nil
}

type RegistrationDataset struct {
	ProviderIdentifier string
}

func (r *IdentityRepository) Register(ds *RegistrationDataset) (*Identity, error) {
	return nil, nil
}
