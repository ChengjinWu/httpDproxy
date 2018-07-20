package letvadx

type BidRequest struct {
	Id     string        `json:"id"`
	Imp    []*Impression `json:"imp"`
	Site   *Site         `json:"site"`
	Device *Device       `json:"device"`
	User   *User         `json:"user"`
	App    *App          `json:"app"`
}

type Impression struct {
	Id       string    `json:"id"`
	Adzoneid int32     `json:"adzoneid"`
	Pmp      *Pmp      `json:"pmp"`
	Bidfloor float64   `json:"bidfloor"`
	Display  int32     `json:"display"`
	Banner   *Banner   `json:"banner"`
	Video    *Video    `json:"video"`
	Native   []*Native `json:"native"`
	Date     string    `json:"date"`
}

type Banner struct {
	W         int32 `json:"w"`
	H         int32 `json:"h"`
	AdModelId int32 `json:"ad_model_id"` // 可投放的广告模版ID
}

type Video struct {
	W           int32    `json:"w"`
	H           int32    `json:"h"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	Mime        []string `json:"mime"` // video/x-flv video/mp4
	AdModelId   int32    `json:"ad_model_id"`
}

type Native struct {
	AdModelId int32         `json:"ad_mode_id"`
	Title     *Native_Title `json:"title"`
	Image     *Native_Image `json:"image"`
	Video     *Native_Video `json:"video"`
	Data      *Native_Data  `json:"data"`
}

type Native_Title struct {
	Len int32 `json:"len"`
}

type Native_Image struct {
	W     int32    `json:"w"`
	H     int32    `json:"h"`
	Mimes []string `json:"mimes"`
}

type Native_Video struct {
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	Mimes       []string `json:"mimes"` // video/x-flv video/mp4
}

type Native_Data struct {
	Len int32 `json:"len"`
}

type Site struct {
	Name    string   `json:"name"`
	Page    string   `json:"page"`
	Ref     string   `json:"ref"`
	Cr      int32    `json:"cr"`
	Content *Content `json:"content"` //only video slot has this info
}

type Content struct {
	Title    string       `json:"title"`
	Keywords string       `json:"keywords"`
	Ext      *Content_Ext `json:"ext"`
}

type Content_Ext struct {
	Channel string `json:"channel"`
	Cs      string `json:"cs"`
}

type Device struct {
	Ua             string      `json:"ua"`
	Ip             string      `json:"ip"`
	Geo            *Geo        `json:"geo"`
	Did            string      `json:"did"`
	Dpid           string      `json:"dpid"`
	Make           string      `json:"make"`
	Model          string      `json:"model"`
	Os             string      `json:"os"`
	Osv            string      `json:"osv"`
	Carrier        string      `json:"carrier"`
	Language       string      `json:"language"`
	Js             int32       `json:"js"`
	Connectiontype int32       `json:"connectiontype"`
	Devicetype     int32       `json:"devicetype"` // 0:mobile 1:pad 2:pc 3:netTV
	Ext            *Device_Ext `json:"ext"`
}

type Geo struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type Device_Ext struct {
	Idfa string `json:"idfa"`
	Mac  string `json:"mac"`
}

type User struct {
	Id     string `json:"id"`
	Gender string `json:"gender"` // M:male F:female
	Yob    int32  `json:"yob"`    // 1880
}

type App struct {
	Name string `json:"name"`
}

type Pmp struct {
	PrivateAuction int32   `json:"private_auction"` // 1:pdb 0:pd
	Deals          []*Deal `json:"deals"`
}

type Deal struct {
	Id       int32 `json:"id"`
	Bidfloor int32 `json:"bidfloor"`
}

type BidResponse struct {
	Id      string     `json:"id"`    // req.id
	Bidid   string     `json:"bidid"` // dsp bid id
	Seatbid []*Seatbid `json:"seatbid"`
}

type Seatbid struct {
	Bid []*Seatbid_Bid `json:"bid"`
}

type Seatbid_Bid struct {
	Id      string           `json:"id"`    // dsp imp id
	Impid   string           `json:"impid"` // letv bidreq.imp.id
	Price   float64          `json:"price"`
	Nurl    string           `json:"nurl,omitempty"` // $(DSP_BID_PRICE)$
	Adm     string           `json:"adm"`
	Admpara string           `json:"admpara,omitempty"`
	Crid    string           `json:"crid"`
	Ext     *Seatbid_Bid_Ext `json:"ext"`
}

type Seatbid_Bid_Ext struct {
	Ldp           string   `json:"ldp"`
	Pm            []string `json:"pm"`
	Cm            []string `json:"cm"`
	Start         []string `json:"start,omitempty"`
	FirstQuartile []string `json:"firstQuartile,omitempty"`
	Midpoint      []string `json:"midpoint,omitempty"`
	ThirdQuartile []string `json:"thirdQuartile,omitempty"`
	Complete      []string `json:"complete,omitempty"`
	IsApp         string   `json:"is_app"`
	AppName       string   `json:"appName,omitempty"`
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	Type          string   `json:"type,omitempty"` // MIME type
}
