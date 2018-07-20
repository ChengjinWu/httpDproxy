package innersdk

import (
	"fmt"
	"thrift-gen/adserver"
	"wisdom/common"
)

/*Request Part*/

const (
	OSANDROID = 0
	OSIOS     = 1
	OSWP      = 2
	OSOTH     = 3

	DEVUNKNOWN = 0
	DEVMOBILE  = 1
	DEVPAD     = 2
	DEVPC      = 3
	DEVTV      = 4

	SMALL  = "3"
	MIDDLE = "2"
	LARGE  = "1"
	XLARGE = "0"

	URL_APP_DOWNLOAD = 0 //app下载
	URL_LANDSCAPE_H5 = 1 //横屏H5
	URL_PORTRAIT_H5  = 2 //竖屏H5
)

var OsMap = map[uint32]string{
	0: "android",
	1: "ios",
	2: "wp",
	3: "others",
}

var ConnTypeMap = map[uint32]adserver.ConnectionType{
	0: adserver.ConnectionType_NA,
	1: adserver.ConnectionType_WIFI,
	2: adserver.ConnectionType_CELL2G,
	3: adserver.ConnectionType_CELL3G,
	4: adserver.ConnectionType_CELL4G,
}

var AdslotTypeMap = map[uint32]string{
	1: "SPLASH",
	2: "CARD",
	3: "NATIVE",
	4: "BANNER",
}

type Request struct {
	SlotID string `json:"slot_id"` //广告位id
	Type   string `json:"type"`    //广告类型
	Count  uint32 `json:"count"`   //请求数量
}

type Geo struct {
	Latitude  float32 `json:"latitude"`  // 纬度,取值从­90到90,南为负值
	Longitude float32 `json:"longitude"` // 经度,取值从­180到180,西为负值
}

type Device struct {
	Geo            *Geo   `json:"geo"`
	ScreenWidth    uint32 `json:"screen_width"`    //屏幕宽度(像素)
	ScreenHeight   uint32 `json:"screen_height"`   //屏幕高度(像素)
	IDFA           string `json:"idfa"`            //ios的idfa
	AndroidID      string `json:"androidid"`       //android的androdid
	MAC            string `json:"mac"`             //mac地址
	IMEI           string `json:"imei"`            //用户端的imei
	Make           string `json:"make"`            //设备品牌
	Model          string `json:"model"`           //设备型号
	OS             uint32 `json:"os"`              //设备操作系统 0:Android,1:IOS,2:WP,3:OTHERS
	OSV            string `json:"osv"`             //操作系统版本
	ConnectionType uint32 `json:"connection_type"` //网络连接类型 0:未知,1:Wifi,2:2G,3:3G,4:4G
	DeviceType     uint32 `json:"device_type"`     //设备类型0:未知,1:手机,2:平板,3:PC,4:TV
}

type App struct {
	ID         string   `json:"id"`         //应用程序id
	Name       string   `json:"name"`       //应用程序名称
	Domain     string   `json:"domain"`     //应用程序域名
	Cat        []string `json:"cat"`        //应用程序分类信息
	SectionCat []string `json:"sectioncat"` //应用程序子分类
}

type AdRequest struct {
	Requests        []*Request `json:"reqs"`
	Device          *Device    `json:"device"`
	App             *App       `json:"app"`
	Signature       string     `json:"signature"`
	Timestamp       string     `json:"timestamp"`
	CachedCreatives []int32    `json:"cached_creatives"` //预加载的创意列表
	SdkVersion      string     `json:"sdk_version"`      //sdk版本号
	Ip              string     `json:"ip,omitempty"`     //请求ip,由adx侧 通过header获取
}

func (req *Request) GetAdType() string {
	if req != nil {
		return req.Type
	}
	return ""
}

func (req *Request) GetSlotId() string {
	if req != nil {
		return req.SlotID
	}
	return ""
}

func (req *Request) GetAdCount() uint32 {
	if req != nil {
		return req.Count
	}

	return 0
}

