package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/nudopnu/scraper/internal/auth"
	"github.com/nudopnu/scraper/internal/config"
	"github.com/nudopnu/scraper/internal/database"
)

func initAdmin(s *State) error {
	_, err := s.db.GetAdmin(context.Background())
	if err == nil {
		return nil
	}
	username, ok := os.LookupEnv("ADMIN_USERNAME")
	if !ok {
		return errors.New("no admin username provided")
	}
	password, ok := os.LookupEnv("ADMIN_PASSWORD")
	if !ok {
		return errors.New("no admin password provided")
	}
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	_, err = s.db.RegisterAdmin(context.Background(), database.RegisterAdminParams{
		Username:       username,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return err
	}
	return nil
}

func initState() (*State, error) {
	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", config.GetDbUrl())
	dbQueries := database.New(db)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %q", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	state := State{
		db:  dbQueries,
		cfg: config,
		ls: &LoopState{
			ctx:         ctx,
			taskQueue:   make(chan Task),
			cancelQueue: cancel,
		},
	}
	err = initAdmin(&state)
	if err != nil {
		return nil, fmt.Errorf("error initializing admin user: %q", err)
	}
	return &state, nil
}
