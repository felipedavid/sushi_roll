package main

import "net/http"

func (a *app) logRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		rw := newResponseWriter(w)

		next.ServeHTTP(rw, r)

		a.infoLog.Printf("%s \"%s %s\" -> %d %s\n",
			r.RemoteAddr, r.Method, r.URL.Path, rw.status, http.StatusText(rw.status))
	}

	return http.HandlerFunc(fn)
}
