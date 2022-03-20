package consume

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BusStop struct {
	StopId    int64       `json:"id"`
	StopName  string      `json:"name"`
	Geometry  []Geometry  `json:"geometry"`
	Forecasts []Forecast `json:"forecast"`
}

type Geometry struct {
	Lat  string `json:"lat"`
	Long string `json:"lon"`
}

type Forecast struct {
	ForcastSeconds float64 `json:"forecast_seconds"`
	VehicleId int64 `json:"vehicle_id"`
}

var BusStops = make(map[string]BusStop)

func ConsumeBusStop(busStopId string) BusStop {
	// log.Println("--- (func start) ConsumeBusStop ---")

	apiEndpoint := fmt.Sprintf("https://dummy.uwave.sg/busstop/%s", busStopId)
	// log.Println("API Endpoint:", apiEndpoint)

	response, err := http.Get(apiEndpoint)

	if err != nil {
		// Handle error
		log.Println(err.Error())
	}
	defer response.Body.Close()

	// 1. Decode JSON data
	// 2. Filter required information through struct
	// 3. Stores struct in map
	var busStopObject BusStop
	json.NewDecoder(response.Body).Decode(&busStopObject)
	BusStops[busStopId] = busStopObject

	// log.Printf("Return struct: %+v\n", busStopObject)
	// log.Println("--- (func end) ConsumeBusStop ---")
	return busStopObject
}

func BusStopAvail(busStopId string) bool{
	// log.Println("--- (func start) BusStopAvail ---")

	_, avail := BusStops[busStopId]
	
	// log.Println("--- (func end) BusStopAvail ---")
	return avail
}