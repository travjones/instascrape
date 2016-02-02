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
	var iwr InstaWebResp

	// unmarshal json into &iwr
	if err := json.Unmarshal([]byte(instaJson[0]), &iwr); err != nil {
		panic(err)
	}

	// print shortcode for latest 12 nodes (media) -- will use with open oembed insta api endpoint
	// for i := 0; i < 12; i++ {
	// 	fmt.Println("https://api.instagram.com/oembed/?url=http://instagr.am/p/" + data.EntryData.Profilepage[0].User.Media.Nodes[i].Code)
	// }

	// slice of InstaOembedReqs
	var ioreqs []InstaOembedReq

	// for each piece of media add shortcode and url and then append to slice
	for i := 0; i < 12; i++ {
		var ior InstaOembedReq
		ior.ShortCode = iwr.EntryData.Profilepage[0].User.Media.Nodes[i].Code
		// omitting script -- make sure you pull it in on the page
		ior.Url = "https://api.instagram.com/oembed/?url=http://instagr.am/p/" + iwr.EntryData.Profilepage[0].User.Media.Nodes[i].Code + "&omitscript=true"
		ioreqs = append(ioreqs, ior)
	}

	// slice of InstaOembedResp
	var ioresps []InstaOembedResp

	// for each piece of media GET JSON
	for i := 0; i < 12; i++ {
		// new GET req using oembed URL
		req, err := http.NewRequest("GET", ioreqs[i].Url, nil)
		if err != nil {
			panic(err)
		}

		// da muthafuckin client
		client := &http.Client{}

		// Do request and drop response in resp
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close() // close dat shit or leak memory

		// fresh IOResp to hold each response
		var ioresp InstaOembedResp

		// decode response into &ioresp
		if err := json.NewDecoder(resp.Body).Decode(&ioresp); err != nil {
			panic(err)
		}

		// append each response to the IOResps slice
		ioresps = append(ioresps, ioresp)
	}

	// print HTML for each insta embed
	for i := 0; i < 12; i++ {
		fmt.Println(ioresps[i].HTML)
	}
}

func main() {
	scraper("http://instagram.com/thrashermag")
}
