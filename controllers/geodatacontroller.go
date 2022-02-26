package GeoDataController

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"fmt"
	pd "github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func GetFeatureVectors(lat float64, lon float64, radius float64) []byte {

	gps := models.GPSCoordinates{Lat: lat, Lon: lon}
	db := database.GetDBConnection()
	rows, err := db.QueryFeatureVectors(gps, radius)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	cols, _ := rows.Columns()

	store := []map[string]interface{}{}

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println(err)
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		store = append(store, m)
	}

	defer rows.Close()

	df := pd.LoadMaps(store, pd.NaNValues([]string{}))
	feature_map := transformDataframeToFeatureVector(df)
	js, _ := json.Marshal(feature_map)

	return js
}

// Transform the dataframe
func transformDataframeToFeatureVector(df pd.DataFrame) map[string]int {
	//final_df := pd.DataFrame{}
	column_map := make(map[string]int)

	var column_names []string
	column_names = df.Names()

	for _, column := range column_names {
		column_without_nil := df.Select(column).Filter(
			pd.F{
				Colname:    column,
				Comparator: series.Neq,
				Comparando: "<nil>",
			})
		column_values := column_without_nil.Records()
		for _, el := range column_values {
			if el[0] != column {
				if val, ok := column_map[column+":"+el[0]]; ok {
					column_map[column+":"+el[0]] = val + 1
				} else {
					column_map[column+":"+el[0]] = 1
				}
			}
		}
	}
	return column_map
}
