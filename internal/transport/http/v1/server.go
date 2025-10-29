package v1

import (
	"context"
	"debez/pkg/logger"
	"net/http"
	"time"
)

const (
	user       = "/api/v1/user"
	userGet    = "GET " + user + "/getUser"
	userGetAll = "GET " + user + "/getUsers"
	userSave   = "POST " + user + "/saveUser"
	userUpdate = "POST " + user + "/updateUser"
	userDelete = "DELETE " + user + "/deleteUser"
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
)

type Server struct {
	srv *http.Server
}

func NewServer(port string) *Server {
	srv := http.Server{
		Addr:              ":" + port,
		Handler:           nil,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
	}

	return &Server{srv: &srv}
}

func (s *Server) RegisterHandlers(log logger.Logger, handler *Handler) error {

	mux := http.NewServeMux()

	mux.HandleFunc(userGet, handler.GetUser)
	mux.HandleFunc(userGetAll, handler.GetUsers)
	mux.HandleFunc(userSave, handler.SaveUser)
	mux.HandleFunc(userUpdate, handler.UpdateUser)
	mux.HandleFunc(userDelete, handler.DeleteUser)

	s.srv.Handler = AddMetadata(RegistLoggerMiddleware(log, LoggingMiddleware(mux)))

	return nil
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
