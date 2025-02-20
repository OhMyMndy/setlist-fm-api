package client

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

// create test for GetSetlist
func TestGetSetlist(t *testing.T) {
	file, err := os.Open("setlist.html")
	if err != nil {
		t.Errorf("Error loading file: %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	setlists, err := parseHtmlSetlist(file)
	if err != nil {
		t.Errorf("GetSetlist returned an error: %v", err)
		return
	}

	expectedSetlist := Setlist{
		Artist:    "Shadow of Intent",
		ArtistURL: "http://localhost:8008/setlists/shadow-of-intent-2bd8d4fa.html",
		//URL:       "setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.html",
		//Id:        "c2V0bGlzdC9zaGFkb3ctb2YtaW50ZW50LzIwMjUvYW1hZ2VyLWJpby1jb3BlbmhhZ2VuLWRlbm1hcmstNjM1MzMyODMuaHRtbA==",
		Venue:    "Amager Bio, Copenhagen, Denmark",
		VenueURL: "http://localhost:8008/venue/amager-bio-copenhagen-denmark-5bd63b24.html",
		Date:     time.Date(2025, 2, 16, 0, 0, 0, 0, time.UTC),
		//URL:       "http://localhost:8008/setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.html",
		Songs: []Song{
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NzNmNzRlODk=",
				Name:   "We Descend...",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=73f74e89",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NmJmNzRlOGU=",
				Name:   "The Horror Within",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=6bf74e8e",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9MjNlNGRjOGY=",
				Name:   "Intensified Genocide",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=23e4dc8f",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9M2VmZjFjMw==",
				Name:   "The Migrant",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=3eff1c3",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9Mzk0NTUwZg==",
				Name:   "Flying the Black Flag",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=394550f",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9N2JmNzRlOGM=",
				Name:   "The Heretic Prevails",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=7bf74e8c",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9M2ZkYTFhNw==",
				Name:   "Melancholy",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=3fda1a7",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NGJlNWQ3ZDI=",
				Name:   "Blood in the Sands of Time",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=4be5d7d2",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9YmZkYTFhNg==",
				Name:   "Barren and Breathless Macrocosm",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=bfda1a6",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NWI5M2YzYTg=",
				Name:   "Drum Solo",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=5b93f3a8",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NjNmNzRlODM=",
				Name:   "The Battle of the Maginot Sphere",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=63f74e83",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9MTNmZGExYTU=",
				Name:   "Malediction",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=13fda1a5",
			},
			{
				Artist: "Shadow of Intent",
				Id:     "aHR0cDovL2xvY2FsaG9zdDo4MDA4L3N0YXRzL3NvbmdzL3NoYWRvdy1vZi1pbnRlbnQtMmJkOGQ0ZmEuaHRtbD9zb25naWQ9NjM5MzgyZjM=",
				Name:   "The Tartarus Impalement",
				URL:    "http://localhost:8008/stats/songs/shadow-of-intent-2bd8d4fa.html?songid=639382f3",
			},
		},
	}

	assert.Equal(t, expectedSetlist, *setlists, "Setlists are not equal")

}

func TestGetSetlists(t *testing.T) {
	file, err := os.Open("setlists.html")
	if err != nil {
		t.Errorf("Error loading file: %v", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	setlists, err := parseHtmlSetlists(file)
	if err != nil {
		t.Errorf("GetSetlist returned an error: %v", err)
		return
	}

	if len(*setlists) == 0 {
		t.Errorf("Length of setlists is 0")
		return
	}

	newSetlist := Setlist{
		Artist:    "Shadow of Intent",
		ArtistURL: "http://localhost:8008/setlists/shadow-of-intent-2bd8d4fa.html",
		Name:      "setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.html",
		Id:        "c2V0bGlzdC9zaGFkb3ctb2YtaW50ZW50LzIwMjUvYW1hZ2VyLWJpby1jb3BlbmhhZ2VuLWRlbm1hcmstNjM1MzMyODMuaHRtbA==",
		Venue:     "Amager Bio, Copenhagen, Denmark",
		VenueURL:  "http://localhost:8008/venue/amager-bio-copenhagen-denmark-5bd63b24.html",
		Date:      time.Date(2025, 2, 16, 0, 0, 0, 0, time.UTC),
		URL:       "http://localhost:8008/setlist/shadow-of-intent/2025/amager-bio-copenhagen-denmark-63533283.html",
		Songs:     []Song{},
	}

	assert.Equal(t, newSetlist, (*setlists)[0], "Setlists are not equal")

}
