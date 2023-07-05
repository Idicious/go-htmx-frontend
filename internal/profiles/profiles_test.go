package profiles

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	t.Run("should return a profile", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		mock, err := pgxmock.NewConn()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close(ctx)

		mock_response := pgxmock.NewRows([]string{"username"}).AddRow("thedcsherman")
		mock.ExpectQuery(`SELECT username FROM profiles`).
			WillReturnRows(mock_response)

		profile_service := NewProfileService(mock)
		profile, err := profile_service.GetProfile(ctx, "1")

		assert.NoError(t, err)
		assert.Equal(t, profile, Profile{Username: "thedcsherman"})
	})

	t.Run("should return an error", func(t *testing.T) {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		mock, err := pgxmock.NewConn()
		if err != nil {
			t.Fatal(err)
		}
		defer mock.Close(ctx)

		mock.ExpectQuery(`SELECT username FROM profiles`).
			WillReturnError(errors.New("db error"))

		profile_service := NewProfileService(mock)
		profile, err := profile_service.GetProfile(ctx, "1")

		assert.Error(t, err, errors.New("db error"))
		assert.Equal(t, profile, Profile{})
	})
}
