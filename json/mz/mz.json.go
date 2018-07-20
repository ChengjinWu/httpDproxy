package mz

type Request_Content_Ext struct {
	Channel   string `json:"channel"`
	Cs        string `json:"cs"`
	CopyRight int    `json:"copyright"`
	Quality   int    `json:"quality"`
}

type Request_Content struct {
	Title    string              `json:"title"`
	Keywords string              `json:"keywords"`
	Ext      Request_Content_Ext `json :"ext"`
}

type Request_Site struct {
	Name    string          `json:"name"`
	Page    string          `json:"page"`
	Ref     string          `json:"ref"`
	Content Request_Content `json:"content"`
	Cat     []string        `json:"cat"`
}

type Request_Impression_Banner struct {
	W     int32    `json:"w"`
	H     int32    `json:"h"`
	Pos   int32    `json:pos`
	Mimes []string `json:mimes`
}

type Request_Impression_Video struct {
	Mimes       []string `json:"mimes"`
	Linearity   int32    `json:"linearity"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	W           int32    `json:"w"`
	H           int32    `json:"h"`
}

type Request_Impression_Pmp struct {
	Deals          []*Request_Impression_Pmp_Deal `json:"deals"`
	PrivateAuction uint32                         `json:"private_auction"`
}

type Request_Impression_Pmp_Deal struct {
	Id       string   `json:"id"`
	Bidfloor int      `json:"bidfloor"`
	Wseat    []string `json:"wseat"`
	At       int      `json:"at"`
}

type Request_Impression_Ext struct {
	ShowType         int `json:"showtype"`
	HasWinNotice     int `json:"has_winnotice"`
	HasClickThrought int `json:"has_clickthrough"`
}

type Native_Request_Asset struct {
	Id       int                         `json:"id"`
	Required bool                        `json:"required"`
	Title    *Native_Request_Asset_Title `json:"title"`
	Img      *Native_Request_Asset_Img   `json:"img"`
	Video    *Native_Request_Asset_Video `json:"video"`
	Data     *Native_Request_Asset_Data  `json:"data"`
}

type Native_Request_Asset_Title struct {
	Len int `json:"len"`
}

type Native_Request_Asset_Img struct {
	Type   int      `json:"type"`
	ImgNum int      `json:"img_num"`
	W      int32    `json:"w"`
	H      int32    `json:"h"`
	Mimes  []string `json:"mimes"`
}

type Native_Request_Asset_Video struct {
	Mimes       []string `json:"mimes"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	W           int32    `json:"w"`
	H           int32    `json:"h"`
}

type Native_Request_Asset_Data struct {
	Type int `json:"type"`
	Len  int `json:"len"`
}

type Native_Request struct {
	Assets []*Native_Request_Asset `json:"assets"`
}

type Native_Response_Asset struct {
	Id       int                          `json:"id"`
	Required bool                         `json:"required"`
	Title    *Native_Response_Asset_Title `json:"title,omitempty"`
	Img      *Native_Response_Asset_Img   `json:"img,omitempty"`
	Video    *Native_Response_Asset_Video `json:"video,omitempty"`
	Data     *Native_Response_Asset_Data  `json:"data,omitempty"`
}

type Native_Response_Asset_Title struct {
	Text string `json:"text"`
}

type Native_Response_Asset_Img struct {
	Type int      `json:"type"`
	Urls []string `json:"urls"`
	W    int32    `json:"w"`
	H    int32    `json:"h"`
}

type Native_Response_Asset_Video struct {
	Url      string `json:"url"`
	Cover    string `json:"cover_img_url"`
	W        string `json:"w"`
	H        string `json:"h"`
	Duration string `json:"duration"`
}

type Native_Response_Asset_Data struct {
	Value string `json:"value"`
}

type Native_Response_Link struct {
	Url           string   `json:"url"`
	Clicktrackers []string `json:"clicktrackers,omitempty"`
	Action        int      `json:"action"`
}

type NativeAd struct {
	Assets      []*Native_Response_Asset `json:"assets"`
	Link        *Native_Response_Link    `json:"link"`
	Imptrackers []string                 `json:"imptrackers"`
}

