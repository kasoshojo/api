package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
)

// MessagesController implements the messages resource.
type MessagesController struct {
	*goa.Controller
	db *gorm.DB
}

// NewMessagesController creates a messages controller.
func NewMessagesController(service *goa.Service, database *gorm.DB) *MessagesController {
	return &MessagesController{Controller: service.NewController("MessagesController"), db: database}
}

// List runs the list action.
func (c *MessagesController) List(ctx *app.ListMessagesContext) error {
	// MessagesController_List: start_implement

	// Put your logic here

	// MessagesController_List: end_implement
	res := app.GoaMessageCollection{}
	return ctx.OK(res)
}

// Read runs the read action.
func (c *MessagesController) Read(ctx *app.ReadMessagesContext) error {
	// MessagesController_Read: start_implement

	// Put your logic here

	// MessagesController_Read: end_implement
	return nil
}
