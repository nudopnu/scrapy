package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nudopnu/scraper/internal/auth"
	"github.com/nudopnu/scraper/internal/customerror"
	"github.com/nudopnu/scraper/internal/database"
)

func (state *State) AuthorizedByAccessToken(role string, handler func(w JSONResponseWriter, r *http.Request, s *State, user database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonWriter := JSONResponseWriter{ResponseWriter: w}
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		id, err := auth.ValidateJWT(token, state.cfg.JwtSecret)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := state.db.GetUserById(r.Context(), id)
		if err != nil || (user.Role != role && user.Role != "admin") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler(jsonWriter, r, state, user)
	}
}

func (state *State) AuthorizedByRefreshToken(handler func(w JSONResponseWriter, r *http.Request, s *State, user database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonWriter := JSONResponseWriter{ResponseWriter: w}
		sessionTokenString, err := auth.GetBearerToken(r.Header)
		if err != nil {
			jsonWriter.error(http.StatusUnauthorized, customerror.New("invalid refresh token", err))
			return
		}
		sessionToken, err := state.db.GetRefreshToken(r.Context(), sessionTokenString)
		if err != nil {
			jsonWriter.error(http.StatusUnauthorized, customerror.New("invalid refresh token", err))
			return
		}
		tokenExpired := time.Now().After(sessionToken.ExpiresAt)
		tokenRevoked := sessionToken.RevokedAt.Valid
		if tokenRevoked || tokenExpired {
			jsonWriter.error(http.StatusUnauthorized, customerror.New("invalid refresh token", fmt.Errorf("%v > %v", sessionToken.ExpiresAt, time.Now())))
			return
		}
		user, err := state.db.GetUserById(r.Context(), sessionToken.UserID)
		if err != nil {
			jsonWriter.error(http.StatusInternalServerError, customerror.New("inavlid user id", err))
		}
		handler(jsonWriter, r, state, user)
	}
}

func (state *State) Default(handler func(w JSONResponseWriter, r *http.Request, s *State)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonWriter := JSONResponseWriter{ResponseWriter: w}
		handler(jsonWriter, r, state)
	}
}
