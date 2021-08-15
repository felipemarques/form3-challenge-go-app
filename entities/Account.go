package entities

type Account struct {
	Title string `json:"title"`
}

func (a Account) setTitle(title string) {
	a.Title = title
}