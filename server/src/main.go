package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nudopnu/scraper/internal/config"
	"github.com/nudopnu/scraper/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
	ls  *LoopState
}

func main() {
	handler, err := initState()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})
	mux.HandleFunc("POST /api/v1/users", handler.Default(HandlerRegister))
	mux.HandleFunc("POST /api/v1/userinfo", handler.AuthorizedByAccessToken("user", HandlerUserInfo))
	mux.HandleFunc("POST /api/v1/login", handler.Default(HandlerLogin))
	mux.HandleFunc("POST /api/v1/refresh", handler.AuthorizedByRefreshToken(HandlerRefresh))

	mux.HandleFunc("GET /api/v1/agents", handler.AuthorizedByAccessToken("user", HandlerListSearchAgents))
	mux.HandleFunc("POST /api/v1/agents", handler.AuthorizedByAccessToken("user", HandlerAddSearchAgent))
	mux.HandleFunc("GET /api/v1/agents/{id}", handler.AuthorizedByAccessToken("user", HandlerAdsByAgent))

	mux.HandleFunc("POST /api/v1/loop/start", handler.AuthorizedByAccessToken("admin", HandlerStartLoop))
	mux.HandleFunc("POST /api/v1/loop/stop", handler.AuthorizedByAccessToken("admin", HandlerStopLoop))

	handler.StartMainLoop()

	fmt.Printf("listening on %s\n", handler.cfg.Host)
	log.Fatal(http.ListenAndServe(handler.cfg.Host, CORSMiddleware(mux)))
}
