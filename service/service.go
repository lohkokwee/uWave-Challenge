package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	consume "github.com/lohkokwee/uwave_challenge/consume"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseWithStopData struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    consume.BusStop `json:"data"`
}

type ResponseWithLineData struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    consume.BusLine `json:"data"`
}

var TimeStamp time.Time = time.Now().Round(time.Second)

func RetrieveStopDetails(w http.ResponseWriter, r *http.Request) {
	// log.Println("--- (func) RetrieveStopDetails ---")

	vars := mux.Vars(r)
	busStopId := vars["busStopId"]

	// Error handling
	var validStops = []string{
		"378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230", "378229", "378228", "378227",
		"382995", "378224", "378226", "383010", "383009", "383006", "383004", "378234", "383003", "378222", "383048",
		"378203", "382999", "378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207",
	}

	w.Header().Set("Content-Type", "application/json")

	for index, validStop := range validStops {
		if busStopId == validStop {
			break
		}

		if index == len(validStops)-1 {
			response := Response{"404", "Invalid bus stop id."}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// API Limit handling (Temp Storage)
	// We only call the API if there is a change in time(seconds)/data not avail
	currentTime := time.Now().Round(time.Second)
	var busStopObject consume.BusStop

	if TimeStamp == currentTime && consume.BusStopAvail(busStopId) {
		busStopObject = consume.BusStops[busStopId]
	} else {
		busStopObject = consume.ConsumeBusStop(busStopId)
		TimeStamp = currentTime
	}

	successResponse := ResponseWithStopData{"200", "OK.", busStopObject}
	json.NewEncoder(w).Encode(successResponse)

	// log.Println("--- (func end) RetrieveStopDetails ---")
}

func RetrieveLineDetails(w http.ResponseWriter, r *http.Request) {
	// log.Println("--- (func) RetrieveLineDetails ---")

	vars := mux.Vars(r)
	busLineId := vars["busLineId"]

	// Error handling
	var validLines = []string{
		"44478", "44479", "44480", "44481",
	}

	w.Header().Set("Content-Type", "application/json")

	for index, validLine := range validLines {
		if string(busLineId) == validLine {
			break
		}

		if index == len(validLines)-1 {
			response := Response{"404", "Invalid bus line id."}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	// API Limit handling (Temp Storage)
	// We only call the API if there is a change in time(seconds)/data not avail
	currentTime := time.Now().Round(time.Second)
	var busLineObject consume.BusLine

	if TimeStamp == currentTime && consume.BusLineAvail(busLineId) {
		busLineObject = consume.BusLines[busLineId]
	} else {
		busLineObject = consume.ConsumeBusLine(busLineId)
		TimeStamp = currentTime
	}

	successResponse := ResponseWithLineData{"200", "OK.", busLineObject}
	json.NewEncoder(w).Encode(successResponse)

	// log.Println("--- (func end) RetrieveLineDetails ---")
}
