package meiyou

type Request_Impression_Ext struct {
	Splash int32 `json:"splash"`
}

type Request_Impression struct {
	Id          string                  `json:"id"`
	Adspaceid   string                  `json:"adspaceid"`
	Banner      *Request_Banner         `json:"banner"`
	Native      *Request_Native         `json:"native"`
	Bidfloor    int32                   `json:"bidfloor"`
	Bidfloorcur string                  `json:"bidfloorcur"`
	Instl       uint32                  `json:"instl"`
	Ext         *Request_Impression_Ext `json:"ext"`
}

type Request_Banner struct {
	W     int32    `json:"w"`
	H     int32    `json:"h"`
	Btype []string `json:"btype"`
	Battr []string `json:"battr"`
	Pos   int32    `json:"pos"`
	Mimes []string `json:"mimes"`
	Api   []int    `josn:"api"`
}
type Request_Native struct {
	Layout   int              `json:"layout"`
	Plcmtcnt uint32           `josn::Plcmtcnt"` //not used now
	Assets   []*Request_Asset `json:"assets"`
}
type Request_Asset struct {
	Id       uint32               `json:"id"`
	Required uint32               `json:"required"`
	Title    *Request_Asset_Title `json:"tile"`
	Img      *Request_Asset_Image `json:"img"`
	Data     *Request_Asset_Data  `json:"data"`
}
type Request_Asset_Title struct {
	Len uint32 `json:"len"`
}
type Request_Asset_Image struct {
	Type   int      `json:"type"`
	W      int32    `json:"w"`
	Wmin   int32    `json;"wmin"`
	H      int32    `json:"h"`
	hmin   int32    `json:"hmin"`
	mimies []string `json:"mimes"`
}
type Request_Asset_Data struct {
	Type int    `json:"type"`
	Len  uint32 `json:"len"`
}

type Request_App struct {
	Id       string               `json:"id"`
	Name     string               `json:"name"`
	Cat      []string             `json:"cat"`
	Ver      string               `json:"ver"`
	Bundle   string               `json:"bundle"`
	Paid     uint32               `json:"paid"`
	Storeurl string               `json:"storeurl"`
	Keywords string               `json:"Keywords"`
	Pub      string               `json:"pub"`
	Content  *Request_App_Content `json:"content"`
	Ext      *Request_App_Ext     `json:"ext"`
}
type Request_App_Content struct {
	Title    string   `json:"title"`
	Cat      []string `json:"cat"`
	Url      string   `json:"url"`
	Keywords string   `json:"Keywords"`
}
type Request_App_Ext struct {
	Itid string `json:"itid"`
}

type Request_Device struct {
	Did            string              `json:"did"`
	Dpid           string              `json:"dpid"`
	Mac            string              `json:"mac"`
	Ua             string              `json:"ua"`
	Geo            *Request_Device_Geo `json:"geo"`
	Dnt            int32               `json:"dnt"`
	Ip             string              `json:"ip"`
	Devicetype     uint32              `json:"devicetype"`
	Make           string              `json:"make"`
	Model          string              `json:"model"`
	Os             string              `json:"os"`
	Osv            string              `json:"OSV"`
	Connectiontype uint32              `json:"connectiontype"`
	Carrier        string              `json:"carrier"`
	Language       string              `json:"language"`
	W              uint32              `json:"w"`
	H              uint32              `json:"h"`
	ppi            int                 `json:"ppi"`
	Pxratio        float32             `json:"pxratio"`
	Ext            *Request_Device_Ext `json:"ext"`
}
type Request_Device_Geo struct {
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Type    int32   `json:"type"`
	Country string  `json:"country"`
	Region  string  `json:"region"`
	City    string  `json:"city"`
	zip     string  `json:"zip"`
}
type Request_Device_Ext struct {
	Orientation uint32 `json:"orientation"`
}

