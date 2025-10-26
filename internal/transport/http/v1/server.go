package v1

import (
	"context"
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

func (s *Server) RegisterHandlers(handler *Handler) error {
	http.HandleFunc(userGet, handler.GetUser)
	http.HandleFunc(userGetAll, handler.GetUsers)
	http.HandleFunc(userSave, handler.SaveUser)
	http.HandleFunc(userUpdate, handler.UpdateUser)
	http.HandleFunc(userDelete, handler.DeleteUser)

	return nil
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
