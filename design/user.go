package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("users", func() {
	BasePath("/users")
	DefaultMedia(UserMedia)

	Security(JWT, func() {
		Scope("api:user")
	})

	Action("update", func() {
		Description("Update user by authorization token")
		Routing(PUT("/me"))
		Payload(UserMedia)
		Response(OK, UserMedia)
		Response(NotFound)
		Response(Conflict)
		Response(Unauthorized)
	})

	Action("view", func() {
		Description("Get user by authorization token")
		Routing(GET("/me"))
		Response(OK, UserMedia)
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("updatepassword", func() {
		Description("Set a new password")
		Routing(POST("/me/password"))
		Payload(func() {
			Member("oldPassword", String)
			Member("password", String)
			Required("password")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("forgotpassword", func() {
		Description("Retrieve a new password")
		Routing(POST("/forgot"))
		Payload(func() {
			Member("username", String)
			Member("securityquestion", String)
			Member("securityanswer", String)
			Required("username")
			Required("securityquestion")
			Required("securityanswer")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("getquestion", func() {
		NoSecurity()
		Description("Get security question")
		Routing(GET("/question"))
		Params(func() {
			Param("username", String, "Username")
		})
		Response(OK, String)
		Response(NotFound)
		Response(Unauthorized)
	})

	Action("register", func() {
		NoSecurity()
		Description("Register a new user")
		Routing(POST("/"))
		Payload(Register)
		Response(NoContent)
		Response(Conflict)
		Response(UnprocessableEntity)
	})

	Action("addcode", func() {
		Description("Add a voting code")
		Routing(POST("/me/codes"))
		Payload(func() {
			Member("code", String)
			Required("code")
		})
		Response(NoContent)
		Response(Unauthorized)
		Response(NotFound)
		Response(Conflict)
	})
})

//UserMedia object
var UserMedia = MediaType("application/vnd.goa.user+json", func() {
	Description("A user")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique ID")
		Attribute("status", Integer, "User status")
		Attribute("given_names", String, "Given names")
		Attribute("username", String, "Username")
		Attribute("last_names", String, "Last names")
		Attribute("location", String, "Location")
		Attribute("mail", String, "E-mail address")
		Attribute("birthdate", DateTime, "Birthdate")
		Attribute("points", Integer, "Points")
		Attribute("codes", ArrayOf(String), "Codes")
		Attribute("created_at", DateTime, "Record created timestamp")
		Attribute("updated_at", DateTime, "Record updated timestamp")
		Required("id", "username")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // have a "default" view.
		Attribute("status")
		Attribute("username")
		Attribute("given_names")
		Attribute("last_names")
		Attribute("location")
		Attribute("mail")
		Attribute("birthdate")
		Attribute("points")
		Attribute("codes")
		Attribute("created_at")
		Attribute("updated_at")
	})
})

var Register = Type("Registerrequest", func() {
	Attribute("user", String, "Login username")
	Attribute("pwd", String, "Login password")
	Attribute("fname", String, "First name")
	Attribute("lname", String, "Last name")
	Attribute("referrer", String, "Referral")
	Attribute("secret_question", String, "Secret question")
	Attribute("secret_answer", String, "Secret answer")
	Required("user")
	Required("pwd")
	Required("secret_question")
	Required("secret_answer")
})
