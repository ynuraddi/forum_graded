package transport

import "net/http"

func initDefaultRouter() http.Handler {
	mux := new(http.ServeMux)

	return mux
}
