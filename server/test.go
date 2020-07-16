package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Data map[string]string

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/demo", decodeJSON)
	mux.HandleFunc("/demo2", testHard)

	err := http.ListenAndServe(":"+os.Getenv("PORT_SERVER"), mux)
	log.Fatal(err)
}

func decodeJSON(w http.ResponseWriter, r *http.Request) {
	var d Data

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server: JSON data ", d)
	fmt.Fprintf(w, "Client: Response JSON data %+v", d)
}

func testHard(w http.ResponseWriter, r *http.Request) {

	data := ""

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server: Hard Testing")
	fmt.Fprintf(w, "Client: Hard Testing: %+v", data)
}
