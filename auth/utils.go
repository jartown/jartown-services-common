package auth

import (
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	common "gitlab.com/diancai/diancai-services-common"
)

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return hex.DecodeString(common.Conf.JWT.Key)
}

func remarshalUserData(userType string, userData map[string]interface{}) (interface{}, error) {
	j, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	switch userType {
	case UserTypeStaff:
		var user UserStaff
		err := json.Unmarshal(j, &user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	return nil, common.ErrInvalidUserType
}

func remarshalUserDataAsMap(userData interface{}) (map[string]interface{}, error) {
	j, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func getTokenTimeoutDuration() time.Duration {
	duration, err := time.ParseDuration(common.Conf.JWT.Timeout)
	if err != nil {
		return time.Duration(1 * time.Hour)
	}
	return duration
}

func GetUserFromToken(tokenString string) (interface{}, error) {
	var claims UserClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, common.ErrUnauthorized
	}

	user, err := remarshalUserData(claims.UserType, claims.UserData)
	if err != nil {
		return nil, err
	}

	return user, nil
}
