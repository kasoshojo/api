// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "samclick": Application Contexts
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
	"net/http"
	"strconv"
	"time"
)

// SecureAuthContext provides the auth secure action context.
type SecureAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewSecureAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller secure action.
func NewSecureAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*SecureAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SecureAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SecureAuthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SecureAuthContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// SigninAuthContext provides the auth signin action context.
type SigninAuthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *SigninAuthPayload
}

// NewSigninAuthContext parses the incoming request URL and body, performs validations and creates the
// context used by the auth controller signin action.
func NewSigninAuthContext(ctx context.Context, r *http.Request, service *goa.Service) (*SigninAuthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := SigninAuthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// signinAuthPayload is the auth signin action payload.
type signinAuthPayload struct {
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *signinAuthPayload) Validate() (err error) {
	if payload.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates SigninAuthPayload from signinAuthPayload
func (payload *signinAuthPayload) Publicize() *SigninAuthPayload {
	var pub SigninAuthPayload
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	if payload.Username != nil {
		pub.Username = *payload.Username
	}
	return &pub
}

// SigninAuthPayload is the auth signin action payload.
type SigninAuthPayload struct {
	Password string `form:"password" json:"password" xml:"password"`
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate runs the validation rules defined in the design.
func (payload *SigninAuthPayload) Validate() (err error) {
	if payload.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *SigninAuthContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *SigninAuthContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// HealthHealthContext provides the health health action context.
type HealthHealthContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHealthHealthContext parses the incoming request URL and body, performs validations and creates the
// context used by the health controller health action.
func NewHealthHealthContext(ctx context.Context, r *http.Request, service *goa.Service) (*HealthHealthContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := HealthHealthContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *HealthHealthContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// ListMessagesContext provides the messages list action context.
type ListMessagesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListMessagesContext parses the incoming request URL and body, performs validations and creates the
// context used by the messages controller list action.
func NewListMessagesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListMessagesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListMessagesContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListMessagesContext) OK(r GoaMessageCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.message+json; type=collection")
	if r == nil {
		r = GoaMessageCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListMessagesContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListMessagesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ReadMessagesContext provides the messages read action context.
type ReadMessagesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewReadMessagesContext parses the incoming request URL and body, performs validations and creates the
// context used by the messages controller read action.
func NewReadMessagesContext(ctx context.Context, r *http.Request, service *goa.Service) (*ReadMessagesContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ReadMessagesContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *ReadMessagesContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ReadMessagesContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ReadMessagesContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListNewsContext provides the news list action context.
type ListNewsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListNewsContext parses the incoming request URL and body, performs validations and creates the
// context used by the news controller list action.
func NewListNewsContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListNewsContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListNewsContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListNewsContext) OK(r GoaNewsCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.news+json; type=collection")
	if r == nil {
		r = GoaNewsCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListNewsContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListNewsContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// GetSurveysContext provides the surveys get action context.
type GetSurveysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewGetSurveysContext parses the incoming request URL and body, performs validations and creates the
// context used by the surveys controller get action.
func NewGetSurveysContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetSurveysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetSurveysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetSurveysContext) OK(r *GoaSurvey) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.survey+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *GetSurveysContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetSurveysContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListSurveysContext provides the surveys list action context.
type ListSurveysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListSurveysContext parses the incoming request URL and body, performs validations and creates the
// context used by the surveys controller list action.
func NewListSurveysContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListSurveysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListSurveysContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListSurveysContext) OK(r GoaSurveyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.survey+json; type=collection")
	if r == nil {
		r = GoaSurveyCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListSurveysContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListSurveysContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// VoteSurveysContext provides the surveys vote action context.
type VoteSurveysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      int
	Payload *SurveyResultPayload
}

// NewVoteSurveysContext parses the incoming request URL and body, performs validations and creates the
// context used by the surveys controller vote action.
func NewVoteSurveysContext(ctx context.Context, r *http.Request, service *goa.Service) (*VoteSurveysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := VoteSurveysContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *VoteSurveysContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *VoteSurveysContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *VoteSurveysContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// AddcodeUsersContext provides the users addcode action context.
type AddcodeUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AddcodeUsersPayload
}

// NewAddcodeUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller addcode action.
func NewAddcodeUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*AddcodeUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := AddcodeUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// addcodeUsersPayload is the users addcode action payload.
type addcodeUsersPayload struct {
	Code *string `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *addcodeUsersPayload) Validate() (err error) {
	if payload.Code == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "code"))
	}
	return
}

// Publicize creates AddcodeUsersPayload from addcodeUsersPayload
func (payload *addcodeUsersPayload) Publicize() *AddcodeUsersPayload {
	var pub AddcodeUsersPayload
	if payload.Code != nil {
		pub.Code = *payload.Code
	}
	return &pub
}

// AddcodeUsersPayload is the users addcode action payload.
type AddcodeUsersPayload struct {
	Code string `form:"code" json:"code" xml:"code"`
}

// Validate runs the validation rules defined in the design.
func (payload *AddcodeUsersPayload) Validate() (err error) {
	if payload.Code == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "code"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *AddcodeUsersContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *AddcodeUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *AddcodeUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// Conflict sends a HTTP response with status code 409.
func (ctx *AddcodeUsersContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// ForgotpasswordUsersContext provides the users forgotpassword action context.
type ForgotpasswordUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ForgotpasswordUsersPayload
}

// NewForgotpasswordUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller forgotpassword action.
func NewForgotpasswordUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*ForgotpasswordUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ForgotpasswordUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// forgotpasswordUsersPayload is the users forgotpassword action payload.
type forgotpasswordUsersPayload struct {
	Securityanswer   *string `form:"securityanswer,omitempty" json:"securityanswer,omitempty" xml:"securityanswer,omitempty"`
	Securityquestion *string `form:"securityquestion,omitempty" json:"securityquestion,omitempty" xml:"securityquestion,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *forgotpasswordUsersPayload) Validate() (err error) {
	if payload.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Securityquestion == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "securityquestion"))
	}
	if payload.Securityanswer == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "securityanswer"))
	}
	return
}

// Publicize creates ForgotpasswordUsersPayload from forgotpasswordUsersPayload
func (payload *forgotpasswordUsersPayload) Publicize() *ForgotpasswordUsersPayload {
	var pub ForgotpasswordUsersPayload
	if payload.Securityanswer != nil {
		pub.Securityanswer = *payload.Securityanswer
	}
	if payload.Securityquestion != nil {
		pub.Securityquestion = *payload.Securityquestion
	}
	if payload.Username != nil {
		pub.Username = *payload.Username
	}
	return &pub
}

// ForgotpasswordUsersPayload is the users forgotpassword action payload.
type ForgotpasswordUsersPayload struct {
	Securityanswer   string `form:"securityanswer" json:"securityanswer" xml:"securityanswer"`
	Securityquestion string `form:"securityquestion" json:"securityquestion" xml:"securityquestion"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate runs the validation rules defined in the design.
func (payload *ForgotpasswordUsersPayload) Validate() (err error) {
	if payload.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	if payload.Securityquestion == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "securityquestion"))
	}
	if payload.Securityanswer == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "securityanswer"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *ForgotpasswordUsersContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ForgotpasswordUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ForgotpasswordUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// RegisterUsersContext provides the users register action context.
type RegisterUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *Registerrequest
}

// NewRegisterUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller register action.
func NewRegisterUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*RegisterUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := RegisterUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *RegisterUsersContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Conflict sends a HTTP response with status code 409.
func (ctx *RegisterUsersContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// UnprocessableEntity sends a HTTP response with status code 422.
func (ctx *RegisterUsersContext) UnprocessableEntity() error {
	ctx.ResponseData.WriteHeader(422)
	return nil
}

// UpdateUsersContext provides the users update action context.
type UpdateUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UpdateUsersPayload
}

// NewUpdateUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller update action.
func NewUpdateUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// updateUsersPayload is the users update action payload.
type updateUsersPayload struct {
	// Birthdate
	Birthdate *time.Time `form:"birthdate,omitempty" json:"birthdate,omitempty" xml:"birthdate,omitempty"`
	// Codes
	Codes []string `form:"codes,omitempty" json:"codes,omitempty" xml:"codes,omitempty"`
	// Record created timestamp
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Given names
	GivenNames *string `form:"given_names,omitempty" json:"given_names,omitempty" xml:"given_names,omitempty"`
	// Unique ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Last names
	LastNames *string `form:"last_names,omitempty" json:"last_names,omitempty" xml:"last_names,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// E-mail address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Points
	Points *int `form:"points,omitempty" json:"points,omitempty" xml:"points,omitempty"`
	// User status
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Record updated timestamp
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUsersPayload) Validate() (err error) {
	if payload.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "id"))
	}
	if payload.Username == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	return
}

