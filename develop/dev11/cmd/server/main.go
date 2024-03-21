package main

import (
	"log"
	"net/http"

	"myapp/dev11/internal/handler"
)

func main() {
	// Инициализация HTTP сервера
	http.HandleFunc("/create_event", handler.CreateEventHandler)
	http.HandleFunc("/update_event", handler.UpdateEventHandler)
	http.HandleFunc("/delete_event", handler.DeleteEventHandler)
	http.HandleFunc("/events_for_day", handler.EventsForDayHandler)
	http.HandleFunc("/events_for_week", handler.EventsForWeekHandler)
	http.HandleFunc("/events_for_month", handler.EventsForMonthHandler)

	// Запуск HTTP сервера
	log.Fatal(http.ListenAndServe(":8080", nil))
}
