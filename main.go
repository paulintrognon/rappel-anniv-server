package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Event struct {
	Title string    `json:"title"`
	Desc  string    `json:"desc"`
	Date  time.Time `json:"date"`
}

type Events []Event

func allEvents(w http.ResponseWriter, r *http.Request) {
	events := Events{
		Event{Title: "Maman", Desc: "Anniversaire de maman", Date: time.Now()},
	}

	fmt.Println("Hit: All Events")
	json.NewEncoder(w).Encode(events)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rappel-anniversaire.fr API")
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/events/all", allEvents)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func main() {
	handleRequests()
}
