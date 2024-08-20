package models

type RegisterUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Utc      string `json:"utc"`
}

func ToRegisterUserResponse(user User) (registerUserResponse RegisterUserResponse) {
	registerUserResponse.Username = user.Username.String
	registerUserResponse.Email = user.Email.String
	registerUserResponse.Utc = user.Utc.String
	return
}
