package auth

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	common "github.com/singkorn/jartown-services-common"
	"github.com/singkorn/jartown-services-common/derror"
)

func MiddlewareAuthentication(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return
	}

	authHeaderSplits := strings.SplitN(authHeader, " ", 2)
	if (len(authHeaderSplits) != 2) || (authHeaderSplits[0] != "Bearer") {
		common.HandlerReturnError(c, common.ErrUnauthorized)
		c.Abort()
		return
	}

	tokenString := authHeaderSplits[1]

	user, err := GetUserFromToken(tokenString)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				common.HandlerReturnError(c, derror.ErrorDebug(common.ErrUnauthorized, err.Error()))
				c.Abort()
				return
			}
		}
		common.HandlerReturnError(c, derror.Error(err))
		c.Abort()
		return
	}

	c.Set("DiancaiUser", user)
}

func MiddlewareAuthRequired(c *gin.Context) {
	_, ok := c.Get("DiancaiUser")
	if !ok {
		common.HandlerReturnError(c, common.ErrUnauthorized)
		c.Abort()
	}
}
