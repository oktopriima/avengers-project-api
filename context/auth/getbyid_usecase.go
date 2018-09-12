package auth

import (
	"github.com/oktopriima/avengers-project-core/model"
	"github.com/oktopriima/avengers-project-core/repository"
)

type GetByIDUsecase interface {
	Get(UserID int) (error, GetByIDResponse)
}

type GetByIDResponse interface {
	GetData() *model.Users
}

type getbyidUsecase struct {
	userRepo repository.UserRepository
}

func NewGetByIDUsecase(userRepo repository.UserRepository) GetByIDUsecase {
	return &getbyidUsecase{userRepo}
}

func (this getbyidUsecase) Get(UserID int) (error, GetByIDResponse) {
	err, resp := this.userRepo.FindByID(UserID)
	if err != nil {
		return err, getbyidResponse{}
	}

	return nil, getbyidResponse{resp}
}
