package handlers

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	Setup(*gin.Engine)
}

var ErrUnexpected error = errors.New("unexpected error ocurred")

func NewErrorRes(err error) ErrorRes {
	return ErrorRes{
		Message: err.Error(),
	}
}

type ErrorRes struct {
	Message string `json:"message"`
}

func (er ErrorRes) ToString() string {
	myString, _ := json.Marshal(er)
	return string(myString)
}
