package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("news", func() {
	BasePath("/news")
	DefaultMedia(NewsMedia)
	NoSecurity()

	Action("list", func() {
		Description("Get news list")
		Routing(GET("/"))
		Response(OK, CollectionOf(NewsMedia))
		Response(NotFound)
		Response(Unauthorized)
	})
})

/*var NewsPageMedia = MediaType("application/vnd.goa.newspage+json", func() {
	Description("The news page")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("video_url", String, "video url")
		Attribute("video_url_en", String, "video url english")
		Attribute("homepage_url", String, "homepage url")
		Attribute("homepage_url_en", String, "homepage url english")
		Attribute("news", CollectionOf(NewsMedia), "news items")
		Required("video_url", "video_url_en", "homepage_url", "homepage_url_en", "news")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("video_url")
		Attribute("video_url_en")
		Attribute("homepage_url")
		Attribute("homepage_url_en")
		Attribute("news")
	})
})*/

var NewsMedia = MediaType("application/vnd.goa.news+json", func() {
	Description("A news item")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique ID")
		Attribute("title_en", String, "Title")
		Attribute("title_ja", String, "Title")
		Attribute("image_url", String, "Image url")
		Attribute("image_url_en", String, "Image url")
		Attribute("image_url_ja", String, "Image url")
		Attribute("url", String, "News url")
		Attribute("url_en", String, "News url")
		Attribute("url_ja", String, "News url")
		Attribute("date", DateTime, "Date posted")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "title_en", "title_ja", "image_url", "url", "date", "created_at")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // have a "default" view.
		Attribute("title_en")
		Attribute("title_ja")
		Attribute("image_url")
		Attribute("image_url_en")
		Attribute("image_url_ja")
		Attribute("url")
		Attribute("url_en")
		Attribute("url_ja")
		Attribute("date")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
