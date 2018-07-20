// Code generated by protoc-gen-go.
// source: adx.proto
// DO NOT EDIT!

/*
Package adx is a generated protocol buffer package.

It is generated from these files:
	adx.proto

It has these top-level messages:
	Request
	Response
	Direct_Response
*/
package adx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 审核接口和免审核接口公用 request
type Request struct {
	Id                 *string                       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	At                 *uint32                       `protobuf:"varint,2,opt,name=at" json:"at,omitempty"`
	Site               *Request_Site                 `protobuf:"bytes,3,opt,name=site" json:"site,omitempty"`
	Device             *Request_Device               `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	User               *Request_User                 `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Impression         []*Request_Impression         `protobuf:"bytes,6,rep,name=impression" json:"impression,omitempty"`
	InnerInfo          *string                       `protobuf:"bytes,7,opt,name=inner_info" json:"inner_info,omitempty"`
	App                *Request_App                  `protobuf:"bytes,8,opt,name=app" json:"app,omitempty"`
	UnlikeOrderList    []string                      `protobuf:"bytes,9,rep,name=unlike_order_list" json:"unlike_order_list,omitempty"`
	Oids               []*Request_PlayOrder          `protobuf:"bytes,10,rep,name=oids" json:"oids,omitempty"`
	RequestBlackOrders []string                      `protobuf:"bytes,11,rep,name=request_black_orders" json:"request_black_orders,omitempty"`
	ChannelBlackOrders []*Request_ChannelBlackOrders `protobuf:"bytes,12,rep,name=channel_black_orders" json:"channel_black_orders,omitempty"`
	IsOrderCache       *bool                         `protobuf:"varint,13,opt,name=is_order_cache" json:"is_order_cache,omitempty"`
	IsFirstTime        *bool                         `protobuf:"varint,14,opt,name=is_first_time" json:"is_first_time,omitempty"`
	XXX_unrecognized   []byte                        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Request) GetAt() uint32 {
	if m != nil && m.At != nil {
		return *m.At
	}
	return 0
}

func (m *Request) GetSite() *Request_Site {
	if m != nil {
		return m.Site
	}
	return nil
}

func (m *Request) GetDevice() *Request_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *Request) GetUser() *Request_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Request) GetImpression() []*Request_Impression {
	if m != nil {
		return m.Impression
	}
	return nil
}

func (m *Request) GetInnerInfo() string {
	if m != nil && m.InnerInfo != nil {
		return *m.InnerInfo
	}
	return ""
}

func (m *Request) GetApp() *Request_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Request) GetUnlikeOrderList() []string {
	if m != nil {
		return m.UnlikeOrderList
	}
	return nil
}

func (m *Request) GetOids() []*Request_PlayOrder {
	if m != nil {
		return m.Oids
	}
	return nil
}

func (m *Request) GetRequestBlackOrders() []string {
	if m != nil {
		return m.RequestBlackOrders
	}
	return nil
}

func (m *Request) GetChannelBlackOrders() []*Request_ChannelBlackOrders {
	if m != nil {
		return m.ChannelBlackOrders
	}
	return nil
}

func (m *Request) GetIsOrderCache() bool {
	if m != nil && m.IsOrderCache != nil {
		return *m.IsOrderCache
	}
	return false
}

func (m *Request) GetIsFirstTime() bool {
	if m != nil && m.IsFirstTime != nil {
		return *m.IsFirstTime
	}
	return false
}

