package cmd

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/q-sw/go-bank-analysis/types"
)

type APIServer struct {
	ListenPort string
}

func NewServer(ListenPort string) *APIServer {

	return &APIServer{
		ListenPort: ListenPort,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthCheck)
	nordiRouter := router.PathPrefix("/nordigen").Subrouter()
	nordiRouter.HandleFunc("/login", loginToNordigen).Methods("POST")
	nordiRouter.HandleFunc("/banks", getFrenchBanks).Methods("GET")
	nordiRouter.HandleFunc("/agree", createUserAgreement).Methods("POST")
	nordiRouter.HandleFunc("/record-bank", recordBank).Methods("POST")
	nordiRouter.HandleFunc("/{recordId}/accounts", listAccounts).Methods("GET")
	nordiRouter.HandleFunc("/{accountId}/transactions", getTransactions).Methods("GET")
	nordiRouter.HandleFunc("/{accountId}/balances", getBalances).Methods("GET")
	http.ListenAndServe(s.ListenPort, router)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
	} else {

		m := types.HealthCheck{
			Message:    "healthy",
			HTTPStatus: http.StatusOK,
		}
		json.NewEncoder(w).Encode(m)
	}
}
