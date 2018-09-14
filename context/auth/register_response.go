package auth

import "github.com/oktopriima/avengers-project-core/model"

type registerResponse struct {
	Data model.Users
}

func (this registerResponse) GetData() model.Users {
	return this.Data
}