func (req *AdRequest) GetRequests() []*Request {
	if req != nil {
		return req.Requests
	}
	return nil
}

func (req *AdRequest) GetDevice() *Device {
	if req != nil {
		return req.Device
	}
	return nil
}

func (req *AdRequest) GetAppId() string {
	if req != nil {
		if req.App != nil {
			return req.App.ID
		}
	}
	return ""
}
func (req *AdRequest) GetAppCats() []string {
	if req != nil {
		if req.App != nil {
			return req.App.Cat
		}
	}
	return nil
}

func (req *AdRequest) GetAppName() string {
	if req != nil {
		if req.App != nil {
			return req.App.Name
		}
	}
	return ""
}

func (req *AdRequest) GetAppSectionCat() []string {
	if req != nil {
		if req.App != nil {
			return req.App.SectionCat
		}
	}
	return nil
}

func (req *AdRequest) GetSignature() string {
	if req != nil {
		return req.Signature
	}
	return ""
}

func (req *AdRequest) GetCachedCreatives() []int32 {
	if req != nil {
		return req.CachedCreatives
	}
	return []int32{}
}

func (req *AdRequest) GetDeviceGeo() *Geo {
	if req != nil && req.Device != nil {
		return req.Device.Geo
	}
	return nil
}

func (req *AdRequest) GetDeviceWidth() uint32 {
	if req != nil && req.Device != nil {
		return req.Device.ScreenWidth
	}
	return 0
}

func (req *AdRequest) GetDeviceHeight() uint32 {
	if req != nil && req.Device != nil {
		return req.Device.ScreenHeight
	}
	return 0
}

func (req *AdRequest) GetDeviceIdfa() string {
	if req != nil && req.Device != nil {
		return req.Device.IDFA
	}
	return ""
}

func (req *AdRequest) GetDeviceAndroidID() string {
	if req != nil && req.Device != nil {
		return req.Device.AndroidID
	}
	return ""
}

func (req *AdRequest) GetDeviceMac() string {
	if req != nil && req.Device != nil {
		return req.Device.MAC
	}
	return ""
}

func (req *AdRequest) GetDeviceIMEI() string {
	if req != nil && req.Device != nil {
		return req.Device.IMEI
	}
	return ""
}

func (req *AdRequest) GetDeviceMake() string {
	if req != nil && req.Device != nil {
		return req.Device.Make
	}
	return ""
}

func (req *AdRequest) GetDeviceModel() string {
	if req != nil && req.Device != nil {
		return req.Device.Model
	}
	return ""
}

func (req *AdRequest) GetDeviceOS() uint32 {
	if req != nil && req.Device != nil {
		return req.Device.OS
	}
	return 0
}

func (req *AdRequest) GetDeviceOSV() string {
	if req != nil && req.Device != nil {
		return req.Device.OSV
	}
	return ""
}

func (req *AdRequest) GetDeviceConnType() uint32 {
	if req != nil && req.Device != nil {
		return req.Device.ConnectionType
	}
	return 0
}

func (req *AdRequest) GetDeviceType() uint32 {
	if req != nil && req.Device != nil {
		return req.Device.DeviceType
	}
	return 0
}

/*Response Part*/

const (
	// event type
	EVNT_VIDEO_VIEW = "video_view"    //播放停止时,传送播放整体数据
	EVNT_AD_START   = "start"         //广告开始播放
	EVNT_AD_25PCT   = "first_quarter" //广告播放到1/4
	EVNT_AD_MIDDLE  = "middle"        //广告播放到1/2
	EVNT_AD_75PCT   = "third_quarter" //广告播放到3/4
	EVNT_AD_100PCT  = "complete"      //广告播放到100%
	EVNT_MUTE       = "mute"          //静音
	EVNT_UNMUTE     = "unmute"        //关闭静音

	//material type
	TEXT_MATERIAL  = 0
	IMAGE_MATERIAL = 1
	VIDEO_MATERIAL = 2
)

