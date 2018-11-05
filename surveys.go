package main

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
)

// SurveysController implements the surveys resource.
type SurveysController struct {
	*goa.Controller
	db *gorm.DB
}

// NewSurveysController creates a surveys controller.
func NewSurveysController(service *goa.Service, database *gorm.DB) *SurveysController {
	return &SurveysController{Controller: service.NewController("SurveysController"), db: database}
}

// List runs the list action.
func (c *SurveysController) List(ctx *app.ListSurveysContext) error {
	res := app.GoaSurveyCollection{}
	return ctx.OK(res)
}

// Get runs the get action.
func (c *SurveysController) Get(ctx *app.GetSurveysContext) error {
	res := app.GoaSurvey{}
	return ctx.OK(&res)
}
