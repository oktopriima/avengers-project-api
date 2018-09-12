package auth

import "github.com/oktopriima/avengers-project-core/model"

type getbyidResponse struct {
	Data *model.Users
}

func (this getbyidResponse) GetData() *model.Users {
	return this.Data
}
