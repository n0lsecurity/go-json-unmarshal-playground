package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestBody struct {
	ID int `json:"id"`
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello from here!"))
}

func printMyId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The method is not supported"))
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte("Media is not JSON type"))
		return
	}

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "There is some error in the json body format %v", err)
		return
	}

	defer r.Body.Close()

	var requestBody RequestBody

	err = json.Unmarshal(bodyData, &requestBody)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error on unmarshaling JSON data")
	}

	id := requestBody.ID

	fmt.Fprintf(w, "The JSON ID is : %v", id)
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific Snippet!"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("The method is not allowed!"))
		return
	}

	w.Write([]byte("The snippet created successfully! hahaha"))
}

func main() {

	//mux is router handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", printMyId)

	log.Print("Server is started on 4000 port!")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
