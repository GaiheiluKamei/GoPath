package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type timeZoneConvertion struct {
	TimeZone string
	CurrentTime string
}

var conversionMap = map[string]string{
	"ASR": "-3h", // North America Atlantic Standard Time
	"EST": "-5h", // North America Eastern Standard Time
	"BST": "+1h", // British Summer Time
	"IST": "+5h30m", // India Standard Time
	"HKT": "+8h", // Hang Kong Time
	"ART": "-3h", // Argentina Time
	"RAE": "MMMMMM", // WILL RAISE AN ERROR
}

func main() {
	http.HandleFunc("/convert", loggingMiddleware(handler))
	http.HandleFunc("/", loggingMiddleware(notFoundHandler))
	log.Printf("%s - Starting server on port: 8080.", time.Now().Format("2009-01-02 15:04:05"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func loggingMiddleware(handler handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", time.Now().Format("2018-11-25 14:32:58"), r.Method, r.URL.String())
		handler(w, r)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404: The requested URL does not exist.")
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")
	if timeZone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error 400: tz parameter is required.")
		return
	}

	timeDifference, ok := conversionMap[timeZone]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 %s not found", timeZone)
		return
	}

	currentTimeConverted, err := getCurrentTimeByTimeDifference(timeDifference)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server error.")
		return
	}

	tzc := new(timeZoneConvertion)
	tzc.TimeZone = timeZone
	tzc.CurrentTime = currentTimeConverted

	jsonResponse, err := json.Marshal(tzc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: Server error.")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonResponse))
	
}

func getCurrentTimeByTimeDifference(timeDifference string) (string, error) {
	now := time.Now().UTC()

	difference, err := time.ParseDuration(timeDifference)
	if err != nil {
		return "", err
	}

	return now.Add(difference).Format("15:04:05"), nil
}