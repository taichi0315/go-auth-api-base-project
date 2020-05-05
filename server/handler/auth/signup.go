package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/model"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	authJson "github.com/nepp-tumsat/documents-api/server/json/auth"
	"github.com/nepp-tumsat/documents-api/server/response"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

func HandleAuthSignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var requestBody authJson.AuthSignUpRequest
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in json: %v", err))
			response.BadRequest(writer, "Can't decode of json")
			return
		}

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		err = authRepo.InsertUser(model.User{UserName: requestBody.UserName})
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		user, err := authRepo.SelectUserByUserName(requestBody.UserName)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		hash, err := passwordToHash(requestBody.Password)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in bcrypt: %v", err))
			return
		}

		err = authRepo.InsertUserAuth(model.UserAuth{UserID: user.UserID, Email: requestBody.Email, Hash: hash})
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		token, err := uuid.NewRandom()
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in uuid: %v", err))
			return
		}

		err = authRepo.InsertAuthToken(model.AuthToken{UserID: user.UserID, Token: token.String()})
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			return
		}

		response.Success(writer, authJson.AuthSignUpResponse{UserName: user.UserName, Token: token.String()})
	}
}

func passwordToHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}
