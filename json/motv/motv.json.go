package motv

type Request_Req struct {
	SlotId string `json:"slot_id"` // 广告位 ID
	Type   uint32 `json:"type"`    // 廣告類型 SPLASH (1), CARD (2), NATIVE (3), BANNER(4)
	Count  uint32 `json:"count"`   // 廣告請求數量
}

type Request_App struct {
	Id         string   `json:"id"`         // 应用程序ID
	Name       string   `json:"name"`       // 应用程序名称
	Domain     string   `json:"domain"`     // 应用程序域名
	Cat        []string `json:"cat"`        // 应用在AppStore 或者GooglePlay 或者其他应用市场的的分类信息
	Sectioncat []string `json:"sectioncat"` // 子分类
}

type Request_Device_Geo struct {
	Latitude  float32 `json:"latitude"`  // 纬度,取值从­90到90,南为负值
	Longitude float32 `json:"longitude"` // 经度,取值从­180到180,西为负值
}

type Request_Device struct {
	Geo            *Request_Device_Geo `json:"geo"`             // 地理位置信息
	ScreenWidth    uint32              `json:"screen_width"`    // 屏幕宽度(像素)
	ScreenHeight   uint32              `json:"screen_height"`   // 屏幕高度(像素)
	Idfa           string              `json:"idfa"`            // IDFA, 适用于iOS,保留原始值
	Androidid      string              `json:"androidid"`       // 适用于 Android ,保留原始值
	Mac            string              `json:"mac"`             // 用户终端硬件地址,适用于 Android 和 iOS,保留原始值
	Imei           string              `json:"imei"`            // 用户终端IMEI,保留原始值
	Make           string              `json:"make"`            // 设备品牌,如"Apple"
	Model          string              `json:"model"`           // 设备型号,如"iPhone"
	Os             uint32              `json:"os"`              // 设备操作系统 0­Android , 1­iOS , 2­WP , 3­Others
	Osv            string              `json:"osv"`             // 设备操作系统版本,如"3.1.2"
	Connectiontype uint32              `json:"connection_type"` //0:未知; 1:Wifi;  2:2G;  3:3G;  4:4G
	Devicetype     uint32              `json:"device_type"`     //1:手机 2:平板  3:PC  4:TV
}

type Request struct {
	Reqs            []*Request_Req  `json:"reqs"`
	App             *Request_App    `json:"app"`              // APP相关信息
	Device          *Request_Device `json:"device"`           // 设备信息
	Signature       string          `json:"signature"`        // 加密签名
	Timestamp       string          `json:"timestamp"`        // 时间戳
	CachedCreatives []int32         `json:"cached_creatives"` // 已经预加载的创意ID列表
}

type Response_Impression struct {
	SlotId             string  `json:"slot_id"`              // 對應的 Req.slot_id
	RespId             string  `json:"resp_id"`              // 廣告回應的隨機編號,用於事件追蹤、Debug
	Adid               uint32  `json:"adid"`                 // Crystal Express 平台的廣告 ID
	CreativeId         uint32  `json:"creative_id"`          // 廣告創意 ID
	CreativeUpdateTime uint32  `json:"creative_update_time"` // 廣告創意更新時間
	StartTime          uint32  `json:"start_time"`           // 廣告開始時間
	EndTime            uint32  `json:"end_time"`             // 廣告結束時間 (SDK 需緩存的期限)
	TzOffset           uint32  `json:"tz_offset"`            // 時區設定, 配合 play_time 判斷, 不給代表使用裝置時區
	PlayTime           []int32 `json:"play_time"`            // 數組長度為7,分別表示週日至週六7天,每個元素從低至高位的24bit分別表示0至23小時,bit位為1表示該小時允許投放
	DeadLine           uint32  `json:"deadline"`             // 此次廣告設定有效期限
	AllowedImps        uint32  `json:"allowed_imps"`         // 在廣告有效期限內可允許的曝光量
	FreqCap            uint32  `json:"freq_cap"`             // 在廣告有效期限內可每小時的頻控
	Priority           uint32  `json:"priority"`             // 廣告播放權重 [1 ~ 64], 越大越重要
	Ext                string  `json:"ext"`
}

type Response struct {
	ImpResps []*Response_Impression `json:"imp_resps"`
}

func (m *Request) GetReqs() []*Request_Req {
	if m != nil {
		return m.Reqs
	}
	return nil
}

func (m *Request) GetApp() *Request_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Request) GetDevice() *Request_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *Request) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Request) GetCachedCreatives() []int32 {
	if m != nil {
		return m.CachedCreatives
	}
	return nil
}

func (m *Request) GetDeviceGeo() *Request_Device_Geo {
	if m == nil || m.Device == nil {
		return nil
	}
	return m.Device.Geo
}

func (m *Request) GetDeviceMake() string {
	if m != nil && m.Device != nil {
		return m.Device.Make
	}
	return ""
}

func (m *Request) GetDeviceModel() string {
	if m != nil && m.Device != nil {
		return m.Device.Model
	}
	return ""
}

func (m *Request) GetDeviceOs() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Os
	}
	return 0
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

func (m *Request) GetDeviceIdfa() string {
	if m != nil && m.Device != nil {
		return m.Device.Idfa
	}
	return ""
}

func (m *Request) GetDeviceImei() string {
	if m != nil && m.Device != nil {
		return m.Device.Imei
	}
	return ""
}

func (m *Request) GetDeviceMac() string {
	if m != nil && m.Device != nil {
		return m.Device.Mac
	}
	return ""
}

func (m *Request) GetDeviceAndroidid() string {
	if m != nil && m.Device != nil {
		return m.Device.Androidid
	}
	return ""
}

func (m *Request) GetDeviceScreenWidth() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.ScreenWidth
	}
	return 0
}

func (m *Request) GetDeviceScreenHeight() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.ScreenHeight
	}
	return 0
}

func (m *Request) GetDeviceConnType() uint32 {
	if m != nil && m.Device != nil {
		return m.Device.Connectiontype
	}
	return 0
}

func (m *Request) GetAppId() string {
	if m != nil && m.App != nil {
		return m.App.Id
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

func (m *Request_Req) GetSlotId() string {
	if m == nil {
		return ""
	}
	return m.SlotId
}

func (m *Request_Req) GetReqType() uint32 {
	if m == nil {
		return 0
	}
	return m.Type
}

func (m *Request_Req) GetReqCount() uint32 {
	if m == nil {
		return 0
	}
	return m.Count
}
