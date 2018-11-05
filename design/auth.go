package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("auth", func() {
	Description("This resource uses JWT to secure its endpoints")
	BasePath("/auth")

	Security(JWT, func() {
		Scope("api:user")
	})

	Action("signin", func() {
		Description("Creates a valid JWT")
		NoSecurity()
		Routing(POST("/signin"))
		Payload(func() {
			Member("username", String)
			Member("password", String)
			Required("username")
			Required("password")
		})
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})

	Action("secure", func() {
		Description("This action will test authentication")
		Routing(GET("/"))
		Response(OK)
		Response(Unauthorized)
	})
})
