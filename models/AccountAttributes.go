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
func (a AccountAttributes) setClassification(classification string) {
	a.AccountClassification = &classification
}

func (a AccountAttributes) getClassification() string {
	return *a.AccountClassification
}

// AccountMatchingOptOut
func (a AccountAttributes) setMatchingOptOut(matchingOptOut bool) {
	a.AccountMatchingOptOut = &matchingOptOut
}

func (a AccountAttributes) getMatchingOptOut() bool {
	return *a.AccountMatchingOptOut
}

// AccountNumber
func (a AccountAttributes) setAccountNumber(accountNumber string) {
	a.AccountNumber = accountNumber
}

func (a AccountAttributes) getAccountNumber() string {
	return a.AccountNumber
}

// AlternativeNames
func (a AccountAttributes) setAlternativeNames(names []string) {
	a.AlternativeNames = names
}

func (a AccountAttributes) getAlternativeNames() []string {
	return a.AlternativeNames
}

// BankID
func (a AccountAttributes) setBankID(bankId string) {
	a.BankID = bankId
}

func (e AccountAttributes) getBankID() string {
	return e.BankID
}

// BankIDCode
func (a AccountAttributes) setBankIDCode(bankIdCode string) {
	a.BankIDCode = bankIdCode
}

func (e AccountAttributes) getBankIDCode() string {
	return e.BankIDCode
}

// BaseCurrency
func (a AccountAttributes) setBaseCurrency(currency string) {
	a.BaseCurrency = currency
}

func (a AccountAttributes) getBaseCurrency() string {
	return a.BaseCurrency
}

// Bic
func (a AccountAttributes) setBic(bic string) {
	a.Bic = bic
}

func (a AccountAttributes) getBic() string {
	return a.Bic
}

// Country
func (a AccountAttributes) setCountry(country string){
	a.Country = &country
}

func (a AccountAttributes) getCountry() string {
	return *a.Country
}

// Iban
func (a AccountAttributes) setIban(iban string) {
	a.Iban = iban
}

func (a AccountAttributes) getIban() string {
	return a.Iban
}

// JointAccount
func (a AccountAttributes) setJointAccount(jointAccount bool) {
	a.JointAccount = &jointAccount
}

func (a AccountAttributes) getJointAccount() bool {
	return *a.JointAccount
}

// Name
func (a AccountAttributes) setName(names []string) {
	a.Name = names
}

func (a AccountAttributes) getName() []string {
	return a.Name
}

// SecondaryIdentification
func (a AccountAttributes) setSecondaryIdentification(identification string) {
	a.SecondaryIdentification = identification
}

func (a AccountAttributes) getSecondaryIdentification() string {
	return a.SecondaryIdentification
}

// Status
func (a AccountAttributes) setStatus(status string){
	a.Status = &status
}

func (a AccountAttributes) getStatus() string {
	return *a.Status
}

// Switched
func (a AccountAttributes) setSwitched(switched bool) {
	a.Switched = &switched
}

func (a AccountAttributes) getSwitched() bool {
	return *a.Switched
}
