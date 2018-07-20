package adview

type Request_Impression struct {
	Id          string                     `json:"id"`
	Banner      *Request_Impression_Banner `json:"banner"`
	Native      *Request_Impression_Native `json:"native"`
	Pmp         *Request_Impression_Pmp    `json:"pmp"`
	Instl       int32                      `json:"instl"`
	Tagid       string                     `json:"tagid"` // 用于发起拍卖的特定广告位置或广告牌的标识。
	Bidfloor    int32                      `json:"bidfloor"`
	BidfloorCur string                     `json:"bidfloorcur"`
}

type Request_Impression_Banner struct {
	W     int32   `json:"w"`
	H     int32   `json:"h"`
	Pos   int32   `json:"pos"`
	Btype []int32 `json:"btype"`
	Battr []int32 `json:"battr"`
	Mimes []int32 `json:"mimes"`
}

type Request_Impression_Native struct {
	Request *Request_Native_Request `json:"request"`
	Ver     string                  `json:"ver"`
	Battr   []int32                 `json:"battr"`
}

type Request_Impression_Pmp struct {
	// 0:PrivateAuction 1:PreferredDeal/PDB
	Privateauction int32               `json:"private_auction"`
	Deals          []*Request_Pmp_Deal `json:"deals"`
}

type Request_Pmp_Deal struct {
	Id          string   `json:"id"`
	Bidfloor    int32    `json:"bidfloor"`
	Bidfloorcur int32    `json:"bidfloorcur"`
	At          int32    `json:"at"`
	Wseat       []string `json:"wseat"`
	Wadomain    []string `json:"wadomain"`
}

type Request_Native_Request struct {
	Ver      string                  `json:"ver"`
	Layout   int32                   `json:"layout"`
	Adunit   int32                   `json:"adunit"`
	Plcmtcnt int32                   `json:"plcmtcnt"`
	Seq      int32                   `json:"seq"`
	Assets   []*Request_Native_Asset `json:"assets"`
}

type Request_Native_Asset struct {
	Id       int32                `json:"id"`
	Required int32                `json:"required"`
	Title    *Request_Asset_Title `json:"title"`
	Img      *Request_Asset_Image `json:"img"`
	Video    *Request_Asset_Video `json:"video"`
	Data     *Request_Asset_Data  `json:"data"`
}

type Request_Asset_Title struct {
	Len int32 `json:"len"`
}

type Request_Asset_Image struct {
	Type  int32    `json:"type"`
	W     int32    `json:"w"`
	Wmin  int32    `json:"wmin"`
	H     int32    `json:"h"`
	Hmin  int32    `json:"hmin"`
	Mimes []string `json:"mimes"`
}

type Request_Asset_Video struct {
	Mimes       []string `json:"mimes"`
	Minduration int32    `json:"minduration"`
	Maxduration int32    `json:"maxduration"`
	Protocols   []int32  `json:"protocols"`
}

type Request_Asset_Data struct {
	Type int32 `json:"type"`
	Len  int32 `json:"len"`
}

type Request_App struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Cat      []int  `json:"cat"`
	Ver      string `json:"ver"`
	Bundle   string `json:"bundle"`
	Paid     int    `json:"paid"`
	Storeurl string `json:"storeurl"`
	Keywords string `json:"Keywords"`
}

type Request_Device struct {
	Dnt            int32               `json:"dnt"`
	Ua             string              `json:"ua"`
	Ip             string              `json:"ip"`
	Geo            *Request_Device_Geo `json:"geo"`
	Idfa           string              `json:"idfa"`
	Didsha1        string              `json:"didsha1"`
	Dpidsha1       string              `json:"dpidsha1"`
	Macsha1        string              `json:"macsha1"`
	Didmd5         string              `json:"didmd5"`
	Dpidmd5        string              `json:"dpidmd5"`
	Macmd5         string              `json:"macmd5"`
	ipv6           string              `json:"ipv6"`
	Carrier        string              `json:"carrier"`
	Language       string              `json:"language"`
	Make           string              `json:"make"`
	Model          string              `json:"model"`
	Os             string              `json:"os"`
	Osv            string              `json:"OSV"`
	Js             int                 `json:"js"`
	Connectiontype int32               `json:"connectiontype"`
	Devicetype     int32               `json:"devicetype"`
	Sw             int32               `json:"sw"`
	Sh             int32               `json:"sh"`
	Orientation    int32               `json:"orientation"`
	Ext            *Request_Device_Ext `json:"ext"`
}

