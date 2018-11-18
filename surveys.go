package main

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"github.com/jinzhu/gorm"
	"github.com/kasoshojo/api/app"
	"github.com/kasoshojo/api/model"
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

func (c *SurveysController) Vote(ctx *app.VoteSurveysContext) error {

	var user model.User
	token := jwt.ContextJWT(ctx)
	claims := token.Claims.(jwtgo.MapClaims)
	userid := claims["user"]
	if err := c.db.Where("id = ?", userid).First(&user).Error; err != nil {
		return ctx.Unauthorized()
	}

	var survey model.Survey
	err := c.db.Joins("JOIN survey_questions ON survey_questions.survey_id = surveys.id").
		Joins("JOIN survey_answers ON survey_answers.survey_question_id = survey_questions.id").
		Where("survey_questions.id = ? AND survey_answers.id = ?", ctx.Payload.QuestionID, ctx.Payload.AnswerID).Find(&survey).Error
	if err != nil {
		return err
	}
	if survey.ID == 0 {
		return ctx.NotFound()
	}
	if ctx.Payload.Weight != nil && user.Points < *ctx.Payload.Weight {
		return ctx.Unauthorized()
	}

	var result model.SurveyResult
	c.db.Where("customer_id = ? AND survey_question_id = ?").Find(&result)
	if result.ID > 0 {
		return ctx.Unauthorized()
	}

	result.AnswerText = ctx.Payload.Text
	result.CustomerID = user.ID
	result.SurveyQuestionID = ctx.Payload.QuestionID
	result.SurveyAnswerID = ctx.Payload.AnswerID
	result.AnswerWeight = ctx.Payload.Weight
	err = c.db.Save(&result).Error
	if err == nil && ctx.Payload.Weight != nil {
		user.Points = user.Points - *result.AnswerWeight
		c.db.Save(&user)
	} else {
		return c.db.Error
	}

	return ctx.NoContent()
}
