// Package classification of Auth API
//
// Documentation for Auth API
//
//	Schemes: http
//	BasePath: /auth
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/ishtiaqhimel/oms/auth-server/models"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body models.GenericError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters loginUser signupUser
type userParamWrapper struct {
	// User data structure to SignUp or Login.
	// in: body
	// required: true
	Body models.User
}
