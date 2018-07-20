package letv

const defaultDevType = 2

type Request_Site_Content_Ext struct {
	Channel string `json:"channel"`
	Cs      string `json:"cs"`
}

type Request_Site_Content struct {
	Title    string                    `json:"title"`
	Keywords string                    `json:"keywords"`
	Ext      *Request_Site_Content_Ext `json :"ext"`
}

type Request_Site struct {
	Name      string                `json:"name"`
	Page      string                `json:"page"`
	Ref       string                `json:"ref"`
	Copyright int                   `json:"cr"`
	Content   *Request_Site_Content `json:"content"`
}

type Request_Impression_Banner struct {
	W uint32 `json:"w"`
	H uint32 `json:"h"`
}

type Request_Impression_Video struct {
	Mimes       []string `json:"mime"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	W           uint32   `json:"w"`
	H           uint32   `json:"h"`
}

type Request_Impression_Pmp_Deal struct {
	Id       int64    `json:"id"`
	BidFloor int32    `json:"bidfloor"`
	Wseat    []string `json:"wseat"`
}

type Request_Impression_Pmp struct {
	PrivateAuction int32                          `json:"private_auction"`
	Deals          []*Request_Impression_Pmp_Deal `json:"deals"`
}

type Request_Impression_Ext struct {
	ExcludeIndustry []int `json:"excluded_ad_industry"`
}

type Request_Impression struct {
	Id       string                     `json:"id"`
	Adzoneid int32                      `json:"adzoneid"`
	Bidfloor float32                    `json:"bidfloor"`
	Banner   *Request_Impression_Banner `json:"banner"`
	Video    *Request_Impression_Video  `json:"video"`
	Pmp      *Request_Impression_Pmp    `json:"pmp"`
	Display  int                        `json:"display"` //通过广告位定向实现
	Ext      *Request_Impression_Ext
}

type Device_Ext struct {
	Idfa         string `json:"idfa"`
	Mac          string `json:"mac"`
	Macmd5       string `json:"macmd5"`
	Ssid         string `json:"ssid"`
	W            uint32 `json:"w"`
	H            uint32 `json:"h"`
	Brk          uint32 `json:"brk"`
	Ts           int    `json:"ts"`
	Interstitial uint32 `json:"interstitial"`
}

type Geo_Ext struct {
	Accuracy int `json:"accuracy"`
}

type Request_Geo struct {
	Lat float64  `json:"lat"`
	Lon float64  `json:"lon"`
	Ext *Geo_Ext `json:"ext"`
}

type Request_Device struct {
	Ua             string       `json:"ua"`
	Ip             string       `json:"ip"`
	Geo            *Request_Geo `json:"geo"`
	Didmd5         string       `json:"didmd5"`
	Dpidmd5        string       `json:"dpidmd5"`
	Did            string       `json:"did"`
	Dpid           string       `json:"dpid"`
	Make           string       `json:"make"`
	Model          string       `json:"model"`
	Os             string       `json:"os"`
	Osv            string       `json:"osv"`
	Carrier        string       `json:"carrier"`
	Language       string       `json:"language"`
	Js             uint32       `json:"js"`
	Connectiontype uint32       `json:"connectiontype"`
	Devicetype     uint32       `json:"devicetype"`
	Ext            *Device_Ext  `json:"ext"`
}

type Request_User struct {
	Id     string `json:"id"`
	Gender string `json:"gender"`
	Yob    uint32 `json:"yob"`
}

type App_Ext struct {
	Sdk    string `json:"sdk"`
	Market uint32 `json:"market"`
	Appid  string `json:"appid"`
	Cat    string `json:"cat"`
	Tag    string `json:"tag"`
}

type Request_App struct {
	Name      string   `json:"name"`
	Copyright int      `json:"cr"`
	Ext       *App_Ext `json:"ext"`
}

type Request struct {
	Id     string                `json:"id"`
	Imp    []*Request_Impression `json:"imp"`
	Site   *Request_Site         `json:"site"`
	Device *Request_Device       `json:"device"`
	User   *Request_User         `json:"user"`
	App    *Request_App          `json:"app"`
}

type Response_Ext struct {
	Ldp  string   `json:"ldp"`
	Pm   []string `json:"pm"`
	Cm   []string `json:"cm"`
	Type string   `json:"type"`
}

type Response_Bid struct {
	Id      string        `json:"id"`
	Impid   string        `json:"impid"`
	Price   float32       `json:"price"`
	Nurl    string        `json:"nurl"`
	Adm     string        `json:"adm"`
	Admpara string        `json:"admpara"`
	Ext     *Response_Ext `json:"ext"`
}

type Response_SeatBid struct {
	Bid []*Response_Bid `json:"bid"`
}

type Response struct {
	Id      string              `json:"id"`
	Bidid   string              `json:"bidid"`
	Seatbid []*Response_SeatBid `json:"seatbid"`
}

func (m *Request) GetDevice() *Request_Device {
	if m != nil {
		return m.Device
	}
	return nil
}
func (m *Request) GetDeviceInstl() uint32 {
	if m != nil && m.Device != nil && m.Device.Ext != nil {
		return m.Device.Ext.Interstitial
	}
	return 0
}

func (m *Request) GetImpression() []*Request_Impression {
	if m != nil {
		return m.Imp
	}
	return nil
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetUserId() string {
	if m != nil && m.User != nil {
		return m.User.Id
	}
	return ""
}

func (m *Request) GetUserGender() string {
	if m != nil && m.User != nil {
		return m.User.Gender
	}
	return ""
}

func (m *Request) GetDeviceIp() string {
	if m != nil && m.Device != nil {
		return m.Device.Ip
	}
	return ""
}

func (m *Request) GetDeviceDid() string {
	if m != nil && m.Device != nil {
		return m.Device.Did
	}
	return ""
}

func (m *Request) GetDeviceDidMd5() string {
	if m != nil && m.Device != nil {
		return m.Device.Didmd5
	}
	return ""
}

func (m *Request) GetDeviceDpid() string {
	if m != nil && m.Device != nil {
		return m.Device.Dpid
	}
	return ""
}

func (m *Request) GetDeviceDpidMd5() string {
	if m != nil && m.Device != nil {
		return m.Device.Dpidmd5
	}
	return ""
}

func (m *Request) GetDeviceMac() string {
	if m != nil && m.Device != nil && m.Device.Ext != nil {
		return m.Device.Ext.Mac
	}
	return ""
}

func (m *Request) GetDeviceCarrier() string {
	if m != nil && m.Device != nil {
		return m.Device.Carrier
	}
	return ""
}

func (m *Request) GetDeviceMake() string {
	if m != nil && m.Device != nil {
		return m.Device.Make
	}
	return ""
}

func (m *Request) GetDeviceModel() string {
	if m != nil && m.Device != nil {
		return m.Device.Model
	}
	return ""
}

func (m *Request) GetDeviceOs() string {
	if m != nil && m.Device != nil {
		return m.Device.Os
	}
	return ""
}

func (m *Request) GetDeviceOsv() string {
	if m != nil && m.Device != nil {
		return m.Device.Osv
	}
	return ""
}

func (m *Request) GetDeviceType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Devicetype
	}
	return defaultDevType
}

func (m *Request) GetDeviceConnType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Connectiontype
	}
	return 0
}

func (m *Request) GetDeviceGeoLat() float64 {
	if m != nil && m.Device != nil && m.Device.Geo != nil {
		return m.Device.Geo.Lat
	}
	return 0.0
}

func (m *Request) GetDeviceGeoLon() float64 {
	if m != nil && m.Device != nil && m.Device.Geo != nil {
		return m.Device.Geo.Lon
	}
	return 0.0
}

func (m *Request) GetAppName() string {
	if m != nil && m.App != nil {
		return m.App.Name
	}
	return ""
}

func (m *Request) GetAppid() string {
	if m != nil && m.App != nil && m.App.Ext != nil {
		return m.App.Ext.Appid
	}
	return ""
}

func (m *Request) GetAppCats() string {
	if m != nil && m.App != nil && m.App.Ext != nil {
		return m.App.Ext.Cat
	}
	return ""
}

func (m *Request) GetAppTag() string {
	if m != nil && m.App != nil && m.App.Ext != nil {
		return m.App.Ext.Tag
	}
	return ""
}
func (m *Request) GetDeviceUa() string {
	if m != nil && m.Device != nil {
		return m.Device.Ua
	}
	return ""
}

func (m *Request) GetSiteRef() string {
	if m != nil && m.Site != nil {
		return m.Site.Ref
	}
	return ""
}

func (m *Request) GetLanguage() string {
	if m != nil && m.Device != nil {
		return m.Device.Language
	}
	return ""
}

func (m *Request) GetSitePage() string {
	if m != nil && m.Site != nil {
		return m.Site.Page
	}
	return ""
}

func (m *Request) GetSiteChannel() string {
	if m != nil && m.Site != nil && m.Site.Content != nil && m.Site.Content.Ext != nil {
		return m.Site.Content.Ext.Channel
	}
	return ""
}

func (m *Request) GetSiteCs() string {
	if m != nil && m.Site != nil && m.Site.Content != nil && m.Site.Content.Ext != nil {
		return m.Site.Content.Ext.Cs
	}
	return ""
}

func (m *Request_Impression) GetBanner() *Request_Impression_Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *Request_Impression) GetVideo() *Request_Impression_Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Request_Impression) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_Impression) GetAdzoneid() int32 {
	if m != nil {
		return m.Adzoneid
	}
	return 0
}

func (m *Request_Impression_Banner) GetWidth() uint32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Impression_Banner) GetHeight() uint32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *Request_Impression_Video) GetWidth() uint32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Impression_Video) GetHeight() uint32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *Request_Impression_Video) GetCreativeType() []string {
	if m != nil && len(m.Mimes) > 0 {
		return m.Mimes
	}
	return nil
}

func (m *Request_Impression_Video) GetVideoMinduration() int32 {
	//comment this because letv said that for the moment minduration == maxduration, and minduration can be ignored.
	//if m != nil {
	//	return m.Minduration
	//}
	return -1
}

func (m *Request_Impression_Video) GetVideoMaxduration() int32 {
	if m != nil {
		return m.Maxduration
	}
	return -1
}

func (m *Request_Impression) GetBidfloor() float32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0.0
}

func (m *Request_Impression) GetPmp() *Request_Impression_Pmp {
	if m != nil {
		return m.Pmp
	}
	return nil
}

func (m *Request_Impression) GetDisplay() int {
	if m != nil {
		return m.Display
	}
	return 0
}

func (m *Request_Impression_Pmp) GetPrivateAuction() int32 {
	if m != nil {
		return m.PrivateAuction
	}
	return 0
}

func (m *Request_Impression_Pmp) GetDeals() []*Request_Impression_Pmp_Deal {
	if m != nil {
		return m.Deals
	}
	return nil
}

func (m *Request_Impression_Pmp_Deal) GetDealId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request_Impression_Pmp_Deal) GetDealBidFloor() int32 {
	if m != nil {
		return m.BidFloor
	}
	return 0
}

func (m *Request_Impression_Pmp_Deal) GetDealWseat() []string {
	if m != nil {
		return m.Wseat
	}
	return nil
}

func (m *Request_Device) GetDeviceExt() *Device_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *Request) GetIosIdfa() string {
	if m != nil && m.Device != nil && m.Device.Ext != nil {
		return m.Device.Ext.Idfa
	}
	return ""
}

func (m *Request_Impression) GetBlockingIndustryId() []int {
	if m != nil && m.Ext != nil {
		return m.Ext.ExcludeIndustry
	}
	return nil
}
