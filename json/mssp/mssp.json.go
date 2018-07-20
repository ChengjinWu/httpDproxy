package mssp

type Request_Impression struct {
	Impid       string   `json:"impid"`
	Bidfloor    int32    `json:"bidfloor"`
	Bidfloorcur string   `json:"bidfloorcur"`
	W           int32    `json:"w"`
	H           int32    `json:"h"`
	Pos         int32    `json:"pos"`
	Btype       []string `json:"btype"`
	Battr       []string `json:"battr"`
	Instl       uint32   `json:"instl"`
}

type Request_App struct {
	Aid      string   `json:"aid"`
	Name     string   `json:"name"`
	Cat      []string `json:"cat"`
	Ver      string   `json:"ver"`
	Bundle   string   `json:"bundle"`
	Itid     string   `json:"itid"`
	Paid     uint32   `json:"paid"`
	Storeurl string   `json:"storeurl"`
	Keywords string   `json:"Keywords"`
	Pid      string   `json:"Pid"`
	Pub      string   `json:"pub"`
}

type Request_Device struct {
	Did            string `json:"did"`
	Dpid           string `json:"dpid"`
	Ua             string `json:"ua"`
	Ip             string `json:"ip"`
	Country        string `json:"country"`
	Carrier        string `json:"carrier"`
	Language       string `json:"language"`
	Brand          string `json:"make"`
	Model          string `json:"model"`
	Os             string `json:"os"`
	Osv            string `json:"OSV"`
	Connectiontype uint32 `json:"connectiontype"`
	Devicetype     uint32 `json:"devicetype"`
	Loc            string `json:"loc"`
	Sw             uint32 `json:"sw"`
	Sh             uint32 `json:"sh"`
	Orientation    uint32 `json:"orientation"`
}

type Request struct {
	Id     string                `json:"id"`
	Imp    []*Request_Impression `json:"imp"`
	App    *Request_App          `json:"app"`
	Device *Request_Device       `json:"device"`
	Bcat   []string              `json:"bcat"`
	Badv   []string              `json:"badv"`
}

type Response_Image struct {
	Url    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Response_Bid struct {
	Impid   string   `json:"impid"`
	Price   int      `json:"price"`
	Adid    string   `json:"adid"`
	Nurl    string   `json:"nurl"`
	Adp     string   `json:"adp"`
	Adm     string   `json:"adm"`
	Adw     uint32   `json:"adw"`
	Adh     uint32   `json:"adh"`
	Iurl    string   `json:"iurl"`
	Curl    string   `json:"curl"`
	Cturl   []string `json:"cturl"`
	Cid     string   `json:"cid"`
	Crid    string   `json:"crid"`
	Ctype   int      `json:"ctype"`
	Cbundle string   `json:"cbundle"`
	Attr    []int    `json:"attr"`
	Adomain string   `json:"adomain"`
	//the rest fields are add by wanggaoming
	Icon    *Response_Image `json:"icon"`
	Main    *Response_Image `json:"main"`
	Title   string          `json:"title"`
	Desc    string          `json:"description"`
	Cta     string          `json:"cta"`
	Rating  string          `json:"rating"`
	CPrice  int             `json:"cprice"`
	IPrice  int             `json:"iprice"`
	Charge  int             `json:"charge"`
	PkgName string          `json:"pkgname"`
	Target  string          `json:"target"`
}

type Response_SeatBid struct {
	Bid  []*Response_Bid `json:"bid"`
	Seat string          `json:"seat"`
}

type Response struct {
	Id      string              `json:"id"`
	Bidid   string              `json:"bidid"`
	Nbr     int                 `json:"nbr"`
	Seatbid []*Response_SeatBid `json:"seatbid"`
}

type WinNotice_Request struct {
	Id    string `json:"id"`
	Bidid string `json:"bidid"`
	Impid string `json:"impid"`
	Price string `json:"price"`
}

func (m *Request) GetDevice() *Request_Device {
	if m != nil {
		return m.Device
	}
	return nil
}
func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetDeviceIp() string {
	if m != nil && m.Device != nil {
		return m.Device.Ip
	}
	return ""
}

func (m *Request) GetDeviceUa() string {
	if m != nil && m.Device != nil {
		return m.Device.Ua
	}
	return ""
}

func (m *Request) GetDeviceLang() string {
	if m != nil && m.Device != nil {
		return m.Device.Language
	}
	return ""
}

func (m *Request) GetImpression() []*Request_Impression {
	if m != nil {
		return m.Imp
	}
	return nil
}

func (m *Request) GetBCat() []string {
	if m != nil {
		return m.Bcat
	}
	return nil
}

func (m *Request) GetBAdv() []string {
	if m != nil {
		return m.Badv
	}
	return nil
}

func (m *Request_Impression) GetImpId() string {
	if m != nil {
		return m.Impid
	}
	return ""
}

func (m *Request_Impression) GetBidFloor() int32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0
}

func (m *Request_Impression) GetWidth() int32 {
	if m != nil {
		return m.W
	}
	return 0
}

func (m *Request_Impression) GetHeight() int32 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *Request_Impression) GetBtype() []string {
	if m != nil {
		return m.Btype
	}
	return nil
}

func (m *Request_Impression) GetInstl() uint32 {
	if m != nil {
		return m.Instl
	}
	return 0
}

func (m *Request_Impression) GetBattr() []string {
	if m != nil {
		return m.Battr
	}
	return nil
}

func (m *Request) GetDeviceDid() string {
	if m != nil && m.Device != nil {
		return m.Device.Did
	}
	return ""
}

func (m *Request) GetDeviceDpId() string {
	if m != nil && m.Device != nil {
		return m.Device.Dpid
	}
	return ""
}

func (m *Request) GetDeviceCountry() string {
	if m != nil && m.Device != nil {
		return m.Device.Country
	}
	return ""
}

func (m *Request) GetDeviceCarrier() string {
	if m != nil && m.Device != nil {
		return m.Device.Carrier
	}
	return ""
}

func (m *Request) GetDeviceBrand() string {
	if m != nil && m.Device != nil {
		return m.Device.Brand
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
	return 0
}

func (m *Request) GetDeviceConnType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Connectiontype
	}
	return 0
}

func (m *Request) GetDeviceLoc() string {
	if m != nil && m.Device != nil {
		return m.Device.Loc
	}
	return ""
}

func (m *Request) GetAppAid() string {
	if m != nil && m.App != nil {
		return m.App.Aid
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

func (m *Request) GetAppPid() string {
	if m != nil && m.App != nil {
		return m.App.Pid
	}
	return ""
}

func (m *Request) GetAppPub() string {
	if m != nil && m.App != nil {
		return m.App.Pub
	}
	return ""
}
