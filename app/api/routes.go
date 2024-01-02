package api

import (
	"crypto-portfolio/app/service"
	"github.com/gorilla/mux"
)

func InitRoutes(currencyService *service.CurrencyService) *mux.Router {
	handler := NewCurrencyHandler(currencyService)

	router := mux.NewRouter()

	router.HandleFunc("/currency", handler.AddCurrencyHandler).Methods("POST")
	router.HandleFunc("/currency/{id}", handler.GetCurrencyByIDHandler).Methods("GET")
	router.HandleFunc("/currency/{id}", handler.EditCurrencyHandler).Methods("PUT")
	router.HandleFunc("/currency/{id}", handler.DeleteCurrencyHandler).Methods("DELETE")
	router.HandleFunc("/currencies", handler.GetAllCurrenciesHandler).Methods("GET")

	return router
}
