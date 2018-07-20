package mediamax

type MMBidRequest struct {
	Id     string               `json:"id"`
	IsTest bool                 `json:"is_test"`
	IsPing bool                 `json:"is_ping"`
	Tmax   int32                `json:"tmax"`
	Site   *MMBidRequest_Site   `json:"site"`
	Device *MMBidRequest_Device `json:"device"`
	User   *MMBidRequest_User   `json:"user"`
	Imp    []*MMBidRequest_Imp  `json:"imp"`
	App    *MMBidRequest_App    `json:"app"`
}

type MMBidRequest_Site struct {
	Domain         string      `json:"domain"`
	Allyesitetype  string      `json:"allyesitetype"`  //媒体类型分类
	Allyespageform string      `json:"allyespageform"` //展现形式分类
	Content        interface{} `json:"content"`        //视频环境描述
	Page           string      `json:"page"`
	Ref            string      `json:"ref"`
}

type MMBidRequest_Device struct {
	Dnt            int32                    `json:"dnt"`
	Ua             string                   `json:"ua"`
	Ip             string                   `json:"ip"`
	Js             int32                    `json:"js"`         //是否支持javascript
	Devicetype     int32                    `json:"devicetype"` //1:pc 2:mob 3:pad 4:tv
	Make           string                   `json:"make"`
	Model          string                   `json:"model"`
	Os             string                   `json:"os"` //ios android
	Osv            string                   `json:"osv"`
	Connectiontype int32                    `json:"connectiontype"` //0:unknow 1:ethernet 2:wifi 3:unkonwcell 4:2g 5:3g 6:4g
	Carrier        string                   `json:"carrier"`        //46000:移动 46001:联通 46003 46020
	Ifa            string                   `json:"ifa"`            //idfa androidid
	Didsha1        string                   `json:"didsha1"`
	Didmd5         string                   `json:"didmd5"`
	Macsha1        string                   `json:"macsha1"`
	H              int32                    `json:"h"` //屏幕高度
	W              int32                    `json:"w"`
	Pxratio        float32                  `json:"pxratio"`
	Geo            *MMBidRequest_Device_Geo `json:"geo"`
}

type MMBidRequest_User struct {
	Id   string `json:"id"`
	Cver int32  `json:"cver"`
}

type MMBidRequest_Imp struct {
	Id              string                   `json:"id"` //默认为1
	Tagid           string                   `json:"tagid"`
	Bidfloor        float32                  `json:"bidfloor"`        // yuan/cpm
	Interactivetype []int32                  `json:"interactivetype"` //只在app请求内包含该字段
	Pmp             *MMBidReuqest_Imp_Pmp    `json:"pmp"`
	Banner          *MMBidRequest_Imp_Banner `json:"banner"`
	Video           *MMBidRequest_Imp_Video  `json:"video"`
	Native          *MMBidRequest_Imp_Native `json:"native"`
	HttpsFlag       int32                    `json:"https_flag"` //等于1返回https
}

type MMBidRequest_App struct {
	Id      string      `json:"id"`
	Cat     []string    `json:"cat"`
	Bundle  string      `json:"bundle"`
	Content interface{} `json:"content"`
}

type MMBidRequest_Device_Geo struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type MMBidRequest_Imp_Banner struct {
	W              int32   `json:"w"`
	H              int32   `json:"h"`
	Allyesadformat []int32 `json:"allyesadformat"`
	Allyesadform   string  `json:"allyesadform"`
}

type MMBidRequest_Imp_Video struct {
	Linearity      int32   `json:"linearity"` //1:线性 2:非线性
	Protocol       int32   `json:"protocol"`  //3, vast3.0
	W              int32   `json:"w"`
	H              int32   `json:"h"`
	Maxduration    int32   `json:"maxduration"`
	Minduration    int32   `json:"minduration"`
	Allyesadformat []int32 `json:"allyesadformat"`
	Allyesadform   string  `json:"allyesadform"`
}

type MMBidRequest_Imp_Native struct {
	W              int32         `json:"w"`
	H              int32         `json:"h"`
	Allyesadformat []int32       `json:"allyesadformat"`
	Allyesadform   string        `json:"allyesadform"` // 600 601 602 603
	Elements       []string      `json:"elements"`
	ElementsNew    *Native_Elems `json:"elements_new"`
}

type Native_Elems struct {
	Element []*Native_Elem_KV  `json:"element"`
	Image   []*Native_Elem_Img `json:"image"`
}

