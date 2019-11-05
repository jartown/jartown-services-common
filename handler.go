package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.com/diancai/diancai-services-common/derror"
)

func HandlerReturnError(c *gin.Context, err error) {
	derr := derror.Error(err)

	var reply interface{}
	if Conf.Debug {
		reply = derr
	} else {
		reply = ErrorReply{ErrorMsg: derr.Error()}
	}

	if derr.Err == ErrUnauthorized || derr.Err == ErrPasswordMismatch {
		c.JSON(http.StatusUnauthorized, reply)
	} else if derr.Err == ErrItemNotFound {
		c.JSON(http.StatusNotFound, reply)
	} else if derr.Err == ErrForbidden {
		c.JSON(http.StatusForbidden, reply)
	} else if derr.Err == ErrInputValidationFailed {
		c.JSON(http.StatusBadRequest, reply)
	} else {
		c.JSON(http.StatusInternalServerError, reply)
	}
}

func HandlerReturnData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func HandlerRegisterSwagger(r *gin.Engine) {
	r.StaticFile("/swagger.yml", "./swagger.yml")

	url := ginSwagger.URL("/swagger.yml")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
