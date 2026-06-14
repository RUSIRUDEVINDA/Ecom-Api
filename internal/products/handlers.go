package products

import (
	"net/http"
	"github.com/RUSIRUDEVINDA/Ecom-Api/internal/json" 
	
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

	json.Header(w, http.StatusOK, products)
}
