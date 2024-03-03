package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"tc-server/config"
	"tc-server/util"
)

// Authorize parses and validates an auth token
// in the form of middleware. If the token is invalid
// or expired the request will be denied.
func Authorize(conf *config.FullConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "

		header := ctx.GetHeader("Authorization")
		if len(header) < 7 {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		var tokenAsString string
		var err error
		tokenAsString = header[len(BearerSchema):]

		if string(tokenAsString[0]) == `"` {
			tokenAsString, err = strconv.Unquote(tokenAsString)
		}

		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := util.ValidateToken(tokenAsString, conf.Auth.AccessTokenPub)
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			fmt.Println("Failed to validate token: " + err.Error())
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		id := claims["accountId"].(string)

		ctx.Set("accountId", id)
		ctx.Next()
	}
}
