// Package classification of Order API
//
// Documentation for Order API
//
//	Schemes: http
//	BasePath: /api
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

import "github.com/ishtiaqhimel/oms/oms-server/models"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body models.GenericError
}

// Data structure representing a single order
// swagger:response orderResponse
type orderResponseWrapper struct {
	// The letest created/updated order
	// in: body
	Body models.Order
}

// Data structure representing list of orders
// swagger:response orderListResponse
type orderListResponseWrapper struct {
	// The letest created/updated list of orders
	// in: body
	Body []models.Order
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters createOrder
type orderParamWrapper struct {
	// Order data structure to update order.
	// in: body
	// required: true
	Body models.Order
}

// swagger:parameters getOrder deleteOrder
type orderIDParamWrapper struct {
	// The id of the order for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateOrder
type orderUpdateParamsWrapper struct {
	// Order data structure to update order.
	// in: body
	// required: true
	Body models.Order
	// The id of the order for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
