// Code generated by protoc-gen-go.
// source: jrtt.proto
// DO NOT EDIT!

/*
Package toutiao_ssp_api is a generated protocol buffer package.

It is generated from these files:
	jrtt.proto

It has these top-level messages:
	Pmp
	AdSlot
	Publisher
	Content
	App
	Geo
	Device
	Data
	User
	BidRequest
	MaterialMeta
	Bid
	SeatBid
	BidResponse
*/
package toutiao_ssp_api

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type AdType int32

const (
	AdType_TOUTIAO_FEED_LP_LARGE     AdType = 1
	AdType_TOUTIAO_FEED_LP_SMALL     AdType = 2
	AdType_TOUTIAO_FEED_APP_LARGE    AdType = 4
	AdType_TOUTIAO_FEED_APP_SMALL    AdType = 3
	AdType_TOUTIAO_DETAIL_LP_GRAPHIC AdType = 5
	AdType_TOUTIAO_DETAIL_LP_BANNER  AdType = 6
	AdType_TUANZI_FEED_APP           AdType = 7
	AdType_TOUTIAO_DETAIL_APP_BANNER AdType = 9
)

var AdType_name = map[int32]string{
	1: "TOUTIAO_FEED_LP_LARGE",
	2: "TOUTIAO_FEED_LP_SMALL",
	4: "TOUTIAO_FEED_APP_LARGE",
	3: "TOUTIAO_FEED_APP_SMALL",
	5: "TOUTIAO_DETAIL_LP_GRAPHIC",
	6: "TOUTIAO_DETAIL_LP_BANNER",
	7: "TUANZI_FEED_APP",
	9: "TOUTIAO_DETAIL_APP_BANNER",
}
var AdType_value = map[string]int32{
	"TOUTIAO_FEED_LP_LARGE":     1,
	"TOUTIAO_FEED_LP_SMALL":     2,
	"TOUTIAO_FEED_APP_LARGE":    4,
	"TOUTIAO_FEED_APP_SMALL":    3,
	"TOUTIAO_DETAIL_LP_GRAPHIC": 5,
	"TOUTIAO_DETAIL_LP_BANNER":  6,
	"TUANZI_FEED_APP":           7,
	"TOUTIAO_DETAIL_APP_BANNER": 9,
}

func (x AdType) Enum() *AdType {
	p := new(AdType)
	*p = x
	return p
}
func (x AdType) String() string {
	return proto.EnumName(AdType_name, int32(x))
}
func (x *AdType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdType_value, data, "AdType")
	if err != nil {
		return err
	}
	*x = AdType(value)
	return nil
}

type AdSlot_Position int32

const (
	AdSlot_SPLASH AdSlot_Position = 1
	AdSlot_FEED   AdSlot_Position = 2
	AdSlot_DETAIL AdSlot_Position = 4
)

var AdSlot_Position_name = map[int32]string{
	1: "SPLASH",
	2: "FEED",
	4: "DETAIL",
}
var AdSlot_Position_value = map[string]int32{
	"SPLASH": 1,
	"FEED":   2,
	"DETAIL": 4,
}

func (x AdSlot_Position) Enum() *AdSlot_Position {
	p := new(AdSlot_Position)
	*p = x
	return p
}
func (x AdSlot_Position) String() string {
	return proto.EnumName(AdSlot_Position_name, int32(x))
}
func (x *AdSlot_Position) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdSlot_Position_value, data, "AdSlot_Position")
	if err != nil {
		return err
	}
	*x = AdSlot_Position(value)
	return nil
}

type Device_ConnectionType int32

const (
	Device_Honeycomb Device_ConnectionType = 1
	Device_WIFI      Device_ConnectionType = 2
	Device_UNKNOWN   Device_ConnectionType = 3
)

var Device_ConnectionType_name = map[int32]string{
	1: "Honeycomb",
	2: "WIFI",
	3: "UNKNOWN",
}
var Device_ConnectionType_value = map[string]int32{
	"Honeycomb": 1,
	"WIFI":      2,
	"UNKNOWN":   3,
}

