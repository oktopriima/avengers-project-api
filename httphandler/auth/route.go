package auth

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/avengers-project-api/context/auth"
	"github.com/oktopriima/avengers-project-core/service"
)

func NewRoute(api *mux.Router, rr httplib.RequestReader, db *sql.DB, signatureKey []byte) {
	userRepo := service.NewUserRepository(db)

	tuc := auth.NewGetTokenUsecase(userRepo, signatureKey)
	gbiuc := auth.NewGetByIDUsecase(userRepo)
	guauc := auth.NewGetAuthUserUsecase()
	ruc := auth.NewRegisterUsecase(userRepo, signatureKey)

	gth := NewGetTokenHandler(tuc, rr)
	gbih := NewGetByIDHandler(gbiuc, rr, guauc)
	rh := NewRegisterHandler(ruc, rr)

	authRoute := api.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("", gth.Handle).Methods("POST")
	authRoute.HandleFunc("", httplib.JWTMid(gbih.Handle)).Methods("GET")
	authRoute.HandleFunc("/register", rh.Handle).Methods("POST")
}
