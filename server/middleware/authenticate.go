package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/nepp-tumsat/documents-api/infrastructure"
	"github.com/nepp-tumsat/documents-api/infrastructure/persistence"
	"github.com/nepp-tumsat/documents-api/server/response"
	"github.com/nepp-tumsat/documents-api/util/dcontext"
	"golang.org/x/xerrors"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		ctx := request.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// リクエストヘッダからtoken(認証トークン)を取得
		token := request.Header.Get("Authorization")
		if len(token) == 0 {
			log.Printf("%+v\n", xerrors.New("token is empty"))
			response.Unauthorized(writer, "token is empty")
			return
		}

		authRepo := persistence.NewAuthDB(infrastructure.DB)

		// データベースから認証トークンに紐づくユーザの情報を取得
		authToken, err := authRepo.SelectAuthTokenByToken(token)
		if err != nil {
			log.Printf("%+v\n", xerrors.Errorf("Error in repository: %v", err))
			response.Unauthorized(writer, "Invalid token")
			return
		}

		// userIdをContextへ保存して以降の処理に利用する
		ctx = dcontext.SetUserID(ctx, string(authToken.UserID))

		// 次の処理
		nextFunc(writer, request.WithContext(ctx))
	}
}