func (x Device_ConnectionType) Enum() *Device_ConnectionType {
	p := new(Device_ConnectionType)
	*p = x
	return p
}
func (x Device_ConnectionType) String() string {
	return proto.EnumName(Device_ConnectionType_name, int32(x))
}
func (x *Device_ConnectionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Device_ConnectionType_value, data, "Device_ConnectionType")
	if err != nil {
		return err
	}
	*x = Device_ConnectionType(value)
	return nil
}

type Device_DeviceType int32

const (
	Device_PHONE  Device_DeviceType = 1
	Device_TABLET Device_DeviceType = 2
)

var Device_DeviceType_name = map[int32]string{
	1: "PHONE",
	2: "TABLET",
}
var Device_DeviceType_value = map[string]int32{
	"PHONE":  1,
	"TABLET": 2,
}

func (x Device_DeviceType) Enum() *Device_DeviceType {
	p := new(Device_DeviceType)
	*p = x
	return p
}
func (x Device_DeviceType) String() string {
	return proto.EnumName(Device_DeviceType_name, int32(x))
}
func (x *Device_DeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Device_DeviceType_value, data, "Device_DeviceType")
	if err != nil {
		return err
	}
	*x = Device_DeviceType(value)
	return nil
}

type User_Gender int32

const (
	User_MALE    User_Gender = 1
	User_FEMALE  User_Gender = 2
	User_UNKNOWN User_Gender = 3
)

var User_Gender_name = map[int32]string{
	1: "MALE",
	2: "FEMALE",
	3: "UNKNOWN",
}
var User_Gender_value = map[string]int32{
	"MALE":    1,
	"FEMALE":  2,
	"UNKNOWN": 3,
}

func (x User_Gender) Enum() *User_Gender {
	p := new(User_Gender)
	*p = x
	return p
}
func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}
func (x *User_Gender) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(User_Gender_value, data, "User_Gender")
	if err != nil {
		return err
	}
	*x = User_Gender(value)
	return nil
}

type Pmp struct {
	PrivateAuction   *uint32     `protobuf:"varint,1,opt,name=private_auction" json:"private_auction,omitempty"`
	Deals            []*Pmp_Deal `protobuf:"bytes,2,rep,name=deals" json:"deals,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Pmp) Reset()         { *m = Pmp{} }
func (m *Pmp) String() string { return proto.CompactTextString(m) }
func (*Pmp) ProtoMessage()    {}

func (m *Pmp) GetPrivateAuction() uint32 {
	if m != nil && m.PrivateAuction != nil {
		return *m.PrivateAuction
	}
	return 0
}

func (m *Pmp) GetDeals() []*Pmp_Deal {
	if m != nil {
		return m.Deals
	}
	return nil
}

type Pmp_Deal struct {
	Id               *uint32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	BidFloor         *uint32  `protobuf:"varint,2,opt,name=bid_floor" json:"bid_floor,omitempty"`
	Wseat            []string `protobuf:"bytes,3,rep,name=wseat" json:"wseat,omitempty"`
	Wadomain         []string `protobuf:"bytes,4,rep,name=wadomain" json:"wadomain,omitempty"`
	At               *uint32  `protobuf:"varint,5,opt,name=at" json:"at,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Pmp_Deal) Reset()         { *m = Pmp_Deal{} }
func (m *Pmp_Deal) String() string { return proto.CompactTextString(m) }
func (*Pmp_Deal) ProtoMessage()    {}

func (m *Pmp_Deal) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Pmp_Deal) GetBidFloor() uint32 {
	if m != nil && m.BidFloor != nil {
		return *m.BidFloor
	}
	return 0
}

func (m *Pmp_Deal) GetWseat() []string {
	if m != nil {
		return m.Wseat
	}
	return nil
}

func (m *Pmp_Deal) GetWadomain() []string {
	if m != nil {
		return m.Wadomain
	}
	return nil
}

func (m *Pmp_Deal) GetAt() uint32 {
	if m != nil && m.At != nil {
		return *m.At
	}
	return 0
}

