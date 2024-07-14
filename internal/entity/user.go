package entity

type User struct {
	ID       string `json:"id"`
	NickName string `json:"first_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