type Native_Response struct {
	NativeAd *NativeAd `json:"nativead"`
}

type Request_Impression struct {
	Id       string                     `json:"id"`
	Tagid    string                     `json:"tagid"`
	Bidfloor float32                    `json:"bidfloor"`
	Banner   *Request_Impression_Banner `json:"banner"`
	Video    *Request_Impression_Video  `json:"video"`
	NativeAd *Native_Request            `json:"nativead"`
	Pmp      *Request_Impression_Pmp    `json:"pmp"`
	Ext      *Request_Impression_Ext    `json:"ext"`
	Secure   int                        `json:"secure"`
}

type Device_Ext struct {
	Idfa         string `json:"idfa"`
	Idfamd5      string `json:"idfamd5"`
	Mac          string `json:"mac"`
	Macmd5       string `json:"macmd5"`
	Macsha1      string `json:"macsha1"`
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
	Didsha1        string       `json:"didsha1"`
	Dpidmd5        string       `json:"dpidmd5"`
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

type Request_User_Ext struct {
	Models []string `json:"models"`
}

type Request_User struct {
	Id  string            `json:"id"`
	Ext *Request_User_Ext `json:"ext"`
}

type App_Ext struct {
	Sdk    string `json:"sdk"`
	Market uint32 `json:"market"`
}

type Request_App struct {
	Name    string           `json:"name"`
	Bundle  string           `json:"bundle"`
	Itid    string           `json:"itid"`
	Content *Request_Content `json:"content"`
	Cat     []string         `json:"cat"`
	Ext     *App_Ext         `json:"ext"`
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
	Id     string        `json:"id"`
	Impid  string        `json:"impid"`
	Price  float32       `json:"price"`
	Nurl   string        `json:"nurl"`
	Adm    string        `json:"adm"`
	Crid   string        `json:"crid"`
	Pvm    []string      `json:"pvm,omitempty"`
	Dealid string        `json:"dealid"`
	Ext    *Response_Ext `json:"ext,omitempty"`
}

type Response_SeatBid struct {
	Bid []*Response_Bid `json:"bid"`
}

type Response struct {
	Id      string              `json:"id"`
	Bidid   string              `json:"bidid"`
	Seatbid []*Response_SeatBid `json:"seatbid"`
}

func (m *Request) GetSite() *Request_Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Request) GetApp() *Request_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Request) GetUser() *Request_User {
	if m != nil {
		return m.User
	}
	return nil
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

func (m *Request_Device) GetUa() string {
	if m != nil {
		return m.Ua
	}
	return ""
}

func (m *Request_Device) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Request_Device) GetMac() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Mac
	}
	return ""
}

func (m *Request_Device) GetMacmd5() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Macmd5
	}
	return ""
}

func (m *Request_Device) GetMacsha1() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Macsha1
	}
	return ""
}

func (m *Request_Device) GetIdfa() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Idfa
	}
	return ""
}

func (m *Request_Device) GetIdfamd5() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Idfamd5
	}
	return ""
}

func (m *Request_Device) GetDidmd5() string {
	if m != nil {
		return m.Didmd5
	}
	return ""
}

func (m *Request_Device) GetDpidmd5() string {
	if m != nil {
		return m.Dpidmd5
	}
	return ""
}

func (m *Request_Device) GetDidsha1() string {
	if m != nil {
		return m.Didsha1
	}
	return ""
}

func (m *Request_Device) GetCarrier() string {
	if m != nil {
		return m.Carrier
	}
	return ""
}

func (m *Request_Device) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

func (m *Request_Device) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Request_Device) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *Request_Device) GetOsv() string {
	if m != nil {
		return m.Osv
	}
	return ""
}

func (m *Request_Device) GetConnType() uint32 {
	if m != nil {
		return m.Connectiontype
	}
	return 0
}

func (m *Request_Device) GetGeoLat() float64 {
	if m != nil && m.Geo != nil {
		return m.Geo.Lat
	}
	return 0.0
}

func (m *Request_Device) GetGeoLon() float64 {
	if m != nil && m.Geo != nil {
		return m.Geo.Lon
	}
	return 0.0
}

func (m *Request_Device) GetDeviceType() int32 {
	if m != nil {
		return int32(m.Devicetype)
	}
	return -1
}

