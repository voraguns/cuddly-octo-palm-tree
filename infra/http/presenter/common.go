package presenter

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidParam is invalid parameter error
	ErrInvalidParam = errors.New("invalid parameter")
	// ErrUnauthorized is unauthorized error
	ErrUnauthorized = errors.New("unauthorized")
	// ErrServer is server error
	ErrServer = errors.New("server error")
)

// ErrResponse is the error response type
type ErrResponse struct {
	Message string `json:"msg"`
}

// SuccessMessage is the success response type
type SuccessMessage struct {
	Message string `json:"msg" example:"ok"`
}

// OkMsg is the default success response for 200 status code
var OkMsg SuccessMessage = SuccessMessage{
	Message: "ok",
}

func ErrorResponse(c *gin.Context, httpCode int, err error) {
	message := err.Error()
	c.JSON(httpCode, &ErrResponse{
		Message: message,
	})
}
