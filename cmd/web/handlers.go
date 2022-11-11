package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"goskurikhin/pkg/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (srv *server) home(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}

func (srv *server) create(c *gin.Context) {
	// Decoding the incoming message
	reqBody := new(models.UrlMapping)
	responseEncoder := json.NewEncoder(c.Writer)
	if err := json.NewDecoder(c.Request.Body).Decode(&reqBody); err != nil {
		SendError(ErrServer, c, responseEncoder, 500)
		return
	}

	// Checking for a shortname in a data source
	if DoesExist, err := srv.dataSource.DoesExistShortname(reqBody.Shortname); DoesExist {
		SendError(ErrIsExist, c, responseEncoder, 400)
		return
	} else if err != nil {
		SendError(ErrServer, c, responseEncoder, 500)
		return
	}

	if !govalidator.IsURL(reqBody.Longurl) {
		SendError(ErrWrongFormat, c, responseEncoder, 400)
		return
	}

	// Create URL structure
	longurl, err := url.Parse(reqBody.Longurl)
	if err != nil {
		SendError(ErrServer, c, responseEncoder, 500)
		return
	}

	if !longurl.IsAbs() {
		longurl.Scheme = "http"
	}

	err = srv.dataSource.InsertLinks(longurl.String(), reqBody.Shortname)
	if err != nil {
		SendError(ErrServer, c, responseEncoder, 500)
		return
	}
	shorturl := srv.addr + "/re/" + reqBody.Shortname
	responseEncoder.Encode(&models.APIResponse{StatusMessage: shorturl})
}

func (srv *server) redirect(c *gin.Context) {
	responseEncoder := json.NewEncoder(c.Writer)
	shortname := c.Param("shortname")
	longurl, err := srv.dataSource.FindLongLink(shortname)
	if err != nil {
		SendError(ErrServer, c, responseEncoder, 500)
		return
	}
	c.Redirect(http.StatusFound, longurl)
}
