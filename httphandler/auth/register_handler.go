package auth

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/avengers-project-api/context/auth"
)

type RegisterHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type registerHandler struct {
	ruc auth.RegisterUsecase
	rr  httplib.RequestReader
}

func NewRegisterHandler(ruc auth.RegisterUsecase, rr httplib.RequestReader) RegisterHandler {
	return &registerHandler{ruc, rr}
}

func (this registerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}
	err := this.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	err, resp := this.ruc.Create(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp.GetData())
}
