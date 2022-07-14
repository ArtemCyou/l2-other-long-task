package api

import (
	"dev11/internals/app/handlers"
	"net/http"
)

func CreateApiHandlers() *http.ServeMux {
	//создаем мультиплексор HTTP-запросов
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event/", handlers.CreateHandler)
	mux.HandleFunc("/update_event", handlers.CreateHandler)
	mux.HandleFunc("/delete_event", handlers.CreateHandler)
	mux.HandleFunc("/events_for_day", handlers.CreateHandler)
	mux.HandleFunc("/events_for_week", handlers.CreateHandler)
	mux.HandleFunc("/events_for_month", handlers.CreateHandler)

	return mux
}
