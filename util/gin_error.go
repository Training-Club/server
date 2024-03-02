package util

import "github.com/gin-gonic/gin"

// CreateError is a more elegant way to generate gin error responses
// when something fails within a request.
//
// TODO: Add functionality to hide messages when in production mode
func CreateError(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, gin.H{"message": message})
}
