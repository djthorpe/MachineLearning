package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

///////////////////////////////////////////////////////////////////////////////

type CitibikeStationData struct {
	LastUpdated time.Time     `json:"last_updated"`
	TTL         time.Duration `json:"ttl"`
}

///////////////////////////////////////////////////////////////////////////////

const (
	CITIBIKE_URL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"
)

///////////////////////////////////////////////////////////////////////////////

/*
// station is used to unmarshal each of the station documents in
stationData.
type station struct {
ID string `json:"station_id"`
NumBikesAvailable int `json:"num_bikes_available"`
NumBikesDisabled int `json:"num_bike_disabled"`
NumDocksAvailable int `json:"num_docks_available"`
NumDocksDisabled int `json:"num_docks_disabled"`
IsInstalled int `json:"is_installed"`
IsRenting int `json:"is_renting"`
IsReturning int `json:"is_returning"`
LastReported int `json:"last_reported"`
HasAvailableKeys bool `json:"eightd_has_available_keys"`
}
*/

func RunMain() int {
	// Get the JSON response from the URL
	if response, err := http.Get(CITIBIKE_URL); err != nil {
		log.Println(err)
		return -1
	} else {
		defer response.Body.Close()
		if body, err := ioutil.ReadAll(response.Body); err != nil {
			log.Println(err)
			return -1
		} else {
			var data CitibikeStationData
			// Unmarshal the JSON data into the variable.
			if err := json.Unmarshal(body, &data); err != nil {
				log.Println(err)
				return -1
			}
			fmt.Println(data)
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	os.Exit(RunMain())
}
