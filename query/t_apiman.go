package t_apiman

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	appconfig "github.com/Blacksmithov/go-hundlers-dynamically/cfg"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	db         *sql.DB
	datasource string = fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s",
		appconfig.GetValue("database.host"),
		appconfig.GetValue("database.db_name"),
		appconfig.GetValue("database.login"),
		appconfig.GetValue("database.password"),
		appconfig.GetValue("database.port"))
)

func ProjectsList() *sql.Rows {
	var err error

	db, err := sql.Open("pgx", datasource)
	PanicOnErr(err)
	defer db.Close()

	err = db.Ping()
	PanicOnErr(err)

	rows, err := db.Query("select project, endpoint, query from endpoints")
	PanicOnErr(err)
	return rows
}

func QueryToJson(w http.ResponseWriter, query string, args ...interface{}) {
	var objects []map[string]interface{}

	db, err := sql.Open("pgx", datasource)
	PanicOnErr(err)
	defer db.Close()

	rows, err := db.Query(query, args...)
	PanicOnErr(err)
	for rows.Next() {
		columns, err := rows.ColumnTypes()
		PanicOnErr(err)

		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			var v interface{}

			switch column.DatabaseTypeName() {
			case "text":
				v = new(string)
			default:
				v = new(interface{})
			}

			object[column.Name()] = v
			values[i] = v
		}

		err = rows.Scan(values...)
		PanicOnErr(err)
		objects = append(objects, object)
	}

	json.NewEncoder(w).Encode(objects)
}

func PanicOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
