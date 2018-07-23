package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"

	"httpDproxy/proto/adinall"
	"httpDproxy/proto/adx"
	"httpDproxy/proto/baidu"
	"httpDproxy/proto/ftx"
	"httpDproxy/proto/gdt"
	"httpDproxy/proto/googleadx"
	"httpDproxy/proto/iax"
	"httpDproxy/proto/qax"
	"httpDproxy/proto/sohudsp"
	"httpDproxy/proto/tanx"
	"httpDproxy/proto/vam"
	"httpDproxy/proto/zplay"
	"flag"
)

//https://adxb.optaim.com/adinall
const (
	Host                   = "adxb.optaim.com"
	KcontentType           = "Content-Type"
	KcontentLength         = "Content-Length"
	KmaxRtbVersionHeader   = "Maxrtb-Version"
	ContentTypeOctetStream = "application/octet-stream"
	ContentTypeJson        = "application/json"
	ContentTypeProtobuf    = "application/x-protobuf"

	kLogTypeReq      = "req"
	kLogTypeRsp      = "rsp"      // 智盒返回数据
	kLogTypeThirdRsp = "thirdrsp" // 三方dsp返回数据

	Adx         = "adx"
	Allyes      = "allyes"
	Sohudsp     = "sohudsp"
	Tanx        = "tanx"
	Youku       = "youku"
	Google      = "google"
	Mogo        = "mogo"
	Mega        = "mega"
	Mssp        = "mssp"
	Qihu        = "qihu"
	Mz          = "mz"
	Adview      = "adview"
	Adinall     = "adinall"
	Gdt         = "gdt"
	Letv        = "letv"
	Jrtt        = "tt"
	Xt          = "xt"
	Zyzn        = "zyzn"
	Meiyou      = "meiyou"
	Mi          = "mi"
	Momo        = "momo"    //陌陌直投
	MomoAdx     = "momoadx" //陌陌ADX
	Sohuserving = "sohuserving"
	Motv        = "motv"
	Innersdk    = "innersdk"
	Openrtb     = "openrtb"

	Zadx              = "zadx"
	ZadxIfly          = "ifly"
	ZadxVamaker       = "vamaker"
	ZadxNativex       = "nativex"
	ZadxWax           = "wax"
	ZadxZplay         = "zplay"
	ZadxAutohome      = "autohome"
	ZadxMediamax      = "mediamax"
	ZadxYoukuRtb      = "youkurtb"
	ZadxFun           = "funadx"
	ZadxIqiyi         = "qax"
	ZadxLetv          = "letvadx"
	ZadxYoukuPdb      = "youkupdb"
	ZadxMgtv          = "mgtv"
	ZadxBaidu         = "bes"
	ZadxIax           = "iax"
	ZadxMMVideo       = "mmvideo"
	ZadxGoogle        = "googleadx"
	ZadxFtxTencentPdb = "adxpdb"
	ZadxFtxQaxPdb     = "qaxpdb"
	ZadxFtxLetvPdb    = "letvadxpdb"
	ZadxFtxSohuPdb    = "sohupdb"
)

var(
	printReq bool
)


func init() {
	flag.BoolVar(&printReq,"d",true,"print req info")
}


func main() {
	engine := gin.New()
	vi := engine.Group("/")
	vi.Any("/*action", WithHeader)
	err := engine.Run(":8341")
	if err != nil {
		fmt.Println(err)
	}
}

func WithHeader(ctx *gin.Context) {
	ctx.Request.Header.Add("requester-uid", "id")
	httpDo(ctx.Request)
}


func httpDo(bidReq *http.Request) {
	if bidReq == nil {
		fmt.Println("ffffffff")
		return
	}
	bidReq.URL.Scheme = "http"
	bidReq.URL.Host = Host
	bidReq.Host = Host
	client := &http.Client{}
	req, _ := http.NewRequest(bidReq.Method, bidReq.URL.String(), bidReq.Body)
	for key, values := range bidReq.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp == nil || resp.Body == nil {
		fmt.Println("data is null")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	handlerSid := resp.Request.URL.Path
	handlerSid = strings.Trim(handlerSid, "/")
	contentType := resp.Header.Get(KcontentType)
	if len(contentType) ==0 {
		contentType = resp.Request.Header.Get(KcontentType)
	}
	// 尽量用容易读的方式输出，json直接输出，protobuf则输出解码后的数据
	switch contentType {
	case ContentTypeJson:
		fmt.Println("json：",string(body))
	case ContentTypeProtobuf, ContentTypeOctetStream:
		rspMessage := getPbMessage(handlerSid, kLogTypeRsp)
		if rspMessage == nil {
			fmt.Println(base64.StdEncoding.EncodeToString(body))
		} else {
			if err := proto.Unmarshal(body, rspMessage); err != nil {
				fmt.Println(err.Error())
			} else {
				jsonBytes, _ := json.Marshal(rspMessage)
				if printReq {
					fmt.Println(bidReq.Method,bidReq.URL.String())
					fmt.Println(base64.StdEncoding.EncodeToString(body))
				}
				fmt.Println("proto：",string(jsonBytes))
			}
		}
	default:
		fmt.Println("unsupported protocol:" + contentType)
	}
}


func getPbMessage(sid string, logType string) proto.Message {
	switch sid {
	case Gdt:
		switch logType {
		case kLogTypeReq:
			return &gdt.BidRequest{}
		case kLogTypeRsp, kLogTypeThirdRsp:
			return &gdt.BidResponse{}
		}
	case Adx:
		switch logType {
		case kLogTypeReq:
			return &adx.Request{}
		case kLogTypeRsp, kLogTypeThirdRsp:
			return &adx.Response{}
		}
	case Sohudsp:
		switch logType {
		case kLogTypeReq:
			return &sohudsp.Request{}
		case kLogTypeRsp, kLogTypeThirdRsp:
			return &sohudsp.Response{}
		}
	case Tanx:
		switch logType {
		case kLogTypeReq:
			return &tanx.BidRequest{}
		case kLogTypeRsp, kLogTypeThirdRsp:
			return &tanx.BidResponse{}
		}
	case ZadxYoukuPdb:
		if logType == kLogTypeReq {
			return &ftx.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &ftx.BidResponse{}
		}
	case ZadxFtxTencentPdb:
		if logType == kLogTypeReq {
			return &ftx.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &ftx.BidResponse{}
		}
	case ZadxFtxQaxPdb:
		if logType == kLogTypeReq {
			return &ftx.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &ftx.BidResponse{}
		}
	case ZadxBaidu:
		if logType == kLogTypeReq {
			return &baidu.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &baidu.BidResponse{}
		}
	case ZadxIax:
		if logType == kLogTypeReq {
			return &iax.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &iax.BidResponse{}
		}
	case ZadxIqiyi:
		if logType == kLogTypeReq {
			return &qax.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &qax.BidResponse{}
		}
	case ZadxVamaker:
		if logType == kLogTypeReq {
			return &vam.VamRequest{}
		}
		if logType == kLogTypeRsp {
			return &vam.VamResponse{}
		}
	case ZadxZplay:
		if logType == kLogTypeReq {
			return &zplay.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &zplay.BidResponse{}
		}
	case ZadxGoogle:
		if logType == kLogTypeReq {
			return &googleadx.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &googleadx.BidResponse{}
		}
	case Adinall:
		if logType == kLogTypeReq {
			return &adinall.BidRequest{}
		}
		if logType == kLogTypeRsp {
			return &adinall.BidResponse{}
		}
	}

	return nil
}
