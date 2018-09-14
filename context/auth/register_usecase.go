package auth

import (
	"github.com/oktopriima/avengers-project-core/model"
	"github.com/oktopriima/avengers-project-core/repository"
)

type RegisterUsecase interface {
	Create(req RegisterRequest) (error, RegisterResponse)
}

type RegisterRequest interface {
	GetName() string
	GetEmail() string
	GetPassword() string
}

type RegisterResponse interface {
	GetData() model.Users
}

type registerUsecase struct {
	uRepo        repository.UserRepository
	signatureKey []byte
}

func NewRegisterUsecase(uRepo repository.UserRepository, signatureKey []byte) RegisterUsecase {
	return &registerUsecase{uRepo, signatureKey}
}

func (this registerUsecase) Create(req RegisterRequest) (error, RegisterResponse) {
	m := model.Users{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	err := this.uRepo.Register(m)

	if err != nil {
		return err, registerResponse{}
	}

	return nil, registerResponse{m}
}
