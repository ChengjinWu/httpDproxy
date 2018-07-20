package mi

type BidRequest struct {
	Id     string                   `json:"id"`
	Imp    []*BidRequest_Impression `json:"imp"`
	Site   *BidRequest_Site         `json:"site"`
	App    *BidRequest_App          `json:"app"`
	Device *BidRequest_Device       `json:"device"`
	User   *BidRequest_User         `json:"user"`
	Test   int                      `json:"test"`
	At     int                      `json:"at"`
	Tmax   int                      `json:"tmax"`
	Bcat   []string                 `json:"bcat"`
	Badv   []string                 `json:"badv"`
	Ext    interface{}              `json:"ext"`
}

type BidRequest_Impression struct {
	Id         string                   `json:"id"`
	TagId      string                   `json:"tagid"`
	Banner     *BidRequest_Banner       `json:"banner"`
	TextLink   *BidRequest_TextLink     `json:"textlink"`
	NativeAd   *BidRequest_Native       `json:"nativead"`
	Video      *BidRequest_Video        `json:"video"`
	Splash     *BidRequest_Splash       `json:"splash"`
	Instl      int                      `json:"instl"`
	AdmType    int                      `json:"admtype"`
	Templates  []*BidRequest_AdTemplate `json:"templates"`
	Pos        int                      `json:"pos"`
	BidFloor   float64                  `json:"bidfloor"`
	DirectDeal *BidRequest_DirectDeal   `json:"directdeal"`
	AdsCount   int                      `json:"adscount"`
	Ext        string                   `json:"ext"`
}

type BidRequest_Banner struct {
	W         int         `json:"w"`
	H         int         `json:"h"`
	Id        string      `json:"id"`
	Btype     []int       `json:"btype"`
	Battr     []int       `json:"battr"`
	Embedding int         `json:"embedding"`
	Mimes     []string    `json:"mines"`
	Ext       interface{} `json:"ext"`
}

type BidRequest_Native struct {
	Request string      `json:"request"`
	Id      string      `json:"id"`
	W       int         `json:"w"`
	H       int         `json:"h"`
	Ver     string      `json:"ver"`
	Api     []int       `json:"api"`
	Battr   []int       `json:"battr"`
	Ext     interface{} `json:"ext"`
}

type BidRequest_TextLink struct {
	Id     string      `json:"id"`
	Battr  []int       `json:"battr"`
	Length int         `json:"length"`
	Size   int         `json:"size"`
	Ext    interface{} `json:"ext"`
}

type BidRequest_Video struct {
	Mimes            []string    `json:"mines"`
	MinDuration      int         `json:"minduration"`
	MaxDuration      int         `json:"maxduration"`
	Protocols        []int       `json:"protocols"`
	W                int         `json:"w"`
	H                int         `json:"h"`
	Size             int         `json:"size"`
	Battr            []int       `json:"battr"`
	FrequencyCapping int         `json:"frequencycapping"`
	Ext              interface{} `json:"ext"`
}

type BidRequest_Splash struct {
	W                int         `json:"w"`
	H                int         `json:"h"`
	Skip             int         `json:"skip"`
	Duration         int         `json:"duration"`
	DetailPos        int         `json:"detailpos"`
	Btype            []int       `json:"btype"`
	Battr            []int       `json:"battr"`
	Mimes            []string    `json:"mines"`
	FrequencyCapping int         `json:"frequencycapping"`
	Ext              interface{} `json:"ext"`
}

type BidRequest_Site struct {
	Id        string                `json:"id"`
	Name      string                `json:"name"`
	Domain    string                `json:"domain"`
	Cat       []string              `json:"cat"`
	Page      string                `json:"page"`
	Ref       string                `json:"ref"`
	Search    string                `json:"search"`
	Mobile    int                   `json:"mobile"`
	Publisher *BidRequest_Publisher `json:"publisher"`
	Keywords  string                `json:"keywords"`
	Ext       interface{}           `json:"ext"`
}
type BidRequest_App struct {
	Id        string                `json:"id"`
	Name      string                `json:"name"`
	Bundle    string                `json:"bundle"`
	Domain    string                `json:"domain"`
	StoreUrl  string                `json:"storeurl"`
	Cat       []string              `json:"cat"`
	Ver       string                `json:"ver"`
	Paid      int                   `json:"paid"`
	Publisher *BidRequest_Publisher `json:"publisher"`
	Keywords  string                `json:"keywords"`
	Ext       interface{}           `json:"ext"`
}

