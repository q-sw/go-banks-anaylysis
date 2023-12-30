package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/q-sw/go-bank-analysis/types"
)

func loginToNordigen(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {

		secretID := os.Getenv("SECRET_ID")
		secretKey := os.Getenv("SECRET_KEY")
		if secretID == "" && secretKey == "" {
			log.Printf("%v - %v - secret ID or key is empty",
				r.URL, http.StatusInternalServerError)
			m := types.HealthCheck{
				Message:    "Missing Value in configuration",
				HTTPStatus: http.StatusInternalServerError,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			log.Printf("%v - %v", r.URL, http.StatusOK)
			s := types.SecretValues{
				SecretKey: secretKey,
				SecretID:  secretID,
			}

			n := GetNordigenToken(s)
			os.Setenv("TOKEN", n.Access)

			m := types.LoginResponse{
				Message:    "Login to Nordigen successful",
				Values:     n,
				HTTPStatus: http.StatusOK,
			}
			json.NewEncoder(w).Encode(m)
		}
	}
}

func getFrenchBanks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			n := ListFrenchBank(t)

			m := types.BankInformationResponse{
				Message:    "List all French Bank successfully",
				Values:     n,
				HTTPStatus: http.StatusOK,
			}
			json.NewEncoder(w).Encode(m)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}
}

func createUserAgreement(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			var b types.UserAgreement
			json.NewDecoder(r.Body).Decode(&b)
			n := CreateUserAgreement(t, b)
			json.NewEncoder(w).Encode(n)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}

}

func recordBank(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			var b types.BankRequisition
			json.NewDecoder(r.Body).Decode(&b)
			n := CreateBankRequisition(t, b)
			json.NewEncoder(w).Encode(n)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}

}

func listAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			vars := mux.Vars(r)
			a := vars["recordId"]
			fmt.Println(a)
			t := os.Getenv("TOKEN")
			n := ListAccounts(t, a)

			json.NewEncoder(w).Encode(n)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			vars := mux.Vars(r)
			a := vars["accountId"]
			fmt.Println(a)
			t := os.Getenv("TOKEN")
			n := GetTransactions(t, a)

			json.NewEncoder(w).Encode(n)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}
}
func getBalances(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := types.HealthCheck{
			Message:    "Bad Request",
			HTTPStatus: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(m)
		log.Printf("%v - %v", r.URL, http.StatusBadRequest)
	} else {
		t := os.Getenv("TOKEN")
		if t == "" {
			log.Printf("%v - %v - You are not login",
				r.URL, http.StatusForbidden)
			m := types.HealthCheck{
				Message:    "You are not login",
				HTTPStatus: http.StatusForbidden,
			}
			json.NewEncoder(w).Encode(m)
		} else {
			vars := mux.Vars(r)
			a := vars["accountId"]
			fmt.Println(a)
			t := os.Getenv("TOKEN")
			n := GetBalance(t, a)

			nbyte, _ := json.Marshal(n)
			SendMessage(nbyte)
			json.NewEncoder(w).Encode(n)
			log.Printf("%v - %v", r.URL, http.StatusOK)
		}
	}
}
