package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/nudopnu/scraper/internal/database"
	"github.com/nudopnu/scraper/internal/ebay"
)

// func HandlerRun(currentUser database.User, s *State, command Command) error {
// 	for {
// 		agent, err := s.db.GetNextAgentToUpdate(context.Background())
// 		if err != nil {
// 			return err
// 		}

// 		err = s.db.MarkAgentUpdated(context.Background(), agent.ID)
// 		if err != nil {
// 			return fmt.Errorf("error marking agent '%s' as updated: %w", agent.Name, err)
// 		}
// 		fmt.Printf("updated agent '%s'\n", agent.Name)

// 		searchParams, err := s.db.GetSearchParamsBySearchAgent(context.Background(), agent.ID)
// 		if err != nil {
// 			return fmt.Errorf("error getting search params for agent '%s': %w", agent.Name, err)
// 		}
// 		fmt.Printf("retrieved %d search params for agent '%s'\n", len(searchParams), agent.Name)

// 		for _, searchParam := range searchParams {
// 			oldResults, err := s.db.GetResultsByParamId(context.Background(), searchParam.ID)
// 			if err != nil && !errors.Is(err, sql.ErrNoRows) {
// 				return fmt.Errorf("error fetching old results: %w", err)
// 			}
// 			oldResultsByEbayId := make(map[string]database.GetResultsByParamIdRow)
// 			for _, result := range oldResults {
// 				oldResultsByEbayId[result.EbayID] = result
// 			}

// 			parsedAds, _, err := fetchAds(searchParam, 0)
// 			if err != nil {
// 				return err
// 			}

// 			for _, parsedAd := range parsedAds {
// 				oldAd, err := s.db.GetAdByEbayId(context.Background(), parsedAd.Id)
// 				if err != nil {
// 					if !errors.Is(err, sql.ErrNoRows) {
// 						return fmt.Errorf("error getting ad: %w", err)
// 					}
// 					ad, err := s.db.CreateAd(context.Background(), fromParsedAd(parsedAd))
// 					if err != nil {
// 						return fmt.Errorf("error creating ad: %w", err)
// 					}
// 					fmt.Printf("INFO: successfully created ad '%s' with id '%d'\n", ad.Title, ad.ID)
// 					_, err = s.db.CreateResult(context.Background(), database.CreateResultParams{
// 						ParamsID: searchParam.ID,
// 						AdID:     ad.ID,
// 						Status:   "new",
// 					})
// 					if err != nil {
// 						return fmt.Errorf("error creating search result: %w", err)
// 					}
// 					continue
// 				}
// 				oldResult, ok := oldResultsByEbayId[oldAd.EbayID]
// 				if ok {
// 					err = s.db.UpdateResultStatus(context.Background(), database.UpdateResultStatusParams{
// 						ID:     oldResult.ID,
// 						Status: "fresh",
// 					})
// 					if err != nil {
// 						return fmt.Errorf("error updating result: %w", err)
// 					}
// 					fmt.Printf("INFO: Successfully updated search result\n")
// 				} else {
// 					_, err = s.db.CreateResult(context.Background(), database.CreateResultParams{
// 						ParamsID: searchParam.ID,
// 						AdID:     oldAd.ID,
// 						Status:   "new",
// 					})
// 					if err != nil {
// 						return fmt.Errorf("error creating search result: %w", err)
// 					}
// 					fmt.Printf("INFO: successfully added existing ad to result for agent '%s'\n", agent.Name)
// 				}
// 			}
// 			time.Sleep(3 * time.Second)
// 		}
// 		<-time.Tick(60 * time.Second)
// 	}
// }

func fetchAds(searchParam database.Param, pageNumber int) ([]ebay.ParsedAd, bool, error) {
	bytes, err := ebay.GetAdsRaw(ebay.GetAdsRequest{
		Keyword:    searchParam.Keyword,
		LocationId: searchParam.LocationID,
		Distance:   int(searchParam.Distance),
		PageNumber: pageNumber,
	})
	if err != nil {
		return nil, false, fmt.Errorf("error fetching ads: %w", err)
	}
	ads, err := ebay.UnmarshalAds(bytes)
	if err != nil {
		return nil, false, fmt.Errorf("error parsing ads: %w", err)
	}
	numFound, err := strconv.Atoi(ads.HTTPWWWEbayclassifiedsgroupCOMSchemaAdV1Ads.Value.Paging.NumFound)
	if err != nil {
		return nil, false, fmt.Errorf("error parsing ads: %w", err)
	}
	log.Printf("Fetched %d/%d items\n", int(math.Min(float64((pageNumber+1)*int(ads.SearchOptions.Size)), float64(numFound))), numFound)
	hasMorePages := (pageNumber+1)*int(ads.SearchOptions.Size) < numFound
	return ebay.ParseAds(ads), hasMorePages, nil
}
