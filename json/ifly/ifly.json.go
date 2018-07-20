package ifly

type IBidRequest struct {
	Id     string             `json:"id"`
	Imp    []*IBidRequest_Imp `json:"imp"`
	App    *IBidRequest_App   `json:"app"`
	Device *IBidRequest_Dev   `json:"device"`
	User   *IBidRequest_User  `json:"user"`
}

type IBidRequest_Imp struct {
	Id       string                  `json:"id"`
	Tagid    string                  `json:"tagid"`
	Banner   *IBidRequest_Imp_Banner `json:"banner"`
	Native   *IBidRequest_Imp_Native `json:"native"`
	Video    *IBidRequest_Imp_Video  `json:"video"`
	Instl    int32                   `json:"instl"` // 1:横幅2:全屏3:插屏5:音频6:视频7:焦点图8:信息流9:内容流10:动态消息12:一图13三图一文
	Bidfloor float32                 `json:"bidfloor"`
	Cur      string                  `json:"cur"`
	Pmp      *IBidRequest_Imp_Pmp    `json:"pmp"`
	Secure   int32                   `json:"secure"`
}

type IBidRequest_Imp_Banner struct {
	Type int32 `json:"type"` // 1:横幅2:全屏3:插屏
	W    int32 `json:"w"`
	H    int32 `json:"h"`
}

type IBidRequest_Imp_Native struct {
	Title *IBidRequest_Imp_Native_Title `json:"title"`
	Desc  *IBidRequest_Imp_Native_Desc  `json:"desc"`
	Icon  *IBidRequest_Imp_Native_Icon  `json:"icon"`
	Img   *IBidRequest_Imp_Native_Img   `json:"img"`
}

type IBidRequest_Imp_Native_Title struct {
	Len int32 `json:"len"`
}

type IBidRequest_Imp_Native_Desc struct {
	Len int32 `json:"len"`
}

type IBidRequest_Imp_Native_Icon struct {
	H int32 `json:"h"`
	W int32 `json:"w"`
}

type IBidRequest_Imp_Native_Img struct {
	H int32 `json:"h"`
	W int32 `json:"w"`
}

type IBidRequest_Imp_Video struct {
	Protocol    int32 `json:"protocol"` // 101:iflytek specific
	W           int32 `json:"w"`
	H           int32 `json:"h"`
	Minduration int32 `json:"minduration"`
	Maxduration int32 `json:"maxduration"`
	Linearity   int32 `json:"linearity"` // 1:线性/视频流
}

type IBidRequest_Imp_Pmp struct {
	Deals []*IBidRequest_Imp_Pmp_Deal `json:"deals"`
}

type IBidRequest_Imp_Pmp_Deal struct {
	Id       string  `json:"id"`
	Bidfloor float32 `json:"bidfloor"`
}

type IBidRequest_App struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Bundle  string   `json:"bundle"`
	Cat     []string `json:"cat"`
	Context []string `json:"context"`
}

type IBidRequest_Dev struct {
	W        int32                `json:"w"`
	H        int32                `json:"h"`
	Ua       string               `json:"ua"`
	Ip       string               `json:"ip"`
	Geo      *IBidRequest_Dev_Geo `json:"geo"`
	Did      string               `json:"did"`
	Didsha1  string               `json:"didsha1"`
	Didmd5   string               `json:"didmd5"`
	Dpid     string               `json:"dpid"`
	Dpidsha1 string               `json:"dpidsha1"`
	Dpidmd5  string               `json:"dpidmd5"`
	Mac      string               `json:"mac"`
	Macsha1  string               `json:"macsha1"`
	Macmd5   string               `json:"macmd5"`
	Ifa      string               `json:"ifa"`
	Make     string               `json:"make"`
	Model    string               `json:"model"`
	Os       string               `json:"os"` // Android, iOS, Windows Phone, other
	Osv      string               `json:"osv"`
	Carrier  string               `json:"carrier"`
	Language string               `json:"language"`
	Js       int32                `json:"js"`             // 1:启用javascript 0:不启用
	Contype  int32                `json:"connectiontype"` // 0:未知,1:Ethernet,2:wifi,3:2G,4:3G,5:4G
	Devtype  int32                `json:"devicetype"`     // 0:手机,1:平板,2:PC,3:互联网电视
	ad       int32                `json:"ad"`
}

type IBidRequest_Dev_Geo struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type IBidRequest_User struct {
	Tags []string `json:"tags"`
}

type IBidResponse struct {
	Id      string                  `json:"id"`    // IBidRequest Id
	Bidid   string                  `json:"bidid"` // dsp bidid
	Seatbid []*IBidResponse_Seatbid `json:"seatbid"`
	Cur     string                  `json:"cur,omitempty"`
}

type IBidResponse_Seatbid struct {
	Bid []*IBidResponse_Seatbid_Bid `json:"bid"`
}

type IBidResponse_Seatbid_Bid struct {
	Impid    string                 `json:"impid"`
	Price    float32                `json:"price"` // cpm
	BannerAd *IBidResponse_BannerAd `json:"banner_ad,omitempty"`
	NativeAd *IBidResponse_NativeAd `json:"native_ad,omitempty"`
	VideoAd  *IBidResponse_VideoAd  `json:"video_ad,omitempty"`
	Lattr    int32                  `json:"lattr"` // 1:网页,2:下载,3:应用市场
	Dealid   string                 `json:"dealid,omitempty"`
}

type IBidResponse_BannerAd struct {
	Mtype        int32    `json:"mtype"` // 1:纯文字,2:纯图片,3:图文,4:html
	Title        string   `json:"title,omitempty"`
	Desc         string   `json:"desc,omitempty"`
	ImageUrl     string   `json:"image_url"`
	Html         string   `json:"html,omitempty"`
	Landing      string   `json:"landing"`
	W            int32    `json:"w"`
	H            int32    `json:"h"`
	Showurl      []string `json:"impress"`
	Clickurl     []string `json:"click"`
	PackageName  string   `json:"package_name,omitempty"`
	IsMarked     int32    `json:"is_marked"` // 0:未知,1:有水印
	AdSourceMark string   `json:"ad_source_mark,omitempty"`
}

type IBidResponse_NativeAd struct {
	Title        string   `json:"title,omitempty"`
	Desc         string   `json:"desc,omitempty"`
	Icon         string   `json:"icon,omitempty"` // url
	Img          string   `json:"img"`
	Landing      string   `json:"landing"`
	Showurl      []string `json:"imptrackers"`
	Clickurl     []string `json:"clicktrackers"`
	PackageName  string   `json:"package_name,omitempty"`
	ImgUrls      []string `json:"img_urls,omitempty"` // 多图文
	IsMarked     int32    `json:"is_marked"`
	AdSourceMark string   `json:"ad_source_mark,omitempty"`
}

type IBidResponse_VideoAd struct {
	Src          string   `json:"src"`
	Duration     int32    `json:"duration"`
	Landing      string   `json:"landing"`
	Strackers    []string `json:"starttrackers"`
	Ctrackers    []string `json:"completetrackers,omitempty"`
	Clickurl     []string `json:"clicktrackers"`
	PackageName  string   `json:"package_name,omitempty"`
	IsMarked     int32    `json:"is_marked"`
	AdSourceMark string   `json:"ad_source_mark,omitempty"`
}
