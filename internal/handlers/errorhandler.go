package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gokultp/gstreamer/pkg/errors"
)

func HandleError(err errors.IError, w http.ResponseWriter) {
	errData := map[string]interface{}{
		"message": err.Error(),
		"code":    err.Code(),
	}
	data, _ := json.Marshal(errData)
	w.WriteHeader(err.HTTPStatus())
	w.Write(data)
}
