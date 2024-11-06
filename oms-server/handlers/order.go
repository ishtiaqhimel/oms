package handlers

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ishtiaqhimel/oms/oms-server/initializers"
	"github.com/ishtiaqhimel/oms/oms-server/models"
)

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