type Request_Site struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Domain     string   `json:"domain"`
	Cat        []string `json:"cat"`
	SectionCat []string `json:sectioncat"`
	PagetCat   []string `json:"pagecat"`
	Page       string   `json:"page"`
	Ref        string   `json:"ref"`
	Search     string   `json:"search"`
	Mobile     int      `json:"mobile"`
	Keywords   string   `json:"keywords"`
	Pub        string   `json:"pub"`
}

type Request struct {
	Id     string                `json:"id"`
	Imp    []*Request_Impression `json:"imp"`
	Site   *Request_Site         `json:"site"`
	App    *Request_App          `json:"app"`
	Device *Request_Device       `json:"device"`
	Test   int32                 `json:"test"`
	At     int32                 `json:"at"`
	Tmax   int32                 `json:"tmax"`
	Bcat   []string              `json:"bcat"`
	Wcat   []string              `json:"wcat"`
	Badv   []string              `json:"badv"`
}

type Response_Bid_Link_Ext struct {
	Type int `json:"type"`
}
type Response_Bid_Link struct {
	Url          string                 `json:"url"`
	ClickTracker []string               `json:"clicktracker,omitempty"`
	Ext          *Response_Bid_Link_Ext `json:"ext"`
}

type Response_Bid_Ext struct {
	Link    *Response_Bid_Link `json:"link"`
	Instl   uint32             `json:"instl,omitempty"`
	Adt     uint32             `json:"adt,omitempty"`
	Ade     string             `json:"ade,omitempty"`
	ApkName string             `json:"apkname,omitempty"`
	Aurl    string             `json:"aurl,omitempty"`
}
type Response_Bid struct {
	Id      string            `json:"id"`
	Impid   string            `json:"impid"`
	Price   int               `json:"price"`
	Adid    string            `json:"adid"`
	Nurl    string            `json:"nurl,omitempty"`
	Adm     interface{}       `json:"adm"` //This might be Response_Banner or Response_Native
	W       uint32            `json:"w,omitempty"`
	H       uint32            `json:"h,omitempty"`
	Iurl    string            `json:"iurl,omitempty"`
	Cid     string            `json:"cid,omitempty"`
	Crid    string            `json:"crid,omitempty"`
	Cat     int               `json:"cat,omitempty"`
	bundle  string            `json:"bundle,omitempty"`
	Attr    []int             `json:"attr,omitempty"`
	Adomain string            `json:"adomain,omitempty"`
	Ext     *Response_Bid_Ext `json:"ext,omitempty"`
}

type Response_Banner struct {
	CrtType int32  `json:"crt_type"`
	Img     string `json:"img,omitempty"`
	Txt     string `json:"txt,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Html    string `json:"html,omitempty"`
	HtmlUrl string `json:"html_url,omitempty"`
}

type Response_Native struct {
	Assets     []*Response_Native_Asset `json:"assets"`
	Link       *Response_Bid_Link       `json:"link,omitempty"`
	ImpTracker []string                 `json:"imptracker,omitempty"`
}

type Response_Native_Asset struct {
	Id    int                          `json:"id"`
	Title *Response_Native_Asset_Title `json:"title,omitempty"`
	Img   *Response_Native_Asset_Image `json:"img,omitempty"`
	Data  *Response_Native_Asset_Data  `json:"data,omitempty"`
	Link  *Response_Bid_Link           `json:"link,omitempty"`
}

type Response_Native_Asset_Title struct {
	Text string `json:"text"`
}
type Response_Native_Asset_Image struct {
	Url string `json:"url"`
	W   uint32 `json:"w,omitempty"`
	H   uint32 `json:"h,omitempty"`
}
type Response_Native_Asset_Data struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value"`
}
type Response_SeatBid struct {
	Bid  []*Response_Bid `json:"bid"`
	Seat string          `json:"seat,omitempty"`
}

type Response struct {
	Id      string              `json:"id"`
	Bidid   string              `json:"bidid,omitempty"`
	Nbr     int                 `json:"nbr,omitempty"`
	Seatbid []*Response_SeatBid `json:"seatbid,omitempty"`
}

