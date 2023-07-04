package profiles

import (
	"context"
	"testing"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	t.Run("should return a profile", func(t *testing.T) {
		mock, err := pgxmock.NewConn()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close(context.Background())

		mock_response := pgxmock.NewRows([]string{"username"}).AddRow("thedcsherman")
		mock.ExpectQuery(`SELECT username FROM profiles`).
			WillReturnRows(mock_response)

		profile_service := NewProfileService(mock)
		profile, err := profile_service.GetProfile(context.Background(), "1")

		assert.NoError(t, err)
		assert.Equal(t, profile, &Profile{Username: "thedcsherman"})
	})
}
