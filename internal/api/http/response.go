package http

import h "net/http"

type Response struct {
	Status int
	Body   interface{}
	Cookie *h.Cookie
}
