package model

type JsonRequest struct {
	Data *AccountData `json:"data"`
}

func (j *JsonRequest) setData(accountData AccountData) {
	j.Data = &accountData
}