package consume

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BusLine struct {
	LineId    int64     `json:"id"`
	Name      string    `json:"name"`
	RouteName string    `json:"routename"`
	Vehicles  []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	VehicleId        int64  `json:"vehicle_id"`
	RegistrationCode string `json:"registration_code"`
	TimeStamp        string `json:"ts"`
	Stats            Stats  `json:"stats"`
}

type Stats struct {
	CurrentSpeed string `json:"speed"`
	AverageSpeed string `json:"avg_speed"`
	Bearing      int64  `json:"bearing"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
}

type Projection struct {
	EdgeId          int64  `json:"edge_id"`
	EdgeDistance    string `json:"edge_distance"`
	EdgeProjection  string `json:"edgeprojection"`
	EdgeStartNodeId string `json:"edge_start_node_id"`
	EdgeStopNodeId  string `json:"edge_stop_node_id"`
}

var BusLines = make(map[string]BusLine)

func ConsumeBusLine(busLineId string) BusLine {
	// Desc: Consume and filter external API endpoint with Go structs
	// Return: BusLine struct - filtered and ready to be encoded as response

	// log.Println("--- (func start) ConsumeBusLine ---")
	apiEndpoint := fmt.Sprintf("https://dummy.uwave.sg/busline/%s", busLineId)
	response, err := http.Get(apiEndpoint)

	if err != nil {
		log.Println(err.Error())
	}
	defer response.Body.Close()

	var busLineObject BusLine
	json.NewDecoder(response.Body).Decode(&busLineObject)
	BusLines[busLineId] = busLineObject
	
	// log.Printf("Return struct: %+v\n", busLineObject)
	// log.Println("--- (func end) ConsumeBusLine ---")
	return busLineObject
}

func BusLineAvail(busLineId string) bool{
	// Desc: Check if BusLine is available
	// Return: Boolean

	// log.Println("--- (func start) BusLineAvail ---")
	_, avail := BusLines[busLineId]
	// log.Println("--- (func end) BusLineAvail ---")
	return avail
}