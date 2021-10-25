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
	identity := Identity{}
	query := "select id, provider_identifier, alive, registered_at from identities where id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&identity.Id, &identity.PrividerIdentifier, &identity.Alive, &identity.RegisterdAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &identity, nil
}

type RegistrationDataset struct {
	ProviderIdentifier string
}

func (r *IdentityRepository) Register(ds *RegistrationDataset) (*Identity, error) {
	return nil, nil
}
