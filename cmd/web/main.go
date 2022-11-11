package main

import (
	"log"
	"os"
)

type params struct {
	addr string
	dsn  string
}

func getParams() (*params, error) {
	var web_host, web_port, dataSource_host, dataSource_port string
	var exists bool
	if web_host, exists = os.LookupEnv("WEB_HOST"); !exists {
		return nil, NoPATHenvironment
	}
	if web_port, exists = os.LookupEnv("WEB_PORT"); !exists {
		return nil, NoPATHenvironment
	}
	if dataSource_host, exists = os.LookupEnv("DATASOURCE_HOST"); !exists {
		return nil, NoPATHenvironment
	}
	if dataSource_port, exists = os.LookupEnv("DATASOURCE_PORT"); !exists {
		return nil, NoPATHenvironment
	}

	return &params{
		addr: web_host + ":" + web_port,
		dsn:  "mongodb://" + dataSource_host + dataSource_port,
	}, nil
}

func main() {
	p, err := getParams()
	if err != nil {
		log.Fatal(err)
	}

	srv := NewServer()
	log.Fatal(srv.Start(p))
	defer srv.Close()

}
