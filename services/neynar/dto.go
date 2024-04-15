package neynar

type ErrorResponse struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Property string `json:"property"`
	Status   int32  `json:"status"`
}

type FarcasterFollowers struct {
	AllRelevantFollowersDehydrated []RelevantFollowersDehydrated `json:"all_relevant_followers_dehydrated"`
}

type RelevantFollowersDehydrated struct {
	Object string         `json:"object"`
	User   UserDehydrated `json:"user"`
}

type UserDehydrated struct {
	// Object string `json:"object"`
	Fid int32 `json:"fid"`
}

type RetrieveCastResponse struct {
	Cast Cast `json:"cast"`
}
type Cast struct {
	Object    string   `json:"object"`
	Hash      string   `json:"hash"`
	Reactions Reaction `json:"reactions"`
	Author   Author   `json:"author"`
}

type Author struct {
	Fid int32 `json:"fid"`
}

type Reaction struct {
	Likes   []Interactor `json:"likes"`
	Recasts []Interactor `json:"recasts"`
}

type Interactor struct {
	Fid   int32  `json:"fid"`
	Fname string `json:"fname"`
}

type ChannelFollowers struct {
	Users []UserDehydrated `json:"users"`
	Next  struct {
		Cursor string `json:"cursor"`
	} `json:"next"`
}

type ThreadCastsResult struct {
	Casts []Cast `json:"casts"`
}

type ThreadCasts struct {
	Result ThreadCastsResult `json:"result"`
}

// type FarcasterUser struct {
// 	Object string `json:"object"`
// 	Fid    int32  `json:"fid"`
// }
