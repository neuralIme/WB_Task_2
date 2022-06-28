package api

import (
	"errors"
	"net/http"
	"time"
)

type eventsForDateRequest struct {
	StartDay time.Time
}

// parse извлекает дату события
func (data *eventsForDateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	if data.StartDay, err = time.Parse("2006-01-02", req.FormValue("start_day")); err != nil {
		return err
	}

	return nil
}

// EventsForDayHandler  обрабатывает запрос на события дня
func (h *Handler) EventsForDayHandler(req *http.Request) APIResponse {
	if req.Method != http.MethodGet {
		return h.Error(http.StatusInternalServerError, errors.New("wrong method"))
	}
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}

// EventsForWeekHandler обрабатывает запрос на события недели
func (h *Handler) EventsForWeekHandler(req *http.Request) APIResponse {
	if req.Method != http.MethodGet {
		return h.Error(http.StatusInternalServerError, errors.New("wrong method"))
	}
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*7))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}

// EventsForMonthHandler обрабатывает запрос на события месяца
func (h *Handler) EventsForMonthHandler(req *http.Request) APIResponse {
	if req.Method != http.MethodGet {
		return h.Error(http.StatusInternalServerError, errors.New("wrong method"))
	}
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}
	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*30))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}