type Request_Device_Geo struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Timestamp int64   `json:"timestamp"`
	Country   string  `json:"country"`
	Region    string  `json:"region"`
	City      string  `json:"city"`
	Geotype   int32   `json:"type"`
}

type Request_Device_Ext struct {
	Idfa      string `json:"idfa"`
	Uuid      string `json:"uuid"`
	AndroidId string `json:"androidId"`
	Mac       string `json:"mac"`
}

type Request_User struct {
	Id       string              `json:"id"`
	Yob      int32               `json:"yob"`
	Gender   string              `json:"gender"`
	Keywords string              `json:"keywords"`
	Geo      *Request_Device_Geo `json:"geo"`
}

type Request struct {
	Id     string                `json:"id"`
	Imp    []*Request_Impression `json:"imp"`
	App    *Request_App          `json:"app"`
	Device *Request_Device       `json:"device"`
	User   *Request_User         `json:"user"`
	At     int32                 `json:"at"`
	tmax   int32                 `json:"tmax"`
	weat   []string              `json:"weat"`
	cur    []string              `json:"cur"`
	Bcat   []string              `json:"bcat"`
	Badv   []string              `json:"badv"`
}

type Response_Bid struct {
	Impid   string              `json:"impid"`
	Price   int32               `json:"price"`
	Adid    string              `json:"adid"`
	Paymode int32               `json:"paymode,omitempty"`
	Adct    int32               `json:"adct"`
	Admt    int32               `json:"admt"`
	Adm     string              `json:"adm,omitempty"`
	Native  *Response_Native    `json:"native,omitempty"`
	Adi     string              `json:"adi,omitempty"`
	Adt     string              `json:"adt,omitempty"`
	Ads     string              `json:"ads,omitempty"`
	Adw     int32               `json:"adw"`
	Adh     int32               `json:"adh"`
	Adlogo  string              `json:"adLogo"`
	Adtm    int32               `json:"adtm,omitempty"`
	Ade     string              `json:"ade,omitempty"`
	Adurl   string              `json:"adurl,omitempty"`
	Adomain string              `json:"adomain"`
	Wurl    string              `json:"wurl,omitempty"`
	Nurl    map[string][]string `json:"nurl,omitempty"`
	Curl    []string            `json:"curl,omitempty"`
	Dealid  string              `json:"dealid,omitempty"`
	Cid     string              `json:"cid,omitempty"`
	Crid    string              `json:"crid,omitempty"`
	Attr    []int               `json:"attr,omitempty"`
}

type Response_Native struct {
	Ver         string                   `json:"ver"`
	Assets      []*Response_Native_Asset `json:"assets"`
	Link        *Response_Native_Link    `json:"link"`
	Imptrackers []string                 `json:"imptrackers"`
	Jstracker   string                   `json:"jstracker,omitempty"`
}

type Response_Native_Asset struct {
	Id       int32                  `json:"id"`
	Required int32                  `json:"required"`
	Title    *Response_Native_Title `json:"title,omitempty"`
	Img      *Response_Native_Image `json:"img,omitempty"`
	Video    *Response_Native_Video `json:"video,omitempty"`
	Data     *Response_Native_Data  `json:"data,omitempty"`
	Link     *Response_Native_Link  `json:"link,omitempty"`
}

type Response_Native_Title struct {
	Text string `json:"text"`
}

type Response_Native_Image struct {
	Url string `json:"url"`
	W   int32  `json:"w"`
	H   int32  `json:"h"`
}

type Response_Native_Video struct {
	Vasttag string `json:"vasttag"`
}

type Response_Native_Data struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Response_Native_Link struct {
	Url           string   `json:"url"`
	Clicktrackers []string `json:"clicktrackers,omitempty"`
	Fallback      string   `json:"fallback,omitempty"`
}

