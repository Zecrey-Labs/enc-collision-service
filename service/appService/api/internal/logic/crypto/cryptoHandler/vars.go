package cryptohandler

import (
	"errors"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound           = sqlx.ErrNotFound
	ErrInvalidCryptoInput = errors.New("[ErrInvalidCryptoInput] Invalid Crypto input")
)
