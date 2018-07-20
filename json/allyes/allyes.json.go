package allyes

type Request_Site_Publisher struct {
	Id string `json:"id"`
}

type Request_Site struct {
	Id             string                 `json:"id"`
	Domain         string                 `json:"domain"`
	Sectioncat     []string               `json:"sectioncat"`
	Page           string                 `json:"page"`
	Publisher      Request_Site_Publisher `json:"publisher"`
	Allyessitetype string                 `json:"allyessitetype"`
	Allyespageform string                 `json:"allyespageform"`
}

type Request_Impression_Banner struct {
	W              uint32 `json:"w"`
	H              uint32 `json:"h"`
	Pos            uint32 `json:"pos"`
	Allyesadformat []int  `json:"allyesadformat"`
	Allyesadform   string `json:"allyesadform"`
}

type Request_Impression struct {
	Id          string                     `json:"id"`
	Banner      *Request_Impression_Banner `json:"banner"`
	Tagid       string                     `json:"tagid"`
	Bidfloor    float32                    `json:"bidfloor"`
	Bidfloorcur string                     `json:"bidfloorcur"`
}

type Request_Device struct {
	Dnt        int    `json:"dnt"`
	Ua         string `json:"ua"`
	Ip         string `json:"ip"`
	Language   string `json:"language"`
	Js         uint32 `json:"js"`
	Devicetype uint32 `json:"devicetype"`
}

type Request_User struct {
	Id   string `json:"id"`
	Cver uint32 `json:"cver"`
}

type Request struct {
	Id     string                `json:"id"`
	Tmax   int                   `json:"tmax"`
	Site   *Request_Site         `json:"site"`
	Device *Request_Device       `json:"device"`
	User   *Request_User         `json:"user"`
	Imp    []*Request_Impression `json:"imp"`
	Cur    []string              `json:"cur"`
	Bcat   []string              `json:"bcat"`
	Badv   []string              `json:"badv"`
}

type Response_Bid struct {
	Id     string  `json:"id"`
	Impid  string  `json:"impid"`
	Price  float32 `json:"price"`
	Curl   string  `json:"curl"`
	Cmflag int     `json:"cmflag"`
	Adid   string  `json:"adid"`
	Surl   string  `json:"surl"`
	Crid   string  `json:"crid"`
}

type Response_SeatBid struct {
	Bid []*Response_Bid `json:"bid"`
}

type Response struct {
	Id      string              `json:"id"`
	Seatbid []*Response_SeatBid `json:"seatbid"`
	Bidid   string              `json:"bidid"`
	Cur     string              `json:"cur"`
}

func (m *Request) GetImpression() []*Request_Impression {
	if m != nil {
		return m.Imp
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

func (m *Request) GetSite() *Request_Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

func (m *Request_User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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

func (m *Request_Site) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *Request_Impression) GetBanner() *Request_Impression_Banner {
	if m != nil {
		return m.Banner
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

func (m *Request_Impression_Banner) GetCreativeType() []int {
	if m != nil {
		return m.Allyesadformat
	}
	return nil
}

func (m *Request_Impression_Banner) GetPos() uint32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *Request_Impression) GetBidfloor() float32 {
	if m != nil {
		return m.Bidfloor
	}
	return 0.0
}

func (m *Request_Device) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}
