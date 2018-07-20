// Code generated by protoc-gen-go.
// source: momo.proto
// DO NOT EDIT!

/*
Package momo is a generated protocol buffer package.

It is generated from these files:
	momo.proto

It has these top-level messages:
	BidRequest
	BidResponse
*/
package momo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 可展示的原生广告样式,不同样式对应的素材规格见广告样式章节.
type NativeFormat int32

const (
	// 大图: 一张520*260图片, 三图:三张250*250图片
	NativeFormat_FEED_LANDING_PAGE_LARGE_IMG NativeFormat = 1
	NativeFormat_FEED_LANDING_PAGE_SMALL_IMG NativeFormat = 2
	NativeFormat_NEARBY_LANDING_PAGE_NO_IMG  NativeFormat = 3
	NativeFormat_FEED_APP_IOS_LARGE_IMG      NativeFormat = 4
	NativeFormat_FEED_APP_IOS_SMALL_IMG      NativeFormat = 5
	NativeFormat_NEARBY_APP_IOS_NO_IMG       NativeFormat = 6
	NativeFormat_FEED_APP_ANDROID_LARGE_IMG  NativeFormat = 7
	NativeFormat_FEED_APP_ANDROID_SMALL_IMG  NativeFormat = 8
	NativeFormat_NEARBY_APP_ANDROID_NO_IMG   NativeFormat = 9
)

var NativeFormat_name = map[int32]string{
	1: "FEED_LANDING_PAGE_LARGE_IMG",
	2: "FEED_LANDING_PAGE_SMALL_IMG",
	3: "NEARBY_LANDING_PAGE_NO_IMG",
	4: "FEED_APP_IOS_LARGE_IMG",
	5: "FEED_APP_IOS_SMALL_IMG",
	6: "NEARBY_APP_IOS_NO_IMG",
	7: "FEED_APP_ANDROID_LARGE_IMG",
	8: "FEED_APP_ANDROID_SMALL_IMG",
	9: "NEARBY_APP_ANDROID_NO_IMG",
}
var NativeFormat_value = map[string]int32{
	"FEED_LANDING_PAGE_LARGE_IMG": 1,
	"FEED_LANDING_PAGE_SMALL_IMG": 2,
	"NEARBY_LANDING_PAGE_NO_IMG":  3,
	"FEED_APP_IOS_LARGE_IMG":      4,
	"FEED_APP_IOS_SMALL_IMG":      5,
	"NEARBY_APP_IOS_NO_IMG":       6,
	"FEED_APP_ANDROID_LARGE_IMG":  7,
	"FEED_APP_ANDROID_SMALL_IMG":  8,
	"NEARBY_APP_ANDROID_NO_IMG":   9,
}

func (x NativeFormat) Enum() *NativeFormat {
	p := new(NativeFormat)
	*p = x
	return p
}
func (x NativeFormat) String() string {
	return proto.EnumName(NativeFormat_name, int32(x))
}
func (x *NativeFormat) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NativeFormat_value, data, "NativeFormat")
	if err != nil {
		return err
	}
	*x = NativeFormat(value)
	return nil
}

// 网络类型
type BidRequest_Device_ConnectionType int32

const (
	BidRequest_Device_CONNECTION_UNKNOWN BidRequest_Device_ConnectionType = 0
	BidRequest_Device_ETHERNET           BidRequest_Device_ConnectionType = 1
	BidRequest_Device_WIFI               BidRequest_Device_ConnectionType = 2
	BidRequest_Device_CELL_UNKNOWN       BidRequest_Device_ConnectionType = 3
	BidRequest_Device_CELL_2G            BidRequest_Device_ConnectionType = 4
	BidRequest_Device_CELL_3G            BidRequest_Device_ConnectionType = 5
	BidRequest_Device_CELL_4G            BidRequest_Device_ConnectionType = 6
)

var BidRequest_Device_ConnectionType_name = map[int32]string{
	0: "CONNECTION_UNKNOWN",
	1: "ETHERNET",
	2: "WIFI",
	3: "CELL_UNKNOWN",
	4: "CELL_2G",
	5: "CELL_3G",
	6: "CELL_4G",
}
var BidRequest_Device_ConnectionType_value = map[string]int32{
	"CONNECTION_UNKNOWN": 0,
	"ETHERNET":           1,
	"WIFI":               2,
	"CELL_UNKNOWN":       3,
	"CELL_2G":            4,
	"CELL_3G":            5,
	"CELL_4G":            6,
}

