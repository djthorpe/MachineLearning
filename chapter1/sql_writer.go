package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	// The table name
	SQL_TABLENAME = "station"
)

var (
	FlagDatabasePath = flag.String("db", "", "Path to the database")

	SQL_COLUMNS = []string{
		"id integer not null primary key",
		"last_reported string not null",
		"is_installed bool",
		"is_renting bool",
		"docks_available integer",
		"docks_disabled integer",
		"bikes_available integer",
		"bikes_disabled integer",
	}
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

func TableExists(db *sql.DB, name string) (bool, error) {
	if resultset, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'"); err != nil {
		return false, err
	} else {
		defer resultset.Close()
		for resultset.Next() {
			var name2 string
			if err := resultset.Scan(&name2); err != nil {
				return false, err
			}
			if name == name2 {
				return true, nil
			}
		}
		if err := resultset.Err(); err != nil {
			return false, err
		}
		return false, nil
	}
}

func CreateTable(db *sql.DB, name string, columns []string) error {
	if _, err := db.Exec(fmt.Sprintf("CREATE TABLE %v (%v)", name, strings.Join(columns, ","))); err != nil {
		return err
	}
	return nil
}

func WriteData(db *sql.DB, data CitibikeStationData) error {
	// Check for SQL table, create as necessary
	if exists, err := TableExists(db, SQL_TABLENAME); err != nil {
		return err
	} else if exists == false {
		if err := CreateTable(db, SQL_TABLENAME, SQL_COLUMNS); err != nil {
			return err
		}
	}

	// Start transaction
	if tx, err := db.Begin(); err != nil {
		return err
	} else {
		// Prepare statement
		if stmt, err := tx.Prepare(fmt.Sprintf("INSERT OR REPLACE INTO %v values(?,?,?,?,?,?,?,?)", SQL_TABLENAME)); err != nil {
			return err
		} else {
			defer stmt.Close()

			// Write out station data
			for _, station := range data.Data.Stations {
				if _, err := stmt.Exec(station.StationId, station.LastReported.String(), station.IsInstalled, station.IsRenting, station.DocksAvailable, station.DocksDisabled, station.BikesAvailable, station.BikesDisabled); err != nil {
					return err
				}
			}
		}

		// Commit transaction
		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func RunMain() int {
	// Check database flag
	if *FlagDatabasePath == "" {
		log.Println("Expected -db flag")
		return -1
	}
	// Open the database
	db, err := sql.Open("sqlite3", *FlagDatabasePath)
	if err != nil {
		log.Println(err)
		return -1
	}
	defer db.Close()

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
			// Write out the station data
			if err := WriteData(db, data); err != nil {
				log.Println(err)
				return -1
			}
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////////

func main() {
	flag.Parse()
	os.Exit(RunMain())
}
