package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nudopnu/scraper/internal/database"
	"github.com/nudopnu/scraper/internal/ebay"
)

func (s *State) NewUpdateAgentTask(id int) *Task {
	name := "Update Agent"
	operation := func() error {
		params, err := s.db.GetSearchParamsBySearchAgent(s.ls.ctx, int32(id))
		if err != nil {
			return err
		}
		for _, searchParam := range params {
			// fetch new ads
			newAds := []ebay.ParsedAd{}
			pageNumber := 0
			for {
				parsedAds, hasMorePages, err := fetchAds(searchParam, pageNumber)
				if err != nil {
					return fmt.Errorf("error fetching new ads: %w", err)
				}
				newIds := make([]string, 0, len(parsedAds))
				for _, ad := range parsedAds {
					newIds = append(newIds, ad.Id)
				}
				numberOfNewIds, err := s.db.GetNumberOfDuplicates(s.ls.ctx, newIds)
				if err != nil {
					return err
				}
				if !hasMorePages || int(numberOfNewIds) == len(parsedAds) {
					break
				}
				newAds = append(newAds, parsedAds...)
				time.Sleep(1 * time.Second)
				pageNumber++
			}
			// update existing results
			newAdEbayIds := make([]string, 0, len(newAds))
			adIds, labels, urls, imgNumbers := []string{}, []string{}, []string{}, []int32{}
			for _, ad := range newAds {
				newAdEbayIds = append(newAdEbayIds, ad.Id)
				for imgNumber, picture := range ad.Pictures {
					for label := range picture {
						labels = append(labels, label)
						urls = append(urls, picture[label])
						imgNumbers = append(imgNumbers, int32(imgNumber))
						adIds = append(adIds, ad.Id)
					}
				}
			}
			err = s.db.UpdateResultExpired(context.Background(), newAdEbayIds)
			if err != nil {
				return fmt.Errorf("error updating results: %w", err)
			}
			err = s.db.UpdateResultUpdated(context.Background(), newAdEbayIds)
			if err != nil {
				return fmt.Errorf("error updating results: %w", err)
			}
			// insert new ads
			ids, titles, prices, descriptions, locations, postal_codes, category_ids, posted_ats, links := extractAdFields(newAds)
			ads, err := s.db.BulkCreateAds(context.Background(), database.BulkCreateAdsParams{
				Column1: ids,
				Column2: titles,
				Column3: prices,
				Column4: descriptions,
				Column5: locations,
				Column6: postal_codes,
				Column7: category_ids,
				Column8: posted_ats,
				Column9: links,
			})
			if err != nil {
				return fmt.Errorf("error creating ads: %w", err)
			}
			// add new images
			_, err = s.db.BulkCreateImages(s.ls.ctx, database.BulkCreateImagesParams{
				Column1: adIds,
				Column2: labels,
				Column3: imgNumbers,
				Column4: urls,
			})
			if err != nil {
				return fmt.Errorf("error creating images: %w", err)
			}
			// insert new results
			paramIds, newAdIds := extractResultFields(ads, searchParam)
			_, err = s.db.BulkCreateResults(context.Background(), database.BulkCreateResultsParams{
				Column1: paramIds,
				Column2: newAdIds,
			})
			if err != nil {
				return fmt.Errorf("error creating results: %w", err)
			}
		}
		err = s.db.MarkAgentUpdated(s.ls.ctx, int32(id))
		if err != nil {
			return err
		}
		return nil
	}
	return &Task{
		name:      name,
		operation: operation,
	}
}

func extractAdFields(ads []ebay.ParsedAd) (ids, titles, prices, descriptions, locations, postal_codes, category_ids, posted_ats, links []string) {
	for _, ad := range ads {
		ids = append(ids, ad.Id)
		titles = append(titles, ad.Title)
		prices = append(prices, ad.Price)
		descriptions = append(descriptions, ad.Description)
		locations = append(locations, ad.Location)
		postal_codes = append(postal_codes, ad.ZipCode)
		category_ids = append(category_ids, ad.CategoryId)
		posted_ats = append(posted_ats, ad.Date)
		links = append(links, ad.Link)
	}
	return
}

func extractResultFields(ads []database.Ad, params database.Param) (paramIds []int32, newAdIds []string) {
	for _, ad := range ads {
		paramIds = append(paramIds, params.ID)
		newAdIds = append(newAdIds, ad.ID)
	}
	return
}

func extractImageFields(images map[string]string) (labels, urls []string) {
	for label := range images {
		labels = append(labels, label)
		urls = append(urls, images[label])
	}
	return
}
