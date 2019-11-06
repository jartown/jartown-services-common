package auth

import (
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"
	common "github.com/singkorn/jartown-services-common"
)

func IssueToken(issuer string, userDataI interface{}) (TokenExpires, error) {
	userType := ""

	switch userDataI.(type) {
	case UserStaff:
		userType = UserTypeStaff
		userDataI = userDataI.(UserStaff).StripPassword()

	default:
		return TokenExpires{}, common.ErrInvalidUserType
	}

	userData, err := remarshalUserDataAsMap(userDataI)
	if err != nil {
		return TokenExpires{}, err
	}

	currentTime := time.Now()
	expireAt := currentTime.Add(getTokenTimeoutDuration())

	claims := UserClaims{}
	claims.UserType = userType
	claims.UserData = userData
	claims.Issuer = issuer
	claims.IssuedAt = currentTime.UTC().Unix()
	claims.NotBefore = currentTime.UTC().Unix()
	claims.ExpiresAt = expireAt.UTC().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	hexKey, err := hex.DecodeString(common.Conf.JWT.Key)
	if err != nil {
		return TokenExpires{}, err
	}
	jwtStr, err := token.SignedString(hexKey)
	if err != nil {
		return TokenExpires{}, err
	}

	return TokenExpires{Token: jwtStr, ExpiresAt: expireAt}, nil
}