type AdSlot struct {
	Id               *string          `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Banner           []*AdSlot_Banner `protobuf:"bytes,2,rep,name=banner" json:"banner,omitempty"`
	AdType           []AdType         `protobuf:"varint,3,rep,name=ad_type,enum=toutiao_ssp.api.AdType" json:"ad_type,omitempty"`
	BidFloor         *uint32          `protobuf:"varint,4,opt,name=bid_floor" json:"bid_floor,omitempty"`
	Pmp              *Pmp             `protobuf:"bytes,5,opt,name=pmp" json:"pmp,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AdSlot) Reset()         { *m = AdSlot{} }
func (m *AdSlot) String() string { return proto.CompactTextString(m) }
func (*AdSlot) ProtoMessage()    {}

func (m *AdSlot) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AdSlot) GetBanner() []*AdSlot_Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *AdSlot) GetAdType() []AdType {
	if m != nil {
		return m.AdType
	}
	return nil
}

func (m *AdSlot) GetBidFloor() uint32 {
	if m != nil && m.BidFloor != nil {
		return *m.BidFloor
	}
	return 0
}

func (m *AdSlot) GetPmp() *Pmp {
	if m != nil {
		return m.Pmp
	}
	return nil
}

type AdSlot_Banner struct {
	Width            *uint32          `protobuf:"varint,1,req,name=width" json:"width,omitempty"`
	Height           *uint32          `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
	Pos              *AdSlot_Position `protobuf:"varint,3,req,name=pos,enum=toutiao_ssp.api.AdSlot_Position" json:"pos,omitempty"`
	Sequence         *string          `protobuf:"bytes,4,opt,name=sequence" json:"sequence,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AdSlot_Banner) Reset()         { *m = AdSlot_Banner{} }
func (m *AdSlot_Banner) String() string { return proto.CompactTextString(m) }
func (*AdSlot_Banner) ProtoMessage()    {}

func (m *AdSlot_Banner) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *AdSlot_Banner) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *AdSlot_Banner) GetPos() AdSlot_Position {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return AdSlot_SPLASH
}

func (m *AdSlot_Banner) GetSequence() string {
	if m != nil && m.Sequence != nil {
		return *m.Sequence
	}
	return ""
}

type Publisher struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Cat              *string `protobuf:"bytes,3,opt,name=cat" json:"cat,omitempty"`
	Domain           *string `protobuf:"bytes,4,opt,name=domain" json:"domain,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Publisher) Reset()         { *m = Publisher{} }
func (m *Publisher) String() string { return proto.CompactTextString(m) }
func (*Publisher) ProtoMessage()    {}

func (m *Publisher) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Publisher) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Publisher) GetCat() string {
	if m != nil && m.Cat != nil {
		return *m.Cat
	}
	return ""
}

func (m *Publisher) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

type Content struct {
	Id               *string           `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Title            *string           `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Series           *string           `protobuf:"bytes,3,opt,name=series" json:"series,omitempty"`
	Url              *string           `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Keywords         *string           `protobuf:"bytes,5,opt,name=keywords" json:"keywords,omitempty"`
	Contentrating    *string           `protobuf:"bytes,6,opt,name=contentrating" json:"contentrating,omitempty"`
	Userrating       *string           `protobuf:"bytes,7,opt,name=userrating" json:"userrating,omitempty"`
	Context          *string           `protobuf:"bytes,8,opt,name=context" json:"context,omitempty"`
	Producer         *Content_Producer `protobuf:"bytes,9,opt,name=producer" json:"producer,omitempty"`
	Language         *string           `protobuf:"bytes,10,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}

func (m *Content) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Content) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Content) GetSeries() string {
	if m != nil && m.Series != nil {
		return *m.Series
	}
	return ""
}

func (m *Content) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *Content) GetKeywords() string {
	if m != nil && m.Keywords != nil {
		return *m.Keywords
	}
	return ""
}

func (m *Content) GetContentrating() string {
	if m != nil && m.Contentrating != nil {
		return *m.Contentrating
	}
	return ""
}

func (m *Content) GetUserrating() string {
	if m != nil && m.Userrating != nil {
		return *m.Userrating
	}
	return ""
}

func (m *Content) GetContext() string {
	if m != nil && m.Context != nil {
		return *m.Context
	}
	return ""
}

func (m *Content) GetProducer() *Content_Producer {
	if m != nil {
		return m.Producer
	}
	return nil
}

func (m *Content) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type Content_Producer struct {
	Id               *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Cat              *string `protobuf:"bytes,3,opt,name=cat" json:"cat,omitempty"`
	Domain           *string `protobuf:"bytes,4,opt,name=domain" json:"domain,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Content_Producer) Reset()         { *m = Content_Producer{} }
