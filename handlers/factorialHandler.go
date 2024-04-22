package handlers

import (
	"calculateFactorial/models"
	"calculateFactorial/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CalculateFactorial(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	request := r.Context().Value("factorialRequest").(models.FactorialRequest)

	results := utils.RunFactorialCalculations(*request.AFactorial, *request.BFactorial)
	json.NewEncoder(w).Encode(results)
}
