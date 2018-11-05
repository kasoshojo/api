// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "samclick": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/kasoshojo/api/design
// --out=$(GOPATH)src/github.com/kasoshojo/api
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthController is the controller interface for the Auth actions.
type AuthController interface {
	goa.Muxer
	Secure(*SecureAuthContext) error
	Signin(*SigninAuthContext) error
}

// MountAuthController "mounts" a Auth resource controller on the given service.
func MountAuthController(service *goa.Service, ctrl AuthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/auth/", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/auth/signin", ctrl.MuxHandler("preflight", handleAuthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSecureAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Secure(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleAuthOrigin(h)
	service.Mux.Handle("GET", "/auth/", ctrl.MuxHandler("secure", h, nil))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Secure", "route", "GET /auth/", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSigninAuthContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*SigninAuthPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Signin(rctx)
	}
	h = handleAuthOrigin(h)
	service.Mux.Handle("POST", "/auth/signin", ctrl.MuxHandler("signin", h, unmarshalSigninAuthPayload))
	service.LogInfo("mount", "ctrl", "Auth", "action", "Signin", "route", "POST /auth/signin")
}

// handleAuthOrigin applies the CORS response headers corresponding to the origin.
func handleAuthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalSigninAuthPayload unmarshals the request body into the context request data Payload field.
func unmarshalSigninAuthPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &signinAuthPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Health(*HealthHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/health_check", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/health_check", ctrl.MuxHandler("health", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Health", "route", "GET /health_check")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// MessagesController is the controller interface for the Messages actions.
type MessagesController interface {
	goa.Muxer
	List(*ListMessagesContext) error
	Read(*ReadMessagesContext) error
}

// MountMessagesController "mounts" a Messages resource controller on the given service.
func MountMessagesController(service *goa.Service, ctrl MessagesController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/messages/me", ctrl.MuxHandler("preflight", handleMessagesOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/messages/me/:id", ctrl.MuxHandler("preflight", handleMessagesOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListMessagesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleMessagesOrigin(h)
	service.Mux.Handle("GET", "/messages/me", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Messages", "action", "List", "route", "GET /messages/me", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewReadMessagesContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Read(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleMessagesOrigin(h)
	service.Mux.Handle("POST", "/messages/me/:id", ctrl.MuxHandler("read", h, nil))
	service.LogInfo("mount", "ctrl", "Messages", "action", "Read", "route", "POST /messages/me/:id", "security", "jwt")
}

// handleMessagesOrigin applies the CORS response headers corresponding to the origin.
func handleMessagesOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// NewsController is the controller interface for the News actions.
type NewsController interface {
	goa.Muxer
	List(*ListNewsContext) error
}

// MountNewsController "mounts" a News resource controller on the given service.
func MountNewsController(service *goa.Service, ctrl NewsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/news/", ctrl.MuxHandler("preflight", handleNewsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListNewsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleNewsOrigin(h)
	service.Mux.Handle("GET", "/news/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "News", "action", "List", "route", "GET /news/", "security", "jwt")
}

// handleNewsOrigin applies the CORS response headers corresponding to the origin.
func handleNewsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SurveysController is the controller interface for the Surveys actions.
type SurveysController interface {
	goa.Muxer
	Get(*GetSurveysContext) error
	List(*ListSurveysContext) error
}

// MountSurveysController "mounts" a Surveys resource controller on the given service.
func MountSurveysController(service *goa.Service, ctrl SurveysController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/surveys/:id", ctrl.MuxHandler("preflight", handleSurveysOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/surveys/", ctrl.MuxHandler("preflight", handleSurveysOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetSurveysContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleSurveysOrigin(h)
	service.Mux.Handle("GET", "/surveys/:id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Surveys", "action", "Get", "route", "GET /surveys/:id", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListSurveysContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleSurveysOrigin(h)
	service.Mux.Handle("GET", "/surveys/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Surveys", "action", "List", "route", "GET /surveys/", "security", "jwt")
}

// handleSurveysOrigin applies the CORS response headers corresponding to the origin.
func handleSurveysOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// UsersController is the controller interface for the Users actions.
type UsersController interface {
	goa.Muxer
	Addcode(*AddcodeUsersContext) error
	Forgotpassword(*ForgotpasswordUsersContext) error
	Register(*RegisterUsersContext) error
	Update(*UpdateUsersContext) error
	Updatepassword(*UpdatepasswordUsersContext) error
	View(*ViewUsersContext) error
}

// MountUsersController "mounts" a Users resource controller on the given service.
func MountUsersController(service *goa.Service, ctrl UsersController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/users/me/codes", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/forgot", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/me", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/users/me/password", ctrl.MuxHandler("preflight", handleUsersOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddcodeUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddcodeUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Addcode(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/users/me/codes", ctrl.MuxHandler("addcode", h, unmarshalAddcodeUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Addcode", "route", "POST /users/me/codes", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewForgotpasswordUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ForgotpasswordUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Forgotpassword(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/users/forgot", ctrl.MuxHandler("forgotpassword", h, unmarshalForgotpasswordUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Forgotpassword", "route", "POST /users/forgot", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*Registerrequest)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Register(rctx)
	}
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/users/", ctrl.MuxHandler("register", h, unmarshalRegisterUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Register", "route", "POST /users/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleUsersOrigin(h)
	service.Mux.Handle("PUT", "/users/me", ctrl.MuxHandler("update", h, unmarshalUpdateUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Update", "route", "PUT /users/me", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdatepasswordUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdatepasswordUsersPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Updatepassword(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleUsersOrigin(h)
	service.Mux.Handle("POST", "/users/me/password", ctrl.MuxHandler("updatepassword", h, unmarshalUpdatepasswordUsersPayload))
	service.LogInfo("mount", "ctrl", "Users", "action", "Updatepassword", "route", "POST /users/me/password", "security", "jwt")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewViewUsersContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.View(rctx)
	}
	h = handleSecurity("jwt", h, "api:user")
	h = handleUsersOrigin(h)
	service.Mux.Handle("GET", "/users/me", ctrl.MuxHandler("view", h, nil))
	service.LogInfo("mount", "ctrl", "Users", "action", "View", "route", "GET /users/me", "security", "jwt")
}

// handleUsersOrigin applies the CORS response headers corresponding to the origin.
func handleUsersOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Expose-Headers", "Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, Authorization")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddcodeUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddcodeUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addcodeUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalForgotpasswordUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalForgotpasswordUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &forgotpasswordUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalRegisterUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &registerrequest{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdatepasswordUsersPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdatepasswordUsersPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updatepasswordUsersPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}