func (m *Request) GetAppName() string {
	if m != nil && m.App != nil {
		return m.App.Name
	}
	return ""
}

func (m *Request) GetAppid() string {
	if m != nil && m.App != nil {
		if m.App.Itid != "" {
			return m.App.Itid
		} else {
			return m.App.Bundle
		}
	}
	return ""
}

func (m *Request) GetAppCats() []string {
	if m != nil && m.App != nil && m.App.Ext != nil {
		return m.App.Cat
	}
	return nil
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

func (m *Request) GetSitePage() string {
	if m != nil && m.Site != nil {
		return m.Site.Page
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

func (m *Request_Impression) GetNativeAd() *Native_Request {
	if m != nil {
		return m.NativeAd
	}
	return nil
}

func (m *Request_Impression) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_Impression) GetTagid() string {
	if m != nil {
		return m.Tagid
	}
	return ""
}

func (m *Request_Impression) GetSecure() bool {
	if m != nil {
		return m.Secure == 1
	}
	return false
}

func (m *Request_Impression_Banner) GetWidth() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Impression_Banner) GetHeight() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *Request_Impression_Banner) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Request_Impression_Banner) GetMimes() []string {
	if m != nil {
		return m.Mimes
	}
	return nil
}

func (m *Request_Impression_Video) GetWidth() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Impression_Video) GetHeight() int32 {
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

func (m *Request_Impression_Video) GetVideoLinearity() int32 {
	if m != nil {
		return m.Linearity
	}
	return 1
}

func (m *Request_Impression_Video) GetMinduration() int32 {
	if m != nil {
		return m.Minduration
	}
	return 0
}

func (m *Request_Impression_Video) GetMaxduration() int32 {
	if m != nil {
		return m.Maxduration
	}
	return 0
}

func (m *Request_Impression) GetBidfloor() float32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0.0
}

func (m *Request_Impression) GetDeals() []*Request_Impression_Pmp_Deal {
	if m != nil && m.Pmp != nil {
		return m.Pmp.Deals
	}
	return nil
}

func (m *Request) GetScreenWidth() int {
	if m != nil && m.Device != nil && m.Device.Ext != nil {
		return int(m.Device.Ext.W)
	}
	return 0
}

func (m *Request) GetScreenHeight() int {
	if m != nil && m.Device != nil && m.Device.Ext != nil {
		return int(m.Device.Ext.H)
	}
	return 0
}

func (m *Request) GetLanguage() string {
	if m != nil && m.Device != nil {
		return m.Device.Language
	}
	return ""
}

func (m *Native_Request) GetAssets() []*Native_Request_Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Native_Request_Asset) GetTitle() *Native_Request_Asset_Title {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *Native_Request_Asset) GetImg() *Native_Request_Asset_Img {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *Native_Request_Asset) GetData() *Native_Request_Asset_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Native_Request_Asset) GetVideo() *Native_Request_Asset_Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Native_Request_Asset) GetID() int {
	if m != nil {
		return m.Id
	}
	return -1
}

func (m *Native_Request_Asset) IsRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (img *Native_Request_Asset_Img) GetWidth() int32 {
	if img != nil {
		return img.W
	}
	return 0
}

func (img *Native_Request_Asset_Img) GetHeight() int32 {
	if img != nil {
		return img.H
	}
	return 0
}

func (v *Native_Request_Asset_Video) GetW() int32 {
	if v != nil {
		return v.W
	}
	return 0
}

func (v *Native_Request_Asset_Video) GetH() int32 {
	if v != nil {
		return v.H
	}
	return 0
}

func (v *Native_Request_Asset_Video) GetMaxduration() int32 {
	if v != nil {
		return v.Maxduration
	}
	return 0
}

func (v *Native_Request_Asset_Video) GetMinduration() int32 {
	if v != nil {
		return v.Minduration
	}
	return 0
}

func (d *Request_Impression_Pmp_Deal) GetId() string {
	if d != nil {
		return d.Id
	}
	return ""
}

func (m *Request) GetUserCategory() []string {
	if m != nil && m.User != nil && m.User.Ext != nil {
		return m.User.Ext.Models
	}
	return nil
}