type Request_Impression struct {
	Id               *string                              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Tagid            *string                              `protobuf:"bytes,2,opt,name=tagid" json:"tagid,omitempty"`
	Bidfloor         *float32                             `protobuf:"fixed32,3,opt,name=bidfloor" json:"bidfloor,omitempty"`
	Banner           *Request_Impression_Banner           `protobuf:"bytes,4,opt,name=banner" json:"banner,omitempty"`
	Video            *Request_Impression_Video            `protobuf:"bytes,5,opt,name=video" json:"video,omitempty"`
	Clientid         *string                              `protobuf:"bytes,6,opt,name=clientid" json:"clientid,omitempty"`
	Tradecode        *string                              `protobuf:"bytes,7,opt,name=tradecode" json:"tradecode,omitempty"`
	SnsLists         *string                              `protobuf:"bytes,8,opt,name=sns_lists" json:"sns_lists,omitempty"`
	Dealid           *string                              `protobuf:"bytes,9,opt,name=dealid" json:"dealid,omitempty"`
	AdmRequire       []*Request_Impression_MaterialFormat `protobuf:"bytes,10,rep,name=adm_require" json:"adm_require,omitempty"`
	ImpExt           *Request_Impression_ImpExt           `protobuf:"bytes,11,opt,name=imp_ext" json:"imp_ext,omitempty"`
	Channel          *string                              `protobuf:"bytes,12,opt,name=channel" json:"channel,omitempty"`
	Date             *string                              `protobuf:"bytes,13,opt,name=date" json:"date,omitempty"`
	TagSeq           *uint32                              `protobuf:"varint,14,opt,name=tag_seq" json:"tag_seq,omitempty"`
	DisplayType      []*Request_Impression_DisplayType    `protobuf:"bytes,15,rep,name=display_type" json:"display_type,omitempty"`
	Dealids          []string                             `protobuf:"bytes,16,rep,name=dealids" json:"dealids,omitempty"`
	MaxOrderCount    *uint32                              `protobuf:"varint,17,opt,name=max_order_count" json:"max_order_count,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *Request_Impression) Reset()         { *m = Request_Impression{} }
func (m *Request_Impression) String() string { return proto.CompactTextString(m) }
func (*Request_Impression) ProtoMessage()    {}

func (m *Request_Impression) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Request_Impression) GetTagid() string {
	if m != nil && m.Tagid != nil {
		return *m.Tagid
	}
	return ""
}

func (m *Request_Impression) GetBidfloor() float32 {
	if m != nil && m.Bidfloor != nil {
		return *m.Bidfloor
	}
	return 0
}

func (m *Request_Impression) GetBanner() *Request_Impression_Banner {
	if m != nil {
		return m.Banner
	}
	return nil
}

func (m *Request_Impression) GetVideo() *Request_Impression_Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Request_Impression) GetClientid() string {
	if m != nil && m.Clientid != nil {
		return *m.Clientid
	}
	return ""
}

func (m *Request_Impression) GetTradecode() string {
	if m != nil && m.Tradecode != nil {
		return *m.Tradecode
	}
	return ""
}

func (m *Request_Impression) GetSnsLists() string {
	if m != nil && m.SnsLists != nil {
		return *m.SnsLists
	}
	return ""
}

func (m *Request_Impression) GetDealid() string {
	if m != nil && m.Dealid != nil {
		return *m.Dealid
	}
	return ""
}

func (m *Request_Impression) GetAdmRequire() []*Request_Impression_MaterialFormat {
	if m != nil {
		return m.AdmRequire
	}
	return nil
}

func (m *Request_Impression) GetImpExt() *Request_Impression_ImpExt {
	if m != nil {
		return m.ImpExt
	}
	return nil
}

func (m *Request_Impression) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *Request_Impression) GetDate() string {
	if m != nil && m.Date != nil {
		return *m.Date
	}
	return ""
}

func (m *Request_Impression) GetTagSeq() uint32 {
	if m != nil && m.TagSeq != nil {
		return *m.TagSeq
	}
	return 0
}

func (m *Request_Impression) GetDisplayType() []*Request_Impression_DisplayType {
	if m != nil {
		return m.DisplayType
	}
	return nil
}

func (m *Request_Impression) GetDealids() []string {
	if m != nil {
		return m.Dealids
	}
	return nil
}

func (m *Request_Impression) GetMaxOrderCount() uint32 {
	if m != nil && m.MaxOrderCount != nil {
		return *m.MaxOrderCount
	}
	return 0
}

type Request_Impression_ImpExt struct {
	Chid             *uint32  `protobuf:"varint,1,opt,name=chid" json:"chid,omitempty"`
	Stype            []uint32 `protobuf:"varint,2,rep,name=stype" json:"stype,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_Impression_ImpExt) Reset()         { *m = Request_Impression_ImpExt{} }
func (m *Request_Impression_ImpExt) String() string { return proto.CompactTextString(m) }
func (*Request_Impression_ImpExt) ProtoMessage()    {}

func (m *Request_Impression_ImpExt) GetChid() uint32 {
	if m != nil && m.Chid != nil {
		return *m.Chid
	}
	return 0
}

func (m *Request_Impression_ImpExt) GetStype() []uint32 {
	if m != nil {
		return m.Stype
	}
	return nil
}

type Request_Impression_Banner struct {
	Width            *uint32  `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height           *uint32  `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Mimes            []string `protobuf:"bytes,3,rep,name=mimes" json:"mimes,omitempty"`
	ExtraStyle       *string  `protobuf:"bytes,4,opt,name=extra_style" json:"extra_style,omitempty"`
	Visibility       *uint32  `protobuf:"varint,5,opt,name=visibility" json:"visibility,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_Impression_Banner) Reset()         { *m = Request_Impression_Banner{} }
func (m *Request_Impression_Banner) String() string { return proto.CompactTextString(m) }
func (*Request_Impression_Banner) ProtoMessage()    {}

func (m *Request_Impression_Banner) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Request_Impression_Banner) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *Request_Impression_Banner) GetMimes() []string {
	if m != nil {
		return m.Mimes
	}
	return nil
}

func (m *Request_Impression_Banner) GetExtraStyle() string {
	if m != nil && m.ExtraStyle != nil {
		return *m.ExtraStyle
	}
	return ""
}

func (m *Request_Impression_Banner) GetVisibility() uint32 {
	if m != nil && m.Visibility != nil {
		return *m.Visibility
	}
	return 0
}

type Request_Impression_Video struct {
	Mimes            []string `protobuf:"bytes,1,rep,name=mimes" json:"mimes,omitempty"`
	Linearity        *uint32  `protobuf:"varint,2,opt,name=linearity" json:"linearity,omitempty"`
	Minduration      *uint32  `protobuf:"varint,3,opt,name=minduration" json:"minduration,omitempty"`
	Maxduration      *uint32  `protobuf:"varint,4,opt,name=maxduration" json:"maxduration,omitempty"`
	Protocol         *uint32  `protobuf:"varint,5,opt,name=protocol" json:"protocol,omitempty"`
	Width            *uint32  `protobuf:"varint,6,opt,name=width" json:"width,omitempty"`
	Height           *uint32  `protobuf:"varint,7,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_Impression_Video) Reset()         { *m = Request_Impression_Video{} }
func (m *Request_Impression_Video) String() string { return proto.CompactTextString(m) }
func (*Request_Impression_Video) ProtoMessage()    {}

func (m *Request_Impression_Video) GetMimes() []string {
	if m != nil {
		return m.Mimes
	}
	return nil
}

func (m *Request_Impression_Video) GetLinearity() uint32 {
	if m != nil && m.Linearity != nil {
		return *m.Linearity
	}
	return 0
}

func (m *Request_Impression_Video) GetMinduration() uint32 {
	if m != nil && m.Minduration != nil {
		return *m.Minduration
	}
	return 0
}

func (m *Request_Impression_Video) GetMaxduration() uint32 {
	if m != nil && m.Maxduration != nil {
		return *m.Maxduration
	}
	return 0
}

func (m *Request_Impression_Video) GetProtocol() uint32 {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return 0
}

func (m *Request_Impression_Video) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Request_Impression_Video) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type Request_Impression_MaterialFormat struct {
	Width            *uint32 `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height           *uint32 `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	Mimes            *string `protobuf:"bytes,3,opt,name=mimes" json:"mimes,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_Impression_MaterialFormat) Reset()         { *m = Request_Impression_MaterialFormat{} }
func (m *Request_Impression_MaterialFormat) String() string { return proto.CompactTextString(m) }
func (*Request_Impression_MaterialFormat) ProtoMessage()    {}

func (m *Request_Impression_MaterialFormat) GetWidth() uint32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Request_Impression_MaterialFormat) GetHeight() uint32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *Request_Impression_MaterialFormat) GetMimes() string {
	if m != nil && m.Mimes != nil {
		return *m.Mimes
	}
	return ""
}

type Request_Impression_DisplayType struct {
	DisplayType      *uint32                              `protobuf:"varint,1,opt,name=display_type" json:"display_type,omitempty"`
	ClickType        []uint32                             `protobuf:"varint,2,rep,name=click_type" json:"click_type,omitempty"`
	AdmRequire       []*Request_Impression_MaterialFormat `protobuf:"bytes,3,rep,name=adm_require" json:"adm_require,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *Request_Impression_DisplayType) Reset()         { *m = Request_Impression_DisplayType{} }
func (m *Request_Impression_DisplayType) String() string { return proto.CompactTextString(m) }
func (*Request_Impression_DisplayType) ProtoMessage()    {}

func (m *Request_Impression_DisplayType) GetDisplayType() uint32 {
	if m != nil && m.DisplayType != nil {
		return *m.DisplayType
	}
	return 0
}

func (m *Request_Impression_DisplayType) GetClickType() []uint32 {
	if m != nil {
		return m.ClickType
	}
	return nil
}

func (m *Request_Impression_DisplayType) GetAdmRequire() []*Request_Impression_MaterialFormat {
	if m != nil {
		return m.AdmRequire
	}
	return nil
}

type Request_Site struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Page             *string `protobuf:"bytes,2,opt,name=page" json:"page,omitempty"`
	Ref              *string `protobuf:"bytes,3,opt,name=ref" json:"ref,omitempty"`
	Channel          *string `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_Site) Reset()         { *m = Request_Site{} }
func (m *Request_Site) String() string { return proto.CompactTextString(m) }
func (*Request_Site) ProtoMessage()    {}

func (m *Request_Site) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Request_Site) GetPage() string {
	if m != nil && m.Page != nil {
		return *m.Page
	}
	return ""
}

func (m *Request_Site) GetRef() string {
	if m != nil && m.Ref != nil {
		return *m.Ref
	}
	return ""
}

func (m *Request_Site) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

type Request_Device struct {
	Ip               *string             `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Ua               *string             `protobuf:"bytes,2,opt,name=ua" json:"ua,omitempty"`
	Geo              *Request_Device_Geo `protobuf:"bytes,3,opt,name=geo" json:"geo,omitempty"`
	Idfa             *string             `protobuf:"bytes,4,opt,name=idfa" json:"idfa,omitempty"`
	IdfaEnc          *uint32             `protobuf:"varint,5,opt,name=idfa_enc" json:"idfa_enc,omitempty"`
	Openudid         *string             `protobuf:"bytes,6,opt,name=openudid" json:"openudid,omitempty"`
	Carrier          *uint32             `protobuf:"varint,7,opt,name=carrier" json:"carrier,omitempty"`
	Make             *string             `protobuf:"bytes,8,opt,name=make" json:"make,omitempty"`
	Model            *string             `protobuf:"bytes,9,opt,name=model" json:"model,omitempty"`
	Os               *string             `protobuf:"bytes,10,opt,name=os" json:"os,omitempty"`
	Osv              *string             `protobuf:"bytes,11,opt,name=Osv" json:"Osv,omitempty"`
	Js               *uint32             `protobuf:"varint,12,opt,name=Js" json:"Js,omitempty"`
	Connectiontype   *uint32             `protobuf:"varint,13,opt,name=connectiontype" json:"connectiontype,omitempty"`
	Devicetype       *uint32             `protobuf:"varint,14,opt,name=devicetype" json:"devicetype,omitempty"`
	Mac              *string             `protobuf:"bytes,15,opt,name=mac" json:"mac,omitempty"`
	Imei             *string             `protobuf:"bytes,16,opt,name=imei" json:"imei,omitempty"`
	Androidid        *string             `protobuf:"bytes,17,opt,name=androidid" json:"androidid,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Request_Device) Reset()         { *m = Request_Device{} }
func (m *Request_Device) String() string { return proto.CompactTextString(m) }
func (*Request_Device) ProtoMessage()    {}

func (m *Request_Device) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Request_Device) GetUa() string {
	if m != nil && m.Ua != nil {
		return *m.Ua
	}
	return ""
}

func (m *Request_Device) GetGeo() *Request_Device_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *Request_Device) GetIdfa() string {
	if m != nil && m.Idfa != nil {
		return *m.Idfa
	}
	return ""
}

func (m *Request_Device) GetIdfaEnc() uint32 {
	if m != nil && m.IdfaEnc != nil {
		return *m.IdfaEnc
	}
	return 0
}

func (m *Request_Device) GetOpenudid() string {
	if m != nil && m.Openudid != nil {
		return *m.Openudid
	}
	return ""
}

func (m *Request_Device) GetCarrier() uint32 {
	if m != nil && m.Carrier != nil {
		return *m.Carrier
	}
	return 0
}

func (m *Request_Device) GetMake() string {
	if m != nil && m.Make != nil {
		return *m.Make
	}
	return ""
}

func (m *Request_Device) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Request_Device) GetOs() string {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return ""
}

func (m *Request_Device) GetOsv() string {
	if m != nil && m.Osv != nil {
		return *m.Osv
	}
	return ""
}

func (m *Request_Device) GetJs() uint32 {
	if m != nil && m.Js != nil {
		return *m.Js
	}
	return 0
}

func (m *Request_Device) GetConnectiontype() uint32 {
	if m != nil && m.Connectiontype != nil {
		return *m.Connectiontype
	}
	return 0
}

func (m *Request_Device) GetDevicetype() uint32 {
	if m != nil && m.Devicetype != nil {
		return *m.Devicetype
	}
	return 0
}

func (m *Request_Device) GetMac() string {
	if m != nil && m.Mac != nil {
		return *m.Mac
	}
	return ""
}

func (m *Request_Device) GetImei() string {
	if m != nil && m.Imei != nil {
		return *m.Imei
	}
	return ""
}

func (m *Request_Device) GetAndroidid() string {
	if m != nil && m.Androidid != nil {
		return *m.Androidid
	}
	return ""
}

type Request_Device_Geo struct {
	Latitude         *float32 `protobuf:"fixed32,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude        *float32 `protobuf:"fixed32,2,opt,name=longitude" json:"longitude,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_Device_Geo) Reset()         { *m = Request_Device_Geo{} }
func (m *Request_Device_Geo) String() string { return proto.CompactTextString(m) }
func (*Request_Device_Geo) ProtoMessage()    {}

func (m *Request_Device_Geo) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Request_Device_Geo) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

