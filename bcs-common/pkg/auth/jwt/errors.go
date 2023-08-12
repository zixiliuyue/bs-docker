

package jwt

import (
	"errors"
)

var (
	// ErrServerNotInited for server not inited error
	ErrServerNotInited = errors.New("server not init")
	// ErrJWtSignKeyEmpty for jwt signKey empty error
	ErrJWtSignKeyEmpty = errors.New("jwt options signKey empty")
	// ErrJWtVerifyKeyEmpty for jwt verify key error
	ErrJWtVerifyKeyEmpty = errors.New("jwt options verifyKey empty")
	// ErrJWtUserNameEmpty for jwt username error
	ErrJWtUserNameEmpty = errors.New("jwt uerInfo userName empty")
	// ErrJWtClientNameEmpty for jwt clientname error
	ErrJWtClientNameEmpty = errors.New("jwt uerInfo clientName empty")
	// ErrJWtSubType for jwt user type error
	ErrJWtSubType = errors.New("jwt subType err: user or client")
	// ErrTokenIsNil parse with claim return token is nil
	ErrTokenIsNil = errors.New("parse with claims return token is nil")
)