// Publicize creates UpdateUsersPayload from updateUsersPayload
func (payload *updateUsersPayload) Publicize() *UpdateUsersPayload {
	var pub UpdateUsersPayload
	if payload.Birthdate != nil {
		pub.Birthdate = payload.Birthdate
	}
	if payload.Codes != nil {
		pub.Codes = payload.Codes
	}
	if payload.CreatedAt != nil {
		pub.CreatedAt = payload.CreatedAt
	}
	if payload.GivenNames != nil {
		pub.GivenNames = payload.GivenNames
	}
	if payload.ID != nil {
		pub.ID = *payload.ID
	}
	if payload.LastNames != nil {
		pub.LastNames = payload.LastNames
	}
	if payload.Location != nil {
		pub.Location = payload.Location
	}
	if payload.Mail != nil {
		pub.Mail = payload.Mail
	}
	if payload.Points != nil {
		pub.Points = payload.Points
	}
	if payload.Status != nil {
		pub.Status = payload.Status
	}
	if payload.UpdatedAt != nil {
		pub.UpdatedAt = payload.UpdatedAt
	}
	if payload.Username != nil {
		pub.Username = *payload.Username
	}
	return &pub
}

// UpdateUsersPayload is the users update action payload.
type UpdateUsersPayload struct {
	// Birthdate
	Birthdate *time.Time `form:"birthdate,omitempty" json:"birthdate,omitempty" xml:"birthdate,omitempty"`
	// Codes
	Codes []string `form:"codes,omitempty" json:"codes,omitempty" xml:"codes,omitempty"`
	// Record created timestamp
	CreatedAt *time.Time `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Given names
	GivenNames *string `form:"given_names,omitempty" json:"given_names,omitempty" xml:"given_names,omitempty"`
	// Unique ID
	ID int `form:"id" json:"id" xml:"id"`
	// Last names
	LastNames *string `form:"last_names,omitempty" json:"last_names,omitempty" xml:"last_names,omitempty"`
	// Location
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// E-mail address
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// Points
	Points *int `form:"points,omitempty" json:"points,omitempty" xml:"points,omitempty"`
	// User status
	Status *int `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Record updated timestamp
	UpdatedAt *time.Time `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Username
	Username string `form:"username" json:"username" xml:"username"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUsersPayload) Validate() (err error) {

	if payload.Username == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "username"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUsersContext) OK(r *GoaUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// Conflict sends a HTTP response with status code 409.
func (ctx *UpdateUsersContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// UpdatepasswordUsersContext provides the users updatepassword action context.
type UpdatepasswordUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *UpdatepasswordUsersPayload
}

// NewUpdatepasswordUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller updatepassword action.
func NewUpdatepasswordUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdatepasswordUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdatepasswordUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// updatepasswordUsersPayload is the users updatepassword action payload.
type updatepasswordUsersPayload struct {
	OldPassword *string `form:"oldPassword,omitempty" json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
	Password    *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updatepasswordUsersPayload) Validate() (err error) {
	if payload.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// Publicize creates UpdatepasswordUsersPayload from updatepasswordUsersPayload
func (payload *updatepasswordUsersPayload) Publicize() *UpdatepasswordUsersPayload {
	var pub UpdatepasswordUsersPayload
	if payload.OldPassword != nil {
		pub.OldPassword = payload.OldPassword
	}
	if payload.Password != nil {
		pub.Password = *payload.Password
	}
	return &pub
}

// UpdatepasswordUsersPayload is the users updatepassword action payload.
type UpdatepasswordUsersPayload struct {
	OldPassword *string `form:"oldPassword,omitempty" json:"oldPassword,omitempty" xml:"oldPassword,omitempty"`
	Password    string  `form:"password" json:"password" xml:"password"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdatepasswordUsersPayload) Validate() (err error) {
	if payload.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "password"))
	}
	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdatepasswordUsersContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdatepasswordUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdatepasswordUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ViewUsersContext provides the users view action context.
type ViewUsersContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewViewUsersContext parses the incoming request URL and body, performs validations and creates the
// context used by the users controller view action.
func NewViewUsersContext(ctx context.Context, r *http.Request, service *goa.Service) (*ViewUsersContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ViewUsersContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ViewUsersContext) OK(r *GoaUser) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.user+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ViewUsersContext) Unauthorized() error {
	ctx.ResponseData.WriteHeader(401)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ViewUsersContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
