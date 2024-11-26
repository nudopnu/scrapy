package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nudopnu/scraper/internal/customerror"
	"github.com/nudopnu/scraper/internal/database"
)

type Agent struct {
	Id            int          `json:"id"`
	Name          string       `json:"name"`
	UserId        int          `json:"user_id"`
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Thumbnail     string       `json:"thumbnail_url,omitempty"`
}

func HandlerListSearchAgents(w JSONResponseWriter, r *http.Request, s *State, u database.User) {
	agents, err := s.db.ListAgentsWithImages(r.Context())
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error fetching search agents", err))
		return
	}
	response := make([]Agent, 0, len(agents))
	for _, a := range agents {
		response = append(response, Agent{
			Id:            int(a.ID),
			Name:          a.Name,
			UserId:        int(a.UserID),
			LastFetchedAt: a.LastFetchedAt,
			CreatedAt:     a.CreatedAt,
			UpdatedAt:     a.UpdatedAt,
			Thumbnail:     a.Thumbnail,
		})
	}
	w.json(http.StatusOK, response)
}

func HandlerAddSearchAgent(w JSONResponseWriter, r *http.Request, s *State, u database.User) {
	request := struct {
		Name       string `json:"name" validate:"required,min=3,max=20"`
		Keyword    string `json:"keyword" validate:"required,min=3,max=20"`
		PostalCode string `json:"postal_code" validate:"required"`
		Distance   int    `json:"distance" validate:"required"`
	}{}
	if err := ParseAndValidate(r.Body, &request); err != nil {
		fmt.Print(err)
		w.error(http.StatusBadRequest, customerror.New("invalid request data", err))
		return
	}
	agent, err := s.db.CreateSearchAgent(context.Background(), database.CreateSearchAgentParams{
		Name:   request.Name,
		UserID: u.ID,
	})
	if err != nil {
		w.error(http.StatusBadRequest, customerror.New("error creating agent", err))
		return
	}
	location, err := s.GetOrCreateLocation(request.PostalCode)
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error getting location id", err))
		return
	}
	searchParams, err := s.db.CreateSearchParams(context.Background(), database.CreateSearchParamsParams{
		Keyword:    request.Keyword,
		LocationID: location.LocationID,
		Distance:   int32(request.Distance),
	})
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error creating search params", err))
		return
	}
	par, err := s.db.AddSearchParamToAgent(context.Background(), database.AddSearchParamToAgentParams{
		AgentID:  agent.ID,
		ParamsID: searchParams.ID,
	})
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error adding search params to agent", err))
		return
	}
	log.Printf("INFO: successfully added agent '%s' %v\n", agent.Name, par)
	w.json(http.StatusOK, Agent{
		Id:            int(agent.ID),
		Name:          agent.Name,
		UserId:        int(agent.UserID),
		LastFetchedAt: agent.LastFetchedAt,
		CreatedAt:     agent.CreatedAt,
		UpdatedAt:     agent.UpdatedAt,
	})
}
