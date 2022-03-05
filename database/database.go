package database

import (
	"backend/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

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

func (db *DBConnection) CheckAvailability() {
	if err := db.Ping(); err != nil {
		_ = db.Close()
		fmt.Println(err)
		panic(err)
	}
	fmt.Printf("Database Running on Port %d \n", Port)
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

type featureEnumerator struct {
	Features []string
}

var AllFeatures featureEnumerator

func (f *featureEnumerator) ImportFeaturesFromJSON() {
	f.Features = getAllColumns()
}

func getAllColumns() []string {
	var output map[string]interface{}
	jsonFile, err := os.Open("database/feature_list.json")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &output)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	all_columns := []string{}
	for _, value := range output {
		iter := reflect.ValueOf(value).MapRange()
		for iter.Next() {
			innerValueInterface := iter.Value().Interface().([]interface{})
			columns := make([]string, len(innerValueInterface))
			for i, v := range innerValueInterface {
				columns[i] = v.(string)
			}
			all_columns = append(all_columns, columns...)
		}
	}
	return all_columns
}

// Establish a Database Connection
func NewDBConnection() *DBConnection {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBName)

	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		_ = conn.Close()
		fmt.Println(err)
		panic(err)
	}

	db := &DBConnection{conn}
	db.CheckAvailability()
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")

	return db
}
