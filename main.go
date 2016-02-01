package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
)

func scraper(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(dump)

	respString := buf.String()
	shards := strings.Split(respString, "window._sharedData = ")
	shards2 := strings.Split(shards[1], ";</script>")

	var data instaResp

	if err := json.Unmarshal([]byte(shards2[0]), &data); err != nil {
		panic(err)
	}

	fmt.Println(data.EntryData.Profilepage[0].User.Media.Nodes[0].Code)
}

func main() {
	scraper("http://instagram.com/habitatskateboards")
}
