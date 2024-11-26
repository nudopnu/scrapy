package main

import (
	"net/http"

	"github.com/nudopnu/scraper/internal/database"
)

func HandlerStartLoop(w JSONResponseWriter, r *http.Request, s *State, u database.User) {
	s.StartMainLoop()
}

func HandlerStopLoop(w JSONResponseWriter, r *http.Request, s *State, u database.User) {
	s.ls.cancelQueue()
}
