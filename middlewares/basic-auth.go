package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(
		gin.Accounts{
			"Arbaz": "Arbaz@786",
		})
}
