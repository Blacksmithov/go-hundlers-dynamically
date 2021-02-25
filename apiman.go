package main

import (
	"fmt"
	"log"
	"net/http"

	t_apiman "github.com/Blacksmithov/go-hundlers-dynamically/query"
)

func main() {

	fmt.Println("http://localhost:8080")

	rows := t_apiman.ProjectsList()

	for rows.Next() {
		var project string
		var endpoint string
		var query string

		err := rows.Scan(&project, &endpoint, &query)
		if err != nil {
			fmt.Print("ошибка")
		}
		fmt.Printf("http://localhost:8080/%s/%s\n", project, endpoint)
		http.HandleFunc("/"+project+"/"+endpoint, func(w http.ResponseWriter, r *http.Request) {
			t_apiman.QueryToJson(w, query)
		})
	}
	t_apiman.ProjectsList().Close()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
