package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
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
	// UsersController_Addcode: start_implement

	// Put your logic here

	// UsersController_Addcode: end_implement
	return nil
}

// Register runs the register action.
func (c *UsersController) Register(ctx *app.RegisterUsersContext) error {
	// UsersController_Register: start_implement

	// Put your logic here

	// UsersController_Register: end_implement
	return nil
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

	// Put your logic here

	// UsersController_View: end_implement
	res := &app.GoaUser{}
	return ctx.OK(res)
}
