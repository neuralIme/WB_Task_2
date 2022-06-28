package domain

import (
	"sync"
	"time"
)

// StorageInterface интерфейс хранилища событий
type StorageInterface interface {
	GetAll() ([]*Event, error)
	GetByID(id string) (*Event, error)
	GetByPeriod(from, to time.Time) ([]*Event, error)
	Save(event *Event) error
	Remove(event *Event) error
}

// Storage структура для хранения событий
type Storage struct {
	sync.Mutex
	events map[string]*Event
}

// NewStorage создает новое хранилище
func NewStorage() *Storage {
	eventStorage := make(map[string]*Event)
	return &Storage{events: eventStorage}
}

// GetAll возвращает все события
func (storage *Storage) GetAll() ([]*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	result := make([]*Event, 0, len(storage.events))
	for _, event := range storage.events {
		result = append(result, event)
	}

	return result, nil
}

// GetByID возвращает событие по ID
func (storage *Storage) GetByID(id string) (*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if event, ok := storage.events[id]; ok {
		return event, nil
	}

	return nil, nil
}

// GetByPeriod возвращает список событий за период
func (storage *Storage) GetByPeriod(from, to time.Time) ([]*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	result := make([]*Event, 0)

	for _, event := range storage.events {
		if event.DateFrom.Before(to) && event.DateTo.After(from) {
			result = append(result, event)
		}
	}

	return result, nil
}

//Save создает или обновляет событие в хранилище
func (storage *Storage) Save(event *Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.events[event.ID.String()] = event

	return nil
}

//Remove удаляет событие из хранилища
func (storage *Storage) Remove(event *Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.events, event.ID.String())

	return nil
}
