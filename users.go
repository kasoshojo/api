package main

import (
	"log"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
	"github.com/kasoshojo/api/model"
	"github.com/kasoshojo/api/util"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *gorm.DB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, database *gorm.DB) *UsersController {
	return &UsersController{Controller: service.NewController("UsersController"), db: database}
}

// Addcode runs the addcode action.
func (c *UsersController) Addcode(ctx *app.AddcodeUsersContext) error {
	var user model.User
	token := jwt.ContextJWT(ctx)
	claims := token.Claims.(jwtgo.MapClaims)
	userid := claims["user"]
	log.Println(userid)
	if err := c.db.Where("id = ?", userid).First(&user).Error; err != nil {
		return ctx.Unauthorized()
	}

	var code model.VoteCode
	err := c.db.Where("code = ?", ctx.Payload.Code).Find(&code).Error
	if err != nil {
		return err
	}
	code.CustomerID = &user.ID
	now := time.Now()
	code.ClaimDate = &now
	c.db.Save(&code)

	if code.Points != nil {
		user.Points = user.Points + *code.Points
		c.db.Save(&user)
	}
	return ctx.NoContent()
}

// Register runs the register action.
func (c *UsersController) Register(ctx *app.RegisterUsersContext) error {
	// UsersController_Register: start_implement

	user := model.User{}

	c.db.Where("username = ?", ctx.Payload.User).First(&user)
	if user.ID > 0 {
		return ctx.Conflict()
	}

	if ctx.Payload.Fname != nil {
		user.GivenNames = *ctx.Payload.Fname
	}
	if ctx.Payload.Lname != nil {
		user.LastNames = *ctx.Payload.Lname
	}
	user.Username = ctx.Payload.User
	user.Location = ""
	user.Password = util.HashPassword(ctx.Payload.Pwd)
	user.SecretAnswer = ctx.Payload.SecretAnswer
	user.SecretQuestion = ctx.Payload.SecretQuestion

	if err := c.db.Create(&user).Error; err != nil {
		return err
	}
	return ctx.NoContent()
}

// Update runs the update action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	// UsersController_Update: start_implement

	// Put your logic here

	// UsersController_Update: end_implement
	res := &app.GoaUser{}
	return ctx.OK(res)
}

// Updatepassword runs the updatepassword action.
func (c *UsersController) Updatepassword(ctx *app.UpdatepasswordUsersContext) error {
	// UsersController_Updatepassword: start_implement

	// Put your logic here

	// UsersController_Updatepassword: end_implement
	return nil
}

// Forgotpassword runs the forgotpassword action.
func (c *UsersController) Forgotpassword(ctx *app.ForgotpasswordUsersContext) error {
	return nil
}

// View runs the view action.
func (c *UsersController) View(ctx *app.ViewUsersContext) error {
	// UsersController_View: start_implement
	var user model.User
	token := jwt.ContextJWT(ctx)
	claims := token.Claims.(jwtgo.MapClaims)
	userid := claims["user"]
	log.Println(userid)
	if err := c.db.Where("id = ?", userid).First(&user).Error; err != nil {
		return ctx.NotFound()
	}
	// Put your logic here

	// UsersController_View: end_implement
	res := user.ToAppUser()
	return ctx.OK(&res)
}
