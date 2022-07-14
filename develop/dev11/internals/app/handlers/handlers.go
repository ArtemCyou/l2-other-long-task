package handlers

import (
	"dev11/internals/app"
	"dev11/internals/app/db"
	"errors"
	"fmt"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		//забираем параметры в PostForm
		if err := r.ParseForm(); err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		err = db.InsertData(event)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}
		WrapOK(w)

	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if r.Method == http.MethodPost {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		// считываем строку
		err := r.ParseForm()
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		// парсим строку параметров
		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// обновляем событие в БД
		err = db.UpdateData(event)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// отправляем ответ об успешном завершении операции
		WrapOK(w)
	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if r.Method == http.MethodPost {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		// считываем строку
		err := r.ParseForm()
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		// парсим строку параметров
		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// удаляем событие из БД
		err = db.DeleteData(event)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// отправляем ответ об успешном завершении операции
		WrapOK(w)

	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}
}

func dayEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if r.Method == http.MethodGet {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		// считываем строку
		err := r.ParseForm()
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		// парсим строку параметров
		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// получаем список событий за один день с указанной даты
		data, err := db.GetData(event, 1)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			fmt.Fprint(w, data)
			return
		}

		// отправляем ответ
		WrapOK(w)
	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}
}

func weekEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if r.Method == http.MethodGet {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		// считываем строку
		err := r.ParseForm()
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		// парсим строку параметров
		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// получаем список событий за один день с указанной даты
		data, err := db.GetData(event, 7)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			fmt.Fprint(w, data)
			return
		}

		// отправляем ответ
		WrapOK(w)
	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}
}

func monthEventsHandler(w http.ResponseWriter, r *http.Request) {
	// проверяем соответствие методу
	if r.Method == http.MethodGet {

		// сверяем  Media Type
		headerContentType := r.Header.Get("Content-Type")
		if headerContentType != "application/x-www-form-urlencoded" {
			err := errors.New("медиа не поддерживается")
			WrapErrorWithStatus(w, err, r.Method, http.StatusUnsupportedMediaType)
			return
		}

		// считываем строку
		err := r.ParseForm()
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, http.StatusNoContent)
			return
		}

		// парсим строку параметров
		event, status, err := app.ParseEvent(r.PostForm)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			return
		}

		// получаем список событий за один день с указанной даты
		data, err := db.GetData(event, 30)
		if err != nil {
			WrapErrorWithStatus(w, err, r.Method, status)
			fmt.Fprint(w, data)
			return
		}

		// отправляем ответ
		WrapOK(w)
	} else {
		err := errors.New("invalid request method")
		WrapErrorWithStatus(w, err, r.Method, http.StatusMethodNotAllowed)
	}
}
