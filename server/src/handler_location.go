package main

import (
	"context"
	"fmt"

	"github.com/nudopnu/scraper/internal/database"
	"github.com/nudopnu/scraper/internal/ebay"
)

func (s *State) GetOrCreateLocation(postalCode string) (database.Location, error) {
	location, err := s.db.GetLocationByPostalCode(context.Background(), postalCode)
	if err == nil {
		return location, nil
	}
	locationId, err := ebay.GetLocationId(postalCode)
	if err != nil {
		return database.Location{}, fmt.Errorf("error fetching location id for '%s': %w", postalCode, err)
	}
	location, err = s.db.AddLocation(context.Background(), database.AddLocationParams{
		PostalCode: postalCode,
		LocationID: locationId,
	})
	if err != nil {
		return database.Location{}, err
	}
	return location, nil
}
