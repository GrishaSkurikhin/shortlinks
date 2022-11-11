package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"goskurikhin/pkg/models"

	"github.com/gin-gonic/gin"
)

var NoPATHenvironment = errors.New("Cannot find PATH environment")
var ErrIsExist = errors.New("URL already exists")
var ErrWrongFormat = errors.New("Wrong format of URL")
var ErrServer = errors.New("Internal server problem")

func SendError(err error, c *gin.Context, responseEncoder *json.Encoder, errCode int) {
	c.Writer.WriteHeader(errCode)

	if err := responseEncoder.Encode(&models.APIResponse{StatusMessage: err.Error()}); err != nil {
		fmt.Fprintf(c.Writer, err.Error())
	}
}
