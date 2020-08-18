package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/whileW/enze-global/utils"
	"time"
)

//获取服务商凭证
//provider_access_token至少保留512字节的存储空间。
const get_provider_token_url = "https://qyapi.weixin.qq.com/cgi-bin/service/get_provider_token"

type resp_get_provider_token struct {
	ProviderAccessToken 		string			`json:"provider_access_token"`
	ExpiresIn					int				`json:"expires_in"`
	Errcode						int				`json:"errcode"`
	Errmsg						string			`json:"errmsg"`
}
type req_get_provider_token struct {
	Corpid 				string			`json:"corpid"`
	ProviderSecret		string			`json:"provider_secret"`
}
type provider_token_struct struct {
	auth 			*resp_get_provider_token
	expires			time.Time
}
var provider_tokens = map[string]*provider_token_struct{}

func GetProviderTokenById(id int) (*resp_get_provider_token,error) {
	provider,err := GetInfoById(id)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("db err:%v",err))
	}
	return GetProviderToken(provider.Corpid,provider.ProviderSecret)
}
func GetProviderToken(corpid,provider_secret string) (*resp_get_provider_token,error) {
	var token *provider_token_struct = nil
	if t,ok := provider_tokens[corpid];ok {
		if t != nil && time.Now().After(t.expires) {
			token = t
		}
	}
	if token == nil {
		resp, err := get_provider_token_req(corpid,provider_secret)
		if err == nil {
			token = &provider_token_struct{
				auth:    resp,
				expires: time.Now().Add(time.Duration(resp.ExpiresIn / 3 * 2)),
			}
			provider_tokens[corpid] = token
		}else {
			return nil,err
		}
	}
	return token.auth,nil
}
func RefreshProviderToken(corpid,provider_secret string) (*resp_get_provider_token,error) {
	provider_tokens[corpid] = nil
	var token *provider_token_struct = nil

	resp, err := get_provider_token_req(corpid,provider_secret)
	if err == nil {
		token = &provider_token_struct{
			auth:    resp,
			expires: time.Now().Add(time.Duration(resp.ExpiresIn / 3 * 2)),
		}
		provider_tokens[corpid] = token
		return token.auth,err
	}
	return nil,err
}
func get_provider_token_req(corpid,provider_secret string) (*resp_get_provider_token,error) {
	req := &req_get_provider_token{
		Corpid:corpid,
		ProviderSecret:provider_secret,
	}
	result,err := utils.PostWithJson(get_provider_token_url,"json",req)
	if err != nil {
		return nil,err
	}

	resp_data := &resp_get_provider_token{}
	err = json.Unmarshal(result,resp_data)
	if err != nil {
		return nil, err
	}
	if resp_data.Errcode != 0 {
		return nil,errors.New(resp_data.Errmsg)
	}
	return resp_data,nil
}