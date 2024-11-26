package main

import (
	"net/http"
	"time"

	"github.com/nudopnu/scraper/internal/auth"
	"github.com/nudopnu/scraper/internal/customerror"
	"github.com/nudopnu/scraper/internal/database"
)

func HandlerRefresh(w JSONResponseWriter, r *http.Request, state *State, user database.User) {
	accessToken, err := auth.MakeJWT(user.Username, int(user.ID), state.cfg.JwtSecret, 1*time.Hour)
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error creating access token", err))
		return
	}
	w.json(http.StatusOK, struct {
		AccessToken string `json:"access_token"`
	}{AccessToken: accessToken})
}
