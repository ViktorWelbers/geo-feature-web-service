package database

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5433
	User     = "postgres"
	Password = "mysecretpassword"
	DBName   = "postgres"
)

type DBConnection struct {
	*sql.DB
}

func (db *DBConnection) QueryFeatureVectors(gps models.GPSCoordinates, radius float64) (feature_vector *sql.Rows, error error) {
	query := fmt.Sprintf("SELECT amenity, barrier, bicycle, boundary, building, construction, highway,  water, waterway, power, motorcar , covered ,cutting , disused , embankment, historic, landuse, leisure, man_made, office, oneway, place, public_transport, railway, religion, route, service, shop, sport, surface, tourism, tunnel  FROM planet_osm_point WHERE ST_DWithin(way, ST_GeogFromText('SRID=4326;POINT(%f %f)') , %f, false);", gps.Lon, gps.Lat, radius)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return rows, nil
}

func ImportFeatures() map[string]interface{} {
	var output map[string]interface{}
	jsonFile, err := os.Open("database/feature_list.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &output)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return output
}

// Establish a Database Connection
func GetDBConnection() *DBConnection {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")

	return &DBConnection{db}
}
