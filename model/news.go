package model

import (
	"time"

	"github.com/kasoshojo/api/app"
)

type News struct {
	ID         int
	TitleEn    string
	TitleJa    string
	ImageURL   string
	ImageURLEn *string
	ImageURLJa *string
	URL        string
	URLEn      *string
	URLJa      *string
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  *time.Time
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
	ret.TitleJa = news.TitleJa
	ret.TitleEn = news.TitleEn
	ret.UpdatedAt = news.UpdatedAt
	ret.URL = news.URL
	ret.ImageURLEn = news.ImageURLEn
	ret.ImageURLJa = news.ImageURLJa
	ret.URLEn = news.URLEn
	ret.URLJa = news.URLJa
	return ret
}
