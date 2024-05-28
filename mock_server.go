package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PaymentRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type PaymentResponse struct {
	PaymentID string    `json:"payment_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Status    string    `json:"status"`
}

var payments = make(map[string]PaymentResponse)

func createPayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	paymentID := uuid.New().String()
	response := PaymentResponse{
		PaymentID: paymentID,
		Amount:    req.Amount,
		CreatedAt: time.Now(),
		CreatedBy: req.AccountID,
		Status:    "Authorized", // Simulação de lógica de autorização
	}

	payments[paymentID] = response

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	paymentID := r.URL.Path[len("/payments/"):]
	response, ok := payments[paymentID]
	if !ok {
		http.Error(w, "Pagamento não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/payments", createPayment)
	http.HandleFunc("/payments/", getPayment)

	fmt.Println("Iniciando mock server na porta 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
