package models

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

// AccountClassification
func (a *AccountAttributes) setClassification(classification string) {
	a.AccountClassification = &classification
}

func (e *AccountAttributes) getClassification() string {
	return *e.AccountClassification
}

// AccountMatchingOptOut
func (e *AccountAttributes) setMatchingOptOut(matchingOptOut bool) {
	e.AccountMatchingOptOut = &matchingOptOut
}

func (e *AccountAttributes) getMatchingOptOut() bool {
	return *e.AccountMatchingOptOut
}

// AccountNumber
func (e *AccountAttributes) setAccountNumber(accountNumber string) {
	e.AccountNumber = accountNumber
}

func (e *AccountAttributes) getAccountNumber() string {
	return e.AccountNumber
}

// AlternativeNames
func (e *AccountAttributes) setAlternativeNames(names []string) {
	e.AlternativeNames = names
}

func (e *AccountAttributes) getAlternativeNames() []string {
	return e.AlternativeNames
}

// BankID
func (e *AccountAttributes) setBankID(bankId string) {
	e.BankID = bankId
}

func (e *AccountAttributes) getBankID() string {
	return e.BankID
}

// BankIDCode
func (e *AccountAttributes) setBankIDCode(bankIdCode string) {
	e.BankIDCode = bankIdCode
}

func (e *AccountAttributes) getBankIDCode() string {
	return e.BankIDCode
}

// BaseCurrency
func (e *AccountAttributes) setBaseCurrency(currency string) {
	e.BaseCurrency = currency
}

func (e *AccountAttributes) getBaseCurrency() string {
	return e.BaseCurrency
}

// Bic
func (e *AccountAttributes) setBic(bic string) {
	e.Bic = bic
}

func (e *AccountAttributes) getBic() string {
	return e.Bic
}

// Country
func (e *AccountAttributes) setCountry(country string){
	e.Country = &country
}

func (e *AccountAttributes) getCountry() string {
	return *e.Country
}

// Iban
func (e *AccountAttributes) setIban(iban string) {
	e.Iban = iban
}

func (e *AccountAttributes) getIban() string {
	return e.Iban
}

// JointAccount
func (e *AccountAttributes) setJointAccount(jointAccount bool) {
	e.JointAccount = &jointAccount
}

func (e *AccountAttributes) getJointAccount() bool {
	return *e.JointAccount
}

// Name
func (e *AccountAttributes) setName(names []string) {
	e.Name = names
}

func (e *AccountAttributes) getName() []string {
	return e.Name
}

// SecondaryIdentification
func (e *AccountAttributes) setSecondaryIdentification(identification string) {
	e.SecondaryIdentification = identification
}

func (e *AccountAttributes) getSecondaryIdentification() string {
	return e.SecondaryIdentification
}

// Status
func (e *AccountAttributes) setStatus(status string){
	e.Status = &status
}

func (e *AccountAttributes) getStatus() string {
	return *e.Status
}

// Switched
func (e *AccountAttributes) setSwitched(switched bool) {
	e.Switched = &switched
}

func (e *AccountAttributes) getSwitched() bool {
	return *e.Switched
}
