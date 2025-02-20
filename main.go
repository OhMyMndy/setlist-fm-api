package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"setlist-fm-api/client"
)

func main() {
	router := mux.NewRouter()

	// Define the API endpoint
	router.HandleFunc("/api/setlists", setlistsHandler).Methods("GET")
	router.HandleFunc("/api/setlist/{id}", setlistHandler).Methods("GET")

	// Start the server
	log.Println("Server is running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func setlistHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setlist, err := client.GetSetlist(vars["id"])

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get URL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(setlist)
	if err != nil {
		http.Error(w, "Failed encode", http.StatusInternalServerError)
		return
	}
}

// https://www.setlist.fm/setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.html
// https://www.setlist.fm/setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.htm
// venuesHandler handles the /api/venues endpoint
func setlistsHandler(w http.ResponseWriter, r *http.Request) {
	artist := r.URL.Query()["artist"][0]
	// url encode artist variable
	setlists, err := client.GetSetlists(artist)
	if err != nil {
		http.Error(w, "Failed to parse HTML", http.StatusInternalServerError)
		return
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(setlists)
	if err != nil {
		http.Error(w, "Failed encode", http.StatusInternalServerError)
		return
	}
}
