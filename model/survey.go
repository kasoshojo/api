package model

import (
	"time"
)

type SurveyResult struct {
	ID               int
	CustomerID       int
	SurveyQuestionID int
	SurveyAnswerID   int
	AnswerWeight     *int
	AnswerText       *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

func (SurveyResult) TableName() string {
	return "survey_results"
}