func (x BidRequest_Device_ConnectionType) Enum() *BidRequest_Device_ConnectionType {
	p := new(BidRequest_Device_ConnectionType)
	*p = x
	return p
}
func (x BidRequest_Device_ConnectionType) String() string {
	return proto.EnumName(BidRequest_Device_ConnectionType_name, int32(x))
}
func (x *BidRequest_Device_ConnectionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Device_ConnectionType_value, data, "BidRequest_Device_ConnectionType")
	if err != nil {
		return err
	}
	*x = BidRequest_Device_ConnectionType(value)
	return nil
}

// 性别
type BidRequest_User_Gender int32

const (
	BidRequest_User_MALE   BidRequest_User_Gender = 1
	BidRequest_User_FEMALE BidRequest_User_Gender = 2
	BidRequest_User_OTHER  BidRequest_User_Gender = 3
)

var BidRequest_User_Gender_name = map[int32]string{
	1: "MALE",
	2: "FEMALE",
	3: "OTHER",
}
var BidRequest_User_Gender_value = map[string]int32{
	"MALE":   1,
	"FEMALE": 2,
	"OTHER":  3,
}

func (x BidRequest_User_Gender) Enum() *BidRequest_User_Gender {
	p := new(BidRequest_User_Gender)
	*p = x
	return p
}
func (x BidRequest_User_Gender) String() string {
	return proto.EnumName(BidRequest_User_Gender_name, int32(x))
}
func (x *BidRequest_User_Gender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_User_Gender_value, data, "BidRequest_User_Gender")
	if err != nil {
		return err
	}
	*x = BidRequest_User_Gender(value)
	return nil
}

// 不参与竞价的原因
type BidResponse_NoBidReason int32

const (
	BidResponse_UNKNOWN_ERROR              BidResponse_NoBidReason = 0
	BidResponse_TECHNICAL_ERROR            BidResponse_NoBidReason = 1
	BidResponse_INVALID_REQUEST            BidResponse_NoBidReason = 2
	BidResponse_KNOWN_WEB_SPIDER           BidResponse_NoBidReason = 3
	BidResponse_SUSPECTED_NONHUMAN_TRAFFIC BidResponse_NoBidReason = 4
	BidResponse_CLOUD_DATACENTER_PROXYIP   BidResponse_NoBidReason = 5
	BidResponse_UNSUPPORTED_DEVICE         BidResponse_NoBidReason = 6
	BidResponse_BLOCKED_PUBLISHER          BidResponse_NoBidReason = 7
	BidResponse_UNMATCHED_USER             BidResponse_NoBidReason = 8
)

var BidResponse_NoBidReason_name = map[int32]string{
	0: "UNKNOWN_ERROR",
	1: "TECHNICAL_ERROR",
	2: "INVALID_REQUEST",
	3: "KNOWN_WEB_SPIDER",
	4: "SUSPECTED_NONHUMAN_TRAFFIC",
	5: "CLOUD_DATACENTER_PROXYIP",
	6: "UNSUPPORTED_DEVICE",
	7: "BLOCKED_PUBLISHER",
	8: "UNMATCHED_USER",
}
var BidResponse_NoBidReason_value = map[string]int32{
	"UNKNOWN_ERROR":              0,
	"TECHNICAL_ERROR":            1,
	"INVALID_REQUEST":            2,
	"KNOWN_WEB_SPIDER":           3,
	"SUSPECTED_NONHUMAN_TRAFFIC": 4,
	"CLOUD_DATACENTER_PROXYIP":   5,
	"UNSUPPORTED_DEVICE":         6,
	"BLOCKED_PUBLISHER":          7,
	"UNMATCHED_USER":             8,
}

func (x BidResponse_NoBidReason) Enum() *BidResponse_NoBidReason {
	p := new(BidResponse_NoBidReason)
	*p = x
	return p
}
func (x BidResponse_NoBidReason) String() string {
	return proto.EnumName(BidResponse_NoBidReason_name, int32(x))
}
func (x *BidResponse_NoBidReason) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidResponse_NoBidReason_value, data, "BidResponse_NoBidReason")
	if err != nil {
		return err
	}
	*x = BidResponse_NoBidReason(value)
	return nil
}

