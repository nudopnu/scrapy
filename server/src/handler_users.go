package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/nudopnu/scraper/internal/auth"
	"github.com/nudopnu/scraper/internal/customerror"
	"github.com/nudopnu/scraper/internal/database"
)

type User struct {
	Id           int       `json:"id"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	AccessToken  string    `json:"access_token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

type UserInfo struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func HandlerRegister(w JSONResponseWriter, r *http.Request, s *State) {
	request := struct {
		Username string `json:"username" validate:"required,alphanum,min=3,max=20"`
		Password string `json:"password" validate:"required,min=8,max=32,password"`
	}{}
	err := ParseAndValidate(r.Body, &request)
	if err != nil {
		w.error(http.StatusBadRequest, customerror.New("invalid register request", err))
		return
	}
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error hashing password", err))
		return
	}
	user, err := s.db.RegisterUser(context.Background(), database.RegisterUserParams{
		Username:       request.Username,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		w.error(http.StatusBadRequest, customerror.New("username already exists", err))
		return
	}
	log.Printf("sucessfully registered user '%s'\n", user.Username)
	w.json(http.StatusCreated, User{
		Id:        int(user.ID),
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	})
}

func HandlerUserInfo(w JSONResponseWriter, r *http.Request, s *State, user database.User) {
	w.json(http.StatusOK, UserInfo{
		Id:        int(user.ID),
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

func HandlerLogin(w JSONResponseWriter, r *http.Request, s *State) {
	request := struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		w.error(http.StatusBadRequest, customerror.New("invalid login request", err))
		return
	}
	username := request.Username
	user, err := s.db.GetUserByUsername(context.Background(), username)
	if err != nil {
		w.error(http.StatusUnauthorized, customerror.New("invalid username or password", err))
		return
	}
	if err = auth.CheckPasswordHash(request.Password, user.HashedPassword); err != nil {
		w.error(http.StatusUnauthorized, customerror.New("invalid username or password", err))
		return
	}
	log.Printf("INFO: successfully logged in as '%s'\n", user.Username)
	jwt, err := auth.MakeJWT(user.Username, int(user.ID), s.cfg.Server.JwtSecret, 1*time.Hour)
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error creating access token", err))
		return
	}
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error creating refresh token", err))
		return
	}
	_, err = s.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	})
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error creating refresh token", err))
		return
	}
	w.json(http.StatusOK, User{
		Id:           int(user.ID),
		Username:     user.Username,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		AccessToken:  jwt,
		RefreshToken: refreshToken,
	})
}
