// handlers/order_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ishtiaqhimel/oms/oms-server/initializers"
	"github.com/ishtiaqhimel/oms/oms-server/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Order{})

	db.Create(&models.Order{Username: "Ishtiaq Islam", ProductName: "tea"})
	db.Create(&models.Order{Username: "Mridul", ProductName: "water"})

	return db
}

func TestCreateOrder(t *testing.T) {
	initializers.DB = setupTestDB()

	newOrder := models.Order{
		Username:    "Ishtiaq",
		ProductName: "tea",
	}

	payload, err := json.Marshal(newOrder)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/order", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CreateOrder(w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var resOrder models.Order
	err = models.FromJSON(&resOrder, rr.Body)
	assert.NoError(t, err)

	assert.Equal(t, "Ishtiaq", resOrder.Username)
	assert.Equal(t, "tea", resOrder.ProductName)
}

func TestListOrders(t *testing.T) {
	initializers.DB = setupTestDB()

	req, err := http.NewRequest("GET", "/api/order", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ListOrders(w, r)
	})

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resOrders []models.Order
	err = models.FromJSON(&resOrders, rr.Body)
	assert.NoError(t, err)

	orders := []models.Order{
		{
			Username:    "Ishtiaq Islam",
			ProductName: "tea",
		},
		{
			Username:    "Mridul",
			ProductName: "water",
		},
	}

	assert.Equal(t, len(orders), len(resOrders))

	for idx := range orders {
		assert.Equal(t, orders[idx].Username, resOrders[idx].Username)
		assert.Equal(t, orders[idx].ProductName, resOrders[idx].ProductName)
	}
}
