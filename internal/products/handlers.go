package products

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//1. call the service -> service.ListProducts()
	//2. return the response to the client(return JSON an http request)

	products := struct {
		Products []string `json:"products"`
	}{}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
