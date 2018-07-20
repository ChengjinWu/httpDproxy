package qihu

type Request_Impression struct {
	Impid       string                     `json:"id"`
	Banner      *Request_Impression_Banner `json:"banner"`
	Instl       uint32                     `json:"instl"`
	Bidfloor    float32                    `json:"bidfloor"`
	Bidfloorcur string                     `json:"bidfloorcur"`
}

type Request_Impression_Banner struct {
	Width  uint32   `json:"w"`
	Height uint32   `json:"h"`
	Btype  []string `json:"btype"`
	Battr  []string `json:"battr"`
}

type Request_App struct {
	Aid      string   `json:"id"`
	Name     string   `json:"name"`
	Cat      []string `json:"cat"`
	Ver      string   `json:"ver"`
	Bundle   string   `json:"bundle"`
	Paid     uint32   `json:"paid"`
	Keywords string   `json:"keywords"`
}

type Request_Device struct {
	Did            string `json:"didmd5"`
	Dpid           string `json:"dpidmd5"`
	Mac            string `json:"macmd5"`
	Ua             string `json:"ua"`
	Ip             string `json:"ip"`
	Lat            string `json:"lat"`
	Long           string `json:"lon"`
	Carrier        string `json:"carrier"`
	Language       string `json:"language"`
	Brand          string `json:"make"`
	Model          string `json:"model"`
	Os             string `json:"os"`
	Osv            string `json:"OSV"`
	Connectiontype uint32 `json:"connectiontype"`
	Devicetype     uint32 `json:"devicetype"`
}

type Request struct {
	Id     string                `json:"id"`
	Imp    []*Request_Impression `json:"imp"`
	App    *Request_App          `json:"app"`
	Device *Request_Device       `json:"device"`
	At     uint32                `json:"at"`
	Tmax   uint32                `json:"tmax"`
	Bcat   []string              `json:"bcat"`
	Badv   []string              `json:"badv"`
}

type Response_Bid struct {
	Impid   string   `json:"impid"`
	Price   float32  `json:"price"`
	Adid    string   `json:"adid"`
	Nurl    string   `json:"nurl"`
	Adm     string   `json:"adm"`
	Adomain []string `json:"adomain"`
	Iurl    string   `json:"iurl"`
	Cid     string   `json:"cid"`
	Crid    string   `json:"crid"`
	Attr    []int    `json:"attr"`
}

type Response_SeatBid struct {
	Bid  []*Response_Bid `json:"bid"`
	Seat string          `json:"seat"`
}

type Response struct {
	Id      string              `json:"id"`
	Seatbid []*Response_SeatBid `json:"seatbid"`
	Bidid   string              `json:"bidid"`
	Nbr     int                 `json:"nbr"`
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

func (m *Request_Impression) GetBannerWidth() uint32 {
	if m != nil && m.Banner != nil {
		return m.Banner.Width
	}
	return 0
}

func (m *Request_Impression) GetBannerHeight() uint32 {
	if m != nil && m.Banner != nil {
		return m.Banner.Height
	}
	return 0
}

func (m *Request_Impression) GetBidFloor() float32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0.0
}

func (m *Request_Impression) GetBtype() []string {
	if m != nil && m.Banner != nil {
		return m.Banner.Btype
	}
	return nil
}

func (m *Request_Impression) GetBattr() []string {
	if m != nil && m.Banner != nil {
		return m.Banner.Battr
	}
	return nil
}

func (m *Request_Impression) GetInstl() uint32 {
	if m != nil {
		return m.Instl
	}
	return 0
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

func (m *Request) GetDeviceLat() string {
	if m != nil && m.Device != nil {
		return m.Device.Lat
	}
	return ""
}

func (m *Request) GetDeviceLon() string {
	if m != nil && m.Device != nil {
		return m.Device.Long
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

func (m *Request) GetAppKeywords() string {
	if m != nil && m.App != nil {
		return m.App.Keywords
	}
	return ""
}
