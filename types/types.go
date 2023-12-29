package types

type HealthCheck struct {
	Message    string `json:"message"`
	HTTPStatus int    `json:"http_status"`
}

type SecretValues struct {
	SecretKey string `json:"secret_key"`
	SecretID  string `json:"secret_id"`
}

type LoginResponse struct {
	Message    string        `json:"message"`
	Values     TokenResponse `json:"response_values"`
	HTTPStatus int           `json:"http_status"`
}

type TokenResponse struct {
	Access           string `json:"access"`
	AccessExpiration int    `json:"access_expires"`
	Refresh          string `json:"refresh"`
	RefreshExpires   int    `json:"refresh_expires"`
}

type BankInformations struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Bic                  string   `json:"bic"`
	TranscationTotalDays string   `json:"transaction_total_days"`
	Countries            []string `json:"countries"`
	Logo                 string   `json:"logo"`
}
type BankInformationResponse struct {
	Message    string             `json:"message"`
	Values     []BankInformations `json:"response_values"`
	HTTPStatus int                `json:"http_status"`
}

type UserAgreement struct {
	InstitutionID         string   `json:"institution_id"`
	MaxHistoryTransaction string   `json:"max_historical_days"`
	AccessTime            string   `json:"access_valid_for_days"`
	AccessScope           []string `json:"access_scope"`
}

type UserAgreementResponse struct {
	ID                    string   `json:"id"`
	Created_at            string   `json:"created"`
	InstitutionID         string   `json:"institution_id"`
	MaxHistoryTransaction int      `json:"max_historical_days"`
	AccessTime            int      `json:"access_valid_for_days"`
	AccessScope           []string `json:"access_scope"`
	Accepted              string   `json:"accepted"`
}
type BankRequisition struct {
	Redirect      string `json:"redirect"`
	InstitutionID string `json:"institution_id"`
	Reference     string `json:"reference"`
	AgreementID   string `json:"agreement"`
	UserLanguage  string `json:"user_language"`
}

type BankRequisitionResponse struct {
	ID           string                `json:"id"`
	Redirect     string                `json:"redirect"`
	Status       BankRequisitionStatus `json:"status"`
	Agreement    string                `json:"agreement"`
	Accounts     []string              `json:"accounts"`
	Reference    string                `json:"reference"`
	UserLanguage string                `json:"user_language"`
	Link         string                `json:"link"`
}

type BankRequisitionStatus struct {
	Short       string `json:"short"`
	Long        string `json:"long"`
	Description string `json:"description"`
}

type AccountResponse struct {
	ID        string   `json:"id"`
	Status    string   `json:"status"`
	Agreement string   `json:"agreement"`
	Accounts  []string `json:"accounts"`
	Reference string   `json:"reference"`
}

type TransactionReponse struct {
	Transaction TransactionResult `json:"transactions"`
}

type TransactionResult struct {
	Booked  []BookedTransaction  `json:"booked"`
	Pending []PendingTransaction `json:"pending"`
}

type BookedTransaction struct {
	TransactionID       string           `json:"internalTransactionId"`
	TransctionAmount    TransctionAmount `json:"transactionAmount"`
	BankTransactionCode string           `json:"bankTransactionCode"`
	BookingDate         string           `json:"bookingDate"`
	Description         []string         `json:"remittanceInformationUnstructuredArray"`
}

type PendingTransaction struct {
	TransctionAmount TransctionAmount `json:"transactionAmount"`
	BookingDate      string           `json:"bookingDate"`
	Description      []string         `json:"remittanceInformationUnstructuredArray"`
}

type DebtorAccount struct {
	Iban string `json:"iban"`
}

type TransctionAmount struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type Balance struct {
	TransctionAmount TransctionAmount `json:"balanceAmount"`
	BalanceType      string           `json:"balanceType"`
}

type BalanceResponse struct {
	Balances []Balance `json:"balances"`
}
