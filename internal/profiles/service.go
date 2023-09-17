package profiles

import (
	"context"
	"fmt"
	"thedcsherman/htmx/internal/db"
)

type ProfileService struct {
	connection db.PgxIface
}

type Profile struct {
	Username string
}

func NewProfileService(connection db.PgxIface) *ProfileService {
	return &ProfileService{connection: connection}
}

func (p *ProfileService) GetProfile(ctx context.Context, id string) (*Profile, error) {
	var profile Profile

	query := fmt.Sprintf(`SELECT username FROM profiles WHERE ID = %s;`, id)
	err := p.connection.QueryRow(ctx, query).Scan(&profile.Username)

	if err != nil {
		return &profile, err
	}

	return &profile, nil
}
