package profiles

import (
	"net/http"
	"thedcsherman/htmx/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateProfileRoutes(r *gin.Engine, profile_service ProfileIface) {
	r.GET("/api/profile", func(c *gin.Context) {
		profile, err := profile_service.GetProfile(c, "1")
		if err != nil {
			utils.ReturnError(c, err, "Unable to get profile")
			return
		}

		c.HTML(http.StatusOK, "data.html", gin.H{
			"Data": profile.Username,
		})
	})
}
