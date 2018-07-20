package xt

type RequestImp struct {
	ImpId  string `json:"impid"`
	Width  int32  `json:"w"`
	Height int32  `json:"h"`
	Pos    int32  `json:"pos"`
}

type RequestApp struct {
	Name     string `json:"name"`
	Cats     string `json:"cat"`
	Ver      string `json:"ver"`
	Keywords string `json:"Keyword"`
}

type RequestDevice struct {
	Did   string `json:"did"`
	Dpid  string `json:"dpid"`
	Ip    string `json:"ip"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Os    string `json:"os"`
	Osv   string `json:"osv"`
}

type RequestBrowser struct {
	Name string `json:"bname"`
	Ua   string `json:"bsoft"`
	Url  string `json:"burl"`
}

type Request struct {
	Bid     string          `json:"requestid"`
	Imp     []*RequestImp   `json:"imp"`
	App     *RequestApp     `json:"app"`
	Browser *RequestBrowser `json:"browser"`
	Device  *RequestDevice  `json:"device"`
	Labels  string          `json:"adselect"`
}

type ResponseAd struct {
	ImpId    string `json:"impid"`
	AdId     string `json:"adid"`
	Width    int32  `json:"w"`
	Height   int32  `json:"h"`
	Curl     string `json:"curl"`
	MediaUrl string `json:"adurl"`
	MId      string `json:"crid"`
	Imurl    string `json:"impression"`
}

type Response struct {
	Bid string      `json:"requestid"`
	Ad  *ResponseAd `json:"adresponse"`
}

func (req *Request) GetId() string {
	return req.Bid
}

func (req *Request) GetUa() string {
	browser := req.Browser
	if browser != nil {
		return browser.Ua
	}
	return ""
}

func (imp *RequestImp) GetId() string {
	if imp != nil {
		return imp.ImpId
	}
	return ""
}

func (imp *RequestImp) GetWidth() int32 {
	if imp != nil {
		return imp.Width
	}
	return 0
}

func (imp *RequestImp) GetHeight() int32 {
	if imp != nil {
		return imp.Height
	}
	return 0
}

func (app *RequestApp) GetName() string {
	if app != nil {
		return app.Name
	}
	return ""
}

func (app *RequestApp) GetCats() string {
	if app != nil {
		return app.Cats
	}
	return ""
}

func (app *RequestApp) GetKeywords() string {
	if app != nil {
		return app.Keywords
	}
	return ""
}

func (req *Request) GetDeviceIp() string {
	if req != nil && req.Device != nil {
		return req.Device.Ip
	}
	return ""
}

func (req *Request) GetDeviceDid() string {
	if req != nil && req.Device != nil {
		return req.Device.Did
	}
	return ""
}

func (req *Request) GetDeviceDpId() string {
	if req != nil && req.Device != nil {
		return req.Device.Dpid
	}
	return ""
}

func (req *Request) GetDeviceMake() string {
	if req != nil && req.Device != nil {
		return req.Device.Make
	}
	return ""
}

func (req *Request) GetDeviceModel() string {
	if req != nil && req.Device != nil {
		return req.Device.Model
	}
	return ""
}

func (req *Request) GetDeviceOs() string {
	if req != nil && req.Device != nil {
		return req.Device.Os
	}
	return ""
}

func (req *Request) GetDeviceOsv() string {
	if req != nil && req.Device != nil {
		return req.Device.Osv
	}
	return ""
}
