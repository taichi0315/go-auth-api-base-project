package response

import (
	"encoding/json"
	"net/http"
)

func Success(writer http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		InternalServerError(writer, "marshal error")
		return
	}
	writer.Write(data)
}

// BadRequest HTTPコード:400 BadRequestを処理する
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

func Unauthorized(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusUnauthorized, message)
}

func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		writer.Write(data)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