type Response_SeatBid struct {
	Bid  []*Response_Bid `json:"bid"`
	Seat string          `json:"seat,omitempty"`
}

type Response struct {
	Id      string              `json:"id"`
	Bidid   string              `json:"bidid"`
	Nbr     int                 `json:"nbr"` // 不出价原因
	Seatbid []*Response_SeatBid `json:"seatbid"`
}

type WinNotice_Request struct {
	Id    string `json:"id"`
	Bidid string `json:"bidid"`
	Impid string `json:"impid"`
	Price string `json:"price"`
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetImpression() []*Request_Impression {
	if m != nil {
		return m.Imp
	}
	return nil
}

func (m *Request) GetApp() *Request_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Request) GetDevice() *Request_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *Request) GetBAdv() []string {
	if m != nil {
		return m.Badv
	}
	return nil
}

func (m *Request) GetBCat() []string {
	if m != nil {
		return m.Bcat
	}
	return nil
}

// 竞拍类型：
// 0 – 最高价格成交(参与竞价)
// 1 – 以次高价格成交（参与竞价）
// 2 – 优先购买(不参与竞价)
func (m *Request) IsPreferredAuctionType() bool {
	if m != nil && m.At == 2 {
		return true
	}
	return false
}

// Request_Device
func (m *Request_Device) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Request_Device) GetUa() string {
	if m != nil {
		return m.Ua
	}
	return ""
}

func (m *Request_Device) GetLang() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Request_Device) GetGeo() *Request_Device_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

// 原始类型id
func (m *Request_Device) GetImei() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Uuid
	}
	return ""
}

func (m *Request_Device) GetAndroidId() string {
	if m != nil && m.Ext != nil {
		return m.Ext.AndroidId
	}
	return ""
}

func (m *Request_Device) GetIdfa() string {
	if m != nil {
		if m.Idfa != "" {
			return m.Idfa
		}

		if m.Ext != nil {
			return m.Ext.Idfa
		}
	}
	return ""
}

func (m *Request_Device) GetMac() string {
	if m != nil && m.Ext != nil {
		return m.Ext.Mac
	}
	return ""
}

// md5类型id
func (m *Request_Device) GetDidMd5() string {
	if m != nil {
		return m.Didmd5
	}
	return ""
}

func (m *Request_Device) GetDpIdMd5() string {
	if m != nil {
		return m.Dpidmd5
	}
	return ""
}

func (m *Request_Device) GetMacMd5() string {
	if m != nil {
		return m.Macmd5
	}
	return ""
}

// sha1类型id
func (m *Request_Device) GetDidSha1() string {
	if m != nil {
		return m.Didsha1
	}
	return ""
}

func (m *Request_Device) GetDpIdSha1() string {
	if m != nil {
		return m.Dpidsha1
	}
	return ""
}

func (m *Request_Device) GetMacSha1() string {
	if m != nil {
		return m.Macsha1
	}
	return ""
}

func (m *Request_Device) GetCarrier() string {
	if m != nil {
		return m.Carrier
	}
	return ""
}

func (m *Request_Device) GetBrand() string {
	if m != nil {
		return m.Model
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

func (m *Request_Device) GetDeviceType() int32 {
	if m != nil {
		return m.Devicetype
	}
	return 0
}

func (m *Request_Device) GetConnType() int32 {
	if m != nil {
		return m.Connectiontype
	}
	return 0
}

// Request_Impression
func (m *Request_Impression) GetBanner() *Request_Impression_Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *Request_Impression) GetNative() *Request_Impression_Native {
	if m != nil {
		return m.Native
	}
	return nil
}

func (m *Request_Impression) GetPmp() *Request_Impression_Pmp {
	if m != nil {
		return m.Pmp
	}
	return nil
}

func (m *Request_Impression) GetImpId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_Impression) GetTagId() string {
	if m != nil {
		return m.Tagid
	}
	return ""
}

func (m *Request_Impression) GetBidFloor() int32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0
}

func (m *Request_Impression) GetSlotType() int32 {
	if m != nil {
		return m.Instl
	}
	return 0
}

// Request_Impression_Banner
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