func (m *Content_Producer) String() string { return proto.CompactTextString(m) }
func (*Content_Producer) ProtoMessage()    {}

func (m *Content_Producer) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Content_Producer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Content_Producer) GetCat() string {
	if m != nil && m.Cat != nil {
		return *m.Cat
	}
	return ""
}

func (m *Content_Producer) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

type App struct {
	Id               *string    `protobuf:"bytes,1,req,name=id,def=11" json:"id,omitempty"`
	Name             *string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Domain           *string    `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Ver              *string    `protobuf:"bytes,4,opt,name=ver" json:"ver,omitempty"`
	Bundle           *string    `protobuf:"bytes,5,opt,name=bundle" json:"bundle,omitempty"`
	Privacypolicy    *uint32    `protobuf:"varint,6,opt,name=privacypolicy" json:"privacypolicy,omitempty"`
	Paid             *uint32    `protobuf:"varint,7,opt,name=paid" json:"paid,omitempty"`
	Publisher        *Publisher `protobuf:"bytes,8,opt,name=publisher" json:"publisher,omitempty"`
	Content          *Content   `protobuf:"bytes,9,opt,name=content" json:"content,omitempty"`
	Keywords         *string    `protobuf:"bytes,10,opt,name=keywords" json:"keywords,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}

const Default_App_Id string = "11"

func (m *App) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return Default_App_Id
}

func (m *App) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *App) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *App) GetVer() string {
	if m != nil && m.Ver != nil {
		return *m.Ver
	}
	return ""
}

func (m *App) GetBundle() string {
	if m != nil && m.Bundle != nil {
		return *m.Bundle
	}
	return ""
}

func (m *App) GetPrivacypolicy() uint32 {
	if m != nil && m.Privacypolicy != nil {
		return *m.Privacypolicy
	}
	return 0
}

func (m *App) GetPaid() uint32 {
	if m != nil && m.Paid != nil {
		return *m.Paid
	}
	return 0
}

func (m *App) GetPublisher() *Publisher {
	if m != nil {
		return m.Publisher
	}
	return nil
}

func (m *App) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *App) GetKeywords() string {
	if m != nil && m.Keywords != nil {
		return *m.Keywords
	}
	return ""
}

