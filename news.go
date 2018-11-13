package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
	"github.com/kasoshojo/api/model"
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
	var news []model.News
	c.db.Find(&news)
	res := app.GoaNewsCollection{}

	for _, n := range news {
		temp := n.ToAppNews()
		res = append(res, &temp)
	}
	return ctx.OK(res)
}
