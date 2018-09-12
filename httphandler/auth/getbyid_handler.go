package auth

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/avengers-project-api/context/auth"
)

type GetByIDHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type getbyidHandler struct {
	gbiuc auth.GetByIDUsecase
	rr    httplib.RequestReader
	guauc auth.GetAuthUserUsecase
}

func NewGetByIDHandler(gbiuc auth.GetByIDUsecase, rr httplib.RequestReader, guauc auth.GetAuthUserUsecase) GetByIDHandler {
	return &getbyidHandler{gbiuc, rr, guauc}
}

func (this getbyidHandler) Handle(w http.ResponseWriter, r *http.Request) {
	UserID := this.guauc.GetUserID(r)

	err, resp := this.gbiuc.Get(UserID)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp.GetData())
}
