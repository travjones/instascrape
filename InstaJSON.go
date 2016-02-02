package main

type InstaOembedReq struct {
	ShortCode string
	Url       string
}

type InstaOembedResp struct {
	ProviderURL     string      `json:"provider_url"`
	MediaID         string      `json:"media_id"`
	AuthorName      string      `json:"author_name"`
	Height          interface{} `json:"height"`
	ThumbnailURL    string      `json:"thumbnail_url"`
	ThumbnailWidth  int         `json:"thumbnail_width"`
	ThumbnailHeight int         `json:"thumbnail_height"`
	ProviderName    string      `json:"provider_name"`
	Title           string      `json:"title"`
	HTML            string      `json:"html"`
	Width           int         `json:"width"`
	Version         string      `json:"version"`
	AuthorURL       string      `json:"author_url"`
	AuthorID        int         `json:"author_id"`
	Type            string      `json:"type"`
}