type BidRequest struct {
	// 竞价请求的唯一标识
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// API的版本号
	Version *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	// 是否是测试请求.测试请求的广告不会计费和展示
	IsTest *bool `protobuf:"varint,3,opt,name=is_test,def=0" json:"is_test,omitempty"`
	// 是否是Ping请求.Ping请求不会触发竞价
	IsPing           *bool              `protobuf:"varint,4,opt,name=is_ping,def=0" json:"is_ping,omitempty"`
	Imp              []*BidRequest_Imp  `protobuf:"bytes,5,rep,name=imp" json:"imp,omitempty"`
	Device           *BidRequest_Device `protobuf:"bytes,6,opt,name=device" json:"device,omitempty"`
	App              *BidRequest_App    `protobuf:"bytes,7,opt,name=app" json:"app,omitempty"`
	User             *BidRequest_User   `protobuf:"bytes,8,opt,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *BidRequest) Reset()         { *m = BidRequest{} }
func (m *BidRequest) String() string { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()    {}

const Default_BidRequest_IsTest bool = false
const Default_BidRequest_IsPing bool = false

func (m *BidRequest) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *BidRequest) GetIsTest() bool {
	if m != nil && m.IsTest != nil {
		return *m.IsTest
	}
	return Default_BidRequest_IsTest
}

func (m *BidRequest) GetIsPing() bool {
	if m != nil && m.IsPing != nil {
		return *m.IsPing
	}
	return Default_BidRequest_IsPing
}

func (m *BidRequest) GetImp() []*BidRequest_Imp {
	if m != nil {
		return m.Imp
	}
	return nil
}

func (m *BidRequest) GetDevice() *BidRequest_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *BidRequest) GetApp() *BidRequest_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *BidRequest) GetUser() *BidRequest_User {
	if m != nil {
		return m.User
	}
	return nil
}

// Impression信息.不同广告位对应不同Imp,一次竞价请求至少包含一个Imp对象
type BidRequest_Imp struct {
	// Imp的ID
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// 广告位ID
	Slotid *string `protobuf:"bytes,2,req,name=slotid" json:"slotid,omitempty"`
	// 底价,单位元,小数点后最多两位
	Bidfloor *float64               `protobuf:"fixed64,3,opt,name=bidfloor" json:"bidfloor,omitempty"`
	Native   *BidRequest_Imp_Native `protobuf:"bytes,4,opt,name=native" json:"native,omitempty"`
	// 屏蔽的广告主类别
	Bcat []string `protobuf:"bytes,5,rep,name=bcat" json:"bcat,omitempty"`
	// 屏蔽的广告主domain
	Badv             []string `protobuf:"bytes,6,rep,name=badv" json:"badv,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BidRequest_Imp) Reset()         { *m = BidRequest_Imp{} }
func (m *BidRequest_Imp) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Imp) ProtoMessage()    {}

func (m *BidRequest_Imp) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest_Imp) GetSlotid() string {
	if m != nil && m.Slotid != nil {
		return *m.Slotid
	}
	return ""
}

func (m *BidRequest_Imp) GetBidfloor() float64 {
	if m != nil && m.Bidfloor != nil {
		return *m.Bidfloor
	}
	return 0
}

func (m *BidRequest_Imp) GetNative() *BidRequest_Imp_Native {
	if m != nil {
		return m.Native
	}
	return nil
}

func (m *BidRequest_Imp) GetBcat() []string {
	if m != nil {
		return m.Bcat
	}
	return nil
}

func (m *BidRequest_Imp) GetBadv() []string {
	if m != nil {
		return m.Badv
	}
	return nil
}

