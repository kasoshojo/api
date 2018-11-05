package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("news", func() {
	BasePath("/news")
	DefaultMedia(NewsMedia)

	Security(JWT, func() {
		Scope("api:user")
	})

	Action("list", func() {
		Description("Get news list")
		Routing(GET("/"))
		Response(OK, CollectionOf(NewsMedia))
		Response(NotFound)
		Response(Unauthorized)
	})
})

var NewsMedia = MediaType("application/vnd.goa.news+json", func() {
	Description("A news item")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "Title")
		Attribute("image_url", String, "Image url")
		Attribute("url", String, "News url")
		Attribute("date", DateTime, "Date posted")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "title", "image_url", "url", "date", "created_at")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // have a "default" view.
		Attribute("title")
		Attribute("image_url")
		Attribute("url")
		Attribute("date")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
