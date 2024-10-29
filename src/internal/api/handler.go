package api

import "net/http"

type ApiHandler interface {
}

type CrudHandler interface {
	ApiHandler
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Detail(w http.ResponseWriter, r *http.Request)
	PagedList(w http.ResponseWriter, r *http.Request)
	PageCount(w http.ResponseWriter, r *http.Request)
}
