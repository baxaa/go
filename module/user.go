package module

type User struct {
	Name     string `json:"name"`
	Id       int    `json:"id"`
	password string `json:"password"`
}
