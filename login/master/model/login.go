package model

type Login struct {
	ID       string `json:"id_user"`
	Username string `json:"username"`
	Paswword string `json:"password"`
}
type Response struct {
	Status   int
	Messages string
	Data     interface{}
}
