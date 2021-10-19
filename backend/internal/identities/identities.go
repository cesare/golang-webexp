package identities

import "time"

type Identity struct {
	Id                 string
	PrividerIdentifier string
	Alive              bool
	RegisterdAt        time.Time
}