type Request_User struct {
	Id               *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Buyerid          *string  `protobuf:"bytes,2,opt,name=buyerid" json:"buyerid,omitempty"`
	Gender           *string  `protobuf:"bytes,3,opt,name=gender" json:"gender,omitempty"`
	Age              *uint32  `protobuf:"varint,4,opt,name=age" json:"age,omitempty"`
	TagId            []string `protobuf:"bytes,5,rep,name=tag_id" json:"tag_id,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_User) Reset()         { *m = Request_User{} }
func (m *Request_User) String() string { return proto.CompactTextString(m) }
func (*Request_User) ProtoMessage()    {}

func (m *Request_User) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Request_User) GetBuyerid() string {
	if m != nil && m.Buyerid != nil {
		return *m.Buyerid
	}
	return ""
}

func (m *Request_User) GetGender() string {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return ""
}

func (m *Request_User) GetAge() uint32 {
	if m != nil && m.Age != nil {
		return *m.Age
	}
	return 0
}

func (m *Request_User) GetTagId() []string {
	if m != nil {
		return m.TagId
	}
	return nil
}

type Request_App struct {
	Id               *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Domain           *string  `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	Cat              []string `protobuf:"bytes,4,rep,name=cat" json:"cat,omitempty"`
	Sectioncat       []string `protobuf:"bytes,5,rep,name=sectioncat" json:"sectioncat,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_App) Reset()         { *m = Request_App{} }
func (m *Request_App) String() string { return proto.CompactTextString(m) }
func (*Request_App) ProtoMessage()    {}

func (m *Request_App) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Request_App) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Request_App) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *Request_App) GetCat() []string {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *Request_App) GetSectioncat() []string {
	if m != nil {
		return m.Sectioncat
	}
	return nil
}

type Request_PlayOrder struct {
	Oid              *string `protobuf:"bytes,1,opt,name=oid" json:"oid,omitempty"`
	Index            *uint32 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request_PlayOrder) Reset()         { *m = Request_PlayOrder{} }
func (m *Request_PlayOrder) String() string { return proto.CompactTextString(m) }
func (*Request_PlayOrder) ProtoMessage()    {}

func (m *Request_PlayOrder) GetOid() string {
	if m != nil && m.Oid != nil {
		return *m.Oid
	}
	return ""
}

func (m *Request_PlayOrder) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type Request_ChannelBlackOrders struct {
	ChannelId        *uint32  `protobuf:"varint,1,opt,name=channel_id" json:"channel_id,omitempty"`
	BlackOrders      []string `protobuf:"bytes,2,rep,name=black_orders" json:"black_orders,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Request_ChannelBlackOrders) Reset()         { *m = Request_ChannelBlackOrders{} }
func (m *Request_ChannelBlackOrders) String() string { return proto.CompactTextString(m) }
func (*Request_ChannelBlackOrders) ProtoMessage()    {}

func (m *Request_ChannelBlackOrders) GetChannelId() uint32 {
	if m != nil && m.ChannelId != nil {
		return *m.ChannelId
	}
	return 0
}

func (m *Request_ChannelBlackOrders) GetBlackOrders() []string {
	if m != nil {
		return m.BlackOrders
	}
	return nil
}

// 审核接口
type Response struct {
	Id               *string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Bidid            *string                  `protobuf:"bytes,2,opt,name=bidid" json:"bidid,omitempty"`
	Seatbid          []*Response_SeatBid      `protobuf:"bytes,3,rep,name=seatbid" json:"seatbid,omitempty"`
	Ext              *Response_BidResponseExt `protobuf:"bytes,4,opt,name=ext" json:"ext,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Response) GetBidid() string {
	if m != nil && m.Bidid != nil {
		return *m.Bidid
	}
	return ""
}

func (m *Response) GetSeatbid() []*Response_SeatBid {
	if m != nil {
		return m.Seatbid
	}
	return nil
}

func (m *Response) GetExt() *Response_BidResponseExt {
	if m != nil {
		return m.Ext
	}
	return nil
}

type Response_BidResponseExt struct {
	UpVersion        *string `protobuf:"bytes,1,opt,name=up_version" json:"up_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response_BidResponseExt) Reset()         { *m = Response_BidResponseExt{} }
func (m *Response_BidResponseExt) String() string { return proto.CompactTextString(m) }
func (*Response_BidResponseExt) ProtoMessage()    {}

func (m *Response_BidResponseExt) GetUpVersion() string {
	if m != nil && m.UpVersion != nil {
		return *m.UpVersion
	}
	return ""
}

type Response_Bid struct {
	Id               *string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Impid            *string  `protobuf:"bytes,2,opt,name=impid" json:"impid,omitempty"`
	Price            *float32 `protobuf:"fixed32,3,opt,name=price" json:"price,omitempty"`
	Adid             *string  `protobuf:"bytes,4,opt,name=adid" json:"adid,omitempty"`
	Nurl             *string  `protobuf:"bytes,5,opt,name=nurl" json:"nurl,omitempty"`
	Adm              *string  `protobuf:"bytes,6,opt,name=adm" json:"adm,omitempty"`
	Ext              *string  `protobuf:"bytes,7,opt,name=ext" json:"ext,omitempty"`
	AdmPara          *string  `protobuf:"bytes,8,opt,name=adm_para" json:"adm_para,omitempty"`
	InnerResp        *string  `protobuf:"bytes,9,opt,name=inner_resp" json:"inner_resp,omitempty"`
	Ext2             *string  `protobuf:"bytes,10,opt,name=ext2" json:"ext2,omitempty"`
	DispExts         []string `protobuf:"bytes,11,rep,name=disp_exts" json:"disp_exts,omitempty"`
	ClickExts        []string `protobuf:"bytes,12,rep,name=click_exts" json:"click_exts,omitempty"`
	Ext3             *string  `protobuf:"bytes,13,opt,name=ext3" json:"ext3,omitempty"`
	DisplayType      *uint32  `protobuf:"varint,14,opt,name=display_type" json:"display_type,omitempty"`
	Stype            *uint32  `protobuf:"varint,15,opt,name=stype" json:"stype,omitempty"`
	Dealid           *string  `protobuf:"bytes,16,opt,name=dealid" json:"dealid,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Response_Bid) Reset()         { *m = Response_Bid{} }
func (m *Response_Bid) String() string { return proto.CompactTextString(m) }
func (*Response_Bid) ProtoMessage()    {}

func (m *Response_Bid) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Response_Bid) GetImpid() string {
	if m != nil && m.Impid != nil {
		return *m.Impid
	}
	return ""
}

func (m *Response_Bid) GetPrice() float32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Response_Bid) GetAdid() string {
	if m != nil && m.Adid != nil {
		return *m.Adid
	}
	return ""
}

func (m *Response_Bid) GetNurl() string {
	if m != nil && m.Nurl != nil {
		return *m.Nurl
	}
	return ""
}

func (m *Response_Bid) GetAdm() string {
	if m != nil && m.Adm != nil {
		return *m.Adm
	}
	return ""
}

func (m *Response_Bid) GetExt() string {
	if m != nil && m.Ext != nil {
		return *m.Ext
	}
	return ""
}

func (m *Response_Bid) GetAdmPara() string {
	if m != nil && m.AdmPara != nil {
		return *m.AdmPara
	}
	return ""
}

func (m *Response_Bid) GetInnerResp() string {
	if m != nil && m.InnerResp != nil {
		return *m.InnerResp
	}
	return ""
}

func (m *Response_Bid) GetExt2() string {
	if m != nil && m.Ext2 != nil {
		return *m.Ext2
	}
	return ""
}

func (m *Response_Bid) GetDispExts() []string {
	if m != nil {
		return m.DispExts
	}
	return nil
}

func (m *Response_Bid) GetClickExts() []string {
	if m != nil {
		return m.ClickExts
	}
	return nil
}

func (m *Response_Bid) GetExt3() string {
	if m != nil && m.Ext3 != nil {
		return *m.Ext3
	}
	return ""
}

func (m *Response_Bid) GetDisplayType() uint32 {
	if m != nil && m.DisplayType != nil {
		return *m.DisplayType
	}
	return 0
}

func (m *Response_Bid) GetStype() uint32 {
	if m != nil && m.Stype != nil {
		return *m.Stype
	}
	return 0
}

func (m *Response_Bid) GetDealid() string {
	if m != nil && m.Dealid != nil {
		return *m.Dealid
	}
	return ""
}

type Response_SeatBid struct {
	Bid              []*Response_Bid `protobuf:"bytes,1,rep,name=bid" json:"bid,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Response_SeatBid) Reset()         { *m = Response_SeatBid{} }
func (m *Response_SeatBid) String() string { return proto.CompactTextString(m) }
func (*Response_SeatBid) ProtoMessage()    {}

func (m *Response_SeatBid) GetBid() []*Response_Bid {
	if m != nil {
		return m.Bid
	}
	return nil
}

// 免审核接口
type Direct_Response struct {
	Id               *string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Bidid            *string                    `protobuf:"bytes,2,opt,name=bidid" json:"bidid,omitempty"`
	Seatbid          []*Direct_Response_SeatBid `protobuf:"bytes,3,rep,name=seatbid" json:"seatbid,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Direct_Response) Reset()         { *m = Direct_Response{} }
func (m *Direct_Response) String() string { return proto.CompactTextString(m) }
func (*Direct_Response) ProtoMessage()    {}

func (m *Direct_Response) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Direct_Response) GetBidid() string {
	if m != nil && m.Bidid != nil {
		return *m.Bidid
	}
	return ""
}

func (m *Direct_Response) GetSeatbid() []*Direct_Response_SeatBid {
	if m != nil {
		return m.Seatbid
	}
	return nil
}

type Direct_Response_MaterialInfo struct {
	Content          *string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Format           *string `protobuf:"bytes,2,opt,name=format" json:"format,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Direct_Response_MaterialInfo) Reset()         { *m = Direct_Response_MaterialInfo{} }
func (m *Direct_Response_MaterialInfo) String() string { return proto.CompactTextString(m) }
func (*Direct_Response_MaterialInfo) ProtoMessage()    {}

func (m *Direct_Response_MaterialInfo) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *Direct_Response_MaterialInfo) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

type Direct_Response_AppInfo struct {
	PkgName  *string `protobuf:"bytes,1,opt,name=pkg_name" json:"pkg_name,omitempty"`
	DeepLink *string `protobuf:"bytes,2,opt,name=deep_link" json:"deep_link,omitempty"`
	// 通用结构如scheme://host/path?query [application]+[openURL]+[options]
	DownLoadUrl      *string `protobuf:"bytes,3,opt,name=down_load_url" json:"down_load_url,omitempty"`
	AppName          *string `protobuf:"bytes,4,opt,name=app_name" json:"app_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Direct_Response_AppInfo) Reset()         { *m = Direct_Response_AppInfo{} }
func (m *Direct_Response_AppInfo) String() string { return proto.CompactTextString(m) }
func (*Direct_Response_AppInfo) ProtoMessage()    {}

func (m *Direct_Response_AppInfo) GetPkgName() string {
	if m != nil && m.PkgName != nil {
		return *m.PkgName
	}
	return ""
}

func (m *Direct_Response_AppInfo) GetDeepLink() string {
	if m != nil && m.DeepLink != nil {
		return *m.DeepLink
	}
	return ""
}

func (m *Direct_Response_AppInfo) GetDownLoadUrl() string {
	if m != nil && m.DownLoadUrl != nil {
		return *m.DownLoadUrl
	}
	return ""
}

func (m *Direct_Response_AppInfo) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

type Direct_Response_Bid struct {
	Id               *string                         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Impid            *string                         `protobuf:"bytes,2,opt,name=impid" json:"impid,omitempty"`
	Materials        []*Direct_Response_MaterialInfo `protobuf:"bytes,3,rep,name=materials" json:"materials,omitempty"`
	Monitorurl       []string                        `protobuf:"bytes,4,rep,name=monitorurl" json:"monitorurl,omitempty"`
	ClickUrl         *string                         `protobuf:"bytes,5,opt,name=click_url" json:"click_url,omitempty"`
	Ext              *string                         `protobuf:"bytes,6,opt,name=ext" json:"ext,omitempty"`
	DispExts         []string                        `protobuf:"bytes,7,rep,name=disp_exts" json:"disp_exts,omitempty"`
	ClickExts        []string                        `protobuf:"bytes,8,rep,name=click_exts" json:"click_exts,omitempty"`
	ClickMonitorUrls []string                        `protobuf:"bytes,9,rep,name=click_monitor_urls" json:"click_monitor_urls,omitempty"`
	Price            *float32                        `protobuf:"fixed32,10,opt,name=price" json:"price,omitempty"`
	AppInfo          *Direct_Response_AppInfo        `protobuf:"bytes,11,opt,name=app_info" json:"app_info,omitempty"`
	DisplayType      *uint32                         `protobuf:"varint,12,opt,name=display_type" json:"display_type,omitempty"`
	CreativeId       *uint32                         `protobuf:"varint,13,opt,name=creative_id" json:"creative_id,omitempty"`
	DisplayId        *uint32                         `protobuf:"varint,14,opt,name=display_id" json:"display_id,omitempty"`
	ClickType        *uint32                         `protobuf:"varint,15,opt,name=click_type" json:"click_type,omitempty"`
	Adid             *string                         `protobuf:"bytes,16,opt,name=adid" json:"adid,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *Direct_Response_Bid) Reset()         { *m = Direct_Response_Bid{} }
func (m *Direct_Response_Bid) String() string { return proto.CompactTextString(m) }
func (*Direct_Response_Bid) ProtoMessage()    {}

func (m *Direct_Response_Bid) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Direct_Response_Bid) GetImpid() string {
	if m != nil && m.Impid != nil {
		return *m.Impid
	}
	return ""
}

func (m *Direct_Response_Bid) GetMaterials() []*Direct_Response_MaterialInfo {
	if m != nil {
		return m.Materials
	}
	return nil
}

func (m *Direct_Response_Bid) GetMonitorurl() []string {
	if m != nil {
		return m.Monitorurl
	}
	return nil
}

func (m *Direct_Response_Bid) GetClickUrl() string {
	if m != nil && m.ClickUrl != nil {
		return *m.ClickUrl
	}
	return ""
}

func (m *Direct_Response_Bid) GetExt() string {
	if m != nil && m.Ext != nil {
		return *m.Ext
	}
	return ""
}

func (m *Direct_Response_Bid) GetDispExts() []string {
	if m != nil {
		return m.DispExts
	}
	return nil
}

func (m *Direct_Response_Bid) GetClickExts() []string {
	if m != nil {
		return m.ClickExts
	}
	return nil
}

func (m *Direct_Response_Bid) GetClickMonitorUrls() []string {
	if m != nil {
		return m.ClickMonitorUrls
	}
	return nil
}

func (m *Direct_Response_Bid) GetPrice() float32 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *Direct_Response_Bid) GetAppInfo() *Direct_Response_AppInfo {
	if m != nil {
		return m.AppInfo
	}
	return nil
}

func (m *Direct_Response_Bid) GetDisplayType() uint32 {
	if m != nil && m.DisplayType != nil {
		return *m.DisplayType
	}
	return 0
}

func (m *Direct_Response_Bid) GetCreativeId() uint32 {
	if m != nil && m.CreativeId != nil {
		return *m.CreativeId
	}
	return 0
}

func (m *Direct_Response_Bid) GetDisplayId() uint32 {
	if m != nil && m.DisplayId != nil {
		return *m.DisplayId
	}
	return 0
}

func (m *Direct_Response_Bid) GetClickType() uint32 {
	if m != nil && m.ClickType != nil {
		return *m.ClickType
	}
	return 0
}

func (m *Direct_Response_Bid) GetAdid() string {
	if m != nil && m.Adid != nil {
		return *m.Adid
	}
	return ""
}

type Direct_Response_SeatBid struct {
	Bid              []*Direct_Response_Bid `protobuf:"bytes,1,rep,name=bid" json:"bid,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Direct_Response_SeatBid) Reset()         { *m = Direct_Response_SeatBid{} }
func (m *Direct_Response_SeatBid) String() string { return proto.CompactTextString(m) }
func (*Direct_Response_SeatBid) ProtoMessage()    {}

func (m *Direct_Response_SeatBid) GetBid() []*Direct_Response_Bid {
	if m != nil {
		return m.Bid
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "adx.Request")
	proto.RegisterType((*Request_Impression)(nil), "adx.Request.Impression")
	proto.RegisterType((*Request_Impression_ImpExt)(nil), "adx.Request.Impression.ImpExt")
	proto.RegisterType((*Request_Impression_Banner)(nil), "adx.Request.Impression.Banner")
	proto.RegisterType((*Request_Impression_Video)(nil), "adx.Request.Impression.Video")
	proto.RegisterType((*Request_Impression_MaterialFormat)(nil), "adx.Request.Impression.MaterialFormat")
	proto.RegisterType((*Request_Impression_DisplayType)(nil), "adx.Request.Impression.DisplayType")
	proto.RegisterType((*Request_Site)(nil), "adx.Request.Site")
	proto.RegisterType((*Request_Device)(nil), "adx.Request.Device")
	proto.RegisterType((*Request_Device_Geo)(nil), "adx.Request.Device.Geo")
	proto.RegisterType((*Request_User)(nil), "adx.Request.User")
	proto.RegisterType((*Request_App)(nil), "adx.Request.App")
	proto.RegisterType((*Request_PlayOrder)(nil), "adx.Request.PlayOrder")
	proto.RegisterType((*Request_ChannelBlackOrders)(nil), "adx.Request.ChannelBlackOrders")
	proto.RegisterType((*Response)(nil), "adx.Response")
	proto.RegisterType((*Response_BidResponseExt)(nil), "adx.Response.BidResponseExt")
	proto.RegisterType((*Response_Bid)(nil), "adx.Response.Bid")
	proto.RegisterType((*Response_SeatBid)(nil), "adx.Response.SeatBid")
	proto.RegisterType((*Direct_Response)(nil), "adx.Direct_Response")
	proto.RegisterType((*Direct_Response_MaterialInfo)(nil), "adx.Direct_Response.MaterialInfo")
	proto.RegisterType((*Direct_Response_AppInfo)(nil), "adx.Direct_Response.AppInfo")
	proto.RegisterType((*Direct_Response_Bid)(nil), "adx.Direct_Response.Bid")
	proto.RegisterType((*Direct_Response_SeatBid)(nil), "adx.Direct_Response.SeatBid")
}
