package api

import (
	"log"
	"net/http"
	"time"
)

// Logging логирует все запросы
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

// MainPage главная страница
func (h *Handler) MainPage(_ *http.Request) APIResponse {
	return h.JSON(http.StatusOK, map[string]string{"main": "server route"})
}

// StartServer
func (h *Handler) StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.Handle(h.MainPage))
	mux.HandleFunc("/create_event", h.Handle(h.EventCreateHandler))
	mux.HandleFunc("/update_event", h.Handle(h.EventUpdateHandler))
	mux.HandleFunc("/delete_event", h.Handle(h.EventDeleteHandler))
	mux.HandleFunc("/events_for_day", h.Handle(h.EventsForDayHandler))
	mux.HandleFunc("/events_for_week", h.Handle(h.EventsForWeekHandler))
	mux.HandleFunc("/events_for_month", h.Handle(h.EventsForMonthHandler))

	handler := Logging(mux)

	log.Fatal(http.ListenAndServe("localhost:4000", handler))

}
