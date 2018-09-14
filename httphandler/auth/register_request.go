package auth

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this registerRequest) GetName() string {
	return this.Name
}

func (this registerRequest) GetEmail() string {
	return this.Email
}

func (this registerRequest) GetPassword() string {
	return this.Password
}