type BidRequest_Publisher struct {
	Id     string      `json:"id"`
	Name   string      `json:"name"`
	Cat    []string    `json:"cat"`
	Domain string      `json:"domain"`
	Ext    interface{} `json:"ext"`
}

type BidRequest_Device struct {
	Ua             string          `json:"ua"`
	Geo            *BidRequest_Geo `json:"geo"`
	Ip             string          `json:"ip"`
	Ipv6           string          `json:"ipv6"`
	DeviceType     int             `json:"devicetype"`
	Make           string          `json:"make"`
	Model          string          `json:"model"`
	Os             string          `json:"os"`
	Osv            string          `json:"osv"`
	Hwv            string          `json:"hwv"`
	H              int             `json:"h"`
	W              int             `json:"w"`
	Ppi            int             `json:"ppi"`
	PxRatio        float64         `json:"pxratio"`
	Js             int             `json:"js"`
	Flashver       string          `json:"flashver"`
	Language       string          `json:"language"`
	Carrier        string          `json:"carrier"`
	ConnectionType int             `json:"connectiontype"`
	Ifa            string          `json:"ifa"`
	DidSha1        string          `json:"didsha1"`
	DidMd5         string          `json:"didmd5"`
	Dpid           string          `json:"dpid"`
	DpidSha1       string          `json:"dpidsha1"`
	DpidMd5        string          `json:"dpidmd5"`
	IdfaSha1       string          `json:"idfasha1"`
	IdfaMd5        string          `json:"idfamd5"`
	MacSha1        string          `json:"macsha1"`
	MacMd5         string          `json:"macmd5"`
	Cookie         string          `json:"cookie"`
	Ext            interface{}     `json:"ext"`
}

type BidRequest_Geo struct {
	Lat     float64     `json:"lat"`
	Lon     float64     `json:"lon"`
	Type    int         `json:"type"`
	Country string      `json:"country"`
	Region  string      `json:"region"`
	City    string      `json:"city"`
	Ext     interface{} `json:"ext"`
}

type BidRequest_DirectDeal struct {
	Id  string      `json:"id"`
	Ext interface{} `json:"ext"`
}

