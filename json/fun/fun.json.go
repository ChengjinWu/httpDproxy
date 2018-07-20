package fun

//The bid request/response struct is built by "风行-FUNADX 对接 Interface Specification.pdf"
//request
type FunBidRequest struct {
	Id      string                `json:"id"`
	Version string                `json:"version"`
	IsTest  int32                 `json:"istest"`
	IsPing  int32                 `json:"isping"`
	Imp     []*FunBidRequest_Imp  `json:"imp"`
	Site    *FunBidRequest_Site   `json:"site"`
	Device  *FunBidRequest_Device `json:"device"`
	User    *FunBidRequest_User   `json:"user"`
	App     *FunBidRequest_App    `json:"app"`
	Ext     interface{}           `json:"ext"`
}

type FunBidRequest_Imp struct {
	Id       string                    `json:"id"`
	Tagid    string                    `json:"tagid"`
	Bidfloor float32                   `json:"bidfloor"`
	Pmp      *FunBidReuqest_Imp_Pmp    `json:"pmp"`
	Banner   *FunBidRequest_Imp_Banner `json:"banner"`
	Video    *FunBidRequest_Imp_Video  `json:"video"`
}

type FunBidReuqest_Imp_Pmp struct {
	Deals []*FunBidRequest_Pmp_Deal `json:"deals"`
}

type FunBidRequest_Pmp_Deal struct {
	Id string `json:"id"`
}

type FunBidRequest_Imp_Banner struct {
	W       int32    `json:"w"`
	H       int32    `json:"h"`
	Isratio int32    `json:"isratio"`
	Wmax    int32    `json:"wmax"`
	Hmax    int32    `json:"hmax"`
	Wmin    int32    `json:"wmin"`
	Hmin    int32    `json:"hmin"`
	Mimes   []string `json:"mimes"`
}

type FunBidRequest_Imp_Video struct {
	Mimes       []string `json:"mimes"`
	Linearity   int32    `json:"linearity"`
	Sequence    int32    `json:"sequence"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	W           int32    `json:"w"`
	H           int32    `json:"h"`
	Isratio     int32    `json:"isratio"`
	Wmax        int32    `json:"wmax"`
	Hmax        int32    `json:"hmax"`
	Wmin        int32    `json:"wmin"`
	Hmin        int32    `json:"hmin"`
}

type FunBidRequest_Site struct {
	Name    string                 `json:"name"`
	Cat     []string               `json:"cat"`
	Page    string                 `json:"page"`
	Ref     string                 `json:"ref"`
	Content *FunBidRequest_Content `json:"content"`
}

type FunBidRequest_Device struct {
	Ua             string                    `json:"ua"`
	Ip             string                    `json:"ip"`
	Geo            *FunBidRequest_Geo        `json:"geo"`
	Did            string                    `json:"did"`
	Didmd5         string                    `json:"didmd5"`
	Dpid           string                    `json:"dpid"`
	Dpidmd5        string                    `json:"dpidmd5"`
	Make           string                    `json:"make"`
	Model          string                    `json:"model"`
	Os             string                    `json:"os"`
	Osv            string                    `json:"osv"`
	Carrier        string                    `json:"carrier"`
	Language       string                    `json:"language"`
	Js             int32                     `json:"js"`
	Connectiontype int32                     `json:"connectiontype"`
	Devicetype     int32                     `json:"devicetype"`
	Ext            *FunBidRequest_Device_Ext `json:"ext"`
}

type FunBidRequest_Device_Ext struct {
	Idfa        string `json:"idfa"`
	Mac         string `json:"mac"`
	Macmd5      string `json:"macmd5"`
	Ssid        string `json:"ssid"`
	W           int32  `json:"w"`
	H           int32  `json:"h"`
	Brk         int32  `json:"brk"`
	Interactive int32  `json:"interactive"`
}

type FunBidRequest_Geo struct {
	Lat float32                `json:"lat"`
	Lon float32                `json:"lon"`
	Ext *FunBidRequest_Geo_Ext `json:"ext"`
}

type FunBidRequest_Geo_Ext struct {
	Accuracy int32 `json:"accuracy"`
}

type FunBidRequest_User struct {
	Id string `json:"id"`
}

type FunBidRequest_App struct {
	Name    string                 `json:"name"`
	Cat     []string               `json:"cat"`
	Content *FunBidRequest_Content `json:"content"`
}

type FunBidRequest_Content struct {
	Id  string                     `json:"id"`
	Ext *FunBidRequest_Content_Ext `json:"ext"`
}

type FunBidRequest_Content_Ext struct {
	Channel string `json:"channel"`
}

type FunBidResponse struct {
	Id      string                    `json:"id"`    //reqid
	Bidid   string                    `json:"bidid"` //dsp respid
	Seatbid []*FunBidResponse_Seatbid `json:"seatbid"`
}

type FunBidResponse_Seatbid struct {
	Bid []*FunBidResponse_Seatbid_Bid `json:"bid"`
}

type FunBidResponse_Seatbid_Bid struct {
	Id    string              `json:"id"`
	Impid string              `json:"impid"`
	Price float32             `json:"price"`
	Nurl  string              `json:"nurl,omitempty"`
	Adm   string              `json:"adm"`
	Crid  string              `json:"crid"`
	Ext   *FunBidResponse_Ext `json:"ext,omitempty"`
}

type FunBidResponse_Ext struct {
	Ldp   string                   `json:"ldp,omitempty"`
	Pm    []*FunBidResponse_Pm     `json:"pm,omitempty"`
	Cm    []string                 `json:"cm,omitempty"`
	CmExt []*FunBidResponse_Cm_Ext `json:"cm_ext,omitempty"`
}
type FunBidResponse_Pm struct {
	Point    int32  `json:"point"`
	Url      string `json:"url"`
	Provider string `json:"provider,omitempty"`
}

type FunBidResponse_Cm_Ext struct {
	Provider string `json:"provider,omitempty"`
}
