package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserType string                 `json:"user_type"`
	UserData map[string]interface{} `json:"user_data"`
	jwt.StandardClaims
}

type TokenExpires struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

const UserTypeStaff = "staff"
const UserTypeWeixin = "weixin"
