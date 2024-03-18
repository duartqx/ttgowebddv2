package http

import (
	"errors"
	"net/http"

	e "github.com/duartqx/ddgobase/common/errors"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	var valError *e.ValidationErrors
	switch {
	case errors.As(err, &valError):
		// w.Header().Set("Content-Type", "application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(valError.Error())
	case errors.Is(err, e.NotFoundError):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, e.BadRequestError):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, e.ForbiddenError):
		http.Error(w, err.Error(), http.StatusForbidden)
	case errors.Is(err, e.Unauthorized):
		http.Error(w, err.Error(), http.StatusUnauthorized)
	default:
		panic(err.Error())
	}
}
