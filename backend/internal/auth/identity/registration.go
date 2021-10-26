package identity

import (
	"database/sql"
	"webexp/internal/identities"
)

type IdentityRegistration struct {
	db         *sql.DB
	identifier string
}

func NewIdentityRegistration(db *sql.DB, providerIdentifier string) *IdentityRegistration {
	return &IdentityRegistration{
		db:         db,
		identifier: providerIdentifier,
	}
}

func (r *IdentityRegistration) Execute() (*identities.Identity, error) {
	repo := identities.NewIdentityRepository(r.db)

	existingIdentity, err := repo.FindByProviderIdentifier(r.identifier)
	if err != nil {
		return nil, err
	}
	if existingIdentity != nil {
		return existingIdentity, nil
	}

	dataset := identities.RegistrationDataset{ProviderIdentifier: r.identifier}
	newIdentity, err := repo.Register(&dataset)
	if err != nil {
		return nil, err
	}

	return newIdentity, nil
}