type AdMeta struct {
	Type        uint32 `json:"type,omitempty"`         //基本素材类型 0-文字 1-图片 2-视频
	Size        uint32 `json:"size,omitempty"`         //图片/视频素材 体积 (单位:B)
	Width       uint32 `json:"width,omitempty"`        //图片/视频素材 width
	Height      uint32 `json:"height,omitempty"`       //图片/视频素材 height
	Url         string `json:"url,omitempty"`          //图片/视频素材 路径
	Md5         string `json:"md5,omitempty"`          //图片/视频素材 Md5编码
	UpdatedTime uint32 `json:"updated_time,omitempty"` //图片/视频素材 更新时间
	Text        string `json:"text,omitempty"`         //文字素材 内容
	Duration    uint32 `json:"duration,omitempty"`     //视频素材 时长
}

type ValidArea struct {
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Width  float32 `json:"width"`
	Height float32 `json:"height"`
}

type Effects struct {
	DisplayTime     uint32     `json:"display_time"`
	AnimateDuration uint32     `json:"animate_duration"`
	EngageArea      *ValidArea `json:"engage_area,omitempty"`
}

type AdMaterial struct {
	Type      string             `json:"type"`      //素材类型
	ClickUrl  string             `json:"click_url"` //点击地址
	DeepLink  *DeepLink          `json:"deep_link,omitempty"`
	MetaGroup map[string]*AdMeta `json:"meta_group"`        //素材内容组合
	Effects   *Effects           `json:"effects,omitempty"` //广告创意效果
}

type Response struct {
	CreativeID           uint32              `json:"creative_id"`            //创意id
	InAppBrowser         bool                `json:"in_app_browser"`         //是否用应用内置浏览器跳转到落地页
	Material             *AdMaterial         `json:"material"`               //广告素材内容
	CreativeUpdateTime   uint32              `json:"creative_update_time"`   //创意更新时间
	StartTime            uint32              `json:"start_time"`             //广告开始时间
	EndTime              uint32              `json:"end_time"`               //广告结束时间
	TzOffset             uint32              `json:"tz_offset"`              //时区 配合playtime
	PlayTime             []int32             `json:"play_time"`              //7位长度数组,表示允许投放的日期(周六->周日) 低24bit 表示允许投放的小时
	DeadLine             uint32              `json:"deadline"`               //此次广告的有效期限
	AllowedImps          uint32              `json:"allowed_imps"`           //有效期内允许的曝光量
	FreqCap              uint32              `json:"freq_limit_hour"`        //每小时的频控
	Priority             uint32              `json:"priority"`               //播放权重
	TkImpUrl             []string            `json:"tk_imp_url"`             //用于广告曝光追踪
	TkCkUrl              []string            `json:"tk_ck_url"`              //用于广告点击追踪
	TkEvntUrlMap         map[string][]string `json:"tk_event_url_map"`       //不同事件类型对应的追踪url（除去常规的曝光点击之外的事件）
	IsDisabled           bool                `json:"is_disabled"`            //广告是否已失效
	DailyImpressionLimit uint32              `json:"daily_impression_limit"` //广告最大日曝光量
}

type AdExtend struct {
	Ext string `json:"ext"`
}

type CreativeExt struct {
	CreativeId uint32 `json:"creative_id"`
	Ext        string `json:"ext"`
}
type SlotAd struct {
	SlotId      string         `json:"slot_id"`
	CreativeExt []*CreativeExt `json:"creative_ext"`
}

type AdResponse struct {
	RequestId  string      `json:"request_id"`
	AdResponse []*Response `json:"ad_resps"` //creative_id -> response
	//SlotAdMap  map[string]map[string]*AdExtend `json:"slot_ad_map"` //slot_id->ad_id
	SlotAdMap []*SlotAd `json:"slot_ad_map"`
}

type MoTvAdm struct {
	CreativeId uint32          `json:"creative_id"`
	List       *MotvAdmaterial `json:"list"`
}

