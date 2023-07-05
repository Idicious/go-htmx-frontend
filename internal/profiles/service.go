package profiles

import (
	"thedcsherman/htmx/internal/db"

	"github.com/gin-gonic/gin"
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

func (p *ProfileService) GetProfile(ctx *gin.Context, id string) (Profile, error) {
	var profile Profile

	query := `SELECT username FROM profiles;`
	err := p.connection.QueryRow(ctx, query).Scan(&profile.Username)

	if err != nil {
		return profile, err
	}

	return profile, nil
}
