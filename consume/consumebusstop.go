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
	// Desc: Consume and filter external API endpoint with Go structs
	// Return: BusStop struct - filtered and ready to be encoded as response

	// log.Println("--- (func start) ConsumeBusStop ---")
	apiEndpoint := fmt.Sprintf("https://dummy.uwave.sg/busstop/%s", busStopId)
	response, err := http.Get(apiEndpoint)

	if err != nil {
		log.Println(err.Error())
	}
	defer response.Body.Close()

	var busStopObject BusStop
	json.NewDecoder(response.Body).Decode(&busStopObject)
	BusStops[busStopId] = busStopObject

	// log.Printf("Return struct: %+v\n", busStopObject)
	// log.Println("--- (func end) ConsumeBusStop ---")
	return busStopObject
}

func BusStopAvail(busStopId string) bool{
	// Desc: Check if BusStop is available
	// Return: Boolean

	// log.Println("--- (func start) BusStopAvail ---")
	_, avail := BusStops[busStopId]
	// log.Println("--- (func end) BusStopAvail ---")
	return avail
}