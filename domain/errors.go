package domain

import "errors"

var (
	ErrNotFound         = errors.New("Requested resource is not found")
	ErrNotEnoughBalance = errors.New("Not enough balance to withdraw from account")
	ErrBadParamInput    = errors.New("Given param is not valid")
	ErrInternal         = errors.New("Internal server error")
)
