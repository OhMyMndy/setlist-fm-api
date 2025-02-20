package client

import "time"

type Setlist struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	Date      time.Time `json:"date"`
	URL       string    `json:"url"`
	Artist    string    `json:"artist"`
	ArtistURL string    `json:"artist_url"`
	Venue     string    `json:"venue"`
	VenueURL  string    `json:"venue_url"`
	Songs     []Song    `json:"songs"`
}

type Song struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	URL    string `json:"url"`
}

type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	MbId string `json:"mbid"`
}
