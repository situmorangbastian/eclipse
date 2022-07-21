package eclipse

import (
	"context"

	"github.com/golang-jwt/jwt"
)

type CustomContext struct {
	context.Context
	user jwt.MapClaims
}

func (c CustomContext) User() jwt.MapClaims {
	return c.user
}

func NewCustomContext(c context.Context, u jwt.MapClaims) CustomContext {
	return CustomContext{Context: c, user: u}
}
