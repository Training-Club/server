package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	GlobalController *GlobalController
	CollectionName   string
}

// ApplyAccountRoutes applies all account routes to the provided
// gin instance.
func (c *GlobalController) ApplyAccountRoutes(router *gin.Engine) {
	ac := AccountController{
		GlobalController: c,
		CollectionName:   "account",
	}

	pub := router.GET("/v1/account")
	{
		pub.GET("/availability/:key/:value", ac.GetAccountAvailability()) // Return if an account field is in available
		pub.GET("/confirm/:confirmId", ac.Confirm())                      // Confirm a confirmation for email or phone

		pub.POST("/", ac.CreateAccount()) // Create a new account
	}

	priv := router.Group("/v1/account")
	{
		priv.GET("/", ac.GetAccountByToken())               // Return account matching request token
		priv.GET("/:key/:value", ac.GetAccountByKeyValue()) // Return simple account info matching the provided key/value combo
	}
}

// GetAccountAvailability will query a key/value field to see
// if there is an existing account in the database matching
// the provided pair. In the event there is a match, the request will
// return a conflict status to inform the client this value is not available.
func (ac *AccountController) GetAccountAvailability() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}

// Confirm will parse a confirmation ID and confirm the provided
// account information matching the confirmation object in cache.
func (ac *AccountController) Confirm() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}

// CreateAccount attempts to create a new Training Club account
func (ac *AccountController) CreateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}

// GetAccountByToken queries the account attached to the requesters
// token stored in their cookies sent within the request.
func (ac *AccountController) GetAccountByToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}

// GetAccountByKeyValue queries basic account information using
// the account username or ID.
func (ac *AccountController) GetAccountByKeyValue() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	}
}
