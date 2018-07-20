package autohome

/*request*/
type BidRequest struct {
	Id      string    `json:"id"`
	Version string    `json:"version"`
	AdSlot  []*AdSlot `json:"adSlot"`
	Site    *Site     `json:"site,omitempty"`
	Mobile  *Mobile   `json:"mobile,omitempty"`
	User    *User     `json:"user,omitempty"`
	IsTest  bool      `json:"is_test"`
}
type AdSlot struct {
	Id           string  `json:"id"`
	SlotId       string  `json:"slotid"`
	BidFloor     int32   `json:"min_cpm_price"`   //底价  cpm  单位分
	DealType     string  `json:"deal_type"`       //交易类型 enum("pdb","pd","pa","rtb")
	SlotPosition int32   `json:"slot_visibility"` //广告位相对位置 0:无数据;1-5:第1-5屏;6:第6屏及其他
	ForBidCat    []int32 `json:"excluded_ad_category,omitempty"`
	Banner       *Banner `json:"banner,omitempty"`
	Video        *Video  `json:"video,omitempty"`
}

type Banner struct {
	Width    int32   `json:"width"`
	Height   int32   `json:"height"`
	ViewType []int32 `json:"view_type"`  //广告位展现形式
	Tempids  []int32 `json:"templateId"` //广告位支持的模板
}

type Video struct{}

type Site struct {
	Url string `json:"url"` //当前页面url
	Ref string `json:"ref"` //指向页url
}

type Mobile struct {
	IsApp   bool    `json:"is_app"` //广告请求是否来自app
	PkgName string  `json:"pkgname"`
	Device  *Device `json:"device"`
}

type Device struct {
	Platform  int     `json:"pm,omitempty"`          //1-ios,2-android,3-other
	Brand     string  `json:"devicebrand,omitempty"` //
	Model     string  `json:"devicemodel,omitempty"` //
	Osv       string  `json:"os_version,omitempty"`
	NetType   int32   `json:"conn,omitempty"`      //0:unknown,1:wifi,2:2g,3:3g,4:4g
	Operator  int     `json:"networkid,omitempty"` //0:unknown,7012:移动,70123:联通,70121:电信
	GeoLong   float32 `json:"lng,omitempty"`       //经度×1000,000
	GeoLat    float32 `json:"lat,omitempty"`       //纬度x1000,000
	SWidth    int32   `json:"screen_width"`
	SHeight   int32   `json:"screen_height"`
	Did       string  `json:"deviceid,omitempty"`           //加密后的设备id,ios:openuuid,android:imei
	Dpi       float64 `json:"screen_density,omitempty"`     //设备像素密度
	ScrOrient int32   `json:"screen_orientation,omitempty"` //0，90，180，270
}

type User struct {
	Id string `json:"id,omitempty"` //adx侧userid
	Ip string `json:"ip,omitempty"`
	Ua string `json:"user_agent,omitempty"`
}

/*response*/
type BidResponse struct {
	BidId    string `json:"id"`
	Version  string `json:"version"`                      //与请求保持一致
	TimeTook int64  `json:"processing_time_ms,omitempty"` //dsp处理消耗时间(ms)
	DoCm     bool   `json:"is_cm"`
	Ads      []*Ads `json:"ads,omitempty"`
}
type Ads struct {
	ImpId        string     `json:"id"`
	SlotId       string     `json:"slotid"`
	MaxPrice     int32      `json:"max_cpm_price"` //最高出价(pdb,pd直接返回底价)
	CreativeId   int64      `json:"creative_id"`
	AdvertiserId int32      `json:"advertiser_id"`
	Width        int32      `json:"width"`
	Height       int32      `json:"height"`
	Category     int32      `json:"category,omitempty"`
	CreativeType int32      `json:"creative_type,omitempty"` //1001:文字链，1002：图片，1003:flash,暂时只开放图片
	TempId       int32      `json:"templateId"`
	AdSnippet    *AdSnippet `json:"adsnippet"`
}
type AdSnippet struct {
	ImgSrc  string   `json:"img,omitempty"`
	Img2Src string   `json:"img2,omitempty"`
	Img3Src string   `json:"img3,omitempty"`
	ImpUrls []string `json:"pv"`   //曝光监控地址, 宏%%SETTLE_PRICE%%
	Link    string   `json:"link"` //点击地址（落地页），与点击宏结合 adx 302跳转到该地址, %%CLICK_URL_UNESC%%
	Text    string   `json:"text,omitempty"`
	SubText string   `json:"subText,omitempty"`
	Content []*Meta  `json:"content,omitempty"` //app端模版
}
type Meta struct {
	Src  string `json:"src"`  //静态图片地址或者文字内容
	Type string `json:"type"` //simg,gif,bimg,text
}
