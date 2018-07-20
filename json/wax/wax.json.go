package wax

type WaxRequest struct {
	Id     string             `json:"id"`     //request Id
	Dealid string             `json:"dealid"` //PD模式,CPM固定
	Imp    []*WaxRequest_Imp  `json:"imp"`
	Device *WaxRequest_Device `json:"device"`
	User   *WaxRequest_User   `json:"user"`
	App    *WaxRequest_App    `json:"app"`
}

type WaxRequest_Imp struct {
	Id          string                 `json:"id"`
	Tagid       string                 `json:"tagid"` //广告位Id
	Bidfloor    int32                  `json:"bidfloor"`
	Bidfloorcur string                 `json:"bidfloorcur"`
	Banner      *WaxRequest_Imp_Banner `json:"banner"`
	Feed        *WaxRequest_Imp_Feed   `json:"feed"`
}

type WaxRequest_Imp_Banner struct {
	Width  int32 `json:"W"`
	Height int32 `json:"H"`
}

type WaxRequest_Imp_Feed struct {
	Type int32 `json:"type"` //0:不区分 1:普通博文 2:card样式
}

type WaxRequest_Device struct {
	UA      string                 `json:"ua"`
	IP      string                 `json:"ip"`
	Geo     *WaxRequest_Device_Geo `json:"geo"`
	Model   string                 `json:"model"`
	Os      string                 `json:"os"` //ios android
	Osv     interface{}            `json:"osv"`
	Contype int32                  `json:"connectiontype"` //0:Unkonwn 1:Ethenet 2:Wifi 3:xG 4:2G 5:3G 6:4G
	Carrier string                 `json:"carrier"`        //46000:移动 46001:联通 46002:电信 46020:铁通
	EXT     *WaxRequest_Device_EXT `json:"ext"`
}

type WaxRequest_Device_Geo struct {
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
	Type int32   `json:"type"` //地理信息来源 0:GSP/位置服务 1:IP地址 2:用户提供
}

type WaxRequest_Device_EXT struct {
	IDFA string `json:"idfa"`
	Imei string `json:"imei"`
}

type WaxRequest_User struct {
	Gender string `json:"gender"`
	Yob    int32  `json:"yob"`
}

type WaxRequest_App struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type WaxResponse struct {
	Id      string                 `json:"id"`    //Request Id
	Bidid   string                 `json:"bidid"` //DSP竞价Id
	Dealid  string                 `json:"dealid"`
	Seatbid []*WaxResponse_SeatBid `json:"seatbid"`
}

type WaxResponse_SeatBid struct {
	Bid []*WaxResponse_SeatBid_Bid `json:"bid"`
}

type WaxResponse_SeatBid_Bid struct {
	DId    string                       `json:"id"`
	Impid  string                       `json:"impid"`
	Price  int32                        `json:"price"`
	WinUrl string                       `json:"nurl"` //win norice url
	Adm    string                       `json:"adm"`  //mid,博文id
	Crid   string                       `json:"crid"`
	EXT    *WaxResponse_SeatBid_Bid_EXT `json:"ext"`
}

type WaxResponse_SeatBid_Bid_EXT struct {
	ShowUrls  []string `json:"pm"`
	ClickUrls []string `json:"cm"`
	Landingid string   `json:"landingid"`
}
