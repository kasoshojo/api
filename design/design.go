package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("samclick", func() {
	Title("The kasoshojo API")
	Description("News and vote data")
	Scheme("http")
	Origin("*", func() {
		Methods("GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS") // Allow all origins to retrieve the Swagger JSON (CORS)
		Credentials()
		Headers("Content-Type", "Accept", "Origin", "Authorization")
		Expose("Authorization")
	})
	Consumes("application/json")
	Produces("application/json")
	Host("api.kasoshojo.com")
})

var _ = Resource("swagger", func() {
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("health", func() {
	Action("health", func() {
		Routing(GET("/health_check"))
		Description("Perform health check.")
		Response(OK)
	})
})

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:user", "User API access")
	Scope("api:admin", "Admin API access")
})
