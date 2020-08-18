package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/whileW/enze-global/utils"
	"time"
)

//获取第三方应用凭证

// 由于第三方服务商可能托管了大量的企业，其安全问题造成的影响会更加严重，故API中除了合法来源IP校验之外，
// 还额外增加了suite_ticket作为安全凭证。
// 获取suite_access_token时，需要suite_ticket参数。suite_ticket由企业微信后台定时推送给“指令回调URL”，
// 每十分钟更新一次，见推送suite_ticket。
// suite_ticket实际有效期为30分钟，可以容错连续两次获取suite_ticket失败的情况，
// 但是请永远使用最新接收到的suite_ticket。
// 通过本接口获取的suite_access_token有效期为2小时，开发者需要进行缓存，不可频繁获取。

const get_suite_token_url = "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"
//const get_suite_token_url = "http://127.0.0.1:8080/user/test"

type resp_get_suite_token struct {
	SuiteAccessToken 		string				`json:"suite_access_token"`
	ExpiresIn					int				`json:"expires_in"`
	Errcode						int				`json:"errcode"`
	Errmsg						string			`json:"errmsg"`
}
type req_get_suite_token struct {
	SuiteId 				string				`json:"suite_id"`
	SuiteSecret				string				`json:"suite_secret"`
	SuiteTicket				string				`json:"suite_ticket"`
}
type suite_token_struct struct {
	auth 			*resp_get_suite_token
	expires			time.Time
}
var suite_tokens = map[string]*suite_token_struct{}

func GetSuiteTokenById(id int) (*resp_get_suite_token,error) {
	provider,err := GetInfoById(id)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("db err:%v",err))
	}
	return GetSuiteToken(provider.Suiteid,provider.SuiteSecret,provider.SuiteTicket)
}
func GetSuiteToken(suiteid,suite_secret,suite_ticket string) (*resp_get_suite_token,error) {
	var token *suite_token_struct = nil
	if t,ok := suite_tokens[suiteid];ok {
		if t != nil && time.Now().After(t.expires) {
			token = t
		}
	}
	if token == nil {
		resp, err := get_suite_token_req(suiteid,suite_secret,suite_ticket)
		if err == nil {
			token = &suite_token_struct{
				auth:    resp,
				expires: time.Now().Add(time.Duration(resp.ExpiresIn / 3 * 2)),
			}
			suite_tokens[suiteid] = token
		}else {
			return nil,err
		}
	}
	return token.auth,nil
}
func RefreshSuiteToken(suiteid,suite_secret,suite_ticket string) (*resp_get_suite_token,error) {
	suite_tokens[suiteid] = nil
	var token *suite_token_struct = nil

	resp, err := get_suite_token_req(suiteid,suite_secret,suite_ticket)
	if err == nil {
		token = &suite_token_struct{
			auth:    resp,
			expires: time.Now().Add(time.Duration(resp.ExpiresIn / 3 * 2)),
		}
		suite_tokens[suiteid] = token
		return token.auth,err
	}
	return nil,err
}
func get_suite_token_req(suiteid,suite_secret,suite_ticket string) (*resp_get_suite_token,error) {
	req := &req_get_suite_token{
		SuiteId:suiteid,
		SuiteSecret:suite_secret,
		SuiteTicket:suite_ticket,
	}
	hc,err := utils.NewHttpClientWithJson(utils.POST,get_suite_token_url,req)
	if err != nil {
		return nil,err
	}
	result,err := hc.SetHeader("Accept","*/*").DelHeader("Accept-Encoding").Do()
	if err != nil {
		return nil,err
	}

	resp_data := &resp_get_suite_token{}
	err = json.Unmarshal(result,resp_data)
	if err != nil {
		return nil, err
	}
	if resp_data.Errcode != 0 {
		return nil,errors.New(resp_data.Errmsg)
	}
	return resp_data,nil
}