// 原生广告对象. 1.0版本API只支持原生广告竞价,此项必填.
type BidRequest_Imp_Native struct {
	NativeFormat     []NativeFormat `protobuf:"varint,1,rep,name=native_format,enum=momo.NativeFormat" json:"native_format,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *BidRequest_Imp_Native) Reset()         { *m = BidRequest_Imp_Native{} }
func (m *BidRequest_Imp_Native) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Imp_Native) ProtoMessage()    {}

func (m *BidRequest_Imp_Native) GetNativeFormat() []NativeFormat {
	if m != nil {
		return m.NativeFormat
	}
	return nil
}

// 设备信息
type BidRequest_Device struct {
	// 操作系统类型, iOS或Android
	Os *string `protobuf:"bytes,1,opt,name=os" json:"os,omitempty"`
	// 操作系统版本,如9.2
	Osv *string `protobuf:"bytes,2,opt,name=osv" json:"osv,omitempty"`
	// 设备型号,例如Nexus,iphone
	Model *string `protobuf:"bytes,3,opt,name=model" json:"model,omitempty"`
	// 手机制造商,例如apple,Samsung
	Make *string `protobuf:"bytes,4,opt,name=make" json:"make,omitempty"`
	// IP地址
	Ip *string `protobuf:"bytes,5,opt,name=ip" json:"ip,omitempty"`
	// MD5加密的设备MAC地址
	Macmd5 *string `protobuf:"bytes,6,opt,name=macmd5" json:"macmd5,omitempty"`
	// 设备号.安卓系统为IMEI, iOS系统为IDFA
	Did *string `protobuf:"bytes,7,opt,name=did" json:"did,omitempty"`
	// MD5加密的设备号.安卓系统为加密的IMEI, iOS系统为加密的IDFA
	Didmd5         *string                           `protobuf:"bytes,8,opt,name=didmd5" json:"didmd5,omitempty"`
	Connectiontype *BidRequest_Device_ConnectionType `protobuf:"varint,9,opt,name=connectiontype,enum=momo.BidRequest_Device_ConnectionType" json:"connectiontype,omitempty"`
	// user agent
	Ua *string `protobuf:"bytes,10,opt,name=ua" json:"ua,omitempty"`
	// 设备当前的位置信息
	Geo              *BidRequest_Geo `protobuf:"bytes,11,opt,name=geo" json:"geo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BidRequest_Device) Reset()         { *m = BidRequest_Device{} }
func (m *BidRequest_Device) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Device) ProtoMessage()    {}

func (m *BidRequest_Device) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *BidRequest_Device) GetOsv() string {
	if m != nil && m.Osv != nil {
		return *m.Osv
	}
	return ""
}

func (m *BidRequest_Device) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *BidRequest_Device) GetMake() string {
	if m != nil && m.Make != nil {
		return *m.Make
	}
	return ""
}

func (m *BidRequest_Device) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *BidRequest_Device) GetMacmd5() string {
	if m != nil && m.Macmd5 != nil {
		return *m.Macmd5
	}
	return ""
}

func (m *BidRequest_Device) GetDid() string {
	if m != nil && m.Did != nil {
		return *m.Did
	}
	return ""
}

func (m *BidRequest_Device) GetDidmd5() string {
	if m != nil && m.Didmd5 != nil {
		return *m.Didmd5
	}
	return ""
}

func (m *BidRequest_Device) GetConnectiontype() BidRequest_Device_ConnectionType {
	if m != nil && m.Connectiontype != nil {
		return *m.Connectiontype
	}
	return BidRequest_Device_CONNECTION_UNKNOWN
}

func (m *BidRequest_Device) GetUa() string {
	if m != nil && m.Ua != nil {
		return *m.Ua
	}
	return ""
}

func (m *BidRequest_Device) GetGeo() *BidRequest_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

