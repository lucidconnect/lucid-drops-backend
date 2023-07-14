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
	Id             string   `json:"id"`
	MembershirUIDs map[string]string `json:"memberships"`
	Name           string   `json:"name"`
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

type PatreonCampaignMembers struct {
	Data []struct {
		Attributes struct {
			FullName                     string `json:"full_name"`
			IsFollower                   bool   `json:"is_follower"`
			LastChargeDate               string `json:"last_charge_date"`
			LastChargeStatus             string `json:"last_charge_status"`
			LifetimeSupportCents         int    `json:"lifetime_support_cents"`
			CurrentlyEntitledAmountCents int    `json:"currently_entitled_amount_cents"`
			PatronStatus                 string `json:"patron_status"`
		} `json:"attributes"`
		ID            string `json:"id"`
		Relationships struct {
			Address struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"address"`
			CurrentlyEntitledTiers struct {
				Data []struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"currently_entitled_tiers"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
	Included []struct {
		Attributes struct {
			Addressee   string    `json:"addressee"`
			City        string    `json:"city"`
			Country     string    `json:"country"`
			CreatedAt   time.Time `json:"created_at"`
			Line1       string    `json:"line_1"`
			Line2       string    `json:"line_2"`
			PhoneNumber string    `json:"phone_number"`
			PostalCode  string    `json:"postal_code"`
			State       string    `json:"state"`
		} `json:"attributes"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"included"`
	Meta struct {
		Pagination struct {
			Cursors struct {
				Next string `json:"next"`
			} `json:"cursors"`
			Total int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}

type PatreonUser struct {
	Data struct {
		Attributes struct {
			Email    string `json:"email"`
			FullName string `json:"full_name"`
		} `json:"attributes"`
		ID            string `json:"id"`
		Relationships struct {
			Campaign struct {
				Data struct {
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"campaign"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
	Included []struct {
		Attributes struct {
			IsMonthly bool   `json:"is_monthly"`
			Summary   string `json:"summary"`
		} `json:"attributes"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"included"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
