// Usage:
//  go run json_reader.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// STRUCTURES

type CitibikeStationData struct {
	LastUpdated *mytime     `json:"last_updated"`
	TTL         *myduration `json:"ttl"`
	Data        *mydata     `json:"data"`
}

// We need to interpret the JSON into a time.Time structure
type mytime struct {
	t time.Time
}

// We need to interpret the JSON duration into a time.Duration structure
type myduration struct {
	d time.Duration
}

// Data in the JSON only contains an array of stations
type mydata struct {
	Stations []*station `json:"stations"`
}

// Data structure which represents a Citibike station
type station struct {
	StationId      string  `json:"station_id"`
	IsInstalled    uint    `json:"is_installed"`
	IsRenting      uint    `json:"is_renting"`
	IsReturning    uint    `json:"is_returning"`
	DocksAvailable uint    `json:"num_docks_available"`
	DocksDisabled  uint    `json:"num_docks_disabled"`
	BikesAvailable uint    `json:"num_bikes_available"`
	BikesDisabled  uint    `json:"num_bikes_disabled"`
	LastReported   *mytime `json:"last_reported"`
}

///////////////////////////////////////////////////////////////////////////////

const (
	// The URL which contains the citibike information
	CITIBIKE_URL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"
)

///////////////////////////////////////////////////////////////////////////////
// PARSERS AND STRINGIFY

// Unmarshall a unixtime into a time.Time structure
func (t *mytime) UnmarshalJSON(j []byte) error {
	if unixtime, err := strconv.ParseInt(string(j), 10, 64); err != nil {
		return err
	} else {
		t.t = time.Unix(unixtime, 0)
		return nil
	}
}

// Unmarshall a pure number into a time.Duration structure, where the
// number represents a second
func (d *myduration) UnmarshalJSON(j []byte) error {
	if seconds, err := strconv.ParseInt(string(j), 10, 64); err != nil {
		return err
	} else {
		d.d = time.Second * time.Duration(seconds)
		return nil
	}
}

func (t *mytime) String() string {
	return t.t.String()
}

func (d *myduration) String() string {
	return d.d.String()
}

func (d *mydata) String() string {
	return fmt.Sprintf("%v", d.Stations)
}

func (s *station) String() string {
	return fmt.Sprintf("station{ id=%v last_reported=%v is_installed=%v is_renting=%v is_returning=%v bikes_available=%v docks_available=%v bikes_disabled=%v docks_disabled=%v }", s.StationId, s.LastReported, s.IsInstalled, s.IsRenting, s.IsReturning, s.BikesAvailable, s.DocksAvailable, s.BikesDisabled, s.DocksDisabled)
}

///////////////////////////////////////////////////////////////////////////////

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
			// Print out the data on stdout
			fmt.Println(data)
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	os.Exit(RunMain())
}