type MotvAdmaterial struct {
	Format              string                 `json:"format"`
	Effects             *MotvEffects           `json:"effects"`
	Assets              map[string]*MoTvAdMeta `json:"assets"`
	AndrodTrackings     map[string][]string    `json:"android_trackings"`
	IosTrackings        map[string][]string    `json:"ios_trackings"`
	AndroidRedirectUrl  *RedirectUrlAndroid    `json:"android_redirect_url"`
	IosRedirectUrl      *RedirectUrlIOS        `json:"ios_redirect_url"`
	DeepLink            map[string]*DeepLink   `json:"Deeplink,omitempty"`
	AndroidInAppBrowser uint32                 `json:"android_inapp_browser"`
	IosInappBrowser     uint32                 `json:"ios_inapp_browser"`
}

type RedirectUrlAndroid struct {
	Link    []string `json:"link"`
	UrlType uint32   `json:"url_type"`
}

type RedirectUrlIOS struct {
	Link    string `json:"link"`
	UrlType uint32 `json:"url_type"`
}

type DeepLink struct {
	Link        string `json:"link"`
	PackageName string `json:"package_name,omitempty"`
}

type MoTvAdMeta struct {
	Type        string `json:"type,omitempty"`         //基本素材类型 0-文字 1-图片 2-视频
	Size        uint32 `json:"size,omitempty"`         //图片/视频素材 体积 (单位:B)
	Width       uint32 `json:"width,omitempty"`        //图片/视频素材 width
	Height      uint32 `json:"height,omitempty"`       //图片/视频素材 height
	Url         string `json:"url,omitempty"`          //图片/视频素材 路径
	Md5         string `json:"md5,omitempty"`          //图片/视频素材 Md5编码
	UpdatedTime uint32 `json:"updated_time,omitempty"` //图片/视频素材 更新时间
	Text        string `json:"text,omitempty"`         //文字素材 内容
	Duration    uint32 `json:"duration,omitempty"`     //视频素材 时长
}

type MotvEffects struct {
	DisplayTime     uint32     `json:"display_time"`
	AnimateDuration uint32     `json:"transition_time"`
	EngageArea      *ValidArea `json:"engage_area"`
}

//[0]:motv [1]:innersdk
//motv素材有些字段与innersdk名字不同,需要做转换
var cvtmap = map[string][][]string{
	"NATIVE_VIDEO":            [][]string{{"icon1", "title1", "desc1"}, {"icon", "title", "desc"}},
	"NATIVE_IMAGE":            [][]string{{"image1", "title1", "desc1", "icon1"}, {"image", "title", "desc", "icon"}},
	"NATIVE_ANIMATION":        [][]string{{"icon1", "title1", "desc1"}, {"icon", "title", "desc"}},
	"CARD_VIDEO_PULLDOWN":     [][]string{{"endcard1"}, {"endcard"}},
	"CARD_IMAGE_GENERAL":      [][]string{{"image1"}, {"image"}},
	"SPLASH2_VIDEO_ROTATE":    [][]string{{"image1"}, {"image"}},
	"SPLASH2_IMAGE_GENERAL_L": [][]string{{"image1"}, {"image"}},
	"SPLASH2_IMAGE_GENERAL_P": [][]string{{"image1"}, {"image"}},
}

