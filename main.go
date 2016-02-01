package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

func scraper(url string) {
	// GET
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// dump response and include body
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	// convert dumped response body to string
	respString := string(dump)

	// split the string
	shards := strings.Split(respString, "window._sharedData = ")
	instaJson := strings.Split(shards[1], ";</script>")

	// instaResp to hold json
	var data instaResp

	// unmarshal json into &data
	if err := json.Unmarshal([]byte(instaJson[0]), &data); err != nil {
		panic(err)
	}

	// print shortcode for first node -- will use with open oembed insta api endpoint
	fmt.Println(data.EntryData.Profilepage[0].User.Media.Nodes[0].Code)
}

func main() {
	scraper("http://instagram.com/habitatskateboards")
}