type Geo struct {
	Lat              *float64 `protobuf:"fixed64,1,opt,name=lat" json:"lat,omitempty"`
	Lon              *float64 `protobuf:"fixed64,2,opt,name=lon" json:"lon,omitempty"`
	Country          *string  `protobuf:"bytes,3,opt,name=country" json:"country,omitempty"`
	Region           *string  `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
	City             *string  `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	Type             *string  `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Geo) Reset()         { *m = Geo{} }
func (m *Geo) String() string { return proto.CompactTextString(m) }
func (*Geo) ProtoMessage()    {}

func (m *Geo) GetLat() float64 {
	if m != nil && m.Lat != nil {
		return *m.Lat
	}
	return 0
}

func (m *Geo) GetLon() float64 {
	if m != nil && m.Lon != nil {
		return *m.Lon
	}
	return 0
}

func (m *Geo) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *Geo) GetRegion() string {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return ""
}

func (m *Geo) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *Geo) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

type Device struct {
	Dnt              *bool                  `protobuf:"varint,1,req,name=dnt" json:"dnt,omitempty"`
	Ua               *string                `protobuf:"bytes,2,req,name=ua" json:"ua,omitempty"`
	Ip               *string                `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Geo              *Geo                   `protobuf:"bytes,4,opt,name=geo" json:"geo,omitempty"`
	DeviceId         *string                `protobuf:"bytes,5,opt,name=device_id" json:"device_id,omitempty"`
	DeviceIdMd5      *string                `protobuf:"bytes,6,opt,name=device_id_md5" json:"device_id_md5,omitempty"`
	Carrier          *string                `protobuf:"bytes,7,opt,name=carrier" json:"carrier,omitempty"`
	Language         *string                `protobuf:"bytes,8,opt,name=language" json:"language,omitempty"`
	Make             *string                `protobuf:"bytes,9,opt,name=make" json:"make,omitempty"`
	Model            *string                `protobuf:"bytes,10,opt,name=model" json:"model,omitempty"`
	Os               *string                `protobuf:"bytes,11,opt,name=os" json:"os,omitempty"`
	Osv              *string                `protobuf:"bytes,12,opt,name=osv" json:"osv,omitempty"`
	Js               *bool                  `protobuf:"varint,13,opt,name=js" json:"js,omitempty"`
	ConnectionType   *Device_ConnectionType `protobuf:"varint,14,opt,name=connection_type,enum=toutiao_ssp.api.Device_ConnectionType" json:"connection_type,omitempty"`
	DeviceType       *Device_DeviceType     `protobuf:"varint,15,opt,name=device_type,enum=toutiao_ssp.api.Device_DeviceType" json:"device_type,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}

func (m *Device) GetDnt() bool {
	if m != nil && m.Dnt != nil {
		return *m.Dnt
	}
	return false
}

func (m *Device) GetUa() string {
	if m != nil && m.Ua != nil {
		return *m.Ua
	}
	return ""
}

func (m *Device) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Device) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Device) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Device) GetDeviceIdMd5() string {
	if m != nil && m.DeviceIdMd5 != nil {
		return *m.DeviceIdMd5
	}
	return ""
}

func (m *Device) GetCarrier() string {
	if m != nil && m.Carrier != nil {
		return *m.Carrier
	}
	return ""
}

func (m *Device) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *Device) GetMake() string {
	if m != nil && m.Make != nil {
		return *m.Make
	}
	return ""
}

func (m *Device) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Device) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *Device) GetOsv() string {
	if m != nil && m.Osv != nil {
		return *m.Osv
	}
	return ""
}

func (m *Device) GetJs() bool {
	if m != nil && m.Js != nil {
		return *m.Js
	}
	return false
}

func (m *Device) GetConnectionType() Device_ConnectionType {
	if m != nil && m.ConnectionType != nil {
		return *m.ConnectionType
	}
	return Device_Honeycomb
}

func (m *Device) GetDeviceType() Device_DeviceType {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return Device_PHONE
}

type Data struct {
	Id               *string         `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name             *string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Segment          []*Data_Segment `protobuf:"bytes,3,rep,name=segment" json:"segment,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}

func (m *Data) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Data) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Data) GetSegment() []*Data_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

type Data_Segment struct {
	Id               *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Data_Segment) Reset()         { *m = Data_Segment{} }
func (m *Data_Segment) String() string { return proto.CompactTextString(m) }
func (*Data_Segment) ProtoMessage()    {}

func (m *Data_Segment) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Data_Segment) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Data_Segment) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type User struct {
	Id               *string      `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	BuyerId          *string      `protobuf:"bytes,2,opt,name=buyer_id" json:"buyer_id,omitempty"`
	Yob              *string      `protobuf:"bytes,3,opt,name=yob" json:"yob,omitempty"`
	Gender           *User_Gender `protobuf:"varint,4,opt,name=gender,enum=toutiao_ssp.api.User_Gender" json:"gender,omitempty"`
	Keywords         *string      `protobuf:"bytes,5,opt,name=keywords" json:"keywords,omitempty"`
	Geo              *Geo         `protobuf:"bytes,6,opt,name=geo" json:"geo,omitempty"`
	Data             []*Data      `protobuf:"bytes,7,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *User) GetBuyerId() string {
	if m != nil && m.BuyerId != nil {
		return *m.BuyerId
	}
	return ""
}

func (m *User) GetYob() string {
	if m != nil && m.Yob != nil {
		return *m.Yob
	}
	return ""
}

func (m *User) GetGender() User_Gender {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return User_MALE
}

func (m *User) GetKeywords() string {
	if m != nil && m.Keywords != nil {
		return *m.Keywords
	}
	return ""
}

func (m *User) GetGeo() *Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *User) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type BidRequest struct {
	RequestId        *string   `protobuf:"bytes,1,req,name=request_id" json:"request_id,omitempty"`
	ApiVersion       *string   `protobuf:"bytes,2,req,name=api_version" json:"api_version,omitempty"`
	Adslots          []*AdSlot `protobuf:"bytes,3,rep,name=adslots" json:"adslots,omitempty"`
	App              *App      `protobuf:"bytes,4,req,name=app" json:"app,omitempty"`
	Device           *Device   `protobuf:"bytes,5,req,name=device" json:"device,omitempty"`
	User             *User     `protobuf:"bytes,6,req,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BidRequest) Reset()         { *m = BidRequest{} }
func (m *BidRequest) String() string { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()    {}

func (m *BidRequest) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *BidRequest) GetApiVersion() string {
	if m != nil && m.ApiVersion != nil {
		return *m.ApiVersion
	}
	return ""
}

func (m *BidRequest) GetAdslots() []*AdSlot {
	if m != nil {
		return m.Adslots
	}
	return nil
}

func (m *BidRequest) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *BidRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *BidRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type MaterialMeta struct {
	AdType           *AdType                    `protobuf:"varint,1,req,name=ad_type,enum=toutiao_ssp.api.AdType" json:"ad_type,omitempty"`
	Nurl             *string                    `protobuf:"bytes,2,req,name=nurl" json:"nurl,omitempty"`
	Title            *string                    `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Source           *string                    `protobuf:"bytes,4,req,name=source" json:"source,omitempty"`
	ImageBanner      *MaterialMeta_ImageMeta    `protobuf:"bytes,5,opt,name=image_banner" json:"image_banner,omitempty"`
	External         *MaterialMeta_ExternalMeta `protobuf:"bytes,6,opt,name=external" json:"external,omitempty"`
	AndroidApp       *MaterialMeta_AndroidApp   `protobuf:"bytes,7,opt,name=android_app" json:"android_app,omitempty"`
	IosApp           *MaterialMeta_IosApp       `protobuf:"bytes,8,opt,name=ios_app" json:"ios_app,omitempty"`
	ShowUrl          []string                   `protobuf:"bytes,9,rep,name=show_url" json:"show_url,omitempty"`
	ClickUrl         []string                   `protobuf:"bytes,10,rep,name=click_url" json:"click_url,omitempty"`
	IsInapp          *bool                      `protobuf:"varint,11,opt,name=is_inapp" json:"is_inapp,omitempty"`
	Ext              *string                    `protobuf:"bytes,12,opt,name=ext" json:"ext,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *MaterialMeta) Reset()         { *m = MaterialMeta{} }
func (m *MaterialMeta) String() string { return proto.CompactTextString(m) }
func (*MaterialMeta) ProtoMessage()    {}

func (m *MaterialMeta) GetAdType() AdType {
	if m != nil && m.AdType != nil {
		return *m.AdType
	}
	return AdType_TOUTIAO_FEED_LP_LARGE
}

func (m *MaterialMeta) GetNurl() string {
	if m != nil && m.Nurl != nil {
		return *m.Nurl
	}
	return ""
}

func (m *MaterialMeta) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *MaterialMeta) GetSource() string {
	if m != nil && m.Source != nil {
		return *m.Source
	}
	return ""
}

func (m *MaterialMeta) GetImageBanner() *MaterialMeta_ImageMeta {
	if m != nil {
		return m.ImageBanner
	}
	return nil
}

func (m *MaterialMeta) GetExternal() *MaterialMeta_ExternalMeta {
	if m != nil {
		return m.External
	}
	return nil
}

func (m *MaterialMeta) GetAndroidApp() *MaterialMeta_AndroidApp {
	if m != nil {
		return m.AndroidApp
	}
	return nil
}

func (m *MaterialMeta) GetIosApp() *MaterialMeta_IosApp {
	if m != nil {
		return m.IosApp
	}
	return nil
}

func (m *MaterialMeta) GetShowUrl() []string {
	if m != nil {
		return m.ShowUrl
	}
	return nil
}

func (m *MaterialMeta) GetClickUrl() []string {
	if m != nil {
		return m.ClickUrl
	}
	return nil
}

func (m *MaterialMeta) GetIsInapp() bool {
	if m != nil && m.IsInapp != nil {
		return *m.IsInapp
	}
	return false
}

func (m *MaterialMeta) GetExt() string {
	if m != nil && m.Ext != nil {
		return *m.Ext
	}
	return ""
}

type MaterialMeta_ImageMeta struct {
	Description      *string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Width            *uint32 `protobuf:"varint,2,req,name=width" json:"width,omitempty"`
	Height           *uint32 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Url              *string `protobuf:"bytes,4,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MaterialMeta_ImageMeta) Reset()         { *m = MaterialMeta_ImageMeta{} }
func (m *MaterialMeta_ImageMeta) String() string { return proto.CompactTextString(m) }
func (*MaterialMeta_ImageMeta) ProtoMessage()    {}

func (m *MaterialMeta_ImageMeta) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *MaterialMeta_ImageMeta) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *MaterialMeta_ImageMeta) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *MaterialMeta_ImageMeta) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type MaterialMeta_ExternalMeta struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MaterialMeta_ExternalMeta) Reset()         { *m = MaterialMeta_ExternalMeta{} }
func (m *MaterialMeta_ExternalMeta) String() string { return proto.CompactTextString(m) }
func (*MaterialMeta_ExternalMeta) ProtoMessage()    {}

func (m *MaterialMeta_ExternalMeta) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type MaterialMeta_AndroidApp struct {
	AppName          *string `protobuf:"bytes,1,req,name=app_name" json:"app_name,omitempty"`
	DownloadUrl      *string `protobuf:"bytes,2,req,name=download_url" json:"download_url,omitempty"`
	OpenUrl          *string `protobuf:"bytes,3,opt,name=open_url" json:"open_url,omitempty"`
	Package          *string `protobuf:"bytes,4,opt,name=package" json:"package,omitempty"`
	WebUrl           *string `protobuf:"bytes,5,opt,name=web_url" json:"web_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MaterialMeta_AndroidApp) Reset()         { *m = MaterialMeta_AndroidApp{} }
func (m *MaterialMeta_AndroidApp) String() string { return proto.CompactTextString(m) }
func (*MaterialMeta_AndroidApp) ProtoMessage()    {}

func (m *MaterialMeta_AndroidApp) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *MaterialMeta_AndroidApp) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *MaterialMeta_AndroidApp) GetOpenUrl() string {
	if m != nil && m.OpenUrl != nil {
		return *m.OpenUrl
	}
	return ""
}

