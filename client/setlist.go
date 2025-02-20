package client

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Kong gateway with caching
var domain = "http://localhost:8008/"

//domain := "https://www.setlist.fm/"

func GetSetlists(artist string) (*[]Setlist, error) {
	artist = url.QueryEscape(artist)

	startDate := time.Now().AddDate(0, 1, 0).Format("2025-01-01")
	endDate := time.Now().Format("2025-01-01")
	htmlURL := fmt.Sprintf("https://www.setlist.fm/search?query=artist:(%s)+date:[%s+TO+%s]+isempty:false", artist, startDate, endDate)

	resp, err := http.Get(htmlURL)
	if err != nil {
		return nil, err
	}
	setlists, _ := parseHtmlSetlists(resp.Body)

	return setlists, nil
}
func parseDate(s *goquery.Selection) (time.Time, error) {
	month := s.Find(".dateBlock .month").First().Text()
	day := s.Find(".dateBlock .day").First().Text()
	year := s.Find(".dateBlock .year").First().Text()

	layout := "Jan 2, 2006 at 3:04pm (UTC)"
	date, err := time.Parse(layout, fmt.Sprintf("%s %s, %s at 0:00am (UTC)", month, day, year))
	if err != nil {
		fmt.Println(err)
		// TODO: return nil instead
		return time.Time{}, err
	}
	return date, nil
}

func parseSetlistPreview(s *goquery.Selection) (*Setlist, error) {
	setlist := Setlist{}
	setlist.Name, _ = s.Find("h2 a").First().Attr("href")

	// TODO: maybe check here?
	setlist.Date, _ = parseDate(s)

	setlist.Artist = s.Find(".details a").First().Text()
	setlist.ArtistURL, _ = s.Find(".details a").First().Attr("href")
	setlist.ArtistURL = domain + setlist.ArtistURL

	setlist.Venue = s.Find(".details a").Last().Text()
	setlist.VenueURL, _ = s.Find(".details a").Last().Attr("href")
	setlist.VenueURL = domain + setlist.VenueURL
	setlist.URL, _ = s.Find("a").Attr("href")
	setlist.Id = b64.StdEncoding.EncodeToString([]byte(setlist.URL))
	setlist.URL = domain + setlist.URL
	setlist.Songs = []Song{}

	return &setlist, nil
}

func parseHtmlSetlists(reader io.Reader) (*[]Setlist, error) {
	var setlists []Setlist
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return &setlists, err
	}

	// Modify the selectors below based on the actual HTML structure
	doc.Find(".setlistPreview").Each(func(i int, s *goquery.Selection) {
		setlist, _ := parseSetlistPreview(s)
		setlists = append(setlists, *setlist)
	})

	return &setlists, nil
}

func parseHtmlSetlist(reader io.Reader) (*Setlist, error) {
	var setlist Setlist

	s, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return &setlist, err
	}

	//setlist.Name = strings.TrimSpace(s.Find(".setlistHeadline h1").First().Text())
	// TODO: fix parsing
	setlist.Date, err = parseDate(s.Find(".setlistInfo"))

	setlist.Artist = s.Find(".setlistHeadline h1 a").First().Text()

	setlist.ArtistURL, _ = s.Find(".setlistHeadline h1 a").First().Attr("href")
	setlist.ArtistURL = strings.ReplaceAll(fmt.Sprintf("%v", setlist.ArtistURL), "../", "")
	setlist.ArtistURL = domain + setlist.ArtistURL

	setlist.Venue = s.Find(".setlistHeadline h1 a").Last().Text()

	setlist.VenueURL, _ = s.Find(".setlistHeadline h1 a").Last().Attr("href")
	setlist.VenueURL = strings.ReplaceAll(fmt.Sprintf("%v", setlist.VenueURL), "../", "")
	// TODO: use current URL
	setlist.VenueURL = domain + setlist.VenueURL

	// TODO: use current url
	//setlist.URL = domain + setlist.URL
	//setlist.URL, _ = s.Find("a").Attr("href")
	//setlist.Id = b64.StdEncoding.EncodeToString([]byte(setlist.URL))
	setlist.Songs = []Song{}
	// Modify the selectors below based on the actual HTML structure
	s.Find(".songsList .song").Each(func(i int, s *goquery.Selection) {

		song := Song{}
		song.Name = s.Find(".songLabel").First().Text()
		song.URL, _ = s.Find(".songLabel").First().Attr("href")
		song.URL = strings.ReplaceAll(fmt.Sprintf("%v", song.URL), "../", "")
		song.URL = domain + song.URL

		song.Id = b64.StdEncoding.EncodeToString([]byte(song.URL))
		song.Artist = setlist.Artist
		// this needs to be done better
		if len(setlist.Artist) == 0 {
			setlist.Artist = song.Artist
		}

		setlist.Songs = append(setlist.Songs, song)
	})

	return &setlist, nil
}

func GetSetlist(id string) (*Setlist, error) {
	setlistUrl, _ := url.QueryUnescape(id)
	setlistUrlDecoded, _ := b64.StdEncoding.DecodeString(setlistUrl)
	setlistUrl = string(setlistUrlDecoded)

	resp, err := http.Get(domain + setlistUrl)

	if err != nil {
		return nil, err
	}
	setlist, err := parseHtmlSetlist(resp.Body)

	return setlist, nil
}
