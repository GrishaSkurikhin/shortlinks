package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"goskurikhin/pkg/dataSource"
	"goskurikhin/pkg/dataSource/mongoDB"
)

type server struct {
	addr       string
	errorLog   *log.Logger
	infoLog    *log.Logger
	router     *gin.Engine
	dataSource dataSource.DataSource
}

func NewServer() *server {
	return &server{
		addr:       "",
		errorLog:   log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:    log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		router:     gin.Default(),
		dataSource: mongoDB.NewMongoDB(),
	}
}

func (srv *server) Start(p *params) error {
	err := srv.dataSource.Open(p.dsn)
	if err != nil {
		return err
	}

	//app.server.ErrorLog = app.errorLog
	srv.addr = p.addr
	srv.createRoutes()
	err = srv.router.Run(srv.addr)
	if err != nil {
		return err
	}
	return nil
}

func (srv *server) Close() {
	srv.dataSource.Close()
}
