package products

import (
	"net/http"
	"log"
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

// ListProducts is a handler function that retrieves a list of products and writes the response as JSON.
func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err:=h.service.ListProducts(r.Context())
	if err!=nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"products"`
	}{}

	json.Header(w, http.StatusOK, products)
}
