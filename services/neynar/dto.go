package neynar

type FarcasterFollowers struct {
	AllRelevantFollowersDehydrated []RelevantFollowersDehydrated `json:"all_relevant_followers_dehydrated"`
}

type RelevantFollowersDehydrated struct {
	Object string         `json:"object"`
	User   UserDehydrated `json:"user"`
}

type UserDehydrated struct {
	Object string `json:"object"`
	Fid    int32  `json:"fid"`
}

type Cast struct {
	Object    string     `json:"object"`
	Hash      string     `json:"hash"`
	Reactions Reaction `json:"reactions"`
	Authour   Author     `json:"author"`
}

type Author struct {
	Fid int32 `json:"fid"`
}

type Reaction struct {
	Likes   []Interactor `json:"likes"`
	Recasts []Interactor `json:"recasts"`
}

type Interactor struct {
	Fid   int32 `json:"fid"`
	Fname string `json:"fname"`
}

type ChannelFollowers struct {
	Users []UserDehydrated `json:"users"`
	Next  struct {
		Cursor string `json:"cursor"`
	} `json:"next"`
}
