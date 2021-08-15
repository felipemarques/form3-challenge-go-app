package model

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func NewAccountAttributes() *AccountAttributes {
	return &AccountAttributes{}
}

// AccountClassification
func (a *AccountAttributes) SetClassification(classification string) {
	a.AccountClassification = &classification
}

func (a *AccountAttributes) GetClassification() string {
	return *a.AccountClassification
}

// AccountMatchingOptOut
func (a *AccountAttributes) SetMatchingOptOut(matchingOptOut bool) {
	a.AccountMatchingOptOut = &matchingOptOut
}

func (a *AccountAttributes) GetMatchingOptOut() bool {
	return *a.AccountMatchingOptOut
}

// AccountNumber
func (a *AccountAttributes) SetAccountNumber(accountNumber string) {
	a.AccountNumber = accountNumber
}

func (a *AccountAttributes) GetAccountNumber() string {
	return a.AccountNumber
}

// AlternativeNames
func (a *AccountAttributes) SetAlternativeNames(names []string) {
	a.AlternativeNames = names
}

func (a *AccountAttributes) GetAlternativeNames() []string {
	return a.AlternativeNames
}

// BankID
func (a *AccountAttributes) SetBankID(bankId string) {
	a.BankID = bankId
}

func (e AccountAttributes) GetBankID() string {
	return e.BankID
}

// BankIDCode
func (a *AccountAttributes) SetBankIDCode(bankIdCode string) {
	a.BankIDCode = bankIdCode
}

func (e AccountAttributes) GetBankIDCode() string {
	return e.BankIDCode
}

// BaseCurrency
func (a *AccountAttributes) SetBaseCurrency(currency string) {
	a.BaseCurrency = currency
}

func (a *AccountAttributes) GetBaseCurrency() string {
	return a.BaseCurrency
}

// Bic
func (a *AccountAttributes) SetBic(bic string) {
	a.Bic = bic
}

func (a *AccountAttributes) GetBic() string {
	return a.Bic
}

// Country
func (a *AccountAttributes) SetCountry(country string){
	a.Country = &country
}

func (a *AccountAttributes) GetCountry() string {
	return *a.Country
}

// Iban
func (a *AccountAttributes) SetIban(iban string) {
	a.Iban = iban
}

func (a *AccountAttributes) GetIban() string {
	return a.Iban
}

// JointAccount
func (a *AccountAttributes) SetJointAccount(jointAccount bool) {
	a.JointAccount = &jointAccount
}

func (a *AccountAttributes) GetJointAccount() bool {
	return *a.JointAccount
}

// Name
func (a *AccountAttributes) SetName(names []string) {
	a.Name = names
}

func (a *AccountAttributes) GetName() []string {
	return a.Name
}

// SecondaryIdentification
func (a *AccountAttributes) SetSecondaryIdentification(identification string) {
	a.SecondaryIdentification = identification
}

func (a *AccountAttributes) GetSecondaryIdentification() string {
	return a.SecondaryIdentification
}

// Status
func (a *AccountAttributes) SetStatus(status string){
	a.Status = &status
}

func (a *AccountAttributes) GetStatus() string {
	return *a.Status
}

// Switched
func (a *AccountAttributes) SetSwitched(switched bool) {
	a.Switched = &switched
}

func (a *AccountAttributes) GetSwitched() bool {
	return *a.Switched
}
