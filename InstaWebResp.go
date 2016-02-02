package main

type InstaWebResp struct {
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Gatekeepers  struct {
		Rhp bool `json:"rhp"`
	} `json:"gatekeepers"`
	Qs         string `json:"qs"`
	StaticRoot string `json:"static_root"`
	Platform   string `json:"platform"`
	Hostname   string `json:"hostname"`
	EntryData  struct {
		Profilepage []struct {
			User struct {
				Username string `json:"username"`
				Follows  struct {
					Count int `json:"count"`
				} `json:"follows"`
				RequestedByViewer bool `json:"requested_by_viewer"`
				FollowedBy        struct {
					Count int `json:"count"`
				} `json:"followed_by"`
				CountryBlock       interface{} `json:"country_block"`
				HasRequestedViewer bool        `json:"has_requested_viewer"`
				FollowedByViewer   bool        `json:"followed_by_viewer"`
				FollowsViewer      bool        `json:"follows_viewer"`
				ProfilePicURL      string      `json:"profile_pic_url"`
				ID                 string      `json:"id"`
				Biography          interface{} `json:"biography"`
				FullName           string      `json:"full_name"`
				Media              struct {
					Count    int `json:"count"`
					PageInfo struct {
						HasPreviousPage bool   `json:"has_previous_page"`
						StartCursor     string `json:"start_cursor"`
						EndCursor       string `json:"end_cursor"`
						HasNextPage     bool   `json:"has_next_page"`
					} `json:"page_info"`
					Nodes []struct {
						Code       string  `json:"code"`
						Date       float64 `json:"date"`
						Dimensions struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"dimensions"`
						Comments struct {
							Count int `json:"count"`
						} `json:"comments"`
						Caption string `json:"caption"`
						Likes   struct {
							Count int `json:"count"`
						} `json:"likes"`
						Owner struct {
							ID string `json:"id"`
						} `json:"owner"`
						ThumbnailSrc string `json:"thumbnail_src"`
						IsVideo      bool   `json:"is_video"`
						ID           string `json:"id"`
						DisplaySrc   string `json:"display_src"`
					} `json:"nodes"`
				} `json:"media"`
				BlockedByViewer  bool   `json:"blocked_by_viewer"`
				IsVerified       bool   `json:"is_verified"`
				HasBlockedViewer bool   `json:"has_blocked_viewer"`
				IsPrivate        bool   `json:"is_private"`
				ExternalURL      string `json:"external_url"`
			} `json:"user"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Qe struct {
		Su struct {
			P struct {
			} `json:"p"`
			G string `json:"g"`
		} `json:"su"`
		Discovery struct {
			P struct {
			} `json:"p"`
			G string `json:"g"`
		} `json:"discovery"`
	} `json:"qe"`
	DisplayPropertiesServerGuess struct {
		ViewportWidth int     `json:"viewport_width"`
		PixelRatio    float64 `json:"pixel_ratio"`
	} `json:"display_properties_server_guess"`
	Config struct {
		Viewer    interface{} `json:"viewer"`
		CsrfToken string      `json:"csrf_token"`
	} `json:"config"`
	EnvironmentSwitcherVisibleServerGuess bool `json:"environment_switcher_visible_server_guess"`
}