type WinNotice_Request struct {
	Id    string `json:"id"`
	Bidid string `json:"bidid"`
	Impid string `json:"impid"`
	Price string `json:"price"`
}

func (m *Request) GetBAdv() []string {
	if m != nil {
		return m.Badv
	}
	return nil
}

func (m *Request) GetImp() []*Request_Impression {
	if m == nil {
		return nil
	}
	return m.Imp
}

func (m *Request_Impression) GetId() string {
	if m == nil {
		return ""
	}
	return m.Id
}
func (m *Request_Impression) GetBanner() *Request_Banner {
	if m == nil {
		return nil
	}
	return m.Banner
}
func (m *Request_Impression) GetBattr() []string {
	if m != nil && m.Banner != nil {
		return m.Banner.Battr
	}
	return nil
}
func (m *Request_Impression) GetBidFloor() int32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0
}

func (m *Request) GetDevice() *Request_Device {
	if m == nil {
		return nil
	}
	return m.Device
}
func (m *Request) GetDeviceType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Devicetype
	}
	return 0
}
func (m *Request) GetDeviceLanguage() string {
	if m != nil && m.Device != nil {
		return m.Device.Language
	}
	return ""
}

func (m *Request) GetDeviceIp() string {
	if m != nil && m.Device != nil {
		return m.Device.Ip
	}
	return ""
}
func (m *Request) GetDeviceLang() string {
	if m != nil && m.Device != nil {
		return m.Device.Language
	}
	return ""
}

func (m *Request) GetDeviceDid() string {
	if m == nil || m.Device == nil {
		return ""
	}
	return m.Device.Dpid
}

func (m *Request) GetDeviceDpid() string {
	if m == nil || m.Device == nil {
		return ""
	}
	return m.Device.Dpid
}
func (m *Request) GetDeviceGeo() *Request_Device_Geo {
	if m == nil || m.Device == nil {
		return nil
	}
	return m.Device.Geo
}

func (m *Request) GetDeviceCountry() string {
	if m != nil && m.Device != nil && m.Device.Geo != nil {
		return m.Device.Geo.Country
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

func (m *Request) GetDeviceConnType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Connectiontype
	}
	return 0
}

func (m *Request) GetAppId() string {
	if m != nil && m.App != nil {
		return m.App.Id
	}
	return ""
}

func (m *Request) GetAppName() string {
	if m != nil && m.App != nil {
		return m.App.Name
	}
	return ""
}

func (m *Request) GetAppCats() []string {
	if m != nil && m.App != nil {
		return m.App.Cat
	}
	return nil
}

func (m *Request) GetAppVersion() string {
	if m != nil && m.App != nil {
		return m.App.Ver
	}
	return ""
}

func (m *Request) GetAppPaid() bool {
	if m != nil && m.App != nil {
		return m.App.Paid == 1
	}
	return false
}

func (m *Request) GetAppStore() string {
	if m != nil && m.App != nil {
		return m.App.Storeurl
	}
	return ""
}

func (m *Request) GetAppKeywords() string {
	if m != nil && m.App != nil {
		return m.App.Keywords
	}
	return ""
}

func (m *Request) GetAppPub() string {
	if m != nil && m.App != nil {
		return m.App.Pub
	}
	return ""
}
func (m *Request_Impression) GetBTypes() []string {
	if m == nil || m.Banner == nil {
		return nil
	}
	return m.Banner.Btype
}
func (m *Request_Impression) GetAdspaceId() string {
	if m == nil {
		return ""
	}
	return m.Adspaceid
}

func (m *Request) GetDeviceUa() string {
	if m != nil || m.Device != nil {
		return m.Device.Ua
	}
	return ""
}

/*Imp    []*Request_Impression `json:"imp"`
Site   *Request_Site         `json:"site"`
App    *Request_App          `json:"app"`
Device *Request_Device       `json:"device"`
Test   int32                 `json:"test"`
*/
