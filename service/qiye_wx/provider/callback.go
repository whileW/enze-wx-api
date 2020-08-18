package provider

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/whileW/enze-global"
	"github.com/whileW/enze-global/utils"
	"github.com/whileW/enze-wx-api/model/request"
	"sort"
)

// 企业微信服务商回调处理
func CallBack(req *request.QiYeWxProviderCallBack) (string,error) {
	provider,err := GetInfoBySuiteId(req.AgentId)
	if err != nil {
		return "",err
	}
	//验证签名
	mysignature := dev_msg_signature(provider.Token,req.Timestamp,req.Nonce,req.MsgSignature)
	if req.MsgSignature != mysignature {
		global.GVA_LOG.Errorw("错误签名",
			"w-sign",req.MsgSignature,
			"m-sign",mysignature)
		return "",errors.New("错误签名")
	}
	//解密报文
	btmpMsg,_,err := qiyewx_aes_decrypt(req.Echostr,provider.EncodingAESKey)
	if err != nil {
		return "",errors.New(fmt.Sprintf("解密失败：%v",err))
	}
	req_struct := utils.StringMap{}
	xml.Unmarshal([]byte(btmpMsg),(*utils.StringMap)(&req_struct))
	//处理消息
	resp_str := ""
	switch req_struct["InfoType"] {
	case "suite_ticket":
		provider.SuiteTicket = req_struct["SuiteTicket"]
		if err := provider.Save();err != nil {
			resp_str = "fail"
			return resp_str,err
		}else {
			resp_str = "success"
		}
		break
	default:
		resp_str = btmpMsg
		break
	}
	return resp_str,nil
}

func dev_msg_signature(token,timestamp,nonce,msg_encrypt string) string {
	signature_arr := []string{token,timestamp,nonce,msg_encrypt}
	sort.Strings(signature_arr)
	raw := ""
	for _,t := range signature_arr {
		raw+=t
	}
	h := sha1.New()
	h.Write([]byte(raw))
	l := fmt.Sprintf("%x", h.Sum(nil))
	return l
}
func qiyewx_aes_decrypt(str,aes_key string) (string,string,error) {
	key,_ := base64.StdEncoding.DecodeString(aes_key+"=")
	iv := key[:16]
	btmpMsg,err := utils.AesDecrypt(str, iv, key)
	if err != nil {
		return "","",err
	}
	if len(btmpMsg)<16 {
		return "","",errors.New("err")
	}
	content := btmpMsg[16:]
	len := int32(binary.BigEndian.Uint32(content[0:4]))
	msg := content[4:len+4]
	corpid := content[len+4:]
	return string(msg),string(corpid),nil
}