func (m *Request_Impression_Banner) GetBtype() []int32 {
	if m != nil {
		return m.Btype
	}
	return nil
}

func (m *Request_Impression_Banner) GetBattr() []int32 {
	if m != nil {
		return m.Battr
	}
	return nil
}

// Request_Impression_Native
func (m *Request_Impression_Native) GetBattr() []int32 {
	if m != nil {
		return m.Battr
	}
	return nil
}

func (m *Request_Impression_Native) GetNativeRequest() *Request_Native_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

// Request_Impression_Pmp
func (m *Request_Impression_Pmp) IsPrivateAuction() bool {
	// 0:PrivateAuction 1:PreferredDeal/PDB
	if m != nil && m.Privateauction == 0 {
		return true
	}
	return false
}

func (m *Request_Impression_Pmp) GetPmpDealIds() []string {
	if m == nil || len(m.Deals) == 0 {
		return nil
	}

	pmpDealIds := make([]string, len(m.Deals))
	for _, deal := range m.Deals {
		pmpDealIds = append(pmpDealIds, deal.Id)
	}
	return pmpDealIds
}

// Request_Device_Geo
func (m *Request_Device_Geo) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0.0
}

func (m *Request_Device_Geo) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0.0
}

func (m *Request_Device_Geo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

// Request_App
func (m *Request_App) GetAppId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request_App) GetAppName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request_App) GetBundle() string {
	if m != nil {
		return m.Bundle
	}
	return ""
}

func (m *Request_App) GetAppCats() []int {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *Request_App) GetAppVersion() string {
	if m != nil {
		return m.Ver
	}
	return ""
}

func (m *Request_App) GetAppPaid() bool {
	if m != nil {
		return m.Paid != 0
	}
	return false
}

func (m *Request_App) GetAppStore() string {
	if m != nil {
		return m.Storeurl
	}
	return ""
}

func (m *Request_App) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Request_App) GetAppKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

// Request_Native_Request
func (m *Request_Native_Request) GetNativeLayout() int32 {
	if m != nil {
		return m.Layout
	}
	return 0
}

func (m *Request_Native_Request) GetNativeAdUnit() int32 {
	if m != nil {
		return m.Adunit
	}
	return 0
}

func (m *Request_Native_Request) GetNativeAssets() []*Request_Native_Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

// Request_Native_Asset
func (m *Request_Native_Asset) GetAssetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 设置为1表示素材必需，此时如果返回的广告缺少对应素材则AdView认为无效。缺省为 0
func (m *Request_Native_Asset) GetAssetRequired() int32 {
	if m != nil {
		return m.Required
	}
	return 0
}

func (m *Request_Native_Asset) IsAssetRequired() bool {
	if m != nil && m.Required == 1 {
		return true
	}
	return false
}

func (m *Request_Native_Asset) GetAssetTitle() *Request_Asset_Title {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *Request_Native_Asset) GetAssetImg() *Request_Asset_Image {
	if m != nil {
		return m.Img
	}
	return nil
}

func (m *Request_Native_Asset) GetAssetVideo() *Request_Asset_Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Request_Native_Asset) GetAssetData() *Request_Asset_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request_Asset_Title
func (m *Request_Asset_Title) GetTitleLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}

// Request_Asset_Image
func (m *Request_Asset_Image) GetImageType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request_Asset_Image) GetWidth() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Asset_Image) GetHeight() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *Request_Asset_Image) GetMimes() []string {
	if m != nil {
		return m.Mimes
	}
	return nil
}

// Request_Asset_Video
func (m *Request_Asset_Video) GetMimes() []string {
	if m != nil {
		return m.Mimes
	}
	return nil
}

func (m *Request_Asset_Video) GetMinDuration() int32 {
	if m != nil {
		return m.Minduration
	}
	return 0
}

func (m *Request_Asset_Video) GetMaxDuration() int32 {
	if m != nil {
		return m.Maxduration
	}
	return 0
}

// Request_Asset_Data
func (m *Request_Asset_Data) GetDataType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request_Asset_Data) GetDataLen() int32 {
	if m != nil {
		return m.Len
	}
	return 0
}
