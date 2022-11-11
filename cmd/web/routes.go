package main

func (srv *server) createRoutes() {
	srv.router.LoadHTMLGlob("./ui/templates/*")
	srv.router.Static("/static", "./ui/static/")

	srv.router.GET("/", srv.home)
	srv.router.GET("/re/:shortname", srv.redirect)
	srv.router.POST("/create", srv.create)
}
