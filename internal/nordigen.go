package nordigen

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/q-sw/go-bank-analysis/types"
	"github.com/q-sw/go-bank-analysis/utils"
)

const (
	nordigenApiURL = "https://bankaccountdata.gocardless.com/api/v2"
)

func GetNordigenToken(v types.SecretValues) types.TokenResponse {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal("Error during Marshal Secret values")
	}
	url := nordigenApiURL + "/token/new/"

	resp := utils.PostRequest(url, data, "")

	var r types.TokenResponse

	json.Unmarshal(resp, &r)
	return r
}

func ListFrenchBank(t string) []types.BankInformations {
	url := nordigenApiURL + "/institutions/?country=FR"

	resp := utils.GetRequest(url, t)
	var r []types.BankInformations
	json.Unmarshal(resp, &r)

	return r

}

func CreateUserAgreement(t string, d types.UserAgreement) types.UserAgreementResponse {
	url := nordigenApiURL + "/agreements/enduser/"

	data, _ := json.Marshal(d)

	resp := utils.PostRequest(url, data, t)
	var r types.UserAgreementResponse
	json.Unmarshal(resp, &r)
	return r

}

func CreateBankRequisition(t string, d types.BankRequisition) types.BankRequisitionResponse {
	url := nordigenApiURL + "/requisitions/"

	data, _ := json.Marshal(d)

	resp := utils.PostRequest(url, data, t)
	fmt.Println(string(resp))
	var r types.BankRequisitionResponse
	json.Unmarshal(resp, &r)
	return r
}

func ListAccounts(t string, accountID string) types.AccountResponse {
	url := nordigenApiURL + "/requisitions/" + accountID
	resp := utils.GetRequest(url, t)

	var r types.AccountResponse
	json.Unmarshal(resp, &r)
	return r
}

func GetTransactions(t string, accountID string) types.TransactionReponse {
	url := nordigenApiURL + "/accounts/" + accountID + "/transactions/"
	resp := utils.GetRequest(url, t)
	var r types.TransactionReponse

	json.Unmarshal(resp, &r)
	return r
}
func GetBalance(t string, accountID string) types.BalanceResponse {
	url := nordigenApiURL + "/accounts/" + accountID + "/balances/"
	resp := utils.GetRequest(url, t)
	fmt.Println(string(resp))
	var r types.BalanceResponse

	json.Unmarshal(resp, &r)
	return r
}
