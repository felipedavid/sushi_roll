package main

import "net/http"

// responseWriter é apenas um wrapper em volta do ResponseWriter entregue pelo pacote http, com o único intuito
// de tornar disponível o status da resposta após a chamada ao método WriteHeader
type responseWriter struct {
	http.ResponseWriter
	status int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: 200}
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}
