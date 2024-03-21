package handler

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"net/http"

	"myapp/dev11/internal/service"
)

// CreateEventHandler обрабатывает запрос на создание события
func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	userID, date, err := parseCreateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для создания события
	result, err := service.CreateEvent(userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, result)
}

// UpdateEventHandler обрабатывает запрос на обновление события
func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	eventID, userID, date, err := parseUpdateEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для обновления события
	result, err := service.UpdateEvent(eventID, userID, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, result)
}

// DeleteEventHandler обрабатывает запрос на удаление события
func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	eventID, err := parseDeleteEventParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для удаления события
	err = service.DeleteEvent(eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, "Event deleted successfully")
}

// EventsForDayHandler обрабатывает запрос на получение событий за день
func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	date, err := parseEventsForDayParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для получения событий за день
	result, err := service.GetEventsForDay(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, result)
}

// EventsForWeekHandler обрабатывает запрос на получение событий за неделю
func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	date, err := parseEventsForWeekParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для получения событий за неделю
	result, err := service.GetEventsForWeek(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, result)
}

// EventsForMonthHandler обрабатывает запрос на получение событий за месяц
func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	date, err := parseEventsForMonthParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызов сервиса для получения событий за месяц
	result, err := service.GetEventsForMonth(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный результат
	sendResponse(w, result)
}

// parseCreateEventParams парсит и валидирует параметры запроса для метода /create_event
func parseCreateEventParams(r *http.Request) (int, time.Time, error) {
	// Получаем параметры из запроса
	userIDStr := r.FormValue("user_id")
	dateStr := r.FormValue("date")

	// Проверяем, что user_id и date присутствуют в запросе
	if userIDStr == "" || dateStr == "" {
		return 0, time.Time{}, errors.New("user_id and date are required")
	}

	// Парсим userID в целое число
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, time.Time{}, errors.New("invalid user ID")
	}

	// Парсим дату из строки
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, time.Time{}, errors.New("invalid date format")
	}

	return userID, date, nil
}

// parseUpdateEventParams парсит и валидирует параметры запроса для метода /update_event
func parseUpdateEventParams(r *http.Request) (int, int, time.Time, error) {
	// Реализация парсинга и валидации параметров
	eventID := 0
	userID := 0
	date := time.Now()

	// Пример парсинга параметров из запроса:
	eventIDParam := r.FormValue("event_id")
	userIDParam := r.FormValue("user_id")
	dateParam := r.FormValue("date")

	// Пример валидации и преобразования параметров
	var err error
	if eventIDParam != "" {
		eventID, err = strconv.Atoi(eventIDParam)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid event_id")
		}
	}

	if userIDParam != "" {
		userID, err = strconv.Atoi(userIDParam)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid user_id")
		}
	}

	if dateParam != "" {
		date, err = time.Parse("2006-01-02", dateParam)
		if err != nil {
			return 0, 0, time.Time{}, errors.New("invalid date")
		}
	}

	return eventID, userID, date, nil
}

// parseDeleteEventParams парсит и валидирует параметры запроса для метода /delete_event
func parseDeleteEventParams(r *http.Request) (int, error) {
	// Получаем параметры из запроса
	eventIDStr := r.FormValue("event_id")

	// Парсим eventID в целое число
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		return 0, errors.New("invalid event ID")
	}

	return eventID, nil
}

// parseEventsForDayParams парсит и валидирует параметры запроса для метода /events_for_day
func parseEventsForDayParams(r *http.Request) (time.Time, error) {
	// Получаем параметры из запроса
	dateStr := r.FormValue("date")

	// Парсим дату из строки
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}

// parseEventsForWeekParams парсит и валидирует параметры запроса для метода /events_for_week
func parseEventsForWeekParams(r *http.Request) (time.Time, error) {
	// Получаем параметры из запроса
	dateStr := r.FormValue("date")

	// Парсим дату из строки
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}

// parseEventsForMonthParams парсит и валидирует параметры запроса для метода /events_for_month
func parseEventsForMonthParams(r *http.Request) (time.Time, error) {
	// Получаем параметры из запроса
	dateStr := r.FormValue("date")

	// Парсим дату из строки
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}

	return date, nil
}

func sendResponse(w http.ResponseWriter, result interface{}) {
	response, err := json.Marshal(map[string]interface{}{"result": result})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