type Native_Elem_KV struct {
	Key string `json:"key"`
	Len int32  `json:"len"`
}

type Native_Elem_Img struct {
	W int32 `json:"w"`
	H int32 `json:"h"`
}

type MMBidReuqest_Imp_Pmp struct {
	PrivateAuction int32                     `json:"private_auction"` //0:rtb 1:pdb
	Deals          []*MMBidRequest_Pmp_Deals `json:"deals"`
}

type MMBidRequest_Pmp_Deals struct {
	Id       int32   `json:"id"`
	Bidfloor float32 `json:"bidfloor"`
}

type MMBidResponse struct {
	Id      string                   `json:"id"`              //reqid
	Bidid   string                   `json:"bidid,omitempty"` //dsp respid
	Cur     string                   `json:"cur"`             // CNY
	Seatbid []*MMBidResponse_Seatbid `json:"seatbid"`
}

type MMBidResponse_Seatbid struct {
	Bid []interface{} `json:"bid"`
}

type MMBidResponse_Seatbid_Bid struct {
	Id     string   `json:"id"`
	Impid  string   `json:"impid"`
	Price  float32  `json:"price"`  // yuan/cpm
	Cmflag int32    `json:"cmflag"` // 1为有效,其他值为无效
	Adid   string   `json:"adid"`
	Nurl   string   `json:"nurl,omitempty"` //winnotice
	Crid   string   `json:"crid"`           //uuid, 动态创意填mid
	Dealid string   `json:"dealid,omitempty"`
	Fmt    int32    `json:"fmt"` //0:jpg 1:gif 2:swf 3:txt
	Cat    []string `json:"cat"`

	/*静态广告回复*/
	Curl string `json:"curl,omitempty"` //点击监测url
	Surl string `json:"surl,omitempty"` //曝光监测url

	/*动态广告回复*/
	Adm string `json:"adm,omitempty"` //创意代码

	/*动态原生广告回复*/
	Advertiserid    string                      `json:"advertiserid,omitempty"` //dsp广告主id
	Adomain         []string                    `json:"adomain,omitempty"`
	Nativead        *MMBidResponse_Bid_Nativead `json:"nativead,omitempty"`
	Interactivetype int32                       `json:"interactivetype,omitempty"`
	Clkurl          string                      `json:"clkurl,omitempty"` //落地页
	Nsurl           []string                    `json:"nsurl,omitempty"`
	Ncurl           []string                    `json:"ncurl,omitempty"`
}

type MMBidResponse_Bid_Nativead struct {
	Nadtype string                        `json:"nadtype"` // 600:feeds 601:content
	Title   string                        `json:"title,omitempty"`
	Stitle  string                        `json:"stitle,omitempty"`
	Desc    string                        `json:"desc,omitempty"`
	Btnname string                        `json:"btnname,omitempty"` //广告点击按钮名称
	Logourl string                        `json:"logourl,omitempty"` //广告主logo
	Image   *MMBidResponse_Nativead_Image `json:"image,omitempty"`
	// todo:文档前后类型不统一,需要更新文档,更新后再完成今日头条的回复
	ImageNew *MMBidResponse_Nativead_ImageNew `json:"mage_new,omitempty"`
}

type MMBidResponse_Nativead_Image struct {
	Url string `json:"url"`
	W   int32  `json:"w"`
	H   int32  `json:"h"`
}

type MMBidResponse_Seatbid_VBid struct {
	Id     string  `json:"id"`
	Impid  string  `json:"impid"`
	Price  float32 `json:"price"`  // yuan/cpm
	Cmflag int32   `json:"cmflag"` // 1为有效,其他值为无效
	Adid   string  `json:"adid"`
	Nurl   string  `json:"nurl"` //winnotice
	Crid   string  `json:"crid"` //uuid, 动态创意填mid
	Dealid string  `json:"dealid,omitempty"`
	/*VAST标准视频广告回复*/
	Vcurl []string                        `json:"vcurl,omitempty"` //视频点击监测url
	Vsurl []string                        `json:"vsurl,omitempty"` //视频曝光监测url
	Event []*MMBidResponse_Bid_VideoEvent `json:"event,omitempty"`
}

type MMBidResponse_Bid_VideoEvent struct {
	Name     string `json:"name"`
	Trackurl string `json:"trackurl"`
	Offset   string `json:"offset,omitempty"` //特定时间点监测
}

type MMBidResponse_Nativead_ImageNew struct {
	Url string `json:"url"`
	W   int32  `json:"w"`
	H   int32  `json:"h"`
}
