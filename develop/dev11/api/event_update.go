package api

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

type eventUpdateRequest struct {
	ID       uuid.UUID
	Title    string
	DateFrom time.Time
	DateTo   time.Time
}

type eventUpdateResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

// EventUpdateHandler обработчик обновления события
func (h *Handler) EventUpdateHandler(req *http.Request) APIResponse {
	if req.Method != http.MethodPost {
		return h.Error(http.StatusInternalServerError, errors.New("wrong method"))
	}
	data := &eventUpdateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	event, err := h.Storage.GetByID(data.ID.String())
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	if event == nil {
		return h.Error(http.StatusNotFound, errors.New("event not found"))
	}

	event.Title = data.Title
	event.DateFrom = data.DateFrom
	event.DateTo = data.DateTo
	if err = h.Storage.Save(event); err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, &eventUpdateResponse{
		ID:       event.ID.String(),
		Title:    event.Title,
		DateFrom: event.DateFrom.Format(time.RFC3339),
		DateTo:   event.DateTo.Format(time.RFC3339),
	})
}

// parse парсит полученные данные в структуру eventUpdateRequest
func (data *eventUpdateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	eventID := req.FormValue("id")
	if eventID == "" {
		return errors.New("event ID is required")
	}

	if data.ID, err = uuid.FromString(eventID); err != nil {
		return err
	}

	data.Title = req.FormValue("title")
	if data.Title == "" {
		return errors.New("event title is required")
	}

	if data.DateFrom, err = time.Parse(time.RFC3339, req.FormValue("date_from")); err != nil {
		return err
	}

	if data.DateTo, err = time.Parse(time.RFC3339, req.FormValue("date_to")); err != nil {
		return err
	}

	if data.DateFrom.After(data.DateTo) {
		return errors.New("date from can't be after date to")
	}

	return nil
}
