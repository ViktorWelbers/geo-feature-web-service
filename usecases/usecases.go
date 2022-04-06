package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"geo-api-backend/entities"
	"geo-api-backend/repositories"
	pd "github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strconv"
)

func GetFeatureVectors(gps entities.GPSCoordinates, db *repositories.DBConnection) map[string]int {

	rows, err := db.QueryFeatureVectors(gps.Lat, gps.Lon, gps.Radius)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	cols, _ := rows.Columns()

	var store []map[string]interface{}

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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	df := pd.LoadMaps(store, pd.NaNValues([]string{}))
	featureMap := transformDfToFeatureMap(&df)

	return featureMap
}

// Transform the dataframe
func transformDfToFeatureMap(df *pd.DataFrame) map[string]int {

	columnMap := make(map[string]int)

	var columnNames []string
	columnNames = df.Names()

	for _, column := range columnNames {
		columnWithoutNil := df.Select(column).Filter(
			pd.F{
				Colname:    column,
				Comparator: series.Neq,
				Comparando: "<nil>",
			})
		columnValues := columnWithoutNil.Records()
		for _, el := range columnValues {
			if el[0] != column {
				if val, ok := columnMap[column+":"+el[0]]; ok {
					columnMap[column+":"+el[0]] = val + 1
				} else {
					columnMap[column+":"+el[0]] = 1
				}
			}
		}
	}
	for _, v := range repositories.AllFeatures.Features {
		if _, ok := columnMap[v]; !ok {
			columnMap[v] = 0
		}
	}

	return columnMap
}

func ParseGeoRequest(latString string, lonString string, radiusString string) (entities.GPSCoordinates, error) {
	lat, err1 := strconv.ParseFloat(latString, 64)
	lon, err2 := strconv.ParseFloat(lonString, 64)
	radius, err3 := strconv.ParseFloat(radiusString, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		return entities.GPSCoordinates{}, errors.New("error when trying to parse values. please provide numerical")
	}
	return entities.GPSCoordinates{Lat: lat, Lon: lon, Radius: radius}, nil
}
