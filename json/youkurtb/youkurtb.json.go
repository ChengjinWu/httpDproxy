package youkurtb

type BidRequest struct {
	Id     string             `json:"id"`
	Imp    []*BidRequest_Imp  `json:"imp"`
	Site   *BidRequest_Site   `json:"site"`
	App    *BidRequest_App    `json:"app"`
	Device *BidRequest_Device `json:"device"`
	User   *BidRequest_User   `json:"user"` // pc中有优酷上的用户唯一id
}

type BidRequest_Imp struct {
	Id       string                 `json:"id"`
	Tagid    string                 `json:"tagid"`
	Bidfloor float32                `json:"bidfloor"` //分cpm
	Banner   *BidRequest_Imp_Banner `json:"banner"`
	Video    *BidRequest_Imp_Video  `json:"video"`
	Native   *BidRequest_Imp_Native `json:"native"`
	Pmp      *BidRequest_Imp_Pmp    `json:"pmp"`
	Ext      *BidRequest_Imp_Ext    `json:"ext"`    //贴片广告个数
	Secure   int32                  `json:"secure"` // 1:https 0:http
}

type BidRequest_Site struct {
	Name    string              `json:"name"`
	Page    string              `json:"page"`
	Ref     string              `json:"ref"`
	Content *BidRequest_Content `json:"content"`
}

type BidRequest_App struct {
	Name    string              `json:"name"`
	Content *BidRequest_Content `json:"content"`
}

type BidRequest_Content struct {
	Title    string                  `json:"title"`
	Keywords string                  `json:"keywords"`
	Ext      *BidRequest_Content_Ext `json:"ext"`
}

type BidRequest_Content_Ext struct {
	Channel string `json:"channel"`
	Cs      string `json:"cs"`
}

type BidRequest_Device struct {
	Ua             string `json:"ua"`
	Ip             string `json:"ip"`
	Didmd5         string `json:"didmd5"` //imei md5
	Os             string `json:"os"`     //Windows Android ios Mac
	Osv            string `json:"osv"`
	Devicetype     int32  `json:"devicetype"` //0:mob 1:pad 2:pc 3:netTV
	Idfa           string `json:"idfa"`       //明文
	Androidid      string `json:"androidid"`  //明文
	Imei           string `json:"imei"`       //明文
	Mac            string `json:"mac"`        //明文
	Make           string `json:"make"`
	Model          string `json:"model"`
	Connectiontype int32  `json:"connectiontype"` //0:unknown 1: ethernet 2:wifi 3:cellunknown 4:2g 5:3g
	Carrier        string `json:"carrier"`        // 0：wifi；1：中国移动；2：中国联通：3：中国电信；4：其他；5：未识别
}

type BidRequest_User struct {
	Id     string `json:"id"`
	Gender string `json:"gender"` // "M"表示男性，"F"表示女性
	Yob    int32  `json:"yob"`    //1966
}

type BidRequest_Imp_Banner struct {
	W int32 `json:"w"`
	H int32 `json:"h"`
}

type BidRequest_Imp_Video struct {
	Linearity   int32 `json:"linearity"` //0:未知 1:线性贴片 2:悬浮 3:暂停 4:全屏悬浮 5:开机图
	Minduration int32 `json:"minduration"`
	Maxduration int32 `json:"maxduration"`
	W           int32 `json:"w"`
	H           int32 `json:"h"`
	Startdelay  int32 `json:"startdelay"` //0:前贴 1:中贴 2:后贴
}

type BidRequest_Imp_Native struct {
	NativeTemplateIds []string                   `json:"native_template_ids"` //原生广告模板ID数组
	Assets            []*BidRequest_Native_Asset `json:"assets"`
}

type BidRequest_Imp_Pmp struct {
	Deals []*BidRequest_Pmp_Deal `json:"deals"`
}

type BidRequest_Pmp_Deal struct {
	Id string `json:"id"`
	At int32  `json:"at"` // 1
}

type BidRequest_Imp_Ext struct {
	Repeat int32 `json:"repeat"`
}

type BidRequest_Native_Asset struct {
	NativeTemplateId string                  `json:"native_template_id"`
	Title            *BidRequest_Asset_Title `json:"title"`
	ImageUrl         *BidRequest_Asset_Image `json:"image_url"`
}

type BidRequest_Asset_Title struct {
	Len int32 `json:"len"`
}

type BidRequest_Asset_Image struct {
	W int32 `json:"w"`
	H int32 `json:"h"`
}

type BidResponse struct {
	Id      string                 `json:"id"`    // bidrequest.id
	Bidid   string                 `json:"bidid"` // dsp id
	Seatbid []*BidResponse_Seatbid `json:"seatbid"`
}

type BidResponse_Seatbid struct {
	Bid []*BidResponse_Seatbid_Bid `json:"bid"`
}

type BidResponse_Seatbid_Bid struct {
	Id     string                  `json:"id"`    //dsp id
	Impid  string                  `json:"impid"` //bidrequest_imp.id
	Native *BidResponse_Bid_Native `json:"native,omitempty"`
	Price  float32                 `json:"price"` //分cpm
	Nurl   string                  `json:"nurl"`
	Adm    string                  `json:"adm,omitempt"` // %%CLICK_URL_ESC%%
	Crid   string                  `json:"crid"`         //mid
	Dealid string                  `json:"dealid,omitempty"`
	Ext    *BidResponse_Bid_Ext    `json:"ext"`
}

type BidResponse_Bid_Ext struct {
	Ldp  string                `json:"ldp"`            //landing page
	Pm   []string              `json:"pm,omitempty"`   //曝光
	Cm   []string              `json:"cm"`             //点击
	Em   []string              `json:"em,omitempty"`   //播放结束
	Tm   []*BidResponse_Ext_TM `json:"tm,omitempty"`   //播放进度
	Type string                `json:"type,omitempty"` //只有当素材为动态HTML snippet时，需要指定"type":"c"，其他类型的素材均不需要提供type字段
	Api  int32                 `json:"api,omitempty"`  //指定素材使用的交互式api，如果不填，表示普通素材；如果为1，表示VPAID格式交互式flv素材；如果为2，表示非交互式swf素材
	Dp   string                `json:"dp,omitempty"`   //deeplink
}

type BidResponse_Ext_TM struct {
	T    int32  `json:"t"` //时间
	Vurl string `json:"url"`
}

type BidResponse_Bid_Native struct {
	NativeTemplateId string `json:"native_template_id,omitempty"`
}
