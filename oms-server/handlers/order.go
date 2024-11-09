package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ishtiaqhimel/oms/oms-server/initializers"
	"github.com/ishtiaqhimel/oms/oms-server/models"
)

// swagger:route POST /order createOrder
// Create a new order
//
// responses:
//		201: orderResponse
//	 400: errorResponse
//   401: errorResponse

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Creating a new order")

	var order models.Order
	if err := models.FromJSON(&order, r.Body); err != nil {
		log.Println("Error deserializing", err)
		w.WriteHeader(http.StatusBadRequest)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	initializers.DB.Create(&order)
	w.WriteHeader(http.StatusCreated)
	models.ToJSON(order, w)
}

// swagger:route GET /order listOrders
// List all orders
//
// responses:
//		200: orderListResponse
//	 400: errorResponse
//   401: errorResponse
//	 500: errorResponse

func ListOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Listing orders")

	var orders []models.Order
	if err := initializers.DB.Find(&orders).Error; err != nil {
		log.Println("Error while listing orders", err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	models.ToJSON(orders, w)
}

// swagger:route GET /order/{id} getOrder
// Get an order by id
//
// responses:
//		200: orderResponse
//	 400: errorResponse
//   401: errorResponse
//	 404: errorResponse

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Getting an order")

	id := chi.URLParam(r, "id")
	var order models.Order
	if err := initializers.DB.First(&order, id).Error; err != nil {
		log.Println("Order not found", err)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	models.ToJSON(order, w)
}

// swagger:route PUT /order/{id} updateOrder
// Update an order by id
//
// responses:
//		200: orderResponse
//	 400: errorResponse
//   401: errorResponse
//	 404: errorResponse

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Updating an order")

	id := chi.URLParam(r, "id")
	var order models.Order
	if err := initializers.DB.First(&order, id).Error; err != nil {
		log.Println("Order not found", err)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	if err := models.FromJSON(&order, r.Body); err != nil {
		log.Println("Error while deserializing", err)
		w.WriteHeader(http.StatusBadRequest)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	initializers.DB.Save(&order)
	models.ToJSON(order, w)
}

// swagger:route DELETE /order/{id} deleteOrder
// Delete an order by id
//
// responses:
//		204: noContentResponse
//	 400: errorResponse
//   401: errorResponse
//	 404: errorResponse

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	log.Println("Deleting an order")

	id := chi.URLParam(r, "id")
	if err := initializers.DB.Delete(&models.Order{}, id).Error; err != nil {
		log.Println("Order not found", err)
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&models.GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
