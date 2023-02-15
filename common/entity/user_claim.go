package entity

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	*jwt.RegisteredClaims
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	IsRefresh bool   `json:"is_refresh"`
}

func (e *UserClaims) String() string {
	return fmt.Sprintf("%+v", *e)
}
