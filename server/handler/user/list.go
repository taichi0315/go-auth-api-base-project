package user

import (
	"net/http"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	userJson "github.com/nepp-tumsat/documents-api/server/json/user"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/xerrors"
)

func HandleUserList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		userRepo := persistence.NewUserDB(infrastructure.DB)

		users, err := userRepo.SelectAll()
		if err != nil {
			err = xerrors.Errorf("Error in repository: %v", err)
		}

		response.Success(writer, userJson.ToUserListResponse(users))
	}
}
