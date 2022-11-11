package models

type UrlMapping struct {
	Longurl   string `json:longurl`
	Shortname string `json:shortname`
}

type APIResponse struct {
	StatusMessage string `json:statusmessage`
}
