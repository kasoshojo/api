package main

import (
	"crypto/rsa"
	"io/ioutil"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
	"github.com/kasoshojo/api/model"
	"github.com/kasoshojo/api/util"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db         *gorm.DB
	privateKey *rsa.PrivateKey
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, database *gorm.DB) *AuthController {
	b, err := ioutil.ReadFile("./key/jwtRS256.key")
	if err != nil {
		return nil
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	return &AuthController{Controller: service.NewController("AuthController"), privateKey: privKey, db: database}
}

// Secure runs the secure action.
func (c *AuthController) Secure(ctx *app.SecureAuthContext) error {
	return ctx.OK(nil)
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	user := ctx.Payload.Username
	hash := util.HashPassword(ctx.Payload.Password)

	var u model.User

	if err := c.db.Where("username = ? AND password = ?", user, hash).First(&u).Error; err != nil {
		return ctx.Unauthorized()
	}
	// Generate JWT

	signedToken := util.GenerateJWTToken(u.ID, &u.Username, c.privateKey, []string{"api:user"})

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	// Send response
	return ctx.NoContent()
}
