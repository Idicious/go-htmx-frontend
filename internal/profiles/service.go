package profiles

import (
	"context"
	"thedcsherman/htmx/internal/db"
)

type ProfileService struct {
	connection db.PgxIface
}

type ProfileIface interface {
	GetProfile(ctx context.Context, id string) (*Profile, error)
}

type Profile struct {
	Username string
}

func NewProfileService(connection db.PgxIface) ProfileIface {
	return &ProfileService{connection: connection}
}

func (p *ProfileService) GetProfile(ctx context.Context, id string) (*Profile, error) {
	var profile Profile

	query := `SELECT username FROM profiles;`
	err := p.connection.QueryRow(ctx, query).Scan(&profile.Username)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
