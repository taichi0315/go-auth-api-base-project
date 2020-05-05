package auth

import (
	"log"
	"net/http"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/xerrors"
)

func HandleAuthSignOut() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		token := request.Header.Get("Authorization")

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		err := authRepo.DeleteAuthTokenByToken(token)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		response.Success(writer, "")
	}
}
