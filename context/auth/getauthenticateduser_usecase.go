package auth

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
)

type GetAuthUserUsecase interface {
	GetUserID(*http.Request) int
}

type getAuthUserUsecase struct {
}

func NewGetAuthUserUsecase() GetAuthUserUsecase {
	return &getAuthUserUsecase{}
}

func (this getAuthUserUsecase) GetUserID(r *http.Request) int {
	userID, _ := httplib.ExtractClaim(r, "user_id")
	return int(userID.(float64))
}
