package httpx

import (
	"log"
	"net/http"

	"github.com/ory/graceful"
)

type GracefulServer struct {
	*http.Server
}

func NewGracefulServer(addr string, handler http.Handler) GracefulServer {
	return GracefulServer{&http.Server{
		Addr:    addr,
		Handler: handler,
	}}
}

func (s *GracefulServer) StartWithGracefulShutdown() {
	if err := graceful.Graceful(s.ListenAndServe, s.Shutdown); err != nil {
		log.Fatal(err)
	}
}
