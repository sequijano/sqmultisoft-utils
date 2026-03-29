package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Success: true, Data: data})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{Success: true, Data: data})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{Success: false, Error: message})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{Success: false, Error: message})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{Success: false, Error: "access denied"})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{Success: false, Error: message})
}

func InternalError(c *gin.Context, err error) {
	log.Printf("[ERROR] %s %s → %v", c.Request.Method, c.Request.URL.Path, err)
	c.JSON(http.StatusInternalServerError, Response{Success: false, Error: err.Error()})
}

func PaymentRequired(c *gin.Context, message string) {
	c.JSON(402, Response{Success: false, Error: message})
}
