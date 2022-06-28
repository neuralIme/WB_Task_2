package api

import (
	"calendar/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	Storage domain.StorageInterface
}

// APIResponse функция-ответ обработчиков
type APIResponse func(resp http.ResponseWriter)

// Handle замыкает полученные данные с ответом
func (h *Handler) Handle(fn func(req *http.Request) APIResponse) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Server", "GoCalendar API")
		fn(req)(resp)
	}
}

// JSON формирует ответ со статус кодом
func (h *Handler) JSON(code int, data interface{}) APIResponse {
	return h.sendJSON(code, map[string]interface{}{"result": data})
}

func (h *Handler) Error(code int, err error) APIResponse {
	return h.sendJSON(code, map[string]string{"error": err.Error()})
}

//
func (h *Handler) sendJSON(code int, data interface{}) APIResponse {
	var encodedData []byte
	var err error

	if data != nil {
		encodedData, err = json.Marshal(data)
		if err != nil {
			return h.Error(http.StatusInternalServerError, err)
		}
	}

	return func(resp http.ResponseWriter) {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.Header().Set("Content-Length", strconv.Itoa(len(encodedData)))
		resp.WriteHeader(code)
		resp.Write(encodedData)
	}
}