//由motv的素材json建立内设部sdk的素材结构(目前前端投放是按照motv的素材结构来投放的)
func BuildAdMeta(moadm *MoTvAdm, dwidth, dheight uint32) *AdMaterial {
	re := &AdMaterial{}
	re.Type = moadm.List.Format
	re.Effects = func(m *MotvEffects) *Effects {
		if m != nil {
			result := &Effects{}
			if m.AnimateDuration != 0 {
				result.AnimateDuration = m.AnimateDuration
			}
			if m.DisplayTime != 0 {
				result.DisplayTime = m.DisplayTime
			}

			if m.EngageArea != nil {
				result.EngageArea = &ValidArea{
					X:      m.EngageArea.X,
					Y:      m.EngageArea.Y,
					Width:  m.EngageArea.Width,
					Height: m.EngageArea.Height,
				}
			} else {
				result.EngageArea = &ValidArea{}
			}

			return result
		} else {
			return &Effects{
				EngageArea: &ValidArea{},
			}
		}
	}(moadm.List.Effects)

	if moadm.List.AndrodTrackings != nil {

		moadm.List.AndrodTrackings["first_quarter"] = moadm.List.AndrodTrackings["first_quartile"]
		moadm.List.AndrodTrackings["third_quarter"] = moadm.List.AndrodTrackings["third_quartile"]
		delete(moadm.List.AndrodTrackings, "first_quartile")
		delete(moadm.List.AndrodTrackings, "third_quartile")
	}

	if moadm.List.IosTrackings != nil {
		moadm.List.IosTrackings["first_quarter"] = moadm.List.IosTrackings["first_quartile"]
		moadm.List.IosTrackings["third_quarter"] = moadm.List.IosTrackings["third_quartile"]
		delete(moadm.List.IosTrackings, "first_quartile")
		delete(moadm.List.IosTrackings, "third_quartile")
	}

	re.MetaGroup = func(ad map[string]*MoTvAdMeta) map[string]*AdMeta {
		ret := map[string]*AdMeta{}
		for k, m := range ad {
			ret[k] = &AdMeta{
				Type: func() uint32 {
					if m.Type == "TEXT" {
						return 0
					}
					if m.Type == "IMAGE" {
						return 1
					}

					if m.Type == "VIDEO" {
						return 2
					}
					return 0
				}(),
				Width:       m.Width,
				Height:      m.Height,
				Size:        m.Size,
				Md5:         m.Md5,
				UpdatedTime: m.UpdatedTime,
				Url:         m.Url,
				Text:        m.Text,
			}
		}
		return ret
	}(moadm.List.Assets)

	//字段名的转换
	rule, ok := cvtmap[re.Type]
	if ok {
		mo := rule[0]
		in := rule[1]
		l := len(mo)
		for i := 0; i < l; i++ {
			if meta, ok := re.MetaGroup[mo[i]]; ok {
				re.MetaGroup[in[i]] = meta
				delete(re.MetaGroup, mo[i])
			}
		}
	}
	return adaptMaterial(moadm.CreativeId, dwidth, dheight, re)
}

func adaptMaterial(cid uint32, dwidth, dheight uint32, m *AdMaterial) *AdMaterial {
	mainmetaname := ""
	maxwidth := uint32(0)

	for name, meta := range m.MetaGroup {
		if meta.Type != 0 {
			if meta.Width > maxwidth {
				mainmetaname = name
				maxwidth = meta.Width
			}
		}
	}

	//获取缩放比例
	mainmeta := m.MetaGroup[mainmetaname]
	num, prop := common.FetchProp(dwidth, dheight,
		mainmeta.Width, mainmeta.Height)

	//素材适配
	for _, meta := range m.MetaGroup {
		if meta.Type != 0 {
			urlnew := common.AdaptSize(num, prop, meta.Url, fmt.Sprintf("%d", cid))
			meta.Url = urlnew
			common.AssetAdaptor(&meta.Width, &meta.Height, prop)
		}
	}

	tmpfunc := func(x, y, prop float32) (_x, _y float32) {
		xtmp := uint32(x)
		ytmp := uint32(y)
		common.AssetAdaptor(&xtmp, &ytmp, prop)
		return float32(xtmp), float32(ytmp)

	}

	if m.Effects != nil {
		if m.Effects.EngageArea != nil {
			xtmp, ytmp := tmpfunc(m.Effects.EngageArea.X, m.Effects.EngageArea.Y, prop)
			m.Effects.EngageArea.X = xtmp
			m.Effects.EngageArea.Y = ytmp

			xtmp, ytmp = tmpfunc(m.Effects.EngageArea.Width, m.Effects.EngageArea.Height, prop)
			m.Effects.EngageArea.Width = xtmp
			m.Effects.EngageArea.Height = ytmp
		}
	}

	return m
}

func ValidVersion(versionnow string, validversions []string) bool {
	for _, v := range validversions {
		if versionnow == v {
			return true
		}
	}
	return false
}
