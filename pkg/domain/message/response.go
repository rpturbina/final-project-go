package message

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code       int        `json:"code"`
	Message    string     `json:"message,omitempty"`
	Type       string     `json:"type,omitempty"`
	Data       any        `json:"data,omitempty"`
	InvalidArg any        `json:"invalid_arg,omitempty"`
	StartTime  *time.Time `json:"start_time,omitempty"`
}

func ErrorResponseSwitcher(ctx *gin.Context, errMsg ErrorMessage) {
	var httpStatusCode int
	var response Response
	switch errMsg.Type {
	case "USERNAME_IS_EMPTY", "EMAIL_IS_EMPTY", "WRONG_EMAIL_FORMAT", "PASSWORD_IS_EMPTY", "INVALID_PASSWORD_FORMAT", "INVALID_PAYLOAD", "USERNAME_REGISTERED", "EMAIL_REGISTERED", "USER_NOT_FOUND", "PHOTO_TITLE_IS_EMPTY", "PHOTO_URL_IS_EMPTY", "INVALID_URL_FORMAT", "PHOTO_NOT_FOUND", "COMMENT_MESSAGE_IS_EMPTY", "PHOTO_ID_IS_EMPTY", "USER_COMMENT_NOT_FOUND", "COMMENT_NOT_FOUND", "SOCIAL_MEDIA_NAME_IS_EMPTY", "SOCIAL_MEDIA_URL_IS_EMPTY", "INVALID_FORMAT", "SOCIAL_MEDIA_NOT_FOUND":
		httpStatusCode = http.StatusBadRequest
		response = Response{
			Code:    96,
			Message: errMsg.Error.Error(),
			Type:    "BAD_REQUEST",
			InvalidArg: gin.H{
				"error_message": errMsg.Error.Error(),
				"error_type":    errMsg.Type,
			},
		}

	case "WRONG_PASSWORD":
		httpStatusCode = http.StatusUnauthorized
		response = Response{
			Code:    97,
			Message: errMsg.Error.Error(),
			Type:    "UNAUTHENTICATED",
			InvalidArg: gin.H{
				"error_message": errMsg.Error.Error(),
				"error_type":    errMsg.Type,
			},
		}

	case "INVALID_SCOPE":
		httpStatusCode = http.StatusForbidden
		response = Response{
			Code:    98,
			Message: errMsg.Error.Error(),
			Type:    "FORBIDDEN",
			InvalidArg: gin.H{
				"error_message": errMsg.Error.Error(),
				"error_type":    errMsg.Type,
			},
		}

	case "INTERNAL_CONNECTION_PROBLEM":
		httpStatusCode = http.StatusInternalServerError
		response = Response{
			Code:    99,
			Message: errMsg.Error.Error(),
			Type:    "INTERNAL_SERVER_ERROR",
			InvalidArg: gin.H{
				"error_message": errMsg.Error.Error(),
				"error_type":    errMsg.Type,
			},
		}

	}

	ctx.AbortWithStatusJSON(httpStatusCode, response)
}
