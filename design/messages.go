package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("messages", func() {
	BasePath("/messages")
	DefaultMedia(MessageMedia)

	Security(JWT, func() {
		Scope("api:user")
	})

	Action("list", func() {
		Description("Get messages list")
		Routing(GET("/me"))
		Response(OK, CollectionOf(MessageMedia))
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("read", func() {
		Description("Mark message as red")
		Routing(POST("/me/:id"))
		Params(func() {
			Param("id", Integer, "Message ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
	})
})

var MessageMedia = MediaType("application/vnd.goa.message+json", func() {
	Description("A message item")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique ID")
		Attribute("title", String, "Title")
		Attribute("body", String, "Body")
		Attribute("date", DateTime, "Date posted")
		Attribute("read", Boolean, "User has read the message")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "title", "body", "date", "read", "created_at")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // have a "default" view.
		Attribute("title")
		Attribute("body")
		Attribute("date")
		Attribute("read")
		Attribute("created_at")
		Attribute("updated_at")
	})
})
