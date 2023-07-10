package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Message
}

// exported function that takes a a gin context and an error and makes the corresponding response
func Respond(c *gin.Context, err *Error) {
	//logger.Errorf("Respond Error: %d", err, err.Status)
	if err.Status == http.StatusMultipleChoices {
		c.Redirect(err.Status, err.Message)
	}
	c.JSON(err.Status, err)
}

func NoRouteHandler(c *gin.Context) {
	Respond(c, NotFound(fmt.Sprintf("Resource %s not found", c.Request.URL.Path)))
}

func NoMethodHandler(c *gin.Context) {
	Respond(c, MethodNotAllowed())
}

func Validation(message string) *Error {
	return &Error{message, http.StatusBadRequest}
}

func BadRequest(message string) *Error {
	return &Error{message, http.StatusBadRequest}
}

func NotFound(message string) *Error {
	return &Error{message, http.StatusNotFound}
}

func MethodNotAllowed() *Error {
	return &Error{"Method not allowed", http.StatusMethodNotAllowed}
}

func Unauthorized(message string) *Error {
	return &Error{message, http.StatusUnauthorized}
}

func InternalServer(message string, err ...error) *Error {
	if len(err) != 0 {
		//logger.Error(message, err[0])
		message = fmt.Sprintf("%s - ERROR: %v", message, err)
	}
	return &Error{message, http.StatusInternalServerError}
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

var (
	NoAutorizado = Unauthorized("usuario no autorizado para esta zona")
)