// 应用信息. 当前所有流量来自陌陌
type BidRequest_App struct {
	// 应用ID
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// 应用名称
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// 版本号,例6.5.0
	Ver *string `protobuf:"bytes,3,opt,name=ver" json:"ver,omitempty"`
	// 应用标识, iOS为iTunes Store ID, 安卓为package name.
	Bundle           *string `protobuf:"bytes,4,opt,name=bundle" json:"bundle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidRequest_App) Reset()         { *m = BidRequest_App{} }
func (m *BidRequest_App) String() string { return proto.CompactTextString(m) }
func (*BidRequest_App) ProtoMessage()    {}

func (m *BidRequest_App) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest_App) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *BidRequest_App) GetVer() string {
	if m != nil && m.Ver != nil {
		return *m.Ver
	}
	return ""
}

func (m *BidRequest_App) GetBundle() string {
	if m != nil && m.Bundle != nil {
		return *m.Bundle
	}
	return ""
}

// 用户信息
type BidRequest_User struct {
	// Momo Exchange提供的用户唯一标识
	Id     *string                 `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Gender *BidRequest_User_Gender `protobuf:"varint,2,opt,name=gender,enum=momo.BidRequest_User_Gender" json:"gender,omitempty"`
	// 年龄范围
	AgeLow  *int32 `protobuf:"varint,3,opt,name=age_low,def=0" json:"age_low,omitempty"`
	AgeHigh *int32 `protobuf:"varint,4,opt,name=age_high,def=999" json:"age_high,omitempty"`
	// 来自不同数据源的用户数据
	Data []*BidRequest_Data `protobuf:"bytes,8,rep,name=data" json:"data,omitempty"`
	// 用户兴趣关键词,逗号分隔
	Keywords *string `protobuf:"bytes,6,opt,name=keywords" json:"keywords,omitempty"`
	// 用户的常住地址,非实时
	Geo              *BidRequest_Geo `protobuf:"bytes,7,opt,name=geo" json:"geo,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BidRequest_User) Reset()         { *m = BidRequest_User{} }
func (m *BidRequest_User) String() string { return proto.CompactTextString(m) }
func (*BidRequest_User) ProtoMessage()    {}

const Default_BidRequest_User_AgeLow int32 = 0
const Default_BidRequest_User_AgeHigh int32 = 999

func (m *BidRequest_User) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest_User) GetGender() BidRequest_User_Gender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return BidRequest_User_MALE
}

func (m *BidRequest_User) GetAgeLow() int32 {
	if m != nil && m.AgeLow != nil {
		return *m.AgeLow
	}
	return Default_BidRequest_User_AgeLow
}

func (m *BidRequest_User) GetAgeHigh() int32 {
	if m != nil && m.AgeHigh != nil {
		return *m.AgeHigh
	}
	return Default_BidRequest_User_AgeHigh
}

func (m *BidRequest_User) GetData() []*BidRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BidRequest_User) GetKeywords() string {
	if m != nil && m.Keywords != nil {
		return *m.Keywords
	}
	return ""
}

func (m *BidRequest_User) GetGeo() *BidRequest_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

type BidRequest_Geo struct {
	// 经纬度
	Lat *float64 `protobuf:"fixed64,1,opt,name=lat" json:"lat,omitempty"`
	Lon *float64 `protobuf:"fixed64,2,opt,name=lon" json:"lon,omitempty"`
	// 省
	Province *string `protobuf:"bytes,4,opt,name=province" json:"province,omitempty"`
	// 市
	City *string `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	// 区
	District *string `protobuf:"bytes,6,opt,name=district" json:"district,omitempty"`
	// 街道
	Street           *string `protobuf:"bytes,7,opt,name=street" json:"street,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidRequest_Geo) Reset()         { *m = BidRequest_Geo{} }
func (m *BidRequest_Geo) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Geo) ProtoMessage()    {}

func (m *BidRequest_Geo) GetLat() float64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *BidRequest_Geo) GetLon() float64 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

func (m *BidRequest_Geo) GetProvince() string {
	if m != nil && m.Province != nil {
		return *m.Province
	}
	return ""
}

func (m *BidRequest_Geo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *BidRequest_Geo) GetDistrict() string {
	if m != nil && m.District != nil {
		return *m.District
	}
	return ""
}

func (m *BidRequest_Geo) GetStreet() string {
	if m != nil && m.Street != nil {
		return *m.Street
	}
	return ""
}

type BidRequest_Data struct {
	// 数据提供商的ID
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// 数据提供商的名称
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Segment数组
	Segment          []*BidRequest_Data_Segment `protobuf:"bytes,3,rep,name=segment" json:"segment,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *BidRequest_Data) Reset()         { *m = BidRequest_Data{} }
func (m *BidRequest_Data) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Data) ProtoMessage()    {}

func (m *BidRequest_Data) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest_Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *BidRequest_Data) GetSegment() []*BidRequest_Data_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

// Key-value形式的Segment数据
type BidRequest_Data_Segment struct {
	// Segment的ID
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Segment的名称
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Segment的值
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidRequest_Data_Segment) Reset()         { *m = BidRequest_Data_Segment{} }
func (m *BidRequest_Data_Segment) String() string { return proto.CompactTextString(m) }
func (*BidRequest_Data_Segment) ProtoMessage()    {}

