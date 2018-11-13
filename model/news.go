package model

import (
	"time"

	"github.com/kasoshojo/api/app"
)

type News struct {
	ID        int
	TitleEn   string
	TitleJa   string
	ImageURL  string
	URL       string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (News) TableName() string {
	return "news"
}

func (news *News) ToAppNews() app.GoaNews {
	ret := app.GoaNews{}
	ret.ID = news.ID
	ret.Date = news.Date
	ret.CreatedAt = news.CreatedAt
	ret.ImageURL = news.ImageURL
	ret.Title = news.TitleJa
	ret.UpdatedAt = news.UpdatedAt
	ret.URL = news.URL
	return ret
}
