package mdl

type UserRegis struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Password2 string `json:"password2"`
}

type User struct {
  Id int `json:"user_id"`
  Username string `json:"username"`
  Password string `json:"user_password"`
}
