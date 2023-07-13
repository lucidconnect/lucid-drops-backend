package structure

import "time"

type PatreonAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	Version      string `json:"version"`
}

type PatreonUserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PatreonCampaignInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PatreonCampaigns struct {
	Data []struct {
		Attributes struct {
			CreatedAt            time.Time   `json:"created_at"`
			CreationName         string      `json:"creation_name"`
			DiscordServerID      string      `json:"discord_server_id"`
			GoogleAnalyticsID    string      `json:"google_analytics_id"`
			HasRss               bool        `json:"has_rss"`
			HasSentRssNotify     bool        `json:"has_sent_rss_notify"`
			ImageSmallURL        string      `json:"image_small_url"`
			ImageURL             string      `json:"image_url"`
			IsChargedImmediately bool        `json:"is_charged_immediately"`
			IsMonthly            bool        `json:"is_monthly"`
			IsNsfw               bool        `json:"is_nsfw"`
			MainVideoEmbed       interface{} `json:"main_video_embed"`
			MainVideoURL         string      `json:"main_video_url"`
			OneLiner             interface{} `json:"one_liner"`
			PatronCount          int         `json:"patron_count"`
			PayPerName           string      `json:"pay_per_name"`
			PledgeURL            string      `json:"pledge_url"`
			PublishedAt          time.Time   `json:"published_at"`
			RssArtworkURL        string      `json:"rss_artwork_url"`
			RssFeedTitle         string      `json:"rss_feed_title"`
			Summary              string      `json:"summary"`
			ThanksEmbed          interface{} `json:"thanks_embed"`
			ThanksMsg            interface{} `json:"thanks_msg"`
			ThanksVideoURL       interface{} `json:"thanks_video_url"`
		} `json:"attributes"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

type PatreonCampaign struct {
	Data struct {
		Attributes struct {
			CreatedAt            time.Time `json:"created_at"`
			CreationName         string    `json:"creation_name"`
			DiscordServerID      string    `json:"discord_server_id"`
			ImageSmallURL        string    `json:"image_small_url"`
			ImageURL             string    `json:"image_url"`
			IsChargedImmediately bool      `json:"is_charged_immediately"`
			IsMonthly            bool      `json:"is_monthly"`
			MainVideoEmbed       string    `json:"main_video_embed"`
			MainVideoURL         string    `json:"main_video_url"`
			OneLiner             string    `json:"one_liner"`
			PatronCount          int       `json:"patron_count"`
			PayPerName           string    `json:"pay_per_name"`
			PledgeURL            string    `json:"pledge_url"`
			PublishedAt          time.Time `json:"published_at"`
			Summary              string    `json:"summary"`
			ThanksEmbed          string    `json:"thanks_embed"`
			ThanksMsg            string    `json:"thanks_msg"`
			ThanksVideoURL       string    `json:"thanks_video_url"`
		} `json:"attributes"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}
