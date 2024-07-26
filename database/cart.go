package database

import "errors"

var (
	ErrCantFindProduct =errors.New("")
	ErrCantDecodeProduct =errors.New("")
	ErrUserIdIsNotValid=errors.New("")
	ErrCantUpdateUser=errors.New("")
	ErrCantRemoveItemCart=errors.New("")
	ErrCantGetItem=errors.New("")
	ErrCantBuyItem=errors.New("")
)