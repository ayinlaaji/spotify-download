package main

import (
	"flag"
	"fmt"
	"github.com/ayinlaaji/spotifyd/spotify"
	"github.com/rylio/ytdl"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Spotifyd struct {
	url string
}

func main() {
	url := flag.String("url", "", "Supply spotify playlist url")
	flag.Parse()

	s := spotify.New()
	body := fetch(url)
	res := s.Parse(body)

	tracks := res.Tracks.Items

	v := "https://www.youtube.com/results?search_query="

	//var htmlChan chan []byte
	//var dwnChan chan []byte

	for _, i := range tracks {
		query := fmt.Sprintf("%s %s", i.Track.Name, i.Track.Artists[0].Name)
		rms := strings.ReplaceAll(query, "-", " ")
		qv := strings.ReplaceAll(rms, " ", "+")

		searchURL := v + qv

		b := fetch(&searchURL)
		videoName := i.Track.Name + "_" + i.Track.Artists[0].Name

		//parse search page for video id
		id := yParse(b)
		getVideo(videoName, id)
	}

}

func fetch(url *string) []byte {
	resp, err := http.Get(*url)

	if err != nil {
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
	}

	return body
}

func yParse(body []byte) string {

	str := string(body)
	exp := `href="\/watch\?v=(.*)\" class`
	data := regexp.MustCompile(exp)

	match := data.FindStringSubmatch(str)

	return match[1]

}

func getVideo(name string, videoID string) {
	fmt.Println(name, videoID)
	vid, err := ytdl.GetVideoInfo("https://www.youtube.com/watch?v=" + videoID)
	if err != nil {
		fmt.Println("Failed to get video info")
		return
	}
	file, _ := os.Create(vid.Title + ".mp4")
	defer file.Close()
	vid.Download(vid.Formats[0], file)
}

func zip() {}
