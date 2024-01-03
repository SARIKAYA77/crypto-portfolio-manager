package api

import (
	"crypto-portfolio/app/model"
	"crypto-portfolio/app/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CurrencyHandler struct {
	CurrencyService *service.CurrencyService
}

func NewCurrencyHandler(currencyService *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{CurrencyService: currencyService}
}

func (h *CurrencyHandler) AddCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	var currency model.Currency

	err := json.NewDecoder(r.Body).Decode(&currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCurrency, err := h.CurrencyService.AddCurrency(currency.Code, currency.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCurrency)
}

func (h *CurrencyHandler) EditCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid currency ID", http.StatusBadRequest)
		return
	}

	var currency model.Currency
	err = json.NewDecoder(r.Body).Decode(&currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	editedCurrency, err := h.CurrencyService.EditCurrency(id, currency.Code, currency.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(editedCurrency)
}

func (h *CurrencyHandler) DeleteCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid currency ID", http.StatusBadRequest)
		return
	}

	err = h.CurrencyService.DeleteCurrency(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CurrencyHandler) GetCurrencyByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid currency ID", http.StatusBadRequest)
		return
	}

	currency, err := h.CurrencyService.GetCurrencyByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currency)
}

func (h *CurrencyHandler) GetAllCurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	currencies, err := h.CurrencyService.GetAllCurrencies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(currencies)
	if err != nil {
		return
	}
}
