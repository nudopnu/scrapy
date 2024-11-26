package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/nudopnu/scraper/internal/customerror"
	"github.com/nudopnu/scraper/internal/database"
)

type Ad struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Price       string    `json:"price"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	PostalCode  string    `json:"postal_code"`
	CategoryID  string    `json:"category_id"`
	PostedAt    string    `json:"posted_at"`
	Link        string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Images      []Image   `json:"images"`
}

type Image struct {
	ImageUrl    string `json:"image_url"`
	ImageNumber int    `json:"image_number"`
}

func HandlerAdsByAgent(w JSONResponseWriter, r *http.Request, s *State, user database.User) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.error(http.StatusBadRequest, customerror.New("error parsing path parameter", err))
		return
	}
	page, err := strconv.Atoi(r.PathValue("page"))
	if err != nil {
		page = 0
	}
	pageSize := 31
	dbAds, err := s.db.ListResultsFromAgent(r.Context(), database.ListResultsFromAgentParams{
		ID:     int32(id),
		Limit:  int32(pageSize),
		Offset: int32(page * pageSize),
	})
	ads := make([]Ad, 0, len(dbAds))
	for i, dbAd := range dbAds {
		ads = append(ads, Ad{
			ID:          dbAd.AdID,
			Title:       dbAd.Title,
			Price:       dbAd.Price.String,
			Description: dbAd.Description.String,
			Location:    dbAd.Location.String,
			PostalCode:  dbAd.PostalCode.String,
			CategoryID:  dbAd.CategoryID.String,
			PostedAt:    dbAd.PostedAt.String,
			Link:        dbAd.Link.String,
			CreatedAt:   dbAd.CreatedAt,
			UpdatedAt:   dbAd.UpdatedAt,
			Images:      []Image{},
		})
		if err := json.Unmarshal(dbAd.Images, &ads[i].Images); err != nil {
			w.error(http.StatusInternalServerError, customerror.New("error parsing image links", err))
			return
		}
	}
	if err != nil {
		w.error(http.StatusInternalServerError, customerror.New("error fetching results", err))
		return
	}
	w.json(http.StatusOK, ads)
}