func (m *BidRequest_Data_Segment) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidRequest_Data_Segment) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *BidRequest_Data_Segment) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type BidResponse struct {
	// BidRequest的ID
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// 不参与竞价的原因. DSP不参与竞价时可以通过该字段告知原因
	Nbr *BidResponse_NoBidReason `protobuf:"varint,2,opt,name=nbr,enum=momo.BidResponse_NoBidReason" json:"nbr,omitempty"`
	// 不参与竞价时,DSP提供的错误信息详细内容,供双方调试用
	Em *string `protobuf:"bytes,3,opt,name=em" json:"em,omitempty"`
	// 来自不同Seat的竞价信息
	Seatbid []*BidResponse_SeatBid `protobuf:"bytes,4,rep,name=seatbid" json:"seatbid,omitempty"`
	// DSP生成的bid id
	Bidid            *string `protobuf:"bytes,5,opt,name=bidid" json:"bidid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidResponse) Reset()         { *m = BidResponse{} }
func (m *BidResponse) String() string { return proto.CompactTextString(m) }
func (*BidResponse) ProtoMessage()    {}

func (m *BidResponse) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidResponse) GetNbr() BidResponse_NoBidReason {
	if m != nil && m.Nbr != nil {
		return *m.Nbr
	}
	return BidResponse_UNKNOWN_ERROR
}

func (m *BidResponse) GetEm() string {
	if m != nil && m.Em != nil {
		return *m.Em
	}
	return ""
}

func (m *BidResponse) GetSeatbid() []*BidResponse_SeatBid {
	if m != nil {
		return m.Seatbid
	}
	return nil
}

func (m *BidResponse) GetBidid() string {
	if m != nil && m.Bidid != nil {
		return *m.Bidid
	}
	return ""
}

type BidResponse_SeatBid struct {
	// 针对Imp的Bid信息
	Bid []*BidResponse_SeatBid_Bid `protobuf:"bytes,1,rep,name=bid" json:"bid,omitempty"`
	// DSP的Seat标识
	Seat             *string `protobuf:"bytes,2,req,name=seat" json:"seat,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidResponse_SeatBid) Reset()         { *m = BidResponse_SeatBid{} }
func (m *BidResponse_SeatBid) String() string { return proto.CompactTextString(m) }
func (*BidResponse_SeatBid) ProtoMessage()    {}

func (m *BidResponse_SeatBid) GetBid() []*BidResponse_SeatBid_Bid {
	if m != nil {
		return m.Bid
	}
	return nil
}

func (m *BidResponse_SeatBid) GetSeat() string {
	if m != nil && m.Seat != nil {
		return *m.Seat
	}
	return ""
}

// 广告实体
type BidResponse_SeatBid_Bid struct {
	// DSP生成的ID
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// BidRequest中的Impression ID
	Impid *string `protobuf:"bytes,2,req,name=impid" json:"impid,omitempty"`
	// 出价,单位元,小数点后最多两位,多余位数截断
	Price *float64 `protobuf:"fixed64,3,req,name=price" json:"price,omitempty"`
	// Campaign ID
	Cid *string `protobuf:"bytes,4,opt,name=cid" json:"cid,omitempty"`
	// 广告 ID
	Adid *string `protobuf:"bytes,5,opt,name=adid" json:"adid,omitempty"`
	// 广告创意 ID
	Crid *string `protobuf:"bytes,6,req,name=crid" json:"crid,omitempty"`
	// 广告的类目category
	Cat []string `protobuf:"bytes,7,rep,name=cat" json:"cat,omitempty"`
	// 广告主的Domain
	Adomain []string `protobuf:"bytes,8,rep,name=adomain" json:"adomain,omitempty"`
	// Win notice的URL,支持宏替换.DSP可不提供该字段,而是通过曝光监测链接宏替换获取win price,具体见宏替换说明
	Nurl *string `protobuf:"bytes,9,opt,name=nurl" json:"nurl,omitempty"`
	// 应用下载广告的应用标识, iOS为iTunes Store ID, 安卓为package name.
	Bundle         *string                                 `protobuf:"bytes,10,opt,name=bundle" json:"bundle,omitempty"`
	NativeCreative *BidResponse_SeatBid_Bid_NativeCreative `protobuf:"bytes,11,opt,name=native_creative" json:"native_creative,omitempty"`
	// 曝光监测地址,广告曝光时客户端会在后台ping该地址.支持宏替换(设备信息,win price等),具体见接口文档
	Imptrackers []string `protobuf:"bytes,12,rep,name=imptrackers" json:"imptrackers,omitempty"`
	// 点击监测地址,广告点击时客户端会在后台ping该地址.支持宏替换(设备信息,win price等),具体见接口文档
	Clicktrackers    []string `protobuf:"bytes,13,rep,name=clicktrackers" json:"clicktrackers,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BidResponse_SeatBid_Bid) Reset()         { *m = BidResponse_SeatBid_Bid{} }
func (m *BidResponse_SeatBid_Bid) String() string { return proto.CompactTextString(m) }
func (*BidResponse_SeatBid_Bid) ProtoMessage()    {}

func (m *BidResponse_SeatBid_Bid) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetImpid() string {
	if m != nil && m.Impid != nil {
		return *m.Impid
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *BidResponse_SeatBid_Bid) GetCid() string {
	if m != nil && m.Cid != nil {
		return *m.Cid
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetAdid() string {
	if m != nil && m.Adid != nil {
		return *m.Adid
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetCrid() string {
	if m != nil && m.Crid != nil {
		return *m.Crid
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetCat() []string {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid) GetAdomain() []string {
	if m != nil {
		return m.Adomain
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid) GetNurl() string {
	if m != nil && m.Nurl != nil {
		return *m.Nurl
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetBundle() string {
	if m != nil && m.Bundle != nil {
		return *m.Bundle
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid) GetNativeCreative() *BidResponse_SeatBid_Bid_NativeCreative {
	if m != nil {
		return m.NativeCreative
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid) GetImptrackers() []string {
	if m != nil {
		return m.Imptrackers
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid) GetClicktrackers() []string {
	if m != nil {
		return m.Clicktrackers
	}
	return nil
}

// 原生广告的素材信息
type BidResponse_SeatBid_Bid_NativeCreative struct {
	// 素材对应的广告样式,必须是BidRequest中所给的样式之一
	NativeFormat *NativeFormat `protobuf:"varint,1,req,name=native_format,enum=momo.NativeFormat" json:"native_format,omitempty"`
	// 广告标题,品牌名或应用名称
	Title *string `protobuf:"bytes,2,req,name=title" json:"title,omitempty"`
	// 广告推广语
	Desc *string `protobuf:"bytes,3,req,name=desc" json:"desc,omitempty"`
	// 广告Logo
	Logo *BidResponse_SeatBid_Bid_NativeCreative_Image `protobuf:"bytes,4,req,name=logo" json:"logo,omitempty"`
	// 广告详情图. 大图样式广告必须返回一个image对象, 三图版返回三个image对象,无图版不需要
	Image []*BidResponse_SeatBid_Bid_NativeCreative_Image `protobuf:"bytes,5,rep,name=image" json:"image,omitempty"`
	// 落地页广告的点击地址，支持中间/监测页跳转
	LandingpageUrl *string `protobuf:"bytes,6,opt,name=landingpage_url" json:"landingpage_url,omitempty"`
	// 应用下载广告的下载地址.
	// 安卓必须为.apk结尾地址,且必须同时提供app_intro_url.点击"安装"按钮请求app_download_url,点击其他区域请求app_intro_url
	// iOS支持中间页跳转.点击广告所有区域均请求该地址.应用下载广告必填
	AppDownloadUrl *string `protobuf:"bytes,7,opt,name=app_download_url" json:"app_download_url,omitempty"`
	// 安卓应用的介绍页面.点击安装按钮请求app_download_url,点击其他区域请求app_intro_url.安卓应用下载广告必填
	AppIntroUrl *string `protobuf:"bytes,8,opt,name=app_intro_url" json:"app_intro_url,omitempty"`
	// 应用版本，只能由半角“.”符号与阿拉伯数字组成，例如12.9.3。应用下载广告必填
	AppVer *string `protobuf:"bytes,9,opt,name=app_ver" json:"app_ver,omitempty"`
	// 应用大小,单位MB,应用下载广告必填
	AppSize          *float64 `protobuf:"fixed64,10,opt,name=app_size" json:"app_size,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) Reset() {
	*m = BidResponse_SeatBid_Bid_NativeCreative{}
}
func (m *BidResponse_SeatBid_Bid_NativeCreative) String() string { return proto.CompactTextString(m) }
func (*BidResponse_SeatBid_Bid_NativeCreative) ProtoMessage()    {}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetNativeFormat() NativeFormat {
	if m != nil && m.NativeFormat != nil {
		return *m.NativeFormat
	}
	return NativeFormat_FEED_LANDING_PAGE_LARGE_IMG
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetLogo() *BidResponse_SeatBid_Bid_NativeCreative_Image {
	if m != nil {
		return m.Logo
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetImage() []*BidResponse_SeatBid_Bid_NativeCreative_Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetLandingpageUrl() string {
	if m != nil && m.LandingpageUrl != nil {
		return *m.LandingpageUrl
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetAppDownloadUrl() string {
	if m != nil && m.AppDownloadUrl != nil {
		return *m.AppDownloadUrl
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetAppIntroUrl() string {
	if m != nil && m.AppIntroUrl != nil {
		return *m.AppIntroUrl
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetAppVer() string {
	if m != nil && m.AppVer != nil {
		return *m.AppVer
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative) GetAppSize() float64 {
	if m != nil && m.AppSize != nil {
		return *m.AppSize
	}
	return 0
}

// 支持PNG/JPG格式的图片
type BidResponse_SeatBid_Bid_NativeCreative_Image struct {
	// 图片的地址
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// 图片的宽高. 在保证宽高比例的情况下,可以大于竞价请求中指定的大小.
	Width            *int32 `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height           *int32 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BidResponse_SeatBid_Bid_NativeCreative_Image) Reset() {
	*m = BidResponse_SeatBid_Bid_NativeCreative_Image{}
}
func (m *BidResponse_SeatBid_Bid_NativeCreative_Image) String() string {
	return proto.CompactTextString(m)
}
func (*BidResponse_SeatBid_Bid_NativeCreative_Image) ProtoMessage() {}

func (m *BidResponse_SeatBid_Bid_NativeCreative_Image) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *BidResponse_SeatBid_Bid_NativeCreative_Image) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *BidResponse_SeatBid_Bid_NativeCreative_Image) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*BidRequest)(nil), "momo.BidRequest")
	proto.RegisterType((*BidRequest_Imp)(nil), "momo.BidRequest.Imp")
	proto.RegisterType((*BidRequest_Imp_Native)(nil), "momo.BidRequest.Imp.Native")
	proto.RegisterType((*BidRequest_Device)(nil), "momo.BidRequest.Device")
	proto.RegisterType((*BidRequest_App)(nil), "momo.BidRequest.App")
	proto.RegisterType((*BidRequest_User)(nil), "momo.BidRequest.User")
	proto.RegisterType((*BidRequest_Geo)(nil), "momo.BidRequest.Geo")
	proto.RegisterType((*BidRequest_Data)(nil), "momo.BidRequest.Data")
	proto.RegisterType((*BidRequest_Data_Segment)(nil), "momo.BidRequest.Data.Segment")
	proto.RegisterType((*BidResponse)(nil), "momo.BidResponse")
	proto.RegisterType((*BidResponse_SeatBid)(nil), "momo.BidResponse.SeatBid")
	proto.RegisterType((*BidResponse_SeatBid_Bid)(nil), "momo.BidResponse.SeatBid.Bid")
	proto.RegisterType((*BidResponse_SeatBid_Bid_NativeCreative)(nil), "momo.BidResponse.SeatBid.Bid.NativeCreative")
	proto.RegisterType((*BidResponse_SeatBid_Bid_NativeCreative_Image)(nil), "momo.BidResponse.SeatBid.Bid.NativeCreative.Image")
	proto.RegisterEnum("momo.NativeFormat", NativeFormat_name, NativeFormat_value)
	proto.RegisterEnum("momo.BidRequest_Device_ConnectionType", BidRequest_Device_ConnectionType_name, BidRequest_Device_ConnectionType_value)
	proto.RegisterEnum("momo.BidRequest_User_Gender", BidRequest_User_Gender_name, BidRequest_User_Gender_value)
	proto.RegisterEnum("momo.BidResponse_NoBidReason", BidResponse_NoBidReason_name, BidResponse_NoBidReason_value)
}