func (m *MaterialMeta_AndroidApp) GetPackage() string {
	if m != nil && m.Package != nil {
		return *m.Package
	}
	return ""
}

func (m *MaterialMeta_AndroidApp) GetWebUrl() string {
	if m != nil && m.WebUrl != nil {
		return *m.WebUrl
	}
	return ""
}

type MaterialMeta_IosApp struct {
	AppName          *string `protobuf:"bytes,1,req,name=app_name" json:"app_name,omitempty"`
	DownloadUrl      *string `protobuf:"bytes,2,req,name=download_url" json:"download_url,omitempty"`
	OpenUrl          *string `protobuf:"bytes,3,opt,name=open_url" json:"open_url,omitempty"`
	Appleid          *string `protobuf:"bytes,4,opt,name=appleid" json:"appleid,omitempty"`
	IpaUrl           *string `protobuf:"bytes,5,opt,name=ipa_url" json:"ipa_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MaterialMeta_IosApp) Reset()         { *m = MaterialMeta_IosApp{} }
func (m *MaterialMeta_IosApp) String() string { return proto.CompactTextString(m) }
func (*MaterialMeta_IosApp) ProtoMessage()    {}

func (m *MaterialMeta_IosApp) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *MaterialMeta_IosApp) GetDownloadUrl() string {
	if m != nil && m.DownloadUrl != nil {
		return *m.DownloadUrl
	}
	return ""
}

func (m *MaterialMeta_IosApp) GetOpenUrl() string {
	if m != nil && m.OpenUrl != nil {
		return *m.OpenUrl
	}
	return ""
}

func (m *MaterialMeta_IosApp) GetAppleid() string {
	if m != nil && m.Appleid != nil {
		return *m.Appleid
	}
	return ""
}

func (m *MaterialMeta_IosApp) GetIpaUrl() string {
	if m != nil && m.IpaUrl != nil {
		return *m.IpaUrl
	}
	return ""
}

type Bid struct {
	Id               *string       `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	AdslotId         *string       `protobuf:"bytes,2,req,name=adslot_id" json:"adslot_id,omitempty"`
	Price            *uint32       `protobuf:"varint,3,req,name=price" json:"price,omitempty"`
	Adid             *uint64       `protobuf:"varint,4,req,name=adid" json:"adid,omitempty"`
	Creative         *MaterialMeta `protobuf:"bytes,5,req,name=creative" json:"creative,omitempty"`
	Dealid           *string       `protobuf:"bytes,6,opt,name=dealid" json:"dealid,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}

func (m *Bid) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Bid) GetAdslotId() string {
	if m != nil && m.AdslotId != nil {
		return *m.AdslotId
	}
	return ""
}

func (m *Bid) GetPrice() uint32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Bid) GetAdid() uint64 {
	if m != nil && m.Adid != nil {
		return *m.Adid
	}
	return 0
}

func (m *Bid) GetCreative() *MaterialMeta {
	if m != nil {
		return m.Creative
	}
	return nil
}

func (m *Bid) GetDealid() string {
	if m != nil && m.Dealid != nil {
		return *m.Dealid
	}
	return ""
}

type SeatBid struct {
	Ads              []*Bid  `protobuf:"bytes,1,rep,name=ads" json:"ads,omitempty"`
	Seat             *string `protobuf:"bytes,2,opt,name=seat" json:"seat,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SeatBid) Reset()         { *m = SeatBid{} }
func (m *SeatBid) String() string { return proto.CompactTextString(m) }
func (*SeatBid) ProtoMessage()    {}

func (m *SeatBid) GetAds() []*Bid {
	if m != nil {
		return m.Ads
	}
	return nil
}

func (m *SeatBid) GetSeat() string {
	if m != nil && m.Seat != nil {
		return *m.Seat
	}
	return ""
}

type BidResponse struct {
	RequestId        *string    `protobuf:"bytes,1,req,name=request_id" json:"request_id,omitempty"`
	Seatbids         []*SeatBid `protobuf:"bytes,2,rep,name=seatbids" json:"seatbids,omitempty"`
	ErrorCode        *uint64    `protobuf:"varint,3,opt,name=error_code" json:"error_code,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *BidResponse) Reset()         { *m = BidResponse{} }
func (m *BidResponse) String() string { return proto.CompactTextString(m) }
func (*BidResponse) ProtoMessage()    {}

func (m *BidResponse) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *BidResponse) GetSeatbids() []*SeatBid {
	if m != nil {
		return m.Seatbids
	}
	return nil
}

func (m *BidResponse) GetErrorCode() uint64 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterEnum("toutiao_ssp.api.AdType", AdType_name, AdType_value)
	proto.RegisterEnum("toutiao_ssp.api.AdSlot_Position", AdSlot_Position_name, AdSlot_Position_value)
	proto.RegisterEnum("toutiao_ssp.api.Device_ConnectionType", Device_ConnectionType_name, Device_ConnectionType_value)
	proto.RegisterEnum("toutiao_ssp.api.Device_DeviceType", Device_DeviceType_name, Device_DeviceType_value)
	proto.RegisterEnum("toutiao_ssp.api.User_Gender", User_Gender_name, User_Gender_value)
}
