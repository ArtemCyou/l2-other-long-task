package app

import (
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	Id int64
	Data time.Time
	Events string
}

// ParseEvent парсим параметры и возвращаем заполненную структуру Event
func ParseEvent(params map[string][]string) (*Event, int, error) {

	// парсим id
	if _, ok := params["id"]; !ok {
		status := http.StatusBadRequest
		err := errors.New("пропущен id")
		return nil, status, err
	}
	id, err := strconv.ParseInt(params["id"][0], 10, 64)
	if err != nil {
		status := http.StatusBadRequest
		err := errors.New("неверный id")
		return nil, status, err
	}

	// парсим date
	if _, ok := params["data"]; !ok {
		status := http.StatusBadRequest
		err := errors.New("пропущена дата")
		return nil, status, err
	}
	date, err := time.Parse("2006-01-02", params["date"][0])
	if err != nil {
		status := http.StatusBadRequest
		err := errors.New("неверная дата")
		return nil, status, err
	}

	// парсим description
	descriptions, ok := params["description"]
	var description string
	if ok {
		description = descriptions[0]
	}

	return &Event{Id: id, Data: date, Events: description}, http.StatusOK, nil
}

