package middlewars

import (
	"calculateFactorial/models"
	"calculateFactorial/utils"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ValidateFactorialRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		var factorialRequest models.FactorialRequest

		err := json.NewDecoder(r.Body).Decode(&factorialRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Invalid JSON"})
			return
		}

		if !utils.ValidateNumbers(factorialRequest) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: "Incorrect input"})
			return
		}

		ctx := context.WithValue(r.Context(), "factorialRequest", factorialRequest)
		w.WriteHeader(http.StatusOK)
		next(w, r.WithContext(ctx), params)
	}
}
