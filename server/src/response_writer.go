package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nudopnu/scraper/internal/customerror"
)

type JSONResponseWriter struct {
	http.ResponseWriter
}

func (w *JSONResponseWriter) json(statusCode int, data interface{}) {
	dat, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
	log.Printf("%d: %+v", statusCode, data)
}

func (w *JSONResponseWriter) error(statusCode int, customError customerror.CustomError) {
	log.Printf("ERROR: %d - %s\n", statusCode, customError.LogMessage)
	errorBody := struct {
		Message string `json:"message"`
	}{
		Message: customError.UserMessage,
	}
	data, err := json.Marshal(errorBody)
	if err != nil {
		log.Printf("error marshalling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