type BidRequest_AdTemplate struct {
	Id     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type BidRequest_User struct {
	Id         string             `json:"id"`
	BuyerId    string             `json:"buyerid"`
	Yob        int                `json:"yob"`
	Gender     string             `json:"gender"`
	Keywords   string             `json:"keywords"`
	CustomData string             `json:"customdata"`
	Data       []*BidRequest_Data `json:"data"`
	Ext        interface{}        `json:"ext"`
}

type BidRequest_Data struct {
	Id      string                `json:"id"`
	Name    string                `json:"name"`
	Segment []*BidRequest_Segment `json:"segment"`
	Ext     interface{}           `json:"ext"`
}

type BidRequest_Segment struct {
	Id    string      `json:"id"`
	Name  string      `json:"name"`
	Value string      `json:"value"`
	Ext   interface{} `json:"ext"`
}

type BidResponse struct {
	Id      string                 `json:"id"`
	SeatBid []*BidResponse_SeatBid `json:"seatbid,omitempty"`
	BidId   string                 `json:"bidid,omitempty"`
	Nbr     int                    `json:"nbr,omitempty"`
	Ext     interface{}            `json:"ext,omitempty"`
}

type BidResponse_SeatBid struct {
	Bid   []*BidResponse_Bid `json:"bid"`
	Seat  string             `json:"seat"`
	Cm    int                `json:"cm"`
	Group int                `json:"group"`
	Ext   interface{}        `json:"ext,omitempty"`
}
type BidResponse_Bid struct {
	Id               string                        `json:"id"`
	ImpId            string                        `json:"impid"`
	Price            float64                       `json:"price"`
	Adid             string                        `json:"adid"`
	Nurl             string                        `json:"nurl,omitempty"`
	Adm              string                        `json:"adm,omitempty"`
	TagId            string                        `json:"tagid,omitempty"`
	TemplateId       string                        `json:"templateid"`
	BillingType      int                           `json:"billingtype,omitempty"`
	Adomain          []string                      `json:"adomain,omitempty"`
	Bundle           string                        `json:"bundle,omitempty"`
	Cid              string                        `json:"cid,omitempty"`
	Crid             string                        `json:"crid,omitempty"`
	Cat              []string                      `json:"cat,omitempty"`
	Attr             []int                         `json:"attr,omitempty"`
	H                int                           `json:"h,omitempty"`
	W                int                           `json:"w,omitempty"`
	LandingUrl       string                        `json:"landingurl,omitempty"`
	ImpUrl           []string                      `json:"impurl,omitempty"`
	Curl             []string                      `json:"curl,omitempty"`
	Durl             []string                      `json:"durl,omitempty"`
	VideoStartUrl    []string                      `json:"videostarturl,omitempty"`
	VideoStopUrl     []string                      `json:"videostopurl,omitempty"`
	VideoEndUrl      []string                      `json:"videoendurl,omitempty"`
	FrequencyCapping *BidResponse_FrequencyCapping `json:"frequencycapping,omitempty"`
	ExtData          string                        `json:"extdata,omitempty"`
	Ext              interface{}                   `json:"ext,omitempty"`
}

type BidResponse_FrequencyCapping struct {
	Global int         `json:"global,omitempty"`
	Weekly int         `json:"weekly,omitempty"`
	daily  int         `json:"daily,omitempty"`
	Hourly int         `json:"hourly,omitempty"`
	Ext    interface{} `json:"ext,omitempty"`
}

func (this *BidRequest) GetId() string {
	if this == nil {
		return ""
	}
	return this.Id
}
func (this *BidRequest) GetImp() []*BidRequest_Impression {
	if this == nil {
		return nil
	}
	return this.Imp
}

func (this *BidRequest) GetSite() *BidRequest_Site {
	if this == nil {
		return nil
	}
	return this.Site
}
func (this *BidRequest) GetApp() *BidRequest_App {
	if this == nil {
		return nil
	}
	return this.App
}
func (this *BidRequest) GetDevice() *BidRequest_Device {
	if this == nil {
		return nil
	}
	return this.Device
}
func (this *BidRequest) GetUser() *BidRequest_User {
	if this == nil {
		return nil
	}
	return this.User
}

func (this *BidRequest) GetBcat() []string {
	if this == nil {
		return nil
	}
	return this.Bcat
}

func (this *BidRequest) GetBadv() []string {
	if this == nil {
		return nil
	}
	return this.Badv
}

func (this *BidRequest_Device) GetCountry() string {
	if this == nil || this.Geo == nil {
		return ""
	}
	return this.Geo.Country
}

func (this *BidRequest_Device) GetCarrier() string {
	if this == nil {
		return ""
	}
	return this.Carrier
}

func (this *BidRequest_Device) GetMake() string {
	if this == nil {
		return ""
	}
	return this.Make
}

func (this *BidRequest_Device) GetModel() string {
	if this == nil {
		return ""
	}
	return this.Model
}

func (this *BidRequest_Device) GetUa() string {
	if this == nil {
		return ""
	}
	return this.Ua
}

func (this *BidRequest_Device) GetIp() string {
	if this == nil {
		return ""
	}
	return this.Ip
}

func (this *BidRequest_Device) GetOs() string {
	if this == nil {
		return ""
	}
	return this.Os
}

func (this *BidRequest_Device) GetOsv() string {
	if this == nil {
		return ""
	}
	return this.Osv
}

func (this *BidRequest_Device) GetConnectionType() int {
	if this == nil {
		return 0
	}
	return this.ConnectionType
}

func (this *BidRequest_Device) GetGeo() *BidRequest_Geo {
	if this == nil {
		return nil
	}
	return this.Geo
}

func (this *BidRequest_Device) GetLanguage() string {
	if this == nil {
		return ""
	}
	return this.Language
}

func (this *BidRequest_Site) GetReferer() string {
	if this == nil {
		return ""
	}
	return this.Ref
}

func (this *BidRequest_User) GetYearOfBorn() int {
	if this == nil {
		return 0
	}
	return this.Yob
}

func (this *BidRequest_Impression) GetBanner() *BidRequest_Banner {
	if this == nil {
		return nil
	}
	return this.Banner
}

func (this *BidRequest_Impression) GetTextLink() *BidRequest_TextLink {
	if this == nil {
		return nil
	}
	return this.TextLink
}

func (this *BidRequest_Impression) GetNativeAd() *BidRequest_Native {
	if this == nil {
		return nil
	}
	return this.NativeAd
}

func (this *BidRequest_Impression) GetVideo() *BidRequest_Video {
	if this == nil {
		return nil
	}
	return this.Video
}

func (this *BidRequest_Impression) GetSplash() *BidRequest_Splash {
	if this == nil {
		return nil
	}
	return this.Splash
}

func (this *BidRequest_Impression) GetTemplates() []*BidRequest_AdTemplate {
	if this == nil {
		return nil
	}
	return this.Templates
}

func (this *BidRequest_Impression) GetDirectDealId() string {
	if this == nil || this.DirectDeal == nil {
		return ""
	}

	return this.DirectDeal.Id
}

func (this *BidRequest_Video) GetMinDuration() int {
	if this == nil {
		return 0
	}
	return this.MinDuration
}

func (this *BidRequest_Video) GetMaxDuration() int {
	if this == nil {
		return 0
	}
	return this.MaxDuration
}

func (this *BidRequest_Banner) GetW() int {
	if this == nil {
		return 0
	}
	return this.W
}

func (this *BidRequest_Banner) GetH() int {
	if this == nil {
		return 0
	}
	return this.H
}

func (this *BidRequest_Video) GetH() int {
	if this == nil {
		return 0
	}
	return this.H
}

func (this *BidRequest_Video) GetW() int {
	if this == nil {
		return 0
	}
	return this.W
}

func (this *BidRequest_Banner) GetBattr() []int {
	if this == nil {
		return nil
	}
	return this.Battr
}

func (this *BidRequest_Native) GetW() int {
	if this == nil {
		return 0
	}
	return this.W
}

func (this *BidRequest_Native) GetH() int {
	if this == nil {
		return 0
	}
	return this.H
}

func (this *BidRequest_Native) GetBattr() []int {
	if this == nil {
		return nil
	}
	return this.Battr
}

func (this *BidRequest_TextLink) GetBattr() []int {
	if this == nil {
		return nil
	}
	return this.Battr
}

func (this *BidRequest_Video) GetBattr() []int {
	if this == nil {
		return nil
	}
	return this.Battr
}

func (this *BidRequest_Splash) GetBattr() []int {
	if this == nil {
		return nil
	}
	return this.Battr
}

func (this *BidRequest_Splash) GetW() int {
	if this == nil {
		return 0
	}
	return this.W
}

func (this *BidRequest_Splash) GetH() int {
	if this == nil {
		return 0
	}
	return this.H
}

func (this *BidRequest_App) GetId() string {
	if this == nil {
		return ""
	}
	return this.Id
}

func (this *BidRequest_App) GetName() string {
	if this == nil {
		return ""
	}
	return this.Name
}

func (this *BidRequest_App) GetCat() []string {
	if this == nil {
		return nil
	}
	return this.Cat
}

func (this *BidRequest_App) GetVer() string {
	if this == nil {
		return ""
	}
	return this.Ver
}

func (this *BidRequest_App) GetPaid() int {
	if this == nil {
		return 0
	}
	return this.Paid
}

func (this *BidRequest_App) GetStoreUrl() string {
	if this == nil {
		return ""
	}
	return this.StoreUrl
}

func (this *BidRequest_App) GetKeywords() string {
	if this == nil {
		return ""
	}
	return this.Keywords
}

func (this *BidRequest_App) GetPublisher() *BidRequest_Publisher {
	if this == nil {
		return nil
	}
	return this.Publisher
}

func (this *BidRequest_Publisher) GetId() string {
	if this == nil {
		return ""
	}
	return this.Id
}

func (this *BidRequest_Publisher) GetName() string {
	if this == nil {
		return ""
	}
	return this.Name
}

func (this *BidRequest_Device) GetW() int {
	if this == nil {
		return 0
	}
	return this.W
}

func (this *BidRequest_Device) GetH() int {
	if this == nil {
		return 0
	}
	return this.H
}

//
