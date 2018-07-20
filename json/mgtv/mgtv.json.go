package mgtv

type BidRequest struct {
	Version     int32                    `json:"version"`
	BId         string                   `json:"bid"`
	RequestType int32                    `json:"request_type"`
	Imp         []*BidRequest_Impression `json:"imp"`
	Device      *BidRequest_Device       `json:"device"`
	Video       *BidRequest_Video        `json:"video"`
}

type BidRequest_Impression struct {
	SpaceId     string            `json:"space_id"`
	Width       int32             `json:"width"`
	Height      int32             `json:"height"`
	MinCpmPrice int32             `json:"min_cpm_price"` // fen cpm
	PlayerId    int32             `json:"player_id"`
	Ctype       []int32           `json:"ctype"`
	Playtime    int32             `json:"playtime"` // sec
	MinPlaytime int32             `json:"min_playtime"`
	MaxPlaytime int32             `json:"max_playtime"`
	Location    int32             `json:"location"`
	Order       int32             `json:"order"`
	Pmp         []*Impression_Pmp `json:"pmp"`
}

type Impression_Pmp struct {
	Id    string `json:"id"`
	At    int32  `json:"at"` // 0:pdb 1:first price deal
	Price int32  `json:"price"`
}

type BidRequest_Device struct {
	Imei           string `json:"imei"`
	Idfa           string `json:"idfa"`
	Anid           string `json:"anid"`
	Mac            string `json:"mac"`
	Os             string `json:"os"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	Ip             string `json:"ip"`
	Ua             string `json:"ua"`
	Connectiontype int32  `json:"connectiontype"` // 0:wifi 1:mob 2:no network
	Type           int32  `json:"type"`
	Duid           string `json:"duid"` // pc user's cookie
	Len            string `json:"len"`  // longitude 经度
	Lon            string `json:"lon"`  // latitude 纬度
}

type BidRequest_Video struct {
	VideoUrl  string `json:"video_url"`
	VideoName string `json:"video_name"`
	ItemIds   string `json:"item_ids"`
	ItemNames string `json:"item_names"`
}

type BidResponse struct {
	Version int32             `json:"version"`
	BId     string            `json:"bid"`
	ErrCode int32             `json:"err_code"` // 200 204
	Ads     []*BidResponse_Ad `json:"ads,omitempty"`
}

type BidResponse_Ad struct {
	SpaceId         string      `json:"space_id"`
	Price           int32       `json:"price"`
	AdUrl           string      `json:"ad_url,omitempty"`
	CreativeId      string      `json:"creative_id"`
	ClickThroughUrl string      `json:"click_through_url"`
	Iurl            []*Ads_Iurl `json:"iurl"`
	Curl            []string    `json:"curl"`
	Title           string      `json:"title,omitempty"`
	Duration        int32       `json:"duration,omitempty"`
	Ctype           int32       `json:"ctype"` // 1:pic 2:video
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Dealid          string      `json:"dealid,omitempty"`
}

type Ads_Iurl struct {
	Event int32  `json:"event,omitempty"` // 0:start
	Url   string `json:"url"`
}
