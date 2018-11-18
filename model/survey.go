package model

import (
	"time"
)

type Survey struct {
	ID     int
	Type   int
	Status *int
}

func (Survey) TableName() string {
	return "surveys"
}

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
