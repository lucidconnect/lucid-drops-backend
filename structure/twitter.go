package structure

type TweetLikesResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		ResultCount   int    `json:"result_count"`
		PreviousToken string `json:"previous_token"`
		NextToken     string `json:"next_token"`
	} `json:"meta"`
}

type TweetRetweetsResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		ResultCount   int    `json:"result_count"`
		PreviousToken string `json:"previous_token"`
		NextToken     string `json:"next_token"`
	} `json:"meta"`
}

type TwitterFollowersResponse struct {
	Data []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
	Meta struct {
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}


type UserDetailsResponse struct {
	Data struct{
		ID string `json:"id"`
		Name     string `json:"name"`
        Username string `json:"username"`
	}
}


type TweetResponse struct {
	Data struct {
		ThreadedConversationWithInjectionsV2 struct {
			Instructions []struct {
				Type    string `json:"type"`
				Entries []struct {
					EntryID   string `json:"entryId"`
					SortIndex string `json:"sortIndex"`
					Content   struct {
						EntryType   string      `json:"entryType"`
						Typename    string      `json:"__typename"`
						Items       []TweetItem `json:"items,omitempty"`
						ItemContent struct {
							ItemType     string `json:"itemType"`
							Typename     string `json:"__typename"`
							TweetResults struct {
								Result struct {
									Typename string `json:"__typename"`
									RestID   string `json:"rest_id"`
									Core     struct {
										UserResults struct {
											Result struct {
												Typename                   string `json:"__typename"`
												ID                         string `json:"id"`
												RestID                     string `json:"rest_id"`
												AffiliatesHighlightedLabel struct {
												} `json:"affiliates_highlighted_label"`
												IsBlueVerified    bool   `json:"is_blue_verified"`
												ProfileImageShape string `json:"profile_image_shape"`
												Legacy            struct {
													CreatedAt           string `json:"created_at"`
													DefaultProfile      bool   `json:"default_profile"`
													DefaultProfileImage bool   `json:"default_profile_image"`
													Description         string `json:"description"`
													Entities            struct {
														Description struct {
															Urls []struct {
																DisplayURL  string `json:"display_url"`
																ExpandedURL string `json:"expanded_url"`
																URL         string `json:"url"`
																Indices     []int  `json:"indices"`
															} `json:"urls"`
														} `json:"description"`
													} `json:"entities"`
													FastFollowersCount      int           `json:"fast_followers_count"`
													FavouritesCount         int           `json:"favourites_count"`
													FollowersCount          int           `json:"followers_count"`
													FriendsCount            int           `json:"friends_count"`
													HasCustomTimelines      bool          `json:"has_custom_timelines"`
													IsTranslator            bool          `json:"is_translator"`
													ListedCount             int           `json:"listed_count"`
													Location                string        `json:"location"`
													MediaCount              int           `json:"media_count"`
													Name                    string        `json:"name"`
													NormalFollowersCount    int           `json:"normal_followers_count"`
													PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
													PossiblySensitive       bool          `json:"possibly_sensitive"`
													ProfileBannerURL        string        `json:"profile_banner_url"`
													ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
													ProfileInterstitialType string        `json:"profile_interstitial_type"`
													ScreenName              string        `json:"screen_name"`
													StatusesCount           int           `json:"statuses_count"`
													TranslatorType          string        `json:"translator_type"`
													Verified                bool          `json:"verified"`
													WithheldInCountries     []interface{} `json:"withheld_in_countries"`
												} `json:"legacy"`
											} `json:"result"`
										} `json:"user_results"`
									} `json:"core"`
									EditControl struct {
										EditTweetIds       []string `json:"edit_tweet_ids"`
										EditableUntilMsecs string   `json:"editable_until_msecs"`
										IsEditEligible     bool     `json:"is_edit_eligible"`
										EditsRemaining     string   `json:"edits_remaining"`
									} `json:"edit_control"`
									IsTranslatable bool `json:"is_translatable"`
									Views          struct {
										Count string `json:"count"`
										State string `json:"state"`
									} `json:"views"`
									Source             string `json:"source"`
									QuotedStatusResult struct {
										Result struct {
											Typename string `json:"__typename"`
											RestID   string `json:"rest_id"`
											Core     struct {
												UserResults struct {
													Result struct {
														Typename                   string `json:"__typename"`
														ID                         string `json:"id"`
														RestID                     string `json:"rest_id"`
														AffiliatesHighlightedLabel struct {
														} `json:"affiliates_highlighted_label"`
														IsBlueVerified    bool   `json:"is_blue_verified"`
														ProfileImageShape string `json:"profile_image_shape"`
														Legacy            struct {
															CreatedAt           string `json:"created_at"`
															DefaultProfile      bool   `json:"default_profile"`
															DefaultProfileImage bool   `json:"default_profile_image"`
															Description         string `json:"description"`
															Entities            struct {
																Description struct {
																	Urls []interface{} `json:"urls"`
																} `json:"description"`
																URL struct {
																	Urls []struct {
																		DisplayURL  string `json:"display_url"`
																		ExpandedURL string `json:"expanded_url"`
																		URL         string `json:"url"`
																		Indices     []int  `json:"indices"`
																	} `json:"urls"`
																} `json:"url"`
															} `json:"entities"`
															FastFollowersCount      int           `json:"fast_followers_count"`
															FavouritesCount         int           `json:"favourites_count"`
															FollowersCount          int           `json:"followers_count"`
															FriendsCount            int           `json:"friends_count"`
															HasCustomTimelines      bool          `json:"has_custom_timelines"`
															IsTranslator            bool          `json:"is_translator"`
															ListedCount             int           `json:"listed_count"`
															Location                string        `json:"location"`
															MediaCount              int           `json:"media_count"`
															Name                    string        `json:"name"`
															NormalFollowersCount    int           `json:"normal_followers_count"`
															PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
															PossiblySensitive       bool          `json:"possibly_sensitive"`
															ProfileBannerURL        string        `json:"profile_banner_url"`
															ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
															ProfileInterstitialType string        `json:"profile_interstitial_type"`
															ScreenName              string        `json:"screen_name"`
															StatusesCount           int           `json:"statuses_count"`
															TranslatorType          string        `json:"translator_type"`
															URL                     string        `json:"url"`
															Verified                bool          `json:"verified"`
															WithheldInCountries     []interface{} `json:"withheld_in_countries"`
														} `json:"legacy"`
														Professional struct {
															RestID           string        `json:"rest_id"`
															ProfessionalType string        `json:"professional_type"`
															Category         []interface{} `json:"category"`
														} `json:"professional"`
													} `json:"result"`
												} `json:"user_results"`
											} `json:"core"`
											EditControl struct {
												EditTweetIds       []string `json:"edit_tweet_ids"`
												EditableUntilMsecs string   `json:"editable_until_msecs"`
												IsEditEligible     bool     `json:"is_edit_eligible"`
												EditsRemaining     string   `json:"edits_remaining"`
											} `json:"edit_control"`
											IsTranslatable bool `json:"is_translatable"`
											Views          struct {
												Count string `json:"count"`
												State string `json:"state"`
											} `json:"views"`
											Source string `json:"source"`
											Legacy struct {
												BookmarkCount       int    `json:"bookmark_count"`
												Bookmarked          bool   `json:"bookmarked"`
												CreatedAt           string `json:"created_at"`
												ConversationControl struct {
													Policy                   string `json:"policy"`
													ConversationOwnerResults struct {
														Result struct {
															Typename string `json:"__typename"`
															Legacy   struct {
																ScreenName string `json:"screen_name"`
															} `json:"legacy"`
														} `json:"result"`
													} `json:"conversation_owner_results"`
												} `json:"conversation_control"`
												ConversationIDStr string `json:"conversation_id_str"`
												DisplayTextRange  []int  `json:"display_text_range"`
												Entities          struct {
													Media []struct {
														DisplayURL    string `json:"display_url"`
														ExpandedURL   string `json:"expanded_url"`
														IDStr         string `json:"id_str"`
														Indices       []int  `json:"indices"`
														MediaURLHTTPS string `json:"media_url_https"`
														Type          string `json:"type"`
														URL           string `json:"url"`
														Features      struct {
															Large struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"large"`
															Medium struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"medium"`
															Small struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"small"`
															Orig struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"orig"`
														} `json:"features"`
														Sizes struct {
															Large struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"large"`
															Medium struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"medium"`
															Small struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"small"`
															Thumb struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"thumb"`
														} `json:"sizes"`
														OriginalInfo struct {
															Height     int `json:"height"`
															Width      int `json:"width"`
															FocusRects []struct {
																X int `json:"x"`
																Y int `json:"y"`
																W int `json:"w"`
																H int `json:"h"`
															} `json:"focus_rects"`
														} `json:"original_info"`
													} `json:"media"`
													UserMentions []interface{} `json:"user_mentions"`
													Urls         []interface{} `json:"urls"`
													Hashtags     []interface{} `json:"hashtags"`
													Symbols      []interface{} `json:"symbols"`
												} `json:"entities"`
												ExtendedEntities struct {
													Media []struct {
														DisplayURL           string `json:"display_url"`
														ExpandedURL          string `json:"expanded_url"`
														IDStr                string `json:"id_str"`
														Indices              []int  `json:"indices"`
														MediaKey             string `json:"media_key"`
														MediaURLHTTPS        string `json:"media_url_https"`
														Type                 string `json:"type"`
														URL                  string `json:"url"`
														ExtMediaAvailability struct {
															Status string `json:"status"`
														} `json:"ext_media_availability"`
														Features struct {
															Large struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"large"`
															Medium struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"medium"`
															Small struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"small"`
															Orig struct {
																Faces []struct {
																	X int `json:"x"`
																	Y int `json:"y"`
																	H int `json:"h"`
																	W int `json:"w"`
																} `json:"faces"`
															} `json:"orig"`
														} `json:"features"`
														Sizes struct {
															Large struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"large"`
															Medium struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"medium"`
															Small struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"small"`
															Thumb struct {
																H      int    `json:"h"`
																W      int    `json:"w"`
																Resize string `json:"resize"`
															} `json:"thumb"`
														} `json:"sizes"`
														OriginalInfo struct {
															Height     int `json:"height"`
															Width      int `json:"width"`
															FocusRects []struct {
																X int `json:"x"`
																Y int `json:"y"`
																W int `json:"w"`
																H int `json:"h"`
															} `json:"focus_rects"`
														} `json:"original_info"`
													} `json:"media"`
												} `json:"extended_entities"`
												FavoriteCount             int    `json:"favorite_count"`
												Favorited                 bool   `json:"favorited"`
												FullText                  string `json:"full_text"`
												IsQuoteStatus             bool   `json:"is_quote_status"`
												Lang                      string `json:"lang"`
												PossiblySensitive         bool   `json:"possibly_sensitive"`
												PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
												QuoteCount                int    `json:"quote_count"`
												ReplyCount                int    `json:"reply_count"`
												RetweetCount              int    `json:"retweet_count"`
												Retweeted                 bool   `json:"retweeted"`
												UserIDStr                 string `json:"user_id_str"`
												IDStr                     string `json:"id_str"`
											} `json:"legacy"`
										} `json:"result"`
									} `json:"quoted_status_result"`
									Legacy struct {
										BookmarkCount     int    `json:"bookmark_count"`
										Bookmarked        bool   `json:"bookmarked"`
										CreatedAt         string `json:"created_at"`
										ConversationIDStr string `json:"conversation_id_str"`
										DisplayTextRange  []int  `json:"display_text_range"`
										Entities          struct {
											UserMentions []interface{} `json:"user_mentions"`
											Urls         []interface{} `json:"urls"`
											Hashtags     []interface{} `json:"hashtags"`
											Symbols      []interface{} `json:"symbols"`
										} `json:"entities"`
										FavoriteCount         int    `json:"favorite_count"`
										Favorited             bool   `json:"favorited"`
										FullText              string `json:"full_text"`
										IsQuoteStatus         bool   `json:"is_quote_status"`
										Lang                  string `json:"lang"`
										QuoteCount            int    `json:"quote_count"`
										QuotedStatusIDStr     string `json:"quoted_status_id_str"`
										QuotedStatusPermalink struct {
											URL      string `json:"url"`
											Expanded string `json:"expanded"`
											Display  string `json:"display"`
										} `json:"quoted_status_permalink"`
										ReplyCount   int    `json:"reply_count"`
										RetweetCount int    `json:"retweet_count"`
										Retweeted    bool   `json:"retweeted"`
										UserIDStr    string `json:"user_id_str"`
										IDStr        string `json:"id_str"`
									} `json:"legacy"`
									QuickPromoteEligibility struct {
										Eligibility string `json:"eligibility"`
									} `json:"quick_promote_eligibility"`
								} `json:"result"`
							} `json:"tweet_results"`
							TweetDisplayType    string `json:"tweetDisplayType"`
							HasModeratedReplies bool   `json:"hasModeratedReplies"`
						} `json:"itemContent"`
					} `json:"content"`
				} `json:"entries,omitempty"`
				Direction string `json:"direction,omitempty"`
			} `json:"instructions"`
		} `json:"threaded_conversation_with_injections_v2"`
	} `json:"data"`
}

type TweetItem struct {
	EntryID string `json:"entryId"`
	Item    struct {
		ItemContent struct {
			ItemType     string `json:"itemType"`
			Typename     string `json:"__typename"`
			TweetResults struct {
				Result struct {
					Typename string `json:"__typename"`
					RestID   string `json:"rest_id"`
					Core     struct {
						UserResults struct {
							Result struct {
								Typename                   string `json:"__typename"`
								ID                         string `json:"id"`
								RestID                     string `json:"rest_id"`
								AffiliatesHighlightedLabel struct {
								} `json:"affiliates_highlighted_label"`
								IsBlueVerified    bool   `json:"is_blue_verified"`
								ProfileImageShape string `json:"profile_image_shape"`
								Legacy            struct {
									CreatedAt           string `json:"created_at"`
									DefaultProfile      bool   `json:"default_profile"`
									DefaultProfileImage bool   `json:"default_profile_image"`
									Description         string `json:"description"`
									Entities            struct {
										Description struct {
											Urls []interface{} `json:"urls"`
										} `json:"description"`
										URL struct {
											Urls []struct {
												DisplayURL  string `json:"display_url"`
												ExpandedURL string `json:"expanded_url"`
												URL         string `json:"url"`
												Indices     []int  `json:"indices"`
											} `json:"urls"`
										} `json:"url"`
									} `json:"entities"`
									FastFollowersCount      int           `json:"fast_followers_count"`
									FavouritesCount         int           `json:"favourites_count"`
									FollowersCount          int           `json:"followers_count"`
									FriendsCount            int           `json:"friends_count"`
									HasCustomTimelines      bool          `json:"has_custom_timelines"`
									IsTranslator            bool          `json:"is_translator"`
									ListedCount             int           `json:"listed_count"`
									Location                string        `json:"location"`
									MediaCount              int           `json:"media_count"`
									Name                    string        `json:"name"`
									NormalFollowersCount    int           `json:"normal_followers_count"`
									PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
									PossiblySensitive       bool          `json:"possibly_sensitive"`
									ProfileBannerURL        string        `json:"profile_banner_url"`
									ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
									ProfileInterstitialType string        `json:"profile_interstitial_type"`
									ScreenName              string        `json:"screen_name"`
									StatusesCount           int           `json:"statuses_count"`
									TranslatorType          string        `json:"translator_type"`
									URL                     string        `json:"url"`
									Verified                bool          `json:"verified"`
									WithheldInCountries     []interface{} `json:"withheld_in_countries"`
								} `json:"legacy"`
							} `json:"result"`
						} `json:"user_results"`
					} `json:"core"`
					EditControl struct {
						EditTweetIds       []string `json:"edit_tweet_ids"`
						EditableUntilMsecs string   `json:"editable_until_msecs"`
						IsEditEligible     bool     `json:"is_edit_eligible"`
						EditsRemaining     string   `json:"edits_remaining"`
					} `json:"edit_control"`
					IsTranslatable bool `json:"is_translatable"`
					Views          struct {
						Count string `json:"count"`
						State string `json:"state"`
					} `json:"views"`
					Source string `json:"source"`
					Legacy struct {
						BookmarkCount     int    `json:"bookmark_count"`
						Bookmarked        bool   `json:"bookmarked"`
						CreatedAt         string `json:"created_at"`
						ConversationIDStr string `json:"conversation_id_str"`
						DisplayTextRange  []int  `json:"display_text_range"`
						Entities          struct {
							UserMentions []struct {
								IDStr      string `json:"id_str"`
								Name       string `json:"name"`
								ScreenName string `json:"screen_name"`
								Indices    []int  `json:"indices"`
							} `json:"user_mentions"`
							Urls     []interface{} `json:"urls"`
							Hashtags []interface{} `json:"hashtags"`
							Symbols  []interface{} `json:"symbols"`
						} `json:"entities"`
						FavoriteCount        int    `json:"favorite_count"`
						Favorited            bool   `json:"favorited"`
						FullText             string `json:"full_text"`
						InReplyToScreenName  string `json:"in_reply_to_screen_name"`
						InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
						InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
						IsQuoteStatus        bool   `json:"is_quote_status"`
						Lang                 string `json:"lang"`
						QuoteCount           int    `json:"quote_count"`
						ReplyCount           int    `json:"reply_count"`
						RetweetCount         int    `json:"retweet_count"`
						Retweeted            bool   `json:"retweeted"`
						UserIDStr            string `json:"user_id_str"`
						IDStr                string `json:"id_str"`
					} `json:"legacy"`
					QuickPromoteEligibility struct {
						Eligibility string `json:"eligibility"`
					} `json:"quick_promote_eligibility"`
				} `json:"result"`
			} `json:"tweet_results"`
			TweetDisplayType string `json:"tweetDisplayType"`
		} `json:"itemContent"`
		ClientEventInfo struct {
			Details struct {
				ConversationDetails struct {
					ConversationSection string `json:"conversationSection"`
				} `json:"conversationDetails"`
				TimelinesDetails struct {
					ControllerData string `json:"controllerData"`
				} `json:"timelinesDetails"`
			} `json:"details"`
		} `json:"clientEventInfo"`
	} `json:"item"`
}
