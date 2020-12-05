package router

import (
	"ebanx/challenge/domain"
	"net/http"
)

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusCreated
	}

	switch err {
	case domain.ErrInternal:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrNotEnoughBalance:
		return http.StatusForbidden
	case domain.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
