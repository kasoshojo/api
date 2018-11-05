package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
)

// NewsController implements the news resource.
type NewsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewNewsController creates a news controller.
func NewNewsController(service *goa.Service, database *gorm.DB) *NewsController {
	return &NewsController{Controller: service.NewController("NewsController"), db: database}
}

// List runs the list action.
func (c *NewsController) List(ctx *app.ListNewsContext) error {
	// NewsController_List: start_implement

	// Put your logic here

	// NewsController_List: end_implement
	res := app.GoaNewsCollection{}
	return ctx.OK(res